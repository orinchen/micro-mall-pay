package service

import (
	"context"
	"fmt"
	"gitee.com/cristiane/micro-mall-pay/model/args"
	"gitee.com/cristiane/micro-mall-pay/model/mysql"
	"gitee.com/cristiane/micro-mall-pay/pkg/code"
	"gitee.com/cristiane/micro-mall-pay/pkg/util"
	"gitee.com/cristiane/micro-mall-pay/proto/micro_mall_pay_proto/pay_business"
	"gitee.com/cristiane/micro-mall-pay/repository"
	"gitee.com/cristiane/micro-mall-pay/vars"
	"gitee.com/kelvins-io/common/errcode"
	"gitee.com/kelvins-io/common/json"
	"gitee.com/kelvins-io/kelvins"
	"github.com/shopspring/decimal"
	"strings"
	"time"
	"xorm.io/xorm"
)

func TradePay(ctx context.Context, req *pay_business.TradePayRequest) (payId string, retCode int) {
	retCode = code.Success
	// 支付状态检查
	retCode = tradePayCheckState(ctx, req)
	if retCode != code.Success {
		return
	}

	// 长事务，多次扣减用户账户在一个事务中完成
	tx := kelvins.XORM_DBEngine.NewSession()
	err := tx.Begin()
	if err != nil {
		kelvins.ErrLogger.Errorf(ctx, "TradePay Begin err: %v", err)
		retCode = code.ErrorServer
		return
	}
	// 检查用户账户余额
	userAccount, retCode := tradePayCheckUserAccount(ctx, tx, req)
	if retCode != code.Success {
		return
	}

	// 依次支付
	payId = util.GetUUID() // 同一批订单支付交易号唯一
	for i := 0; i < len(req.EntryList); i++ {
		retCode = tradePayOne(ctx, payId, req, i, tx, userAccount)
		if retCode != code.Success {
			return
		}
	}

	err = tx.Commit()
	if err != nil {
		kelvins.ErrLogger.Errorf(ctx, "TradePay Commit err: %v", err)
		retCode = code.ErrorServer
		return
	}
	// 触发支付事件通知
	go tradeEventNotice(ctx, req, payId)

	return
}

func tradeEventNotice(ctx context.Context, req *pay_business.TradePayRequest, payId string) {
	// 触发支付消息
	pushSer := NewPushNoticeService(vars.TradePayQueueServer, PushMsgTag{
		DeliveryTag:    args.TaskNameTradePayNotice,
		DeliveryErrTag: args.TaskNameTradePayNoticeErr,
		RetryCount:     kelvins.QueueAMQPSetting.TaskRetryCount,
		RetryTimeout:   kelvins.QueueAMQPSetting.TaskRetryTimeout,
	})
	businessMsg := args.CommonBusinessMsg{
		Type: args.TradePayEventTypeCreate,
		Tag:  args.GetMsg(args.TradePayEventTypeCreate),
		UUID: util.GetUUID(),
		Msg: json.MarshalToStringNoError(args.TradePayNotice{
			Uid:    req.OpUid,
			Time:   util.ParseTimeOfStr(time.Now().Unix()),
			PayId:  payId,
			TxCode: req.OutTxCode,
		}),
	}
	taskUUID, retCode := pushSer.PushMessage(ctx, businessMsg)
	if retCode != code.Success {
		kelvins.ErrLogger.Errorf(ctx, "trade pay businessMsg: %+v  notice send err: ", businessMsg, errcode.GetErrMsg(retCode))
	} else {
		kelvins.BusinessLogger.Infof(ctx, "trade pay businessMsg businessMsg: %+v  taskUUID :%v", businessMsg, taskUUID)
	}
}

