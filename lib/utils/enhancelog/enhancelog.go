package enhancelog

import (
	"fmt"
	"strings"

	"context"

	"code.byted.org/gopkg/logs"
)

type EncryptModel int32
type FuncSwitch int32

// 敏感词列表
var sensitiveList []string

// 终止符
var separatorSet map[rune]bool

// 加密模式，默认加密一部分
var model EncryptModel

// 脱敏功能开关
var encryptSwitch FuncSwitch

// stage功能开关
var stageSwitch FuncSwitch

// 加密内容
var encryption = "******"

const (
	// EncryptAll 敏感字段内容全部加密
	EncryptAll EncryptModel = 0
	// EncryptPart 敏感字段内容中间部分加密
	EncryptPart EncryptModel = 1

	// Off 功能关闭
	Off FuncSwitch = 0
	// On 功能开启
	On FuncSwitch = 1
)

func init() {
	separatorSet = map[rune]bool{
		' ': true,
		',': true,
		'}': true,
		')': true,
	}
	model = EncryptPart
	encryptSwitch = On
	stageSwitch = Off
}

// InitEnhanceLog 配置敏感词和加密模式
func InitEnhanceLog(slist []string, m EncryptModel) {
	sensitiveList = slist
	var newSlist []string
	for _, v := range slist {
		newSlist = append(newSlist, v+"\"")
	}
	sensitiveList = append(sensitiveList, newSlist...)

	model = m
	// 开启这行则原来使用logs打印日志的，堆栈会到更上一层
	logs.SetCallDepth(4)
	logs.Debug("do SetCallDepth")
}

// ConfigFuncSwitch 配置功能开关
func ConfigFuncSwitch(encryptFunc FuncSwitch, stageFunc FuncSwitch) {
	encryptSwitch = encryptFunc
	stageSwitch = stageFunc
}

func ConfigEncryptModel(m EncryptModel) {
	model = m
}

func encryptLog(logContent string) string {
	// 0.开关判断
	if encryptSwitch != On {
		return logContent
	}

	// 1.遍历敏感词列表
	for _, sensitive := range sensitiveList {
		// 2. 寻找敏感内容开头
		fieldIndex := strings.Index(logContent, sensitive+":")
		if fieldIndex != -1 {
			// 3. 查找敏感内容结尾
			valueBeginIndex := fieldIndex + len(sensitive) + 1
			valueEndIndex := len(logContent)

			for tempIndex, value := range logContent[valueBeginIndex:] {
				if separatorSet[value] && !(tempIndex == 0 && value == ' ') { // 后经常出现空格的情况
					valueEndIndex = tempIndex + valueBeginIndex
					break
				}
			}
			// logs.Debug("enhancelog do replace of sensitive word:%s on filed:%s", logContent[valueBeginIndex:valueEndIndex], sensitive)
			// 4. 替换敏感内容
			switch model {
			case EncryptAll:
				logContent = logContent[:valueBeginIndex] + encryption + logContent[valueEndIndex:]
			default:
				// 为支持中文，需要转换成rune slice进行处理
				sensitiveContentRunes := []rune(logContent[valueBeginIndex:valueEndIndex])
				if len(sensitiveContentRunes) >= 10 {
					// 长度大于等于10则替换中间6个
					encryptBeginIndex := len(sensitiveContentRunes)/2 - 3
					for i := 0; i < 6; i++ {
						sensitiveContentRunes[i+encryptBeginIndex] = '*'
					}
					logContent = logContent[:valueBeginIndex] + string(sensitiveContentRunes) + logContent[valueEndIndex:]
				} else {
					// 长度小于10则替换中间一半的内容
					contentLength := len(sensitiveContentRunes)
					encryptionLength := contentLength / 2
					// 长度为0不需要处理
					if contentLength != 0 {
						if encryptionLength == 0 {
							// 对待加密长度为1的进行特殊处理
							encryptionLength = 1
						}
						// 起始位置为中间位置减去加密长度的一半，
						encryptBeginIndex := contentLength/2 - encryptionLength/2

						for i := 0; i < encryptionLength; i++ {
							sensitiveContentRunes[encryptBeginIndex+i] = '*'
						}
						logContent = logContent[:valueBeginIndex] + string(sensitiveContentRunes) + logContent[valueEndIndex:]
					}
				}
			}

		}
	}
	return logContent
}

