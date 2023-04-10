package util

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"reflect"
	"strings"
	"unicode"

	"code.byted.org/ad/gromore/lib/utils/enhancelog"
)

const (
	UNDERLINE = 1
	CAMEL     = 2
	MinNumber = math.MinInt64
)

func GetNumberValue(value interface{}) (v int64) {
	k := reflect.ValueOf(value).Kind()
	switch k {
	case reflect.Int:
		return int64(value.(int))
	case reflect.Int8:
		return int64(value.(int8))
	case reflect.Int16:
		return int64(value.(int16))
	case reflect.Int32:
		return int64(value.(int32))
	case reflect.Int64:
		return reflect.ValueOf(value).Int()
	default:
		return MinNumber
	}
}

// Camel2Case name from camel to underline
func Camel2Case(name string) string {
	var buffer strings.Builder
	// replace geo term. egg. change ID -> Id
	for oldStr, newStr := range getCommonTerm() {
		name = strings.ReplaceAll(name, oldStr, newStr)
	}
	for i, r := range name {
		if unicode.IsUpper(r) {
			// if current rune is not first letter, add '_' before current rune
			if i != 0 {
				buffer.WriteString("_")
			}
			buffer.WriteRune(unicode.ToLower(r))
		} else {
			buffer.WriteRune(r)
		}
	}
	return buffer.String()
}

func getCommonTerm() map[string]string {
	return map[string]string{
		"ID":   "Id",
		"HTTP": "Http",
		"JSON": "Json",
		"URL":  "Url",
		"IP":   "Ip",
		"SQL":  "Sql",
	}
}

// ConvertStruct 将当前对象的字段值赋值到目标结构体的相同字段上
/***
oldAdSlot := util.ConvertStruct(ctx, adSlot, &model.AdSlotInfo{}, nil).(*model.AdSlotInfo)
*/
func ConvertStruct(ctx context.Context, originValue interface{}, targetValue interface{}, ignoreFields []string, ignoreZero ...bool) interface{} {
	// 默认是true，也就是忽略转换为零值的字段。（如果不需要关心为零值的字段，那么可以不传该参数）
	ignoreZeroValue := true
	if len(ignoreZero) > 0 {
		// 如果传了值，默认转换成false，认为需要转换为零值的字段
		ignoreZeroValue = false
	}
	if !reflect.ValueOf(originValue).IsValid() {
		return targetValue
	}
	reflectOrigin := reflect.ValueOf(originValue).Elem()
	targetType := reflect.TypeOf(targetValue).Elem()
	reflectTarget := reflect.ValueOf(targetValue).Elem()

	if !reflectOrigin.IsValid() {
		enhancelog.CtxWarn(ctx, "originValue: %+v is invalid", originValue)
		return targetValue
	}
	for i := 0; i < targetType.NumField(); i++ {
		fieldName := targetType.Field(i).Name
		if !ContainsString(ignoreFields, fieldName) {
			setReflectValue(reflectOrigin.FieldByName(fieldName), reflectTarget.Field(i), ignoreZeroValue)
		}
	}
	return targetValue
}

// ConvertArr 将当前对象的字段值赋值到目标结构体的相同字段上 (支持数组)
func ConvertArr(ctx context.Context, originArr []interface{}, targetArr []interface{}, ignoreFields []string) []interface{} {
	refArrValue := reflect.ValueOf(targetArr)
	refArrItemValueType := refArrValue.Elem().Type()
	newArrValue := reflect.MakeSlice(refArrValue.Type(), 0, 0)
	newArr := newArrValue.Interface().([]interface{})
	for _, originItem := range originArr {
		newItem := reflect.New(refArrItemValueType)
		ConvertStruct(ctx, originItem, newItem, ignoreFields)
		newArr = append(newArr, newItem)
	}
	return newArr
}