func tradePayOne(ctx context.Context, payId string, req *pay_business.TradePayRequest, i int, tx *xorm.Session, userAccount *mysql.Account) int {
	// 生成支付记录
	payRecord := mysql.PayRecord{
		TxId:        payId,
		OutTradeNo:  req.EntryList[i].OutTradeNo,
		TimeExpire:  time.Now().Add(30 * time.Minute),
		NotifyUrl:   req.EntryList[i].NotifyUrl,
		Description: req.EntryList[i].Description,
		Merchant:    req.EntryList[i].Merchant,
		Attach:      req.EntryList[i].Attach,
		User:        req.Account,
		Amount:      req.EntryList[i].Detail.Amount,
		Reduction:   req.EntryList[i].Detail.Reduction,
		CoinType:    int(req.CoinType),
		PayType:     1,
		PayState:    3,
		CreateTime:  time.Now(),
		UpdateTime:  time.Now(),
	}
	err := repository.CreatePayRecord(tx, &payRecord)
	if err != nil {
		errRollback := tx.Rollback()
		if errRollback != nil {
			kelvins.ErrLogger.Errorf(ctx, "GetAccount Rollback err: %v", errRollback)
		}
		kelvins.ErrLogger.Errorf(ctx, "CreatePayRecord err: %v, payRecord: %v", err, payRecord)
		return code.ErrorServer
	}
	reqAmount, err := decimal.NewFromString(req.EntryList[i].Detail.Amount)
	if err != nil {
		errRollback := tx.Rollback()
		if errRollback != nil {
			kelvins.ErrLogger.Errorf(ctx, "GetAccount Rollback err: %v", errRollback)
		}
		kelvins.ErrLogger.Errorf(ctx, "NewFromString err: %v, amount: %v", err, req.EntryList[i].Detail.Amount)
		return code.DecimalParseErr
	}
	reduction, err := decimal.NewFromString(req.EntryList[i].Detail.Reduction)
	if err != nil {
		errRollback := tx.Rollback()
		if errRollback != nil {
			kelvins.ErrLogger.Errorf(ctx, "GetAccount Rollback err: %v", errRollback)
		}
		kelvins.ErrLogger.Errorf(ctx, "NewFromString err: %v, amount: %v", err, req.EntryList[i].Detail.Reduction)
		return code.DecimalParseErr
	}
	userBalance, err := decimal.NewFromString(userAccount.Balance)
	if err != nil {
		errRollback := tx.Rollback()
		if errRollback != nil {
			kelvins.ErrLogger.Errorf(ctx, "TradePay Rollback err: %v", errRollback)
		}
		kelvins.ErrLogger.Errorf(ctx, "TradePay NewFromString err: %v, number: %v", err, userAccount.Balance)
		return code.DecimalParseErr
	}
	merchantAccount, err := repository.GetAccountByTx(tx, req.EntryList[i].Merchant, args.AccountTypeCompany, int(req.CoinType))
	if err != nil {
		errRollback := tx.Rollback()
		if errRollback != nil {
			kelvins.ErrLogger.Errorf(ctx, "GetAccount Rollback err: %v", errRollback)
		}
		kelvins.ErrLogger.Errorf(ctx, "GetAccount err: %v, owner: %v", err, req.EntryList[i].Merchant)
		return code.ErrorServer
	}
	if merchantAccount.Owner == "" {
		errRollback := tx.Rollback()
		if errRollback != nil {
			kelvins.ErrLogger.Errorf(ctx, "GetAccount Rollback err: %v", errRollback)
		}
		return code.MerchantAccountNotExist
	}
	if merchantAccount.State != 3 {
		errRollback := tx.Rollback()
		if errRollback != nil {
			kelvins.ErrLogger.Errorf(ctx, "GetAccount Rollback err: %v", errRollback)
		}
		return code.MerchantAccountStateLock
	}
	merchantBalance, err := decimal.NewFromString(merchantAccount.Balance)
	if err != nil {
		errRollback := tx.Rollback()
		if errRollback != nil {
			kelvins.ErrLogger.Errorf(ctx, "GetAccount Rollback err: %v", errRollback)
		}
		kelvins.ErrLogger.Errorf(ctx, "GetAccount err: %v, owner: %v", err, merchantAccount.Balance)
		return code.ErrorServer
	}
	// 生成交易流水
	fromBalance := util.DecimalSub(userBalance, util.DecimalSub(reqAmount, reduction))
	toBalance := util.DecimalAdd(merchantBalance, util.DecimalSub(reqAmount, reduction))
	transaction := mysql.Transaction{
		FromAccountCode: userAccount.AccountCode,
		FromBalance:     fromBalance.String(),
		ToAccountCode:   merchantAccount.AccountCode,
		ToBalance:       toBalance.String(),
		Amount:          util.DecimalSub(reqAmount, reduction).String(),
		Meta:            req.EntryList[i].Description,
		Scene:           req.EntryList[i].Description,
		OpUid:           req.OpUid,
		OpIp:            req.OpIp,
		TxId:            payId,
		Fingerprint:     time.Now().String(),
		PayType:         0,
		PayDesc:         "交易支付",
		CreateTime:      time.Now(),
		UpdateTime:      time.Now(),
	}
	err = repository.CreateTransaction(tx, &transaction)
	if err != nil {
		errRollback := tx.Rollback()
		if errRollback != nil {
			kelvins.ErrLogger.Errorf(ctx, "GetAccount Rollback err: %v", errRollback)
		}
		kelvins.ErrLogger.Errorf(ctx, "CreateTransaction err: %v, transaction: %+v", err, transaction)
		return code.ErrorServer
	}
	// 扣减用余额，增加商余额
	whereUserAccount := map[string]interface{}{
		"owner":      userAccount.Owner,
		"balance":    userAccount.Balance,
		"last_tx_id": userAccount.LastTxId, // 防止更新期间账户变更
	}
	userAccountChange := map[string]interface{}{
		"balance":     fromBalance.String(),
		"update_time": time.Now(),
		"last_tx_id":  payId, // 记录本次支付事务ID，对标支付记录
	}
	rowsAffected, err := repository.ChangeAccount(tx, whereUserAccount, userAccountChange)
	if err != nil {
		errRollback := tx.Rollback()
		if errRollback != nil {
			kelvins.ErrLogger.Errorf(ctx, "GetAccount Rollback err: %v", errRollback)
		}
		kelvins.ErrLogger.Errorf(ctx, "ChangeAccount err: %v, userAccountQ: %+v, userAccountChange: %+v", err, whereUserAccount, userAccountChange)
		return code.ErrorServer
	}
	fmt.Println("更新用户余额 rowsAffected ==", rowsAffected)
	// 没有符合条件的数据行，说明没有更新成功
	if rowsAffected != 1 {
		errRollback := tx.Rollback()
		if errRollback != nil {
			kelvins.ErrLogger.Errorf(ctx, "GetAccount Rollback err: %v", errRollback)
		}
		return code.TransactionFailed
	}
	// 更新扣减了余额后的用户账户
	userAccount.Balance = fromBalance.String() // 用户账户剩余金额
	userAccount.LastTxId = payId

	// 增加商户账户余额-，增加商户用户余额应该放在事务最后阶段
	whereMerchantAccount := map[string]interface{}{
		"owner":      merchantAccount.Owner,
		"balance":    merchantAccount.Balance,
		"last_tx_id": merchantAccount.LastTxId, // 防止更新期间账户变更
	}
	merchantAccountChange := map[string]interface{}{
		"balance":     toBalance.String(),
		"update_time": time.Now(),
		"last_tx_id":  payId, // 记录本次支付事务ID，对标支付记录
	}
	rowsAffected, err = repository.ChangeAccount(tx, whereMerchantAccount, merchantAccountChange)
	if err != nil {
		errRollback := tx.Rollback()
		if errRollback != nil {
			kelvins.ErrLogger.Errorf(ctx, "GetAccount Rollback err: %v", errRollback)
		}
		kelvins.ErrLogger.Errorf(ctx, "ChangeAccount err: %v, userAccountQ: %+v, userAccountChange: %+v", err, whereMerchantAccount, userAccountChange)
		return code.ErrorServer
	}
	fmt.Println("更新商户余额 rowsAffected ==", rowsAffected)
	// 没有符合条件的数据行，说明没有更新成功
	if rowsAffected != 1 {
		errRollback := tx.Rollback()
		if errRollback != nil {
			kelvins.ErrLogger.Errorf(ctx, "GetAccount Rollback err: %v", errRollback)
		}
		return code.TransactionFailed
	}

	return code.Success
}

