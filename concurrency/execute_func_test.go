package concurrency

import (
	"fmt"
	"testing"
)

type DDD int
type Map1[KEY int | string, VALUE string | float64] map[KEY]VALUE // 声明范型map
type Slice[T int | float64] []T                                   // 声明范型切片
type Struct1[T string | int | float64] struct {
	Title   string
	Content T
} // 声明范型结构体

func Test1(t *testing.T) {
	var d DDD = 123
	var P Slice[int] = []int{123, 223} // 还不支持诶性推导 - 这里需要直接生命变量
	var M Map1[string, string] = map[string]string{
		"t": "t",
	}
	var STR = Struct1[string]{
		Title:   "",
		Content: "",
	}

	d2 := DDD(123)             // 简写
	P2 := Slice[int]{222, 333} // 简写
	M2 := Map1[string, string]{
		"T": "t",
	}
	// 匿名结构体不支持
	fmt.Printf("%d , %v , %v ， %v\n", d, P, M, STR)
	fmt.Printf("%d , %v , %v ， %v\n", d2, P2, M2)
}

// 嵌套范型结构体 - 注意点： 子类型的范型类型需要在前面声明过
type MyStruct[S int | string, P map[S]string] struct {
	Name    string
	Content S
	Job     P
}

func Test2(t *testing.T) {
	ppp := MyStruct[int, map[int]string]{
		Name:    "",
		Content: 12,
		Job: map[int]string{
			1: "123",
		},
	}
	fmt.Sprintf("%v", ppp)
}
