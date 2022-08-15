package util

import (
	"code.byted.org/ad/gromore/lib/utils/enhancelog"

	"crypto/sha1"
	"fmt"
	"reflect"
	"sort"
	"strings"
)

var EncryptionUtilIns = NewEncryptionUtil()

type (
	IEncryptionUtil interface {
		BuildSha1WithSalt(signMap map[string]interface{}, salt string) string
	}

	EncryptionUtil struct{}
)

func NewEncryptionUtil() IEncryptionUtil {
	return &EncryptionUtil{}
}

func (e *EncryptionUtil) BuildSha1WithSalt(signMap map[string]interface{}, salt string) string {
	return BuildSha1WithSalt(signMap, salt)
}

func BuildSha1WithSalt(signMap map[string]interface{}, salt string) string {
	signStr := GenSignStr(signMap)
	finStr := signStr + salt
	sign := fmt.Sprintf("%x", sha1.Sum([]byte(finStr)))
	enhancelog.Debug("BuildSha1WithSalt get sign[%s] from sign map[%#v] salt[%s] sign str[%s]",
		sign, signMap, salt, finStr)
	return sign
}

// GenSignStr 去掉请求参数中的字节类型字段（如文件、字节流）、sign字段、值为空的字段，目前可认为仅有string，int等简单类型
func GenSignStr(signMap map[string]interface{}) string {
	signArr := make([]string, 0, len(signMap))
	signArr = traversalMap(signArr, signMap)

	sort.Strings(signArr)
	ret := strings.Join(signArr, "&")
	return ret
}

func traversalMap(signArr []string, params map[string]interface{}) []string {
	for key, value := range params {
		if key == "sign" {
			continue
		}
		if reflect.ValueOf(value).Kind() == reflect.Array || reflect.ValueOf(value).Kind() == reflect.Slice ||
			reflect.ValueOf(value).Kind() == reflect.Map {
			signArr = appendValue(signArr, key, JSONMarshal(value))
			continue
		}
		signArr = appendValue(signArr, key, value)
	}
	return signArr
}

func appendValue(signArr []string, key string, val interface{}) []string {
	finalVal := fmt.Sprintf("%v", val)
	if finalVal == "" {
		return signArr
	}
	return append(signArr, fmt.Sprintf("%v=%v", key, finalVal))
}
