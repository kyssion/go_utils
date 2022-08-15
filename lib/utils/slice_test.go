package util

import (
	"reflect"
	"testing"
)

func TestContainsNumber(t *testing.T) {
	type args struct {
		slice   interface{}
		element interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "测试数字类型不一致",
			args: args{
				slice:   []int64{1, 3},
				element: int32(3),
			},
			want: true,
		},
		{
			name: "测试数字一致",
			args: args{
				slice:   []int64{1, 3},
				element: int64(3),
			},
			want: true,
		},
		{
			name: "测试float64",
			args: args{
				slice:   []float64{1, 3},
				element: float64(3),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsNumber(tt.args.slice, tt.args.element); got != tt.want {
				t.Errorf("ContainsNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsString(t *testing.T) {
	type args struct {
		slice   []string
		element string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "测试字符串",
			args: args{
				slice:   []string{"1", "3"},
				element: "3",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsString(tt.args.slice, tt.args.element); got != tt.want {
				t.Errorf("ContainsString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSliceSubString(t *testing.T) {
	type args struct {
		a []string
		b []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "测试数组减法",
			args: args{
				a: []string{"1", "3", "4"},
				b: []string{"1", "3"},
			},
			want: []string{"4"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SliceSubString(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SliceSubString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSliceSubInt64(t *testing.T) {
	type args struct {
		a []int64
		b []int64
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{
			name: "测试数组减法",
			args: args{
				a: []int64{1, 3, 4},
				b: []int64{1, 3},
			},
			want: []int64{4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SliceSubInt64(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SliceSubInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsFloat(t *testing.T) {
	type args struct {
		slice   interface{}
		element float64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "包含float类型的数字",
			args: args{
				slice:   []interface{}{float64(123)},
				element: float64(123),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsFloat(tt.args.slice, tt.args.element); got != tt.want {
				t.Errorf("ContainsFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSliceInt64ToStrings(t *testing.T) {
	type args struct {
		slices []int64
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "int64数组转成字符串",
			args: args{
				slices: []int64{123},
			},
			want: []string{"123"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SliceInt64ToStrings(tt.args.slices); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SliceInt64ToStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
