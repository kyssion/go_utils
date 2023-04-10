package util

import (
	"reflect"
	"testing"

	. "code.byted.org/gopkg/mockito"
	. "github.com/smartystreets/goconvey/convey"
)

func TestDecimal(t *testing.T) {
	type args struct {
		value float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "float保留两位小数",
			args: args{
				value: 1.234,
			},
			want: 1.23,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Decimal(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decimal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecimalV2(t *testing.T) {
	type args struct {
		value float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "取两位小数，四舍五入",
			args: args{
				value: 2.345,
			},
			want: 2.35,
		},
		{
			name: "取两位小数",
			args: args{
				value: 2.344,
			},
			want: 2.34,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DecimalV2(tt.args.value); got != tt.want {
				t.Errorf("DecimalV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTruncWithRound(t *testing.T) {
	PatchConvey("四舍五入后取百分位的值", t, func() {
		So(GetPercentValueWithRound(0.727), ShouldEqual, 73)
	})
}

func TestFloatIsEqual(t *testing.T) {
	PatchConvey("测试浮点数相等", t, func() {
		Convey("相等", func() {
			So(FloatIsEqual(1.6, 1.6), ShouldEqual, true)
		})

		Convey("不相等", func() {
			So(FloatIsEqual(1.6, 1.7), ShouldEqual, false)
		})
	})
}

func TestSafeDivision(t *testing.T) {
	PatchConvey("安全的除法", t, func() {
		Convey("正常除数", func() {
			result := SafeDivision(1, 2)
			So(result, ShouldEqual, 0.5)
		})

		Convey("除数是0", func() {
			result := SafeDivision(1, 0)
			So(result, ShouldEqual, 0)
		})
	})
}

func TestPercent2Decimal(t *testing.T) {
	PatchConvey("百分数转小数", t, func() {
		num := Percent2Decimal(23.785)
		So(num, ShouldEqual, 0.2379)
	})
}
