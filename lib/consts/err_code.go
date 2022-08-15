package consts

type GMErr struct {
	code string
	msg  string
}

func NewGMErr(code, msg string) *GMErr {
	return &GMErr{code: code, msg: msg}
}

func (gm *GMErr) GetCode() string {
	return gm.code
}

func (gm *GMErr) GetMsg() string {
	return gm.msg
}

func (gm *GMErr) SetMsg(msg string) {
	gm.msg = msg
}

func (gm *GMErr) IsSuccess() bool {
	return gm == nil
}

func IsSuccess(err error) bool {
	return err == nil
}

func IsGMSuccess(err *GMErr) bool {
	return err == nil
}

const (
	AuroraSuccess = "AU0000"
	CodeStatusErr = "AU3011"
	PangleSuccess = "PG0000"
)

var (
	// Success 成功错误码（应该只在处理响应时使用）
	Success       = &GMErr{code: "GM0000", msg: "成功"}
	ErrDB         = &GMErr{code: "GM0001", msg: "数据库错误"}
	ErrDBNotFound = &GMErr{code: "GM0002", msg: "record not found"}
	// ErrNotLogin // 从 GM1000 到 GM1099 为通用错误码
	ErrNotLogin            = &GMErr{code: "GM1000", msg: "用户未登录"}
	ErrInvalidParam        = &GMErr{code: "GM1001", msg: "非法的参数"}
	ErrSysException        = &GMErr{code: "GM1002", msg: "内部系统异常"}
	ErrMediaAccount        = &GMErr{code: "GM1003", msg: "媒体账户异常"}
	ErrMediaSubAccount     = &GMErr{code: "GM1004", msg: "媒体子账户异常"}
	ErrMediaPermission     = &GMErr{code: "GM1005", msg: "媒体权限异常"}
	ErrMediaFunction       = &GMErr{code: "GM1006", msg: "媒体白名单异常"}
	ErrRespParam           = &GMErr{code: "GM1007", msg: "响应参数异常"}
	ErrInvalidStatus       = &GMErr{code: "GM1008", msg: "状态非法"}
	ErrJSONParse           = &GMErr{code: "GM1009", msg: "JSON解析异常"}
	ErrHTTPRequest         = &GMErr{code: "GM1010", msg: "HTTP请求异常"}
	ErrGeoLocationRequest  = &GMErr{code: "GM1011", msg: "地理中台请求异常"}
	ErrRPCCall             = &GMErr{code: "GM1012", msg: "RPC请求异常"}
	ErrInterfaceAccess     = &GMErr{code: "GM1013", msg: "接口权限异常"}
	ErrRequestFrequency    = &GMErr{code: "GM1014", msg: "请求过于频繁"}
	ErrMediaSubAccountSite = &GMErr{code: "GM1015", msg: "子账号没有相关应用权限"}
	ErrParseFile           = &GMErr{code: "GM1016", msg: "解析文件失败"}
	ErrGetAccountInfo      = &GMErr{code: "GM1017", msg: "获取用户信息异常"}
	ErrTccQueryError       = &GMErr{code: "GM1018", msg: "查询TCC配置失败"}
	ErrMediaFind           = &GMErr{code: "GM1019", msg: "用户信息查询失败"}

	// 广告位相关 GM1100 到 GM1199
	ErrAdUnitNameRepeat      = &GMErr{code: "GM1100", msg: "广告位名称重复"}
	ErrAdUnitNotExists       = &GMErr{code: "GM1101", msg: "广告位不存在"}
	ErrAdUnitIsDisable       = &GMErr{code: "GM1102", msg: "广告位禁用时不能修改下面层级的状态"}
	ErrQueryAdUnit           = &GMErr{code: "GM1103", msg: "查询广告位异常"}
	ErrAdUnitIsNotEnable     = &GMErr{code: "GM1104", msg: "广告位需要是运行状态"}
	ErrAdUnitType            = &GMErr{code: "GM1105", msg: "广告位类型不合法"}
	ErrTestAppNotRecommended = &GMErr{code: "GM1106", msg: "测试应用暂不推荐配置"}
	ErrHasNoRecommended      = &GMErr{code: "GM1107", msg: "没有可推荐的配置"}
	ErrAdUnitNameEmpty       = &GMErr{code: "GM1108", msg: "广告位名称不能为空"}
	ErrAdUnitNumLimit        = &GMErr{code: "GM1109", msg: "批量创建的广告位数量不能超过50个"}

	// 瀑布流相关 GM1200 到 GM1299
	ErrWaterfallNotExists = &GMErr{code: "GM1200", msg: "瀑布流不存在"}

	// 代码位相关 GM1300 到 GM1399
	ErrAtLeastOnePangleRit             = &GMErr{code: "GM1300", msg: "至少包含一个穿山甲rit"}
	ErrAdSlotNumLimit                  = &GMErr{code: "GM1301", msg: "瀑布流总层数超出限制"}
	ErrValidAdSlotNumLimit             = &GMErr{code: "GM1302", msg: "瀑布流有效层数超出限制"}
	ErrRevealLimitOne                  = &GMErr{code: "GM1303", msg: "只能有一个兜底代码位"}
	ErrInvalidAdSlotStatus             = &GMErr{code: "GM1304", msg: "代码位状态非法"}
	ErrInvalidNetwork                  = &GMErr{code: "GM1305", msg: "广告网络非法"}
	ErrSortType                        = &GMErr{code: "GM1306", msg: "排序方式不合法"}
	ErrInvalidBiddingType              = &GMErr{code: "GM1307", msg: "竞价类型非法"}
	ErrNetworkNotSupportBiddingType    = &GMErr{code: "GM1308", msg: "当前network不支持当前的竞价类型"}
	ErrBiddingTypeServer               = &GMErr{code: "GM1309", msg: "没有穿山甲服务端竞价白名单"}
	ErrBiddingTypeClient               = &GMErr{code: "GM1310", msg: "没有客户端竞价白名单"}
	ErrAppConfigInvalid                = &GMErr{code: "GM1311", msg: "Network配置缺失"}
	ErrPriceNotValid                   = &GMErr{code: "GM1312", msg: "价格非法"}
	ErrRitNameTooLong                  = &GMErr{code: "GM1313", msg: "代码位名称过长"}
	ErrDuplicateAdSlotName             = &GMErr{code: "GM1314", msg: "代码位名称重复"}
	ErrDuplicateAdSlotID               = &GMErr{code: "GM1315", msg: "代码位ID重复"}
	ErrPriorityNotPermission           = &GMErr{code: "GM1316", msg: "没有手动排序权限"}
	ErrNetworkNotSupport               = &GMErr{code: "GM1317", msg: "该 network 不支持当前的广告位类型"}
	ErrBiddingAndStandardNotTogether   = &GMErr{code: "GM1318", msg: "穿山甲下竞价代码位和普通代码位不能共存"}
	ErrBiddingOnlyOne                  = &GMErr{code: "GM1319", msg: "一个network只能有一个bidding代码位"}
	ErrPangleAdSlotRelatedError        = &GMErr{code: "GM1320", msg: "当前穿山甲代码位不属于该应用"}
	ErrPangleAdSlotAdUnitTypeError     = &GMErr{code: "GM1321", msg: "当前代码位的广告位类型与在穿山甲的不一致"}
	ErrCodeNotMatchBiddingType         = &GMErr{code: "GM1322", msg: "竞价类型和代码位不匹配"}
	ErrNoUseMediation                  = &GMErr{code: "GM1323", msg: "穿山甲代码位创建时未选择使用穿山甲聚合"}
	ErrAdSlotNotExist                  = &GMErr{code: "GM1324", msg: "代码位不存在"}
	ErrOriginTypeNotSupport            = &GMErr{code: "GM1325", msg: "此广告网络不支持此渲染类型"}
	ErrOriginTypeConflict              = &GMErr{code: "GM1326", msg: "与穿山甲代码位渲染类型冲突"}
	ErrMultilevelPriceLevel            = &GMErr{code: "GM1327", msg: "多阶低价层级数量异常"}
	ErrMultilevelPriceTag              = &GMErr{code: "GM1328", msg: "多阶底价tag类型错误"}
	ErrMultilevelPriceMinNumber        = &GMErr{code: "GM1329", msg: "多阶底价价格异常不满足最小差距"}
	ErrMultilevelPriceInterval         = &GMErr{code: "GM1330", msg: "多阶底价竞价价格不满足区间"}
	ErrMultilevelPriceNotPangleBidding = &GMErr{code: "GM1331", msg: "多阶底价穿山甲bidding依赖不存在"}
	ErrMultilevelPriceLevelNumber      = &GMErr{code: "GM1332", msg: "多阶底价底价格异常"}
	ErrMultilevelPriceLevelDecrease    = &GMErr{code: "GM1333", msg: "多阶底价底价格不满足递减"}
	ErrMultilevelPriceDecrease         = &GMErr{code: "GM1334", msg: "多阶底价竞价价格不满足递减"}
	ErrMultilevelPriceType             = &GMErr{code: "GM1335", msg: "多阶底价 价格格式异常"}
	ErrMultilevelTagNil                = &GMErr{code: "GM1336", msg: "多阶底价 tag 标签不能为空"}
	ErrMultilevelTagRepeat             = &GMErr{code: "GM1337", msg: "多阶底价 tag 标签不能重复"}
	ErrDragSort                        = &GMErr{code: "GM1338", msg: "拖拽排序错误"}
	ErrPangleCodeStatus                = &GMErr{code: "GM1339", msg: "穿山甲代码位状态异常，不可修改信息"}
	ErrBiddingNeedPangleBidding        = &GMErr{code: "GM1340", msg: "其他adn bidding 强依赖 穿山甲bidding"}
	ErrOriginTypeNotSet                = &GMErr{code: "GM1341", msg: "此广告网络未设置渲染类型"}
	ErrNetworkAndRitNotSet             = &GMErr{code: "GM1342", msg: "广告网络和代码位ID是必填项"}
	ErrInvalidPangleStatus             = &GMErr{code: "GM1343", msg: "代码位状态异常，禁止操作"}
	ErrUseTemplateNotSupportCopied     = &GMErr{code: "GM1344", msg: "使用自定义模板时不支持复制"}
	ErrCopyCodeFailed                  = &GMErr{code: "GM1345", msg: "复制穿山甲代码位失败"}
	ErrBiddingSupport                  = &GMErr{code: "GM1346", msg: "当前adn没有bidding权限"}
	ErrCodeDisplayRuleWhiteList        = &GMErr{code: "GM1347", msg: "当前用户没有代码位展示控制规则白名单"}
	ErrCodeDisplayRuleInvalidParam     = &GMErr{code: "GM1348", msg: "非法的代码位展示控制规则值"}
	ErrCodeRitIDWithChinese            = &GMErr{code: "GM1349", msg: "当前ADN的代码位ID包含中文字符，请修改后再操作"}
	ErrBatchHasMultiStatus             = &GMErr{code: "GM1350", msg: "当前批量操作包含多种代码位状态，请保证源状态一致"}
	ErrNotSupportBatchUpdate           = &GMErr{code: "GM1351", msg: "当前代码位不支持批量更新"}
	ErrAtLeastOneEnableRit             = &GMErr{code: "GM1352", msg: "启用的瀑布流下至少有一个启用的代码位"}
	ErrAdSlotIDs                       = &GMErr{code: "GM1353", msg: "要批量编辑的代码位ID不存在或已删除"}
	ErrAdSlotIDNotBelongWaterfall      = &GMErr{code: "GM1354", msg: "要编辑的代码位ID不属于当前瀑布流"}
	ErrFetchDaysOutOfRange             = &GMErr{code: "GM1355", msg: "自动排序取数周期不在指定范围内"}
	ErrSmartSortNotSupport             = &GMErr{code: "GM1356", msg: "自动排序仅支持价格层和兜底层"}
	ErrSmartSortSwitch                 = &GMErr{code: "GM1357", msg: "自动排序开关不合法"}
	ErrCodeDisplayRuleNotSupportPrice  = &GMErr{code: "GM1358", msg: "代码位展示控制规则不支持信息流、draw信息流以及Banner混信息流场景下在价格层使用"}

	// Ab实验相关 GM1400 到 GM1499
	ErrExperimentNotExists                   = &GMErr{code: "GM1400", msg: "AB实验不存在"}
	ErrQueryExperimentCondition              = &GMErr{code: "GM1401", msg: "查找实验条件失败"}
	ErrNotFoundExperimentCondition           = &GMErr{code: "GM1402", msg: "未查到合法的实验条件"}
	ErrQueryExperimentDetail                 = &GMErr{code: "GM1403", msg: "查找实验详情失败"}
	ErrNotFoundExperimentDetail              = &GMErr{code: "GM1404", msg: "未查到合法的实验详情"}
	ErrQueryExperimentRecord                 = &GMErr{code: "GM1405", msg: "查找实验记录失败"}
	ErrNotFoundExperimentRecord              = &GMErr{code: "GM1406", msg: "未查到合法的实验记录"}
	ErrExperimentIsPending                   = &GMErr{code: "GM1407", msg: "B组实验未开启时时不能修改下面层级的状态"}
	ErrHasEnableAB                           = &GMErr{code: "GM1408", msg: "已存在运行中AB"}
	ErrExperimentStatusMustPending           = &GMErr{code: "GM1409", msg: "实验状态必须是待开启"}
	ErrExperimentBNeedAtLeastOneEnableAdSlot = &GMErr{code: "GM1410", msg: "B组至少需要一个可用代码位"}
	ErrExperimentStatusMustRunning           = &GMErr{code: "GM1411", msg: "实验状态必须是运行中"}
	ErrNoAvailableCodes                      = &GMErr{code: "GM1412", msg: "创建实验前至少要有一个启用的代码位"}

	// 流量分组相关 GM1500 到 GM1599
	ErrQuerySegmentTemplate       = &GMErr{code: "GM1501", msg: "查询分组模板异常"}
	ErrSegmentNotExists           = &GMErr{code: "GM1502", msg: "分组不存在"}
	ErrQuerySegment               = &GMErr{code: "GM1503", msg: "查询分组异常"}
	ErrSegmentIsDisable           = &GMErr{code: "GM1504", msg: "流量分组禁用时不能修改下面层级的状态"}
	ErrAtLeastOneInDefaultSegment = &GMErr{code: "GM1505", msg: "默认分组最后一个代码位不能关闭/删除"}
	ErrDeleteSegmentOnExperiment  = &GMErr{code: "GM1506", msg: "分组有组内实验时不能删除"}
	ErrDeleteDefaultSegment       = &GMErr{code: "GM1507", msg: "默认分组不能删除"}

	// ErrDiagnosisConfCreate 诊断工具 GM1600 到 GM1699
	ErrDiagnosisConfCreate     = &GMErr{code: "GM1601", msg: "诊断工具配置创建错误"}
	ErrDiagnosisResolveErrCode = &GMErr{code: "GM1602", msg: "诊断工具错误码数据解析失败"}

	// ErrNotMediationSite 应用相关 GM1700 到 GM1799
	ErrNotMediationSite  = &GMErr{code: "GM1700", msg: "当前app未添加到聚合"}
	ErrAtLeastOneNetwork = &GMErr{code: "GM1701", msg: "除穿山甲外至少配置一个广告网络，才能生成推荐配置"}

	// adn 操作相关错误码

	ErrUpdateCustomAdn          = &GMErr{code: "GM1703", msg: "更新adn信息失败"}
	ErrFindSiteAdn              = &GMErr{code: "GM1704", msg: "查询应用adn信息失败"}
	ErrCreateSiteAdn            = &GMErr{code: "GM1705", msg: "创建应用adn信息失败"}
	ErrUpdateSiteAdn            = &GMErr{code: "GM1706", msg: "更新应用adn信息失败"}
	ErrCustomAdnNotConfig       = &GMErr{code: "GM1707", msg: "自定义adn 没有配置信息"}
	ErrCustomAdnAdapter         = &GMErr{code: "GM1708", msg: "自定义adn adapter 信息配置错误"}
	ErrNoCustomID               = &GMErr{code: "GM1709", msg: "自定义adn 参数非法"}
	ErrFindCurrency             = &GMErr{code: "GM1710", msg: "adn 结算币种没有找到"}
	ErrCustomExtra              = &GMErr{code: "GM1711", msg: "自定义adn extra 字段异常"}
	ErrCustomAdapter            = &GMErr{code: "GM1712", msg: "自定义adn adapter 类型重复"}
	ErrCustomAdapterNoInitClass = &GMErr{code: "GM1713", msg: "自定义adn adapter 没有配置初始化类"}
	ErrCustomAdnNameRepeat      = &GMErr{code: "GM1714", msg: "自定义adn adn 名称重复"}
	ErrCustomAdnNameType        = &GMErr{code: "GM1715", msg: "自定义adn adn 名称格式错误"}
	ErrCustomAdnNameLen         = &GMErr{code: "GM1716", msg: "自定义adn adn 名称长度大于20 错误"}
	ErrCustomAdnSupport         = &GMErr{code: "GM1717", msg: "暂不支持自定adn配置"}
	ErrQueryCustomAdn           = &GMErr{code: "GM1718", msg: "查询adn信息失败"}
	ErrCreateCustomAdn          = &GMErr{code: "GM1719", msg: "创建adn信息失败"}
	ErrNoCustomAdnAdapter       = &GMErr{code: "GM1720", msg: "自定义adn adapter 没有配置"}

	ErrScreenAdtype = &GMErr{code: "GM1801", msg: "插全屏合并adtype非法"}

	ErrDateInvalid             = &GMErr{code: "GM1901", msg: "日期与指定的格式不符"}
	ErrReportDataAllNotMatched = &GMErr{code: "GM1902", msg: "您上传的数据均未匹配到对应的代码位，请检查数据是否上传正确"}
	ErrReportDataInvalid       = &GMErr{code: "GM1903", msg: "您上传的数据存在非法字符，非法数据将被过滤，其余数据将会正常上传"}
	ErrHasNotSupportDate       = &GMErr{code: "GM1904", msg: "仅支持上传(T-1)~(T-7)日内的数据，您上传的数据中存在超出时间范围内的数据，该部分数据会被过滤，其余数据将会正常上传"}
	ErrUploadDuplication       = &GMErr{code: "GM1905", msg: "您已配置通过自动拉取该广告网络API数据，无需手动回传数据"}
	ErrStartDateAfterEnd       = &GMErr{code: "GM1906", msg: "开始日期不能晚于结束日期"}
	ErrDateOutOfRange          = &GMErr{code: "GM1907", msg: "当前查询的日期不在支持范围内"}

	ErrDataCenterQueryIndicators = &GMErr{code: "GM2000", msg: "数据指标读取异常，请您刷新后重试或联系客服"}

	ErrAuroraSuccess = &GMErr{code: "AU0000", msg: "Success"}
)