func tradePayCheckUserAccount(ctx context.Context, tx *xorm.Session, req *pay_business.TradePayRequest) (*mysql.Account, int) {
	userAccount, err := repository.GetAccountByTx(tx, req.Account, args.AccountTypePerson, int(req.CoinType))
	if err != nil {
		errRollback := tx.Rollback()
		if errRollback != nil {
			kelvins.ErrLogger.Errorf(ctx, "GetAccount Rollback err: %v", errRollback)
		}
		kelvins.ErrLogger.Errorf(ctx, "GetAccount err: %v, owner: %v", err, req.Account)
		return userAccount, code.ErrorServer
	}
	if userAccount.Owner == "" {
		errRollback := tx.Rollback()
		if errRollback != nil {
			kelvins.ErrLogger.Errorf(ctx, "GetAccount Rollback err: %v", errRollback)
		}
		return userAccount, code.UserAccountNotExist
	}
	if userAccount.State != 3 {
		errRollback := tx.Rollback()
		if errRollback != nil {
			kelvins.ErrLogger.Errorf(ctx, "GetAccount Rollback err: %v", errRollback)
		}
		return userAccount, code.UserAccountStateLock
	}
	// 检查用户账户余额
	userBalance, err := decimal.NewFromString(userAccount.Balance)
	if err != nil {
		errRollback := tx.Rollback()
		if errRollback != nil {
			kelvins.ErrLogger.Errorf(ctx, "GetAccount Rollback err: %v", errRollback)
		}
		kelvins.ErrLogger.Errorf(ctx, "NewFromString err: %v, number: %v", err, userAccount.Balance)
		return userAccount, code.DecimalParseErr
	}
	totalAmount, _ := decimal.NewFromString("0")
	for i := 0; i < len(req.EntryList); i++ {
		amount := req.EntryList[i].Detail.Amount
		amountDecimal, err := decimal.NewFromString(amount)
		if err != nil {
			errRollback := tx.Rollback()
			if errRollback != nil {
				kelvins.ErrLogger.Errorf(ctx, "GetAccount Rollback err: %v", errRollback)
			}
			kelvins.ErrLogger.Errorf(ctx, "NewFromString err: %v, amount: %v", err, amount)
			return userAccount, code.DecimalParseErr
		}
		totalAmount = util.DecimalAdd(totalAmount, amountDecimal)
	}
	if !util.DecimalGreaterThanOrEqual(userBalance, totalAmount) {
		errRollback := tx.Rollback()
		if errRollback != nil {
			kelvins.ErrLogger.Errorf(ctx, "GetAccount Rollback err: %v", errRollback)
		}
		return userAccount, code.UserAccountNotEnough
	}
	return userAccount, code.Success
}

