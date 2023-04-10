package util

import (
	"reflect"
	"testing"

	. "code.byted.org/gopkg/mockito"

	"github.com/magiconair/properties/assert"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCheckStringHasChineseChar(t *testing.T) {

	Convey("纯字母", t, func() {
		data := "abrsardas"
		So(CheckStringHasChineseChar(data), ShouldEqual, false)
	})

	Convey("纯数字", t, func() {
		data := "125313531513"
		So(CheckStringHasChineseChar(data), ShouldEqual, false)
	})

	Convey("英文字符", t, func() {
		data := "@#￥&&￥*%￥"
		So(CheckStringHasChineseChar(data), ShouldEqual, false)
	})

	Convey("英文字符混合", t, func() {
		data := "abc112vv,@#$ has %+v"
		So(CheckStringHasChineseChar(data), ShouldEqual, false)
	})

	Convey("带中文", t, func() {
		data := "周详 has"
		So(CheckStringHasChineseChar(data), ShouldEqual, true)
	})

	Convey("中文标点", t, func() {
		data := "13，。；"
		So(CheckStringHasChineseChar(data), ShouldEqual, true)
	})

	Convey("全混合", t, func() {
		data := "abrwef额13！；"
		So(CheckStringHasChineseChar(data), ShouldEqual, true)
	})
}

func TestCheckIsNumberAndInSection(t *testing.T) {

	Convey("空串", t, func() {
		So(CheckIsNumberAndInSection("", 1, 10), ShouldEqual, false)
	})

	Convey("中值", t, func() {
		So(CheckIsNumberAndInSection("5", 1, 10), ShouldEqual, true)
	})

	Convey("下边界值", t, func() {
		So(CheckIsNumberAndInSection("5", 1, 10), ShouldEqual, true)
	})

	Convey("上边界值", t, func() {
		So(CheckIsNumberAndInSection("10", 1, 10), ShouldEqual, true)
	})

	Convey("越界值", t, func() {
		So(CheckIsNumberAndInSection("12", 1, 10), ShouldEqual, false)
		So(CheckIsNumberAndInSection("0", 1, 10), ShouldEqual, false)
	})

	Convey("非法值", t, func() {
		So(CheckIsNumberAndInSection("10&zxgf", 1, 10), ShouldEqual, false)
	})
}

func TestMatchByReg(t *testing.T) {

	Convey("", t, func() {
		So(MatchByReg("QuWorker", "Worker$"), ShouldEqual, true)
	})
}

func TestCheckStringIsFloat(t *testing.T) {
	type args struct {
		val string
	}

	tests := []struct {
		name  string
		args  args
		wantV bool
	}{
		{
			name: "合法",
			args: args{
				val: "15.0",
			},
			wantV: true,
		},
		{
			name: "0",
			args: args{
				val: "0.0",
			},
			wantV: true,
		},
		{
			name: "-10086.12663",
			args: args{
				val: "-10086.12663",
			},
			wantV: true,
		},
		{
			name: "s1t0",
			args: args{
				val: "s1t0",
			},
			wantV: false,
		},
		{
			name: "#q1",
			args: args{
				val: "#q1",
			},
			wantV: false,
		},
		{
			name: "[;",
			args: args{
				val: "[;",
			},
			wantV: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotV := CheckStringIsFloat(tt.args.val); gotV != tt.wantV {
				t.Errorf("Int32ToString() = %v, want %v", gotV, tt.wantV)
			}
		})
	}
}

func TestCheckStringIsNumberAndInSection(t *testing.T) {
	type args struct {
		val  string
		low  int
		high int
	}

	tests := []struct {
		name  string
		args  args
		wantV bool
	}{
		{
			name: "合法",
			args: args{
				val:  "15",
				low:  -20,
				high: 25,
			},
			wantV: true,
		},
		{
			name: "-5",
			args: args{
				val:  "-5",
				low:  -20,
				high: 25,
			},
			wantV: true,
		},
		{
			name: "100",
			args: args{
				val:  "100",
				low:  -20,
				high: 25,
			},
			wantV: false,
		},
		{
			name: "-200",
			args: args{
				val:  "-200",
				low:  -20,
				high: 25,
			},
			wantV: false,
		},
		{
			name: "-10086",
			args: args{
				val:  "-10086",
				low:  -20,
				high: 25,
			},
			wantV: false,
		},
		{
			name: "s1t0",
			args: args{
				val:  "s1t0",
				low:  -20,
				high: 25,
			},
			wantV: false,
		},
		{
			name: "#q1",
			args: args{
				val:  "#q1",
				low:  -20,
				high: 25,
			},
			wantV: false,
		},
		{
			name: "[;",
			args: args{
				val:  "[;",
				low:  -20,
				high: 25,
			},
			wantV: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotV := CheckStringIsNumberAndInSection(tt.args.val, tt.args.low, tt.args.high); gotV != tt.wantV {
				t.Errorf("Int32ToString() = %v, want %v", gotV, tt.wantV)
			}
		})
	}
}

