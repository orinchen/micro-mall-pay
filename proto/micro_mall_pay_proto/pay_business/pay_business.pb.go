// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/micro_mall_pay_proto/pay_business/pay_business.proto

package pay_business

import (
	context "context"
	fmt "fmt"
	_ "gitee.com/kelvins-io/common/proto/google/api"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type RetCode int32

const (
	RetCode_SUCCESS                     RetCode = 0
	RetCode_ERROR                       RetCode = 500
	RetCode_USER_NOT_EXIST              RetCode = 50001
	RetCode_USER_EXIST                  RetCode = 50002
	RetCode_MERCHANT_NOT_EXIST          RetCode = 50003
	RetCode_MERCHANT_EXIST              RetCode = 50004
	RetCode_SHOP_NOT_EXIST              RetCode = 50005
	RetCode_SHOP_EXIST                  RetCode = 50006
	RetCode_SKU_NOT_EXIST               RetCode = 50007
	RetCode_SKU_EXIST                   RetCode = 50008
	RetCode_SKU_AMOUNT_NOT_ENOUGH       RetCode = 50009
	RetCode_USER_BALANCE_NOT_ENOUGH     RetCode = 600000
	RetCode_MERCHANT_BALANCE_NOT_ENOUGH RetCode = 6000001
	RetCode_ACCOUNT_LOCK                RetCode = 6000002
)

var RetCode_name = map[int32]string{
	0:       "SUCCESS",
	500:     "ERROR",
	50001:   "USER_NOT_EXIST",
	50002:   "USER_EXIST",
	50003:   "MERCHANT_NOT_EXIST",
	50004:   "MERCHANT_EXIST",
	50005:   "SHOP_NOT_EXIST",
	50006:   "SHOP_EXIST",
	50007:   "SKU_NOT_EXIST",
	50008:   "SKU_EXIST",
	50009:   "SKU_AMOUNT_NOT_ENOUGH",
	600000:  "USER_BALANCE_NOT_ENOUGH",
	6000001: "MERCHANT_BALANCE_NOT_ENOUGH",
	6000002: "ACCOUNT_LOCK",
}

var RetCode_value = map[string]int32{
	"SUCCESS":                     0,
	"ERROR":                       500,
	"USER_NOT_EXIST":              50001,
	"USER_EXIST":                  50002,
	"MERCHANT_NOT_EXIST":          50003,
	"MERCHANT_EXIST":              50004,
	"SHOP_NOT_EXIST":              50005,
	"SHOP_EXIST":                  50006,
	"SKU_NOT_EXIST":               50007,
	"SKU_EXIST":                   50008,
	"SKU_AMOUNT_NOT_ENOUGH":       50009,
	"USER_BALANCE_NOT_ENOUGH":     600000,
	"MERCHANT_BALANCE_NOT_ENOUGH": 6000001,
	"ACCOUNT_LOCK":                6000002,
}

func (x RetCode) String() string {
	return proto.EnumName(RetCode_name, int32(x))
}

func (RetCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d451b8767e35fec5, []int{0}
}

type OperationType int32

const (
	OperationType_CREATE   OperationType = 0
	OperationType_UPDATE   OperationType = 1
	OperationType_DELETE   OperationType = 2
	OperationType_AUDIT    OperationType = 3
	OperationType_PUT_AWAY OperationType = 4
)

var OperationType_name = map[int32]string{
	0: "CREATE",
	1: "UPDATE",
	2: "DELETE",
	3: "AUDIT",
	4: "PUT_AWAY",
}

var OperationType_value = map[string]int32{
	"CREATE":   0,
	"UPDATE":   1,
	"DELETE":   2,
	"AUDIT":    3,
	"PUT_AWAY": 4,
}

func (x OperationType) String() string {
	return proto.EnumName(OperationType_name, int32(x))
}

func (OperationType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d451b8767e35fec5, []int{1}
}

type CoinType int32

const (
	CoinType_CNY CoinType = 0
	CoinType_USD CoinType = 1
)

var CoinType_name = map[int32]string{
	0: "CNY",
	1: "USD",
}

var CoinType_value = map[string]int32{
	"CNY": 0,
	"USD": 1,
}

func (x CoinType) String() string {
	return proto.EnumName(CoinType_name, int32(x))
}

func (CoinType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d451b8767e35fec5, []int{2}
}

type CommonResponse struct {
	Code                 RetCode  `protobuf:"varint,1,opt,name=code,proto3,enum=pay_business.RetCode" json:"code,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommonResponse) Reset()         { *m = CommonResponse{} }
func (m *CommonResponse) String() string { return proto.CompactTextString(m) }
func (*CommonResponse) ProtoMessage()    {}
func (*CommonResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d451b8767e35fec5, []int{0}
}

func (m *CommonResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommonResponse.Unmarshal(m, b)
}
func (m *CommonResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommonResponse.Marshal(b, m, deterministic)
}
func (m *CommonResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommonResponse.Merge(m, src)
}
func (m *CommonResponse) XXX_Size() int {
	return xxx_messageInfo_CommonResponse.Size(m)
}
func (m *CommonResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CommonResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CommonResponse proto.InternalMessageInfo

func (m *CommonResponse) GetCode() RetCode {
	if m != nil {
		return m.Code
	}
	return RetCode_SUCCESS
}

func (m *CommonResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type TradePayRequest struct {
	OutTradeNo           string            `protobuf:"bytes,1,opt,name=out_trade_no,json=outTradeNo,proto3" json:"out_trade_no,omitempty"`
	TimeExpire           string            `protobuf:"bytes,2,opt,name=time_expire,json=timeExpire,proto3" json:"time_expire,omitempty"`
	NotifyUrl            string            `protobuf:"bytes,3,opt,name=notify_url,json=notifyUrl,proto3" json:"notify_url,omitempty"`
	Description          string            `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Merchant             string            `protobuf:"bytes,5,opt,name=merchant,proto3" json:"merchant,omitempty"`
	Attach               string            `protobuf:"bytes,6,opt,name=attach,proto3" json:"attach,omitempty"`
	Account              string            `protobuf:"bytes,7,opt,name=account,proto3" json:"account,omitempty"`
	Detail               *TradeGoodsDetail `protobuf:"bytes,199,opt,name=detail,proto3" json:"detail,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *TradePayRequest) Reset()         { *m = TradePayRequest{} }
func (m *TradePayRequest) String() string { return proto.CompactTextString(m) }
func (*TradePayRequest) ProtoMessage()    {}
func (*TradePayRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d451b8767e35fec5, []int{1}
}

func (m *TradePayRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TradePayRequest.Unmarshal(m, b)
}
func (m *TradePayRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TradePayRequest.Marshal(b, m, deterministic)
}
func (m *TradePayRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TradePayRequest.Merge(m, src)
}
func (m *TradePayRequest) XXX_Size() int {
	return xxx_messageInfo_TradePayRequest.Size(m)
}
func (m *TradePayRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TradePayRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TradePayRequest proto.InternalMessageInfo

func (m *TradePayRequest) GetOutTradeNo() string {
	if m != nil {
		return m.OutTradeNo
	}
	return ""
}

func (m *TradePayRequest) GetTimeExpire() string {
	if m != nil {
		return m.TimeExpire
	}
	return ""
}

func (m *TradePayRequest) GetNotifyUrl() string {
	if m != nil {
		return m.NotifyUrl
	}
	return ""
}

func (m *TradePayRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *TradePayRequest) GetMerchant() string {
	if m != nil {
		return m.Merchant
	}
	return ""
}

func (m *TradePayRequest) GetAttach() string {
	if m != nil {
		return m.Attach
	}
	return ""
}

func (m *TradePayRequest) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *TradePayRequest) GetDetail() *TradeGoodsDetail {
	if m != nil {
		return m.Detail
	}
	return nil
}

type TradeGoodsDetail struct {
	Amount               int64    `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
	CoinType             CoinType `protobuf:"varint,2,opt,name=coin_type,json=coinType,proto3,enum=pay_business.CoinType" json:"coin_type,omitempty"`
	Reduction            int64    `protobuf:"varint,3,opt,name=reduction,proto3" json:"reduction,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TradeGoodsDetail) Reset()         { *m = TradeGoodsDetail{} }
func (m *TradeGoodsDetail) String() string { return proto.CompactTextString(m) }
func (*TradeGoodsDetail) ProtoMessage()    {}
func (*TradeGoodsDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_d451b8767e35fec5, []int{2}
}

func (m *TradeGoodsDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TradeGoodsDetail.Unmarshal(m, b)
}
func (m *TradeGoodsDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TradeGoodsDetail.Marshal(b, m, deterministic)
}
func (m *TradeGoodsDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TradeGoodsDetail.Merge(m, src)
}
func (m *TradeGoodsDetail) XXX_Size() int {
	return xxx_messageInfo_TradeGoodsDetail.Size(m)
}
func (m *TradeGoodsDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_TradeGoodsDetail.DiscardUnknown(m)
}

var xxx_messageInfo_TradeGoodsDetail proto.InternalMessageInfo

func (m *TradeGoodsDetail) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *TradeGoodsDetail) GetCoinType() CoinType {
	if m != nil {
		return m.CoinType
	}
	return CoinType_CNY
}

func (m *TradeGoodsDetail) GetReduction() int64 {
	if m != nil {
		return m.Reduction
	}
	return 0
}

type TradePayResponse struct {
	Common               *CommonResponse `protobuf:"bytes,1,opt,name=common,proto3" json:"common,omitempty"`
	TradeId              string          `protobuf:"bytes,2,opt,name=trade_id,json=tradeId,proto3" json:"trade_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *TradePayResponse) Reset()         { *m = TradePayResponse{} }
func (m *TradePayResponse) String() string { return proto.CompactTextString(m) }
func (*TradePayResponse) ProtoMessage()    {}
func (*TradePayResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d451b8767e35fec5, []int{3}
}

func (m *TradePayResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TradePayResponse.Unmarshal(m, b)
}
func (m *TradePayResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TradePayResponse.Marshal(b, m, deterministic)
}
func (m *TradePayResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TradePayResponse.Merge(m, src)
}
func (m *TradePayResponse) XXX_Size() int {
	return xxx_messageInfo_TradePayResponse.Size(m)
}
func (m *TradePayResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TradePayResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TradePayResponse proto.InternalMessageInfo

func (m *TradePayResponse) GetCommon() *CommonResponse {
	if m != nil {
		return m.Common
	}
	return nil
}

func (m *TradePayResponse) GetTradeId() string {
	if m != nil {
		return m.TradeId
	}
	return ""
}

type TradeRefundRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TradeRefundRequest) Reset()         { *m = TradeRefundRequest{} }
func (m *TradeRefundRequest) String() string { return proto.CompactTextString(m) }
func (*TradeRefundRequest) ProtoMessage()    {}
func (*TradeRefundRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d451b8767e35fec5, []int{4}
}

func (m *TradeRefundRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TradeRefundRequest.Unmarshal(m, b)
}
func (m *TradeRefundRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TradeRefundRequest.Marshal(b, m, deterministic)
}
func (m *TradeRefundRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TradeRefundRequest.Merge(m, src)
}
func (m *TradeRefundRequest) XXX_Size() int {
	return xxx_messageInfo_TradeRefundRequest.Size(m)
}
func (m *TradeRefundRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TradeRefundRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TradeRefundRequest proto.InternalMessageInfo

type TradeRefundResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TradeRefundResponse) Reset()         { *m = TradeRefundResponse{} }
func (m *TradeRefundResponse) String() string { return proto.CompactTextString(m) }
func (*TradeRefundResponse) ProtoMessage()    {}
func (*TradeRefundResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d451b8767e35fec5, []int{5}
}

func (m *TradeRefundResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TradeRefundResponse.Unmarshal(m, b)
}
func (m *TradeRefundResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TradeRefundResponse.Marshal(b, m, deterministic)
}
func (m *TradeRefundResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TradeRefundResponse.Merge(m, src)
}
func (m *TradeRefundResponse) XXX_Size() int {
	return xxx_messageInfo_TradeRefundResponse.Size(m)
}
func (m *TradeRefundResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TradeRefundResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TradeRefundResponse proto.InternalMessageInfo

type TradeQueryRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TradeQueryRequest) Reset()         { *m = TradeQueryRequest{} }
func (m *TradeQueryRequest) String() string { return proto.CompactTextString(m) }
func (*TradeQueryRequest) ProtoMessage()    {}
func (*TradeQueryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d451b8767e35fec5, []int{6}
}

func (m *TradeQueryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TradeQueryRequest.Unmarshal(m, b)
}
func (m *TradeQueryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TradeQueryRequest.Marshal(b, m, deterministic)
}
func (m *TradeQueryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TradeQueryRequest.Merge(m, src)
}
func (m *TradeQueryRequest) XXX_Size() int {
	return xxx_messageInfo_TradeQueryRequest.Size(m)
}
func (m *TradeQueryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TradeQueryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TradeQueryRequest proto.InternalMessageInfo

type TradeQueryResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TradeQueryResponse) Reset()         { *m = TradeQueryResponse{} }
func (m *TradeQueryResponse) String() string { return proto.CompactTextString(m) }
func (*TradeQueryResponse) ProtoMessage()    {}
func (*TradeQueryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d451b8767e35fec5, []int{7}
}

func (m *TradeQueryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TradeQueryResponse.Unmarshal(m, b)
}
func (m *TradeQueryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TradeQueryResponse.Marshal(b, m, deterministic)
}
func (m *TradeQueryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TradeQueryResponse.Merge(m, src)
}
func (m *TradeQueryResponse) XXX_Size() int {
	return xxx_messageInfo_TradeQueryResponse.Size(m)
}
func (m *TradeQueryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TradeQueryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TradeQueryResponse proto.InternalMessageInfo

type TradeCancelRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TradeCancelRequest) Reset()         { *m = TradeCancelRequest{} }
func (m *TradeCancelRequest) String() string { return proto.CompactTextString(m) }
func (*TradeCancelRequest) ProtoMessage()    {}
func (*TradeCancelRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d451b8767e35fec5, []int{8}
}

func (m *TradeCancelRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TradeCancelRequest.Unmarshal(m, b)
}
func (m *TradeCancelRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TradeCancelRequest.Marshal(b, m, deterministic)
}
func (m *TradeCancelRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TradeCancelRequest.Merge(m, src)
}
func (m *TradeCancelRequest) XXX_Size() int {
	return xxx_messageInfo_TradeCancelRequest.Size(m)
}
func (m *TradeCancelRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TradeCancelRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TradeCancelRequest proto.InternalMessageInfo

type TradeCancelResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TradeCancelResponse) Reset()         { *m = TradeCancelResponse{} }
func (m *TradeCancelResponse) String() string { return proto.CompactTextString(m) }
func (*TradeCancelResponse) ProtoMessage()    {}
func (*TradeCancelResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d451b8767e35fec5, []int{9}
}

func (m *TradeCancelResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TradeCancelResponse.Unmarshal(m, b)
}
func (m *TradeCancelResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TradeCancelResponse.Marshal(b, m, deterministic)
}
func (m *TradeCancelResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TradeCancelResponse.Merge(m, src)
}
func (m *TradeCancelResponse) XXX_Size() int {
	return xxx_messageInfo_TradeCancelResponse.Size(m)
}
func (m *TradeCancelResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TradeCancelResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TradeCancelResponse proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("pay_business.RetCode", RetCode_name, RetCode_value)
	proto.RegisterEnum("pay_business.OperationType", OperationType_name, OperationType_value)
	proto.RegisterEnum("pay_business.CoinType", CoinType_name, CoinType_value)
	proto.RegisterType((*CommonResponse)(nil), "pay_business.CommonResponse")
	proto.RegisterType((*TradePayRequest)(nil), "pay_business.TradePayRequest")
	proto.RegisterType((*TradeGoodsDetail)(nil), "pay_business.TradeGoodsDetail")
	proto.RegisterType((*TradePayResponse)(nil), "pay_business.TradePayResponse")
	proto.RegisterType((*TradeRefundRequest)(nil), "pay_business.TradeRefundRequest")
	proto.RegisterType((*TradeRefundResponse)(nil), "pay_business.TradeRefundResponse")
	proto.RegisterType((*TradeQueryRequest)(nil), "pay_business.TradeQueryRequest")
	proto.RegisterType((*TradeQueryResponse)(nil), "pay_business.TradeQueryResponse")
	proto.RegisterType((*TradeCancelRequest)(nil), "pay_business.TradeCancelRequest")
	proto.RegisterType((*TradeCancelResponse)(nil), "pay_business.TradeCancelResponse")
}

func init() {
	proto.RegisterFile("proto/micro_mall_pay_proto/pay_business/pay_business.proto", fileDescriptor_d451b8767e35fec5)
}

var fileDescriptor_d451b8767e35fec5 = []byte{
	// 815 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x94, 0x4d, 0x6f, 0xdc, 0x44,
	0x18, 0xc7, 0xbb, 0xeb, 0x64, 0x5f, 0x9e, 0xbc, 0xd4, 0x3c, 0x69, 0x8a, 0x49, 0x93, 0x12, 0x59,
	0x1c, 0x4a, 0x24, 0xb2, 0x22, 0x45, 0x42, 0xea, 0x01, 0xc9, 0xf5, 0x5a, 0x4d, 0xd4, 0x64, 0x77,
	0x19, 0xdb, 0x82, 0x9e, 0x2c, 0xd7, 0x9e, 0xa6, 0x23, 0x6c, 0x8f, 0xb1, 0xc7, 0x11, 0x7b, 0xe0,
	0x40, 0x0f, 0xfb, 0x05, 0x38, 0x96, 0xaf, 0x82, 0xc4, 0x8d, 0x23, 0x82, 0xf2, 0x7a, 0xe7, 0xdc,
	0x13, 0x1f, 0x00, 0x79, 0x3c, 0x9b, 0xee, 0x46, 0xe9, 0x6d, 0x9e, 0xdf, 0xf3, 0xd7, 0xff, 0xf1,
	0xcc, 0xfc, 0x3d, 0xf0, 0x20, 0x2f, 0xb8, 0xe0, 0x83, 0x94, 0x45, 0x05, 0x0f, 0xd2, 0x30, 0x49,
	0x82, 0x3c, 0x9c, 0x06, 0x0d, 0xac, 0x57, 0x4f, 0xab, 0x92, 0x65, 0xb4, 0x2c, 0x97, 0x8a, 0x43,
	0xd9, 0xc7, 0xf5, 0x45, 0xb6, 0xf3, 0xd9, 0x39, 0x13, 0x94, 0x1e, 0x46, 0x3c, 0x1d, 0x7c, 0x45,
	0x93, 0x0b, 0x96, 0x95, 0x1f, 0x31, 0x3e, 0x88, 0x78, 0x9a, 0xf2, 0x6c, 0xd0, 0x18, 0x9e, 0x73,
	0x7e, 0x9e, 0xd0, 0x41, 0x98, 0xb3, 0x41, 0x98, 0x65, 0x5c, 0x84, 0x82, 0xf1, 0x4c, 0xb9, 0x99,
	0x67, 0xb0, 0x69, 0x4b, 0x31, 0xa1, 0x65, 0xce, 0xb3, 0x92, 0xe2, 0x87, 0xb0, 0x12, 0xf1, 0x98,
	0x1a, 0xad, 0xfd, 0xd6, 0xbd, 0xcd, 0xa3, 0xed, 0xc3, 0xa5, 0x4f, 0x20, 0x54, 0xd8, 0x3c, 0xa6,
	0x44, 0x4a, 0x50, 0x07, 0x2d, 0x2d, 0xcf, 0x8d, 0xf6, 0x7e, 0xeb, 0x5e, 0x9f, 0xd4, 0x4b, 0xf3,
	0x87, 0x36, 0xdc, 0xf4, 0x8a, 0x30, 0xa6, 0x93, 0x70, 0x4a, 0xe8, 0xd7, 0x15, 0x2d, 0x05, 0xee,
	0xc3, 0x3a, 0xaf, 0x44, 0x20, 0x6a, 0x1c, 0x64, 0x5c, 0x1a, 0xf7, 0x09, 0xf0, 0x4a, 0x48, 0xe5,
	0x88, 0xe3, 0xfb, 0xb0, 0x26, 0x58, 0x4a, 0x03, 0xfa, 0x4d, 0xce, 0x0a, 0xaa, 0xfc, 0xa0, 0x46,
	0x8e, 0x24, 0xb8, 0x07, 0x90, 0x71, 0xc1, 0x9e, 0x4d, 0x83, 0xaa, 0x48, 0x0c, 0x4d, 0xf6, 0xfb,
	0x0d, 0xf1, 0x8b, 0x04, 0xf7, 0x61, 0x2d, 0xa6, 0x65, 0x54, 0xb0, 0xbc, 0xde, 0x9a, 0xb1, 0x22,
	0xfb, 0x8b, 0x08, 0x77, 0xa0, 0x97, 0xd2, 0x22, 0x7a, 0x1e, 0x66, 0xc2, 0x58, 0x95, 0xed, 0xcb,
	0x1a, 0x6f, 0x43, 0x27, 0x14, 0x22, 0x8c, 0x9e, 0x1b, 0x1d, 0xd9, 0x51, 0x15, 0x1a, 0xd0, 0x0d,
	0xa3, 0x88, 0x57, 0x99, 0x30, 0xba, 0xb2, 0x31, 0x2f, 0xf1, 0x53, 0xe8, 0xc4, 0x54, 0x84, 0x2c,
	0x31, 0x7e, 0xae, 0x37, 0xb3, 0x76, 0x74, 0x77, 0xf9, 0x94, 0xe4, 0xbe, 0x1e, 0x71, 0x1e, 0x97,
	0x43, 0x29, 0x23, 0x4a, 0x6e, 0x7e, 0x0b, 0xfa, 0xd5, 0x9e, 0x1c, 0x9f, 0xca, 0x29, 0xb5, 0x97,
	0x46, 0x54, 0x85, 0xf7, 0xa1, 0x1f, 0x71, 0x96, 0x05, 0x62, 0x9a, 0x37, 0x47, 0xb2, 0x79, 0x74,
	0x7b, 0x79, 0x8c, 0xcd, 0x59, 0xe6, 0x4d, 0x73, 0x4a, 0x7a, 0x91, 0x5a, 0xe1, 0x2e, 0xf4, 0x0b,
	0x1a, 0x57, 0x91, 0x3c, 0x07, 0x4d, 0xfa, 0xbd, 0x01, 0x66, 0xa4, 0xc6, 0xcb, 0xcb, 0x51, 0xd7,
	0xfd, 0x09, 0x74, 0x9a, 0xb4, 0x18, 0xcd, 0x56, 0x76, 0xaf, 0xce, 0x58, 0x0c, 0x07, 0x51, 0x5a,
	0x7c, 0x0f, 0x7a, 0xcd, 0x7d, 0xb2, 0x58, 0x5d, 0x57, 0x57, 0xd6, 0x27, 0xb1, 0x79, 0x0b, 0x50,
	0x0e, 0x21, 0xf4, 0x59, 0x95, 0xc5, 0x2a, 0x04, 0xe6, 0x36, 0x6c, 0x2d, 0xd1, 0xc6, 0xcf, 0xdc,
	0x82, 0x77, 0x24, 0xfe, 0xbc, 0xa2, 0xc5, 0x3c, 0x30, 0x97, 0x0e, 0x0a, 0x2a, 0xe9, 0x9c, 0xda,
	0x61, 0x16, 0xd1, 0xe4, 0xaa, 0xef, 0x9c, 0x36, 0xe2, 0x83, 0x1f, 0xdb, 0xd0, 0x55, 0x59, 0xc5,
	0x35, 0xe8, 0xba, 0xbe, 0x6d, 0x3b, 0xae, 0xab, 0xdf, 0x40, 0x80, 0x55, 0x87, 0x90, 0x31, 0xd1,
	0xff, 0xd3, 0xf0, 0x16, 0x6c, 0xfa, 0xae, 0x43, 0x82, 0xd1, 0xd8, 0x0b, 0x9c, 0x2f, 0x4f, 0x5c,
	0x4f, 0xff, 0x75, 0xa6, 0xa1, 0x0e, 0x20, 0x69, 0x43, 0x7e, 0x9b, 0x69, 0x68, 0x00, 0x9e, 0x39,
	0xc4, 0x3e, 0xb6, 0x46, 0xde, 0x82, 0xf6, 0xd5, 0x4c, 0x3a, 0x5c, 0x76, 0x1a, 0xfa, 0x7b, 0x43,
	0xdd, 0xe3, 0xf1, 0x64, 0x41, 0xfb, 0x47, 0xe3, 0x2b, 0x69, 0x43, 0xfe, 0x9c, 0x69, 0xb8, 0x05,
	0x1b, 0xee, 0x63, 0x7f, 0x41, 0xf6, 0xd7, 0x4c, 0xc3, 0x9b, 0xd0, 0xaf, 0x61, 0x03, 0xfe, 0x9e,
	0x69, 0x78, 0x07, 0xb6, 0x6b, 0x60, 0x9d, 0x8d, 0xfd, 0xf9, 0xfc, 0xd1, 0xd8, 0x7f, 0x74, 0xac,
	0xff, 0x33, 0xd3, 0x70, 0x0f, 0xde, 0x95, 0x1f, 0xfb, 0xd0, 0x3a, 0xb5, 0x46, 0xb6, 0xb3, 0xd8,
	0xfe, 0xe9, 0x97, 0x0f, 0xd0, 0x84, 0x3b, 0x97, 0xdf, 0x77, 0x8d, 0xe4, 0xbb, 0x97, 0xaf, 0xdb,
	0xb8, 0x05, 0xeb, 0x96, 0x6d, 0x4b, 0xf3, 0xd3, 0xb1, 0xfd, 0x58, 0x7f, 0xf1, 0xf2, 0x75, 0xfb,
	0xe0, 0x14, 0x36, 0xc6, 0x39, 0x2d, 0xe4, 0x53, 0x21, 0x83, 0x05, 0xd0, 0xb1, 0x89, 0x63, 0x79,
	0x8e, 0x3c, 0xc3, 0x8e, 0x3f, 0x19, 0xd6, 0xeb, 0x56, 0xbd, 0x1e, 0x3a, 0xa7, 0x8e, 0xe7, 0xe8,
	0x6d, 0xec, 0xc3, 0xaa, 0xe5, 0x0f, 0x4f, 0x3c, 0x5d, 0xc3, 0x75, 0xe8, 0x4d, 0x7c, 0x2f, 0xb0,
	0xbe, 0xb0, 0x9e, 0xe8, 0x2b, 0x07, 0xbb, 0xd0, 0x9b, 0x67, 0x15, 0xbb, 0xa0, 0xd9, 0xa3, 0x27,
	0xfa, 0x8d, 0x7a, 0xe1, 0xbb, 0x43, 0xbd, 0x75, 0x34, 0x05, 0x9c, 0x84, 0xd3, 0x87, 0x2a, 0x71,
	0x2e, 0x2d, 0x2e, 0x58, 0x44, 0x31, 0x82, 0xde, 0x3c, 0xab, 0xb8, 0x77, 0xcd, 0xef, 0xf5, 0xe6,
	0x81, 0xd9, 0xb9, 0xfb, 0xb6, 0xb6, 0x4a, 0x8e, 0xf1, 0xe2, 0xd5, 0xbf, 0xdf, 0xb7, 0xd1, 0xdc,
	0x18, 0x5c, 0x7c, 0x3c, 0x90, 0x31, 0xad, 0xdf, 0xd5, 0x07, 0xad, 0x83, 0xa7, 0x1d, 0xf9, 0x08,
	0xde, 0xff, 0x3f, 0x00, 0x00, 0xff, 0xff, 0xc7, 0xb8, 0x43, 0x7d, 0x90, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PayBusinessServiceClient is the client API for PayBusinessService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PayBusinessServiceClient interface {
	// 统一收单支付
	TradePay(ctx context.Context, in *TradePayRequest, opts ...grpc.CallOption) (*TradePayResponse, error)
}

type payBusinessServiceClient struct {
	cc *grpc.ClientConn
}

func NewPayBusinessServiceClient(cc *grpc.ClientConn) PayBusinessServiceClient {
	return &payBusinessServiceClient{cc}
}

func (c *payBusinessServiceClient) TradePay(ctx context.Context, in *TradePayRequest, opts ...grpc.CallOption) (*TradePayResponse, error) {
	out := new(TradePayResponse)
	err := c.cc.Invoke(ctx, "/pay_business.PayBusinessService/TradePay", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PayBusinessServiceServer is the server API for PayBusinessService service.
type PayBusinessServiceServer interface {
	// 统一收单支付
	TradePay(context.Context, *TradePayRequest) (*TradePayResponse, error)
}

// UnimplementedPayBusinessServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPayBusinessServiceServer struct {
}

func (*UnimplementedPayBusinessServiceServer) TradePay(ctx context.Context, req *TradePayRequest) (*TradePayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TradePay not implemented")
}

func RegisterPayBusinessServiceServer(s *grpc.Server, srv PayBusinessServiceServer) {
	s.RegisterService(&_PayBusinessService_serviceDesc, srv)
}

func _PayBusinessService_TradePay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TradePayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayBusinessServiceServer).TradePay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pay_business.PayBusinessService/TradePay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayBusinessServiceServer).TradePay(ctx, req.(*TradePayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PayBusinessService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pay_business.PayBusinessService",
	HandlerType: (*PayBusinessServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TradePay",
			Handler:    _PayBusinessService_TradePay_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/micro_mall_pay_proto/pay_business/pay_business.proto",
}
