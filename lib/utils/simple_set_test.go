package util

import (
	"reflect"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func initSet() SimpleSet {
	set := NewSimpleSet()
	set.Add(1)
	return set
}

func TestSimpleSet_Add(t *testing.T) {
	type args struct {
		i interface{}
	}

	set := initSet()
	tests := []struct {
		name string
		set  SimpleSet
		args args
		want bool
	}{
		{
			name: "测试set的添加一个不存在的元素",
			set:  set,
			args: args{
				i: 3,
			},
			want: true,
		},
		{
			name: "测试set的添加一个已存在的元素",
			set:  set,
			args: args{
				i: 1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.set.Add(tt.args.i); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSimpleSet_Contains(t *testing.T) {
	type args struct {
		i []interface{}
	}
	set := initSet()

	tests := []struct {
		name string
		set  SimpleSet
		args args
		want bool
	}{
		{
			name: "测试一个已存在的元素",
			set:  set,
			args: args{
				i: []interface{}{1},
			},
			want: true,
		},
		{
			name: "测试一个不存在的元素",
			set:  set,
			args: args{
				i: []interface{}{3},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.set.Contains(tt.args.i...); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSimpleSet_Remove(t *testing.T) {
	type args struct {
		i interface{}
	}

	set := initSet()

	tests := []struct {
		name string
		set  SimpleSet
		args args
		want bool
	}{
		{
			name: "测试移除一个已存在元素",
			set:  set,
			args: args{
				i: 1,
			},
			want: true,
		},
		{
			name: "测试移除一个不存在元素",
			set:  set,
			args: args{
				i: 3,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.set.Remove(tt.args.i); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSimpleSet_Values(t *testing.T) {
	tests := []struct {
		name string
		set  SimpleSet
		want []interface{}
	}{
		{
			name: "获取set中的元素,返回值类型为interface",
			set:  initSet(),
			want: []interface{}{
				1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.set.Values(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Values() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSimpleSet_ToInt64(t *testing.T) {
	tests := []struct {
		name string
		set  SimpleSet
		want []int64
	}{
		{
			name: "获取set中的元素,返回值类型为int64",
			set:  initSet(),
			want: []int64{
				1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.set.ToInt64(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSimpleSet_AddList(t *testing.T) {
	type args struct {
		values interface{}
	}

	set := NewSimpleSet()

	tests := []struct {
		name string
		set  SimpleSet
		args args
	}{
		{
			name: "添加数组",
			set:  set,
			args: args{
				[]int64{12, 14},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			convey.Convey(tt.name, t, func() {
				tt.set.AddList(tt.args.values)
				convey.So(len(tt.set.ToInt64()), convey.ShouldEqual, 2)
			})
		})
	}
}
