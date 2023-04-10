package util

import (
	"fmt"
	"math"
	"strconv"
)

const (
	MinTolerant = 0.000001 // float精度丢失的最小容忍度
)

func MinInt32(a, b int32) int32 {
	if a <= b {
		return a
	}
	return b
}

// Decimal float保留两位小数（四舍五入）
func Decimal(value float64) float64 {
	num := math.Trunc(value*1e2+0.5) * 1e-2
	return num
}

// DecimalV2 float保留两位小数（四舍五入）避免精度问题
func DecimalV2(value float64) float64 {
	num, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return num
}

// Decimal4 float保留四位小数（四舍五入）避免精度问题
func Decimal4(value float64) float64 {
	num, _ := strconv.ParseFloat(fmt.Sprintf("%.4f", value), 64)
	return num
}

// GetPercentValueWithRound 四舍五入后取百分位的值
func GetPercentValueWithRound(value float64) int64 {
	return int64(value*100 + 0.5)
}

// GetPtrFloat64 获取float64数值的指针
func GetPtrFloat64(value float64) *float64 {
	return &value
}

// GetPtrInt 获取int数值的指针
func GetPtrInt(value int) *int {
	return &value
}

// GetPtrInt32 获取int32数值的指针
func GetPtrInt32(value int32) *int32 {
	return &value
}

// GetPtrInt64 获取int64数值的指针
func GetPtrInt64(value int64) *int64 {
	return &value
}

// GetPtrInt16 获取int32数值的指针
func GetPtrInt16(value int16) *int16 {
	return &value
}

func FloatIsEqual(a, b float64) bool {
	return math.Abs(a-b) < MinTolerant
}

// Percent2Decimal 百分数转成小数，并保留4位小数
func Percent2Decimal(num float64) float64 {
	return Decimal4(num / 100)
}

// SafeDivision = dividend / divisor
func SafeDivision(dividend float64, divisor float64) (result float64) {
	if divisor == 0 {
		return 0
	}
	return dividend / divisor
}