func tradePayCheckState(ctx context.Context, req *pay_business.TradePayRequest) (retCode int) {
	retCode = code.Success
	// 参数验证
	outTradeNoList := make([]string, len(req.EntryList))
	for i := 0; i < len(req.EntryList); i++ {
		outTradeNoList[i] = req.EntryList[i].OutTradeNo
	}
	where := map[string]interface{}{}
	payRecordList, _, err := repository.GetPayRecordList("out_trade_no,pay_state", where, outTradeNoList, nil, nil, 0, 0)
	if err != nil {
		kelvins.ErrLogger.Errorf(ctx, "GetPayRecordList err: %v, outTradeNoList: %v", err, outTradeNoList)
		retCode = code.ErrorServer
		return
	}
	for i := 0; i < len(payRecordList); i++ {
		if payRecordList[i].PayState == 1 {
			retCode = code.TradePayRun
			return
		}
		if payRecordList[i].PayState == 3 {
			retCode = code.TradePaySuccess
			return
		}
		// 超过支付过期时间
		if time.Now().Sub(payRecordList[i].TimeExpire) >= 0 {
			retCode = code.TradePayExpire
			return
		}
	}
	return
}

func CreateAccount(ctx context.Context, req *pay_business.CreateAccountRequest) (accountCode string, retCode int) {
	retCode = code.Success
	accountCode = util.GetUUID()
	account := mysql.Account{
		AccountCode: accountCode,
		Owner:       req.Owner,
		Balance:     req.Balance,
		CoinType:    int(req.CoinType),
		CoinDesc:    "CNY",
		State:       3,
		AccountType: int(req.AccountType) + 1,
		LastTxId:    accountCode, // 初始值等于AccountCode
		CreateTime:  time.Now(),
		UpdateTime:  time.Now(),
	}
	err := repository.CreateAccount(&account)
	if err != nil {
		if strings.Contains(err.Error(), errcode.GetErrMsg(code.DBDuplicateEntry)) {
			retCode = code.AccountExist
			return
		}
		kelvins.ErrLogger.Errorf(ctx, "CreateAccount err: %v, account: %+v", err, account)
		retCode = code.ErrorServer
	}
	return
}