func Fatalf(format string, v ...interface{}) {
	// 1.format输出内容，获得加密前日志
	logContent := format
	if len(v) != 0 {
		logContent = fmt.Sprintf(format, v...)
	}
	// 2.遍历敏感词列表，寻找并替换敏感信息
	logContent = encryptLog(logContent)
	// 3.输出进行加密处理后的日志
	logs.Fatalf("%s", logContent)
}

func Errorf(format string, v ...interface{}) {
	// 1.format输出内容，获得加密前日志
	logContent := format
	if len(v) != 0 {
		logContent = fmt.Sprintf(format, v...)
	}
	// 2.遍历敏感词列表，寻找并替换敏感信息
	logContent = encryptLog(logContent)
	// 3.输出进行加密处理后的日志
	logs.Errorf("%s", logContent)
}

func Warnf(format string, v ...interface{}) {
	// 1.format输出内容，获得加密前日志
	logContent := format
	if len(v) != 0 {
		logContent = fmt.Sprintf(format, v...)
	}
	// 2.遍历敏感词列表，寻找并替换敏感信息
	logContent = encryptLog(logContent)
	// 3.输出进行加密处理后的日志
	logs.Warnf("%s", logContent)
	println(logContent)
}

func Noticef(format string, v ...interface{}) {
	// 1.format输出内容，获得加密前日志
	logContent := format
	if len(v) != 0 {
		logContent = fmt.Sprintf(format, v...)
	}
	// 2.遍历敏感词列表，寻找并替换敏感信息
	logContent = encryptLog(logContent)
	// 3.输出进行加密处理后的日志
	logs.Noticef("%s", logContent)
}

func Infof(format string, v ...interface{}) {
	// 1.format输出内容，获得加密前日志
	logContent := format
	if len(v) != 0 {
		logContent = fmt.Sprintf(format, v...)
	}
	// 2.遍历敏感词列表，寻找并替换敏感信息
	logContent = encryptLog(logContent)
	// 3.输出进行加密处理后的日志
	logs.Infof("%s", logContent)
}

func Debugf(format string, v ...interface{}) {
	// 1.format输出内容，获得加密前日志
	logContent := format
	if len(v) != 0 {
		logContent = fmt.Sprintf(format, v...)
	}
	// 2.遍历敏感词列表，寻找并替换敏感信息
	logContent = encryptLog(logContent)
	// 3.输出进行加密处理后的日志
	logs.Debugf("%s", logContent)
}

func Tracef(format string, v ...interface{}) {
	// 1.format输出内容，获得加密前日志
	logContent := format
	if len(v) != 0 {
		logContent = fmt.Sprintf(format, v...)
	}
	// 2.遍历敏感词列表，寻找并替换敏感信息
	logContent = encryptLog(logContent)
	// 3.输出进行加密处理后的日志
	logs.Tracef("%s", logContent)
}

func Fatal(format string, v ...interface{}) {
	// 1.format输出内容，获得加密前日志
	logContent := format
	if len(v) != 0 {
		logContent = fmt.Sprintf(format, v...)
	}
	// 2.遍历敏感词列表，寻找并替换敏感信息
	logContent = encryptLog(logContent)
	// 3.输出进行加密处理后的日志
	logs.Fatal("%s", logContent)
}

func Error(format string, v ...interface{}) {
	// 1.format输出内容，获得加密前日志
	logContent := format
	if len(v) != 0 {
		logContent = fmt.Sprintf(format, v...)
	}
	// 2.遍历敏感词列表，寻找并替换敏感信息
	logContent = encryptLog(logContent)
	// 3.输出进行加密处理后的日志
	logs.Error("%s", logContent)
}

func Warn(format string, v ...interface{}) {
	// 1.format输出内容，获得加密前日志
	logContent := format
	if len(v) != 0 {
		logContent = fmt.Sprintf(format, v...)
	}
	// 2.遍历敏感词列表，寻找并替换敏感信息
	logContent = encryptLog(logContent)
	// 3.输出进行加密处理后的日志
	logs.Warn("%s", logContent)
}

func Notice(format string, v ...interface{}) {
	// 1.format输出内容，获得加密前日志
	logContent := format
	if len(v) != 0 {
		logContent = fmt.Sprintf(format, v...)
	}
	// 2.遍历敏感词列表，寻找并替换敏感信息
	logContent = encryptLog(logContent)
	// 3.输出进行加密处理后的日志
	logs.Notice("%s", logContent)
}

