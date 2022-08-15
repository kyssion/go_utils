# enhanceLog
在公司code.byted.org/gopkg/logs模块的基础上，做了一些增强的功能，目前支持对日志中的敏感信息进行抹除和RPC调用阶段记录功能。

# 功能
- 支持对日志中的敏感字段内容进行抹除 （默认开启）
- 支持日志中输出rpc调用阶段（默认关闭）

# 设计
详见https://bytedance.feishu.cn/space/doc/doccn9Udjkzmvyfolxm2xKqrDSh

# 用法
# 日志脱敏
对包含敏感信息输出的日志，在日志输出前找出敏感内容，并使用`*`将敏感内容进行替换
## 1.初始化
在项目初始化时，需要调用InitEnhanceLog函数设置敏感词列表以及抹除模式。抹除模式包含Encrypt_ALL（敏感内容全部抹除）、Encrypt_PART（敏感内容部分抹除）两种。
```
    enhancelog.InitEnhanceLog([]string{"CardNo","Pwd"}, Encrypt_PART)
```
## 2.打印日志
使用enhancelog包函数替换需要进行日志加密的logs包的函数
```
    // logs.CtxInfo(ctx, "request CardNo:%v,Pwd:%v", cardNo,pwd)
    enhancelog.CtxInfo(ctx, "request CardNo:%v,Pwd:%v", cardNo,pwd)
```
在enhancelog内部，将通过 "敏感字段名:" 来作为敏感字段内容的起始位置，以',' ' ' '}' ')'作为该字段内容的结尾。因此支持以json格式、%+v格式或任何符合该规则的格式输出的内容进行日志加密。

## 效果演示
下面的日志，配置了Mobile敏感词后，自动将 `Mobile:18854466644`替换成了`Mobile:18******644`
```
    // Debug 2019-09-27 15:24:42,125 v1(6) enhancelog.go:371 10.225.126.73 yyyyy.tp.cashdesk - default canary enhancelog do replace of sensitive word:18854466644 on filed:Mobile
    Debug 2019-09-27 15:24:42,125 v1(6) trade_create_flow.go:329 10.225.126.73 yyyyy.tp.cashdesk 201909271524410102250680800100BC48 default canary Stage=7 RemoteAddr=TradeCreate TradeCreateFlow GetUserInfo uid[201907260027] userInfo[UserInfo({Uid:201907260027 Mid:116738603955272846644 MName:*晨鸣 AuthStatus:1 CertificateNum:1****************7 CertificateType:1 UidType:13 AuthUrl:http://pay-boe.snssdk.com/usercenter/member/guide PwdStatus:1 FindPwdUrl:http://pay-boe.snssdk.com/usercenter/setpass PwdLevel:2 Mobile:18******644 PayIdState:0xc420b3985c BindUrl:http://pay-boe.snssdk.com/usercenter/bindphone ChangeMobileTips: MobileShowStrategy:1 UlpayAuthFlag:1 DecliveUrl: PwdCheckWay:0 IsLogin:true EnableBindCard:1 EnableBindCardMsg: DisableCreditCard:false IsComplete:<nil> DisplayInfos:map[] ShowPortal:true Aid:0})]
```

# RPC调用阶段记录
该功能引自：https://code.byted.org/weigaofeng/extlog/tree/master
为了能使用统一的包所以整合在了一个工程。该功能允许在日志中添加Stage（RPC调用到哪个阶段）和RemoteAddr信息，在日志中输出RPC调用阶段，方便进行问题排查
## 1.启动时初始化
该功能默认关闭，要开启需要在项目初始化时配置启用该功能
```
    //第一个参数配置日志脱敏开关，第二个参数配置调用阶段记录开关
    enhancelog.ConfigFuncSwitch(enhancelog.On,enhancelog.On) 
```
## 2.接口入口初始化context
该功能需要在 rpc服务/web服务 接口入口处初始化context
### 在KITE框架中使用
在各个idl handle入口处添加，用新的ctx往下层调用传递
```
    orgC := context.Background()
    c := InitInRPC(orgC, "interface name")
```
### 在GIN框架中使用
初始化，无返回值
```
    InitInGIN(c *gin.Context, addr string) 
```
## 3.在关键环节更新stage
```
    func subCall(client *kitc.KitcClient, ctx context.Context, id int32, message string) {
        enhancelog.IncrLogStage(ctx)
    }
```
## 4.使用enhancelog输出日志
```
    enhancelog.CtxInfo(c,"finish call1")
```
## 效果演示
下面的日志，配置了RPC调用阶段记录后，自动在日志中输出了`Stage=7 RemoteAddr=TradeCreate`,可以知道这是在TradeCreate接口在调用了7次RPC后输出的日志
```
    Debug 2019-09-27 15:24:42,125 v1(6) trade_create_flow.go:329 10.225.126.73 yyyyy.tp.cashdesk 201909271524410102250680800100BC48 default canary Stage=7 RemoteAddr=TradeCreate TradeCreateFlow GetUserInfo uid[201907260027] userInfo[UserInfo({Uid:201907260027 Mid:116738603955272846644 MName:*晨鸣 AuthStatus:1 CertificateNum:1****************7 CertificateType:1 UidType:13 AuthUrl:http://pay-boe.snssdk.com/usercenter/member/guide PwdStatus:1 FindPwdUrl:http://pay-boe.snssdk.com/usercenter/setpass PwdLevel:2 Mobile:18******644 PayIdState:0xc420b3985c BindUrl:http://pay-boe.snssdk.com/usercenter/bindphone ChangeMobileTips: MobileShowStrategy:1 UlpayAuthFlag:1 DecliveUrl: PwdCheckWay:0 IsLogin:true EnableBindCard:1 EnableBindCardMsg: DisableCreditCard:false IsComplete:<nil> DisplayInfos:map[] ShowPortal:true Aid:0})]
```
