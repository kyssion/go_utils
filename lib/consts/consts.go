package consts

const (
	OpUnitUpdate = 1 // 广告位更新
	OpCodeAdd    = 2 // 代码位增加
	OpCodeUpdate = 3 // 代码位更新
	OpCodeDelete = 4 // 代码位删除
)

const (
	PSM = "ad.pangle.gromore"
)

const (
	EmptyArrayListLength = 0
	EmptyStringValue     = ""
	ZeroIntValue         = 0
	ZeroFloat64Value     = float64(0)
	ZeroStringValue      = ""
	MoneyUnit            = 100000 // 千分之一分转成元
)

type _CtxUserInfo string

const (
	CtxUserInfo = _CtxUserInfo("UserInfo")
)

const (
	Yes = 1
	No  = 2
)

const (
	// MediationBasicFunction 聚合基础白名单
	MediationBasicFunction = "mediation_sdk"
	// FunctionKs 快手白名单
	FunctionKs = "mediation_kuaishou"
	// FunctionMBidding 已全量
	FunctionMBidding = "m_bidding"
	// FunctionManualPriority 已全量
	FunctionManualPriority = "manual_priority"
	// FunctionGdtBidding gdt bidding白名单
	FunctionGdtBidding = "gdt_server_side_bidding"
	// FunctionMultilevelPrice 多阶低价白名单
	FunctionMultilevelPrice = "mediation_multilevel_price"
	// FunctionBaiDuBidding 百度cb 白名单
	FunctionBaiDuBidding = "baidu_client_bidding"
	// FunctionCustomAdn 自定义adn 白名单
	FunctionCustomAdn = "mediation_custom_adn"
	// FunctionKlevin 游可赢 白名单
	FunctionKlevin = "mediation_klevin"
	// FunctionPgMixBidding pangle bidding 和标准代码位可共存 白名单
	FunctionPgMixBidding = "mediation_mix_bidding_target"
	// FunctionMediationServerReward 服务端激励回调白名单
	FunctionMediationServerReward = "mediation_server_reward"
	// FunctionCustomServerBidding 自定义adn server bidding
	FunctionCustomServerBidding = "mediation_custom_server_bidding"
)

// Metrics指标相关
const (
	MetricsPanic               = "panic"
	MetricsQueryFromDataCenter = "query_from_data_center"
)

const (
	RunningStatus         = 1
	ShutdownStatus        = 2
	SegmentDefaultTag     = 1 // 默认分组标签
	SegmentNonDefaultTag  = 2 // 非默认分组标签
	SegmentTemplateTag    = 1 // 模板标签
	SegmentNonTemplateTag = 2 // 非模板标签
)

const (
	ActionCreate            = 1
	ActionUpdate            = 2
	ActionDelete            = 3
	ActionDeleteAndCreate   = 4
	ActionChangeBiddingType = 5
)

const (
	USD = "usd"
	CNY = "cny"

	CountryChina   = "cn"
	TimezoneNorth8 = 8
)

// 时间格式
const (
	DateFormat              = "2006-01-02"
	HourDateFormat          = "2006-01-02 15"
	TimeFormatISO           = "2006-01-02 15:04:05"
	HivePartitionDateFormat = "20060102"
	Month2MinuteFormat      = "01_02_15:04"
)

var ValidSegmentFilter = map[int64]string{
	1: "=",
	2: "!=",
	3: ">",
	4: "<",
	5: "in",
	6: "not_in",
	7: "between",
	8: "or",
	9: "and",
}

var CountryList = map[int64]string{
	1: "CN",
}

const (
	BatchCreateSize = 100 // db批量写规模
)

// db order
const (
	DBOrderTypeDescend = "desc"
)

const (
	CommonErrPrefix           = "common.error_"
	RecommendSettingsSnapshot = "recommendSettingsSnapshot"
	RecommendSettingsType     = "recommended_settings_type"
	UnknownEnum               = "<UNSET>"
	SpanTypeExecuteFunction   = "ExecuteFunction"
)

const (
	CreateWaterfallWithCopyCode    = true
	CreateWaterfallWithoutCopyCode = false
)
