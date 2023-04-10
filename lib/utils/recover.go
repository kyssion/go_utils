package util

import (
	"context"
	"runtime"

	"code.byted.org/gopkg/logs"
)

func RecoverPanic(ctx context.Context, errName string) {
	if e := recover(); e != nil {
		buf := make([]byte, 64<<10)
		buf = buf[:runtime.Stack(buf, false)]
		logs.CtxError(ctx, "[%s] error: %v stack: %v", errName, e, string(buf))
		EmitCounter("panic", 1, map[string]string{"func": errName})
	}
}