func TestCheckStringIsNumber(t *testing.T) {
	type args struct {
		val string
	}

	tests := []struct {
		name  string
		args  args
		wantV bool
	}{
		{
			name: "合法",
			args: args{
				val: "15",
			},
			wantV: true,
		},
		{
			name: "0",
			args: args{
				val: "0",
			},
			wantV: true,
		},
		{
			name: "-10086",
			args: args{
				val: "-10086",
			},
			wantV: true,
		},
		{
			name: "s1t0",
			args: args{
				val: "s1t0",
			},
			wantV: false,
		},
		{
			name: "#q1",
			args: args{
				val: "#q1",
			},
			wantV: false,
		},
		{
			name: "[;",
			args: args{
				val: "[;",
			},
			wantV: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotV := CheckStringIsNumber(tt.args.val); gotV != tt.wantV {
				t.Errorf("Int32ToString() = %v, want %v", gotV, tt.wantV)
			}
		})
	}
}

func TestSplitByReg(t *testing.T) {

	type args struct {
		val  string
		flag string
	}

	tests := []struct {
		name  string
		args  args
		wantV []string
	}{
		{
			name: "合法",
			args: args{
				val:  "abcdefghi#jkl#mno###pqr#stuv#wxyz",
				flag: "[#]+",
			},
			wantV: []string{"abcdefghi", "jkl", "mno", "pqr", "stuv", "wxyz"},
		},
		{
			name: "空输入",
			args: args{
				val:  "",
				flag: "[#]+",
			},
			wantV: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotV := SplitByReg(tt.args.val, tt.args.flag); !StringArrayEqual(gotV, tt.wantV) {
				t.Errorf("Int32ToString() = %v, want %v", gotV, tt.wantV)
			}
		})
	}
}

func StringArrayEqual(src []string, target []string) bool {
	if (src == nil && target != nil) || (src != nil && target == nil) {
		return false
	}
	if len(src) != len(target) {
		return false
	}

	existMap := map[string]bool{}
	for i := 0; i < len(src); i++ {
		existMap[src[i]] = true
	}
	for i := 0; i < len(target); i++ {
		if _, exist := existMap[target[i]]; !exist {
			return false
		}
	}

	return true
}

func TestSubStringByLength(t *testing.T) {

	type args struct {
		val    string
		low    int
		length int
	}

	tests := []struct {
		name  string
		args  args
		wantV string
	}{
		{
			name: "合法",
			args: args{
				val:    "abcdefghijklmnopqrstuvwxyz",
				low:    0,
				length: 4,
			},
			wantV: "abcd",
		},
		{
			name: "非法下界",
			args: args{
				val:    "abcdefghijklmnopqrstuvwxyz",
				low:    -5,
				length: 4,
			},
			wantV: "",
		},
		{
			name: "非法上界",
			args: args{
				val:    "abcdefghijklmnopqrstuvwxyz",
				low:    200,
				length: 5,
			},
			wantV: "",
		},
		{
			name: "非法上下界",
			args: args{
				val:    "abcdefghijklmnopqrstuvwxyz",
				low:    -5,
				length: 149,
			},
			wantV: "",
		},
		{
			name: "非法长度",
			args: args{
				val:    "abcdefghijklmnopqrstuvwxyz",
				low:    20,
				length: -9,
			},
			wantV: "",
		},
		{
			name: "非法长度",
			args: args{
				val:    "abcdefghijklmnopqrstuvwxyz",
				low:    20,
				length: 25,
			},
			wantV: "",
		},
		{
			name: "合法长度",
			args: args{
				val:    "abcdefghijklmnopqrstuvwxyz",
				low:    25,
				length: 1,
			},
			wantV: "z",
		},
		{
			name: "非法长度",
			args: args{
				val:    "abcdefghijklmnopqrstuvwxyz",
				low:    25,
				length: 2,
			},
			wantV: "",
		},
		{
			name: "非法起点&长度",
			args: args{
				val:    "abcdefghijklmnopqrstuvwxyz",
				low:    50,
				length: -9,
			},
			wantV: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotV := SubStringByLength(tt.args.val, tt.args.low, tt.args.length); gotV != tt.wantV {
				t.Errorf("Int32ToString() = %v, want %v", gotV, tt.wantV)
			}
		})
	}
}

