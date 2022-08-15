package enhancelog

import (
	"context"
	"fmt"
	"sort"

	"code.byted.org/gopkg/logs"
)

const (
	logmark = "_WRAP_LOG_MARK_"

	Stage      = "Stage"
	RemoteAddr = "RemoteAddr"

	isOutputSorted = false // 用来控制输出字段顺序是否按key排序
)

type extLogMark map[string]interface{}

// InitInRPC 适用于kite等rpc框架，传递参数为context.Context的场景
// 有返回值，要用此替换掉上下文传递的原context
func InitInRPC(c context.Context, addr string) context.Context {
	if p := getMarkPtr(c); p != nil {
		return c
	}

	ext := extLogMark{Stage: 0, RemoteAddr: addr}

	return context.WithValue(c, logmark, &ext)
}

// IncreStage incr stage
func IncreStage(c context.Context) {
	if p := getMarkPtr(c); p != nil {
		if v, ok := (*p)[Stage]; ok {
			if ind, ok := v.(int); ok {
				(*p)[Stage] = ind + 1
			}
		}
	} else { // should assert
		fmt.Errorf("this context should be inited first, ctx:%v", c)
	}
}

// 获取注入的额外信息指针
func getMarkPtr(c context.Context) *extLogMark {
	if v := c.Value(logmark); v != nil {
		if p, ok := v.(*extLogMark); ok {
			return p
		}
	}
	return nil
}

// 可支持更多字段，如果字段少考虑效率起见，可以直接硬编码
func output(ext extLogMark) []interface{} {
	kvs := make([]interface{}, 0)
	if isOutputSorted {
		sKeys := make([]string, 0)
		for k := range ext {
			sKeys = append(sKeys, k)
		}
		sort.Strings(sKeys)

		for _, k := range sKeys {
			kvs = append(kvs, k, fmt.Sprintf("%v", ext[k]))
		}
	} else {
		for k, v := range ext {
			kvs = append(kvs, k, fmt.Sprintf("%v", v))
		}
	}

	return kvs
}

// modify context
func wrapContext(c context.Context) context.Context {
	if stageSwitch != On {
		return c
	}

	// logs包支持传递nil context
	if c == nil {
		return c
	}

	if p := getMarkPtr(c); p != nil {
		kvs := output(*p)
		return logs.CtxAddKVs(c, kvs...)
	}

	return c
}