func DeepCopyMap(ctx context.Context, src, dst interface{}) error {
	if reflect.ValueOf(dst).Kind() != reflect.Ptr {
		panic(fmt.Sprintf("dst should be reference, dst: %v", reflect.ValueOf(dst).Kind()))
	}
	bytes, err := json.Marshal(src)
	if err != nil {
		enhancelog.CtxError(ctx, "src json Marshal failed: %v, src: %+v", err, src)
		return err
	}
	if err = json.Unmarshal(bytes, &dst); err != nil {
		enhancelog.CtxError(ctx, "dst json Unmarshal failed: %v, data: %s", err, string(bytes))
		return err
	}
	return nil
}

func DeepCopyByConvert(ctx context.Context, src, dst interface{}) {
	ConvertStruct(ctx, src, dst, nil, false)
}

// AssignWithUpdateValue 先深拷贝一份旧值，然后遍历更新值中的有效字段，将有效更新字段覆盖到新值的字段中
func AssignWithUpdateValue(ctx context.Context, oldValue, updateValue interface{}, ignoreFields []string) interface{} {
	enhancelog.CtxInfo(ctx, "enter: oldValue: %+v, updateValue: %+v", oldValue, updateValue)
	newValue := reflect.New(reflect.TypeOf(oldValue).Elem())
	DeepCopyByConvert(ctx, oldValue, newValue.Interface())

	updateReflectValue := reflect.ValueOf(updateValue).Elem()
	updateReflectType := reflect.TypeOf(updateValue).Elem()

	for i := 0; i < updateReflectValue.NumField(); i++ {
		fieldName := updateReflectType.Field(i).Name
		if !ContainsString(ignoreFields, fieldName) {
			setReflectValue(updateReflectValue.Field(i), newValue.Elem().FieldByName(fieldName), true)
		}
	}
	enhancelog.CtxInfo(ctx, "exit: newValue: %+v", newValue.Interface())
	return newValue.Interface()
}

func GetMapFromStructWithFields(ctx context.Context, value interface{}, fields []string) map[string]interface{} {
	enhancelog.CtxInfo(ctx, "enter: value: %+v, need convert fields: %v", value, fields)
	reflectValue := reflect.ValueOf(value)
	reflectType := reflect.TypeOf(value)
	if reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
		reflectType = reflectType.Elem()
	}

	fieldMap := make(map[string]struct{})
	for _, fieldName := range fields {
		fieldMap[Camel2Case(fieldName)] = struct{}{}
	}

	resultValue := make(map[string]interface{})

	for i := 0; i < reflectValue.NumField(); i++ {
		fieldName := Camel2Case(reflectType.Field(i).Name)
		if _, ok := fieldMap[fieldName]; ok {
			if reflectValue.Field(i).IsZero() {
				resultValue[fieldName] = reflect.Zero(reflectValue.Field(i).Type()).Interface()
				continue
			}
			resultValue[fieldName] = reflectValue.Field(i).Interface()
			if reflectValue.Field(i).Kind() == reflect.Ptr {
				resultValue[fieldName] = reflectValue.Field(i).Elem().Interface()
			}
		}
	}
	enhancelog.CtxInfo(ctx, "exit: result: %v", resultValue)
	return resultValue
}

func GetStructFromMap(ctx context.Context, originValue map[string]interface{}, targetValue interface{}, namePattern int) interface{} {
	targetType := reflect.TypeOf(targetValue).Elem()
	reflectTarget := reflect.ValueOf(targetValue).Elem()

	if originValue == nil {
		enhancelog.CtxWarn(ctx, "[GetStructFromMap] originValue: %+v is invalid", originValue)
		return targetValue
	}
	for i := 0; i < targetType.NumField(); i++ {
		fieldName := targetType.Field(i).Name
		if namePattern == UNDERLINE {
			fieldName = Camel2Case(fieldName)
		}
		if value, ok := originValue[fieldName]; ok {
			setReflectValue(reflect.ValueOf(value), reflectTarget.Field(i), false)
		}
	}
	return targetValue
}

