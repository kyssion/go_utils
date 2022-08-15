package util

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"

	"code.byted.org/ad/gromore/lib/consts"
)

func MD5(s string) string {
	h := md5.New()
	io.WriteString(h, s)
	return hex.EncodeToString(h.Sum(nil))
}

func CheckIsNumberAndInSection(str string, low int, high int) bool {
	number, err := strconv.Atoi(str)
	if err != nil {
		return false
	}
	if !(low <= number && number <= high) {
		return false
	}
	return true
}

// StringToInt 字符串转int，不支持科学计数法，例如输入1e3，返回为0
func StringToInt(numStr string) int {
	if num, ok := strconv.Atoi(numStr); ok == nil {
		return num
	}
	return 0
}

func IntToString(num int) string {
	if num == 0 {
		return "0"
	}
	return strconv.Itoa(num)
}

func StringToInt32(numStr string) int32 {
	if num, ok := strconv.Atoi(numStr); ok == nil {
		return int32(num)
	}
	return 0
}

func Int32ToString(num int32) string {
	return strconv.Itoa(int(num))
}

func StringToInt64(numStr string) int64 {
	if num, ok := strconv.ParseInt(numStr, 10, 64); ok == nil {
		return num
	}
	return 0
}

// StringToFloat64 convert string to float64 type,return 0 when convert failed
func StringToFloat64(numStr string) (float64, error) {
	num, err := strconv.ParseFloat(numStr, 64)
	if err != nil {
		return 0, err
	}
	return num, nil
}

func Int64ToString(num int64) string {
	return strconv.FormatInt(num, 10)
}

func SubString(str string, start int, end int) string {
	rs := []rune(str)
	length := len(rs)

	if start < 0 || start > length {
		return ""
	}

	if end < 0 || end > length {
		return ""
	}
	if start >= end {
		return ""
	}

	return string(rs[start:end])
}

func SubStringByLength(str string, start int, length int) string {
	rs := []rune(str)
	lengthReal := len(rs)
	end := start + length

	if start < 0 || start > lengthReal || length < 0 {
		return ""
	}
	if end > lengthReal {
		return ""
	}

	return string(rs[start:end])
}

func SplitByReg(str, split string) []string {
	if str == "" {
		return []string{}
	}

	reg := regexp.MustCompile(split)
	indexes := reg.FindAllStringIndex(str, -1)
	laststart := 0
	result := make([]string, len(indexes)+1)
	for i, element := range indexes {
		result[i] = str[laststart:element[0]]
		laststart = element[1]
	}
	result[len(indexes)] = str[laststart:len(str)]

	return result

}

func CheckStringIsNumber(numStr string) bool {
	if _, err := strconv.Atoi(numStr); err != nil {
		return false
	}
	return true
}

func CheckStringIsNumberAndInSection(str string, low int, high int) bool {
	number, err := strconv.Atoi(str)
	if err != nil || high < low {
		return false
	}
	if !(low <= number && number <= high) {
		return false
	}
	return true
}

func CheckStringIsFloat(numStr string) bool {
	if _, err := strconv.ParseFloat(numStr, 64); err != nil {
		return false
	}
	return true
}

func ToString(v interface{}) string {
	switch vv := v.(type) {
	case string:
		return vv
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return strconv.FormatInt(reflect.ValueOf(vv).Int(), 10)
	case float32, float64:
		value := strconv.FormatFloat(reflect.ValueOf(vv).Float(), 'f', 2, 64)
		if strings.HasSuffix(value, ".00") {
			return strings.ReplaceAll(value, ".00", "")
		}
		return value
	case bool:
		return strconv.FormatBool(vv)
	default:
		return ""
	}
}
func TransStringArray2String(srcData []string, mark string) string {
	return strings.Join(srcData, mark)
}

func SplitToStrings(data string, mark string) []string {
	return strings.Split(data, mark)
}

func TransArrayFromString2Int64(srcData []string) []int64 {
	var target = make([]int64, 0)
	for _, data := range srcData {
		target = append(target, StringToInt64(data))
	}
	return target
}

func TransArray2String(srcData interface{}) ([]string, error) {
	var targets = make([]string, 0)
	reflectValue := reflect.ValueOf(srcData)
	for i := 0; i < reflectValue.Len(); i++ {
		currentValue := reflectValue.Index(i)
		if currentValue.Kind() == reflect.String {
			targets = append(targets, currentValue.String())
		} else {
			num := GetNumberValue(currentValue.Interface())
			if num == MinNumber {
				return nil, fmt.Errorf("type: %v is invalid", currentValue.Kind().String())
			}
			targets = append(targets, Int64ToString(num))
		}
	}
	return targets, nil
}

// GetPtrString 获取string值的指针
func GetPtrString(value string) *string {
	return &value
}

func FormatTimeStr(timeStr string) *time.Time {
	if timeStr == "" {
		return nil
	}
	loc, _ := time.LoadLocation("Local")
	if isoTime, err := time.ParseInLocation(consts.TimeFormatISO, timeStr, loc); err == nil {
		return &isoTime
	}
	if isoTime, err := time.ParseInLocation(consts.HourDateFormat, timeStr, loc); err == nil {
		return &isoTime
	}
	if data, err := time.ParseInLocation(consts.DateFormat, timeStr, loc); err == nil {
		return &data
	}
	return nil
}

func MatchByReg(target string, pattern string) bool {
	match, err := regexp.MatchString(pattern, target)
	if err != nil {
		return false
	}
	return match
}

func Int64ToStringPtr(num int64) *string {
	s := strconv.FormatInt(num, 10)
	return &s
}

func CheckStringHasChineseChar(str string) bool {
	for _, r := range str {
		if unicode.Is(unicode.Scripts["Han"], r) || (regexp.MustCompile("[\u3002\uff1b\uff0c\uff1a\u201c\u201d\uff08\uff09\u3001\uff1f\u300a\u300b]").MatchString(string(r))) {
			return true
		}
	}
	return false
}
