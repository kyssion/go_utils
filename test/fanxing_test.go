package test

import (
	"fmt"
	"testing"
)

type TestStruct struct {
	ID int64
}

// 1. PrintSlice 声明一个函数范型
func PrintSlice[T any](s []T) {
	for _, item := range s {
		print(fmt.Sprintf("%v ", item))
	}
	println()
}

func TestPrintSlice(t *testing.T) {
	t.Run("测试范型函数", func(t *testing.T) {
		PrintSlice([]string{"one", "two"})
		PrintSlice([]int{1, 2, 3})
		PrintSlice([]float32{1.1, 2, 2})
		PrintSlice([]*TestStruct{
			{
				ID: 123,
			}, {
				ID: 124,
			},
		})
	})
}

// 2. 声明一个范型切片 - 声明一个类型为T的 数组范型， 并给他重命名为Vector
type Vector[T any] []T

func TestVector(t *testing.T) {
	t.Run("testVector", func(t *testing.T) {
		vec1 := Vector[int]{1, 2, 3, 4, 5}
		vec2 := Vector[TestStruct]{{ID: 1}, {ID: 2}}
		PrintSlice(vec1)
		PrintSlice(vec2)
	})
}

// 3. 声明一个Map切片
type M[K string, V any] map[K]V // 注意这个地方 ， golang的K 并不是支持任意的类型

func TestMap(t *testing.T) {
	t.Run("test map", func(t *testing.T) {
		m1 := make(M[string, int])
		m1["one"] = 123
		m2 := make(M[string, string])
		m2["one"] = "two"
	})
}
