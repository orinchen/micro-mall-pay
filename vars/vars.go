package vars

import "gitee.com/kelvins-io/common/queue"

var (
	EmailConfigSetting  *EmailConfigSettingS
	AppName             = ""
	TradePayQueueServer *queue.MachineryQueue
)