func setReflectValue(originValue, targetField reflect.Value, ignoreZero bool) {
	if !originValue.IsValid() {
		return
	}
	if ignoreZero && originValue.IsZero() {
		return
	}
	if !targetField.IsValid() {
		return
	}

	var numValue int64
	if originValue.Kind() == reflect.Ptr {
		if originValue.IsZero() {
			numValue = GetNumberValue(reflect.Zero(originValue.Type()).Interface())
		} else {
			numValue = GetNumberValue(originValue.Elem().Interface())
		}
	} else {
		numValue = GetNumberValue(originValue.Interface())
	}

	if numValue != MinNumber {
		if targetField.Kind() == reflect.Ptr {
			setPtrInt(targetField, numValue)
		} else {
			targetField.SetInt(numValue)
		}
	} else {
		setValueForKind(originValue, targetField)
	}
}

func setValueForKind(originValue, targetField reflect.Value) {
	switch {
	case targetField.Kind() == reflect.Ptr && originValue.Kind() != reflect.Ptr:
		if originValue.CanAddr() {
			targetField.Set(originValue.Addr())
		} else {
			newValue := reflect.New(originValue.Type()).Elem()
			newValue.Set(originValue)
			targetField.Set(newValue.Addr())
		}
	case targetField.Kind() != reflect.Ptr && originValue.Kind() == reflect.Ptr:
		targetField.Set(originValue.Elem())
	case targetField.Kind() == reflect.String && (originValue.Kind() == reflect.Struct || isArr(originValue.Kind())):
		if !originValue.IsNil() && originValue.IsValid() {
			jsonStr := getJSONFromInterface(originValue.Interface())
			targetField.SetString(jsonStr)
		}
	case isArr(targetField.Kind()) && originValue.Kind() == reflect.String:
		z := reflect.MakeSlice(targetField.Type(), 0, 0)
		targetField.Set(z)
		itemAddr := targetField.Addr().Interface()
		_ = json.Unmarshal([]byte(originValue.String()), itemAddr)
	case targetField.Kind() == reflect.Struct && originValue.Kind() == reflect.String:
		newOne := reflect.New(targetField.Type())
		targetField.Set(newOne)
		itemAddr := targetField.Addr().Interface()
		_ = json.Unmarshal([]byte(originValue.String()), itemAddr)
	case isArr(targetField.Kind()) && isArr(originValue.Kind()):
		if !originValue.IsNil() && originValue.IsValid() {
			jsonStr := getJSONFromInterface(originValue.Interface())
			newOne := reflect.MakeSlice(targetField.Type(), 0, 0)
			targetField.Set(newOne)
			itemAddr := targetField.Addr().Interface()
			_ = json.Unmarshal([]byte(jsonStr), itemAddr)
		}
	default:
		targetField.Set(originValue)
	}
}

// isArr 判断这个是不是数组类型 包括 arr 和slice
func isArr(kind reflect.Kind) bool {
	return kind == reflect.Slice || kind == reflect.Array
}

func getJSONFromInterface(i interface{}) string {
	if i == nil {
		return ""
	}
	strByte, err := json.Marshal(i)
	if err != nil {
		return ""
	}
	return string(strByte)
}

func setPtrInt(v reflect.Value, num int64) {
	switch k := v.Type().String(); k {
	default:
		panic(&reflect.ValueError{Method: "reflect.Value.SetInt", Kind: v.Kind()})
	case "*int":
		value := int(num)
		v.Set(reflect.ValueOf(&value))
	case "*int8":
		value := int8(num)
		v.Set(reflect.ValueOf(&value))
	case "*int16":
		value := int16(num)
		v.Set(reflect.ValueOf(&value))
	case "*int32":
		value := int32(num)
		v.Set(reflect.ValueOf(&value))
	case "*int64":
		value := num
		v.Set(reflect.ValueOf(&value))
	}
}

func GetLetterByNumber(num int) string {
	return []string{
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N",
		"O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
	}[num]
}
