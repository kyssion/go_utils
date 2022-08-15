package util

import (
	"encoding/json"
	"reflect"
	"strconv"
	"time"

	"code.byted.org/ad/gromore/lib/consts"

	"code.byted.org/ad/gromore/lib/utils/enhancelog"
)

// JSONMarshal 将对象序列化为json字符串
func JSONMarshal(v interface{}) string {
	if v == nil {
		return ""
	}
	b, err := json.Marshal(v)
	if err != nil {
		enhancelog.Warn("JSONMarshal failed for v: %+v, err: %+v", v, err)
		return `{"err_msg": "not json format"}`
	}
	return string(b)
}

// TimeStringToUnix 将时间String形式转换为Unix
func TimeStringToUnix(date string) (int64, error) {
	loc, _ := time.LoadLocation("Local")
	item, err := time.ParseInLocation(consts.DateFormat, date, loc)
	if err != nil {
		return 0, err
	}

	return item.Unix(), nil
}

// TimeHourStringToUnix 将小时级时间String形式转换为Unix
func TimeHourStringToUnix(date string) (int64, error) {

	loc, _ := time.LoadLocation("Local")
	item, err := time.ParseInLocation(consts.TimeFormatISO, date, loc)
	if err != nil {
		return 0, err
	}

	return item.Unix(), nil
}

// TimeAddOneDayStringToUnix 加一天，然后将String转换为Unix
func TimeAddOneDayStringToUnix(date string) (int64, error) {

	loc, _ := time.LoadLocation("Local")
	item, err := time.ParseInLocation("2006-01-02", date, loc)
	if err != nil {
		return 0, err
	}

	return item.AddDate(0, 0, 1).Unix(), nil
}

func SliceToMap(slice interface{}) map[interface{}]struct{} {

	set := map[interface{}]struct{}{}
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice && v.Kind() != reflect.Array {
		return set
	}

	for i := 0; i < v.Len(); i++ {
		set[v.Index(i).Interface()] = struct{}{}
	}
	return set
}

func InStringList(list []string, s string) bool {
	for _, l := range list {
		if l == s {
			return true
		}
	}
	return false
}

func InIntList(list []int, item int) bool {
	for _, i := range list {
		if i == item {
			return true
		}
	}
	return false
}

func GetFloat64FromStr(num string) (float64, error) {
	floatItem, err := strconv.ParseFloat(num, 64)
	if err != nil {
		enhancelog.Warn("get price error , num : %v", num)
		return 0.0, err
	}
	return floatItem, nil
}
