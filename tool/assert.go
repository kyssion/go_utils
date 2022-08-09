package tool

import (
	"fmt"
	"reflect"
)

// Assert 结构化使用校验数据信息，前面的表达式 ， 后面的是对应的输出参数
// Assert
func Assert(b bool, fmtInfo ...interface{}) {
	if !b {
		var fmtStr string
		if len(fmtInfo) == 0 {
			fmtStr = "unexpected happened"
		} else if _, ok := fmtInfo[0].(string); !ok {
			fmtStr = "%+v"
		} else {
			fmtStr = fmtInfo[0].(string)
			fmtInfo = fmtInfo[1:]
		}
		panic(fmt.Sprintf(fmtStr, fmtInfo...))
	}
}

func AssertFunc(target interface{}) {
	Assert(reflect.TypeOf(target).Kind() == reflect.Func, "'%v' is not a function")
}

func AssertPtr(ptr interface{}) {
	Assert(reflect.TypeOf(ptr).Kind() == reflect.Ptr, "'%v' is not a pointer")
}