func Info(format string, v ...interface{}) {
	// 1.format输出内容，获得加密前日志
	logContent := format
	if len(v) != 0 {
		logContent = fmt.Sprintf(format, v...)
	}
	// 2.遍历敏感词列表，寻找并替换敏感信息
	logContent = encryptLog(logContent)
	// 3.输出进行加密处理后的日志
	logs.Info("%s", logContent)
}

func Debug(format string, v ...interface{}) {
	// 1.format输出内容，获得加密前日志
	logContent := format
	if len(v) != 0 {
		logContent = fmt.Sprintf(format, v...)
	}
	// 2.遍历敏感词列表，寻找并替换敏感信息
	logContent = encryptLog(logContent)
	// 3.输出进行加密处理后的日志
	logs.Debug("%s", logContent)
}

func Trace(format string, v ...interface{}) {
	// 1.format输出内容，获得加密前日志
	logContent := format
	if len(v) != 0 {
		logContent = fmt.Sprintf(format, v...)
	}
	// 2.遍历敏感词列表，寻找并替换敏感信息
	logContent = encryptLog(logContent)
	// 3.输出进行加密处理后的日志
	logs.Trace("%s", logContent)
}

func CtxFatal(ctx context.Context, format string, v ...interface{}) {
	// 1.format输出内容，获得加密前日志
	logContent := format
	if len(v) != 0 {
		logContent = fmt.Sprintf(format, v...)
	}
	// 2.遍历敏感词列表，寻找并替换敏感信息
	logContent = encryptLog(logContent)
	// 3.输出进行加密处理后的日志
	logs.CtxFatal(wrapContext(ctx), "%s", logContent)
}

func CtxError(ctx context.Context, format string, v ...interface{}) {
	// 1.format输出内容，获得加密前日志
	logContent := format
	if len(v) != 0 {
		logContent = fmt.Sprintf(format, v...)
	}
	// 2.遍历敏感词列表，寻找并替换敏感信息
	logContent = encryptLog(logContent)
	// 3.输出进行加密处理后的日志
	logs.CtxError(wrapContext(ctx), "%s", logContent)
}

func CtxWarn(ctx context.Context, format string, v ...interface{}) {
	// 1.format输出内容，获得加密前日志
	logContent := format
	if len(v) != 0 {
		logContent = fmt.Sprintf(format, v...)
	}
	// 2.遍历敏感词列表，寻找并替换敏感信息
	logContent = encryptLog(logContent)
	// 3.输出进行加密处理后的日志
	logs.CtxWarn(wrapContext(ctx), "%s", logContent)
}

func CtxNotice(ctx context.Context, format string, v ...interface{}) {
	// 1.format输出内容，获得加密前日志
	logContent := format
	if len(v) != 0 {
		logContent = fmt.Sprintf(format, v...)
	}
	// 2.遍历敏感词列表，寻找并替换敏感信息
	logContent = encryptLog(logContent)
	// 3.输出进行加密处理后的日志
	logs.CtxNotice(wrapContext(ctx), "%s", logContent)
}

func CtxInfo(ctx context.Context, format string, v ...interface{}) {
	// 1.format输出内容，获得加密前日志
	logContent := format
	if len(v) != 0 {
		logContent = fmt.Sprintf(format, v...)
	}
	// 2.遍历敏感词列表，寻找并替换敏感信息
	logContent = encryptLog(logContent)
	// 3.输出进行加密处理后的日志
	if len(logContent) > 10240 {
		logContent = logContent[0:10240]
	}
	logs.CtxInfo(wrapContext(ctx), "%s", logContent)
}

func CtxDebug(ctx context.Context, format string, v ...interface{}) {
	// 1.format输出内容，获得加密前日志
	logContent := format
	if len(v) != 0 {
		logContent = fmt.Sprintf(format, v...)
	}
	// 2.遍历敏感词列表，寻找并替换敏感信息
	logContent = encryptLog(logContent)
	// 3.输出进行加密处理后的日志
	logs.CtxDebug(wrapContext(ctx), "%s", logContent)
}

func CtxTrace(ctx context.Context, format string, v ...interface{}) {
	// 1.format输出内容，获得加密前日志
	logContent := format
	if len(v) != 0 {
		logContent = fmt.Sprintf(format, v...)
	}
	// 2.遍历敏感词列表，寻找并替换敏感信息
	logContent = encryptLog(logContent)
	// 3.输出进行加密处理后的日志
	logs.CtxTrace(wrapContext(ctx), "%s", logContent)
}

func Stop() {
	logs.Stop()
}

func Flush() {
	logs.Flush()
}