func TestSubString(t *testing.T) {

	type args struct {
		val  string
		low  int
		high int
	}

	tests := []struct {
		name  string
		args  args
		wantV string
	}{
		{
			name: "合法",
			args: args{
				val:  "abcdefghijklmnopqrstuvwxyz",
				low:  0,
				high: 4,
			},
			wantV: "abcd",
		},
		{
			name: "非法下界",
			args: args{
				val:  "abcdefghijklmnopqrstuvwxyz",
				low:  -5,
				high: 4,
			},
			wantV: "",
		},
		{
			name: "非法上界",
			args: args{
				val:  "abcdefghijklmnopqrstuvwxyz",
				low:  0,
				high: 200,
			},
			wantV: "",
		},
		{
			name: "非法上下界",
			args: args{
				val:  "abcdefghijklmnopqrstuvwxyz",
				low:  -5,
				high: 49,
			},
			wantV: "",
		},
		{
			name: "非法low&high",
			args: args{
				val:  "abcdefghijklmnopqrstuvwxyz",
				low:  15,
				high: 9,
			},
			wantV: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotV := SubString(tt.args.val, tt.args.low, tt.args.high); gotV != tt.wantV {
				t.Errorf("Int32ToString() = %v, want %v", gotV, tt.wantV)
			}
		})
	}
}

func TestInt64ToString(t *testing.T) {
	type args struct {
		val int64
	}

	tests := []struct {
		name  string
		args  args
		wantV string
	}{
		{
			name: "1",
			args: args{
				1,
			},
			wantV: "1",
		},
		{
			name: "-1",
			args: args{
				val: -1,
			},
			wantV: "-1",
		},
		{
			name: "1000",
			args: args{
				val: 1000,
			},
			wantV: "1000",
		},
		{
			name: "1e5",
			args: args{
				val: 1e5,
			},
			wantV: "100000",
		},
		{
			name: "0",
			args: args{
				val: 0,
			},
			wantV: "0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotV := Int64ToString(tt.args.val); gotV != tt.wantV {
				t.Errorf("Int32ToString() = %v, want %v", gotV, tt.wantV)
			}
		})
	}
}

func TestInt32ToString(t *testing.T) {
	type args struct {
		val int32
	}

	tests := []struct {
		name  string
		args  args
		wantV string
	}{
		{
			name: "1",
			args: args{
				1,
			},
			wantV: "1",
		},
		{
			name: "-1",
			args: args{
				val: -1,
			},
			wantV: "-1",
		},
		{
			name: "1000",
			args: args{
				val: 1000,
			},
			wantV: "1000",
		},
		{
			name: "1e5",
			args: args{
				val: 1e5,
			},
			wantV: "100000",
		},
		{
			name: "0",
			args: args{
				val: 0,
			},
			wantV: "0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotV := Int32ToString(tt.args.val); gotV != tt.wantV {
				t.Errorf("Int32ToString() = %v, want %v", gotV, tt.wantV)
			}
		})
	}
}

func TestMD5(t *testing.T) {

	type args struct {
		val string
	}

	tests := []struct {
		name  string
		args  args
		wantV string
	}{
		{
			name: "简单测试",
			args: args{
				"123456",
			},
			wantV: "e10adc3949ba59abbe56e057f20f883e",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotV := MD5(tt.args.val); gotV != tt.wantV {
				t.Errorf("MD5() = %v, want %v", gotV, tt.wantV)
			}
		})
	}
}

func TestStringToInt(t *testing.T) {

	type args struct {
		val string
	}

	tests := []struct {
		name  string
		args  args
		wantV int
	}{
		{
			name: "1",
			args: args{
				"1",
			},
			wantV: 1,
		},
		{
			name: "a",
			args: args{
				val: "a",
			},
			wantV: 0,
		},
		{
			name: "-1",
			args: args{
				val: "-1",
			},
			wantV: -1,
		},
		{
			name: "#4",
			args: args{
				val: "#4",
			},
			wantV: 0,
		},
		{
			name: "1e3",
			args: args{
				val: "1e3",
			},
			wantV: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotV := StringToInt(tt.args.val); gotV != tt.wantV {
				t.Errorf("StringToInt() = %v, want %v", gotV, tt.wantV)
			}
		})
	}
}

func TestStringToInt32(t *testing.T) {

	type args struct {
		val string
	}

	tests := []struct {
		name  string
		args  args
		wantV int32
	}{
		{
			name: "1",
			args: args{
				"1",
			},
			wantV: 1,
		},
		{
			name: "a",
			args: args{
				val: "a",
			},
			wantV: 0,
		},
		{
			name: "-1",
			args: args{
				val: "-1",
			},
			wantV: -1,
		},
		{
			name: "#4",
			args: args{
				val: "#4",
			},
			wantV: 0,
		},
		{
			name: "1e3",
			args: args{
				val: "1e3",
			},
			wantV: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotV := StringToInt32(tt.args.val); gotV != tt.wantV {
				t.Errorf("StringToInt() = %v, want %v", gotV, tt.wantV)
			}
		})
	}
}

func TestStringToInt64(t *testing.T) {

	type args struct {
		val string
	}

	tests := []struct {
		name  string
		args  args
		wantV int64
	}{
		{
			name: "1",
			args: args{
				"1",
			},
			wantV: 1,
		},
		{
			name: "a",
			args: args{
				val: "a",
			},
			wantV: 0,
		},
		{
			name: "-1",
			args: args{
				val: "-1",
			},
			wantV: -1,
		},
		{
			name: "#4",
			args: args{
				val: "#4",
			},
			wantV: 0,
		},
		{
			name: "1e3",
			args: args{
				val: "1e3",
			},
			wantV: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotV := StringToInt64(tt.args.val); gotV != tt.wantV {
				t.Errorf("StringToInt() = %v, want %v", gotV, tt.wantV)
			}
		})
	}
}

func TestStringToFloat64(t *testing.T) {

	type args struct {
		val string
	}

	tests := []struct {
		name  string
		args  args
		wantV float64
	}{
		{
			name: "1",
			args: args{
				"1",
			},
			wantV: 1.0,
		},
		{
			name: "a",
			args: args{
				val: "a",
			},
			wantV: 0,
		},
		{
			name: "-1",
			args: args{
				val: "-1",
			},
			wantV: -1,
		},
		{
			name: "#4",
			args: args{
				val: "#4",
			},
			wantV: 0,
		},
		{
			name: "1e3",
			args: args{
				val: "1e3",
			},
			wantV: 1000,
		},
		{
			name: "1e-5",
			args: args{
				val: "1e-5",
			},
			wantV: 0.00001,
		},
		{
			name: "-0.31415916",
			args: args{
				val: "-0.31415916",
			},
			wantV: -0.31415916,
		},
		{
			name: "3141.5916",
			args: args{
				val: "3141.5916",
			},
			wantV: 3141.5916,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotV, _ := StringToFloat64(tt.args.val); gotV != tt.wantV {
				t.Errorf("StringToFloat64() = %v, want %v", gotV, tt.wantV)
			}
		})
	}
}

func TestIntToString(t *testing.T) {

	type args struct {
		val int
	}

	tests := []struct {
		name  string
		args  args
		wantV string
	}{
		{
			name: "1",
			args: args{
				1,
			},
			wantV: "1",
		},
		{
			name: "-1",
			args: args{
				val: -1,
			},
			wantV: "-1",
		},
		{
			name: "1000",
			args: args{
				val: 1000,
			},
			wantV: "1000",
		},
		{
			name: "1e5",
			args: args{
				val: 1e5,
			},
			wantV: "100000",
		},
		{
			name: "0",
			args: args{
				val: 0,
			},
			wantV: "0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotV := IntToString(tt.args.val); gotV != tt.wantV {
				t.Errorf("StringToInt() = %v, want %v", gotV, tt.wantV)
			}
		})
	}
}

func TestToString(t *testing.T) {
	var number1 int32 = 123
	var number2 int64 = 1234
	var number3 float32 = 134.1
	var number4 float64 = 123344.1
	assert.Equal(t, ToString(number1), "123")
	assert.Equal(t, ToString(number2), "1234")
	assert.Equal(t, ToString(number3), "134.10")
	assert.Equal(t, ToString(number4), "123344.10")
}
func TestGetPtrString(t *testing.T) {
	type args struct {
		value string
	}

	value := "1"

	tests := []struct {
		name string
		args args
		want *string
	}{
		{
			name: "字符串转指针",
			args: args{
				value: value,
			},
			want: &value,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetPtrString(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPtrString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatTimeStr(t *testing.T) {
	Convey("Test FormatTimeStr", t, func() {
		So(FormatTimeStr(""), ShouldBeNil)
		So(FormatTimeStr("2021-08-01 12:00:00"), ShouldNotBeNil)
		So(FormatTimeStr("2021-08-01 12"), ShouldNotBeNil)
		So(FormatTimeStr("2021-08-01"), ShouldNotBeNil)
	})
}

func TestTransArrayFromString2Int64(t *testing.T) {

}

func TestTransArray2String(t *testing.T) {
	PatchConvey("转换成字符串数组", t, func() {
		Convey("原始数据是int", func() {
			target, err := TransArray2String([]int{1})
			So(err, ShouldBeNil)
			So(target, ShouldResemble, []string{"1"})
		})

		Convey("原始数据是string", func() {
			target, err := TransArray2String([]string{"1"})
			So(err, ShouldBeNil)
			So(target, ShouldResemble, []string{"1"})
		})

		Convey("原始数据是自定义类型，报错", func() {
			type test1 struct {
				a string
			}
			target, err := TransArray2String([]test1{{a: "1"}})
			So(target, ShouldBeNil)
			So(err.Error(), ShouldEqual, "type: struct is invalid")
		})
	})
}
