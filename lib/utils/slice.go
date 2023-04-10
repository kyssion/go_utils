package util

import "reflect"

func ContainsString(slice []string, element string) bool {
	for i := 0; i < len(slice); i++ {
		if slice[i] == element {
			return true
		}
	}

	return false
}

func SliceSubString(a []string, b []string) []string {
	m := make(map[string]struct{}, len(b))
	for _, v := range b {
		m[v] = struct{}{}
	}

	var ret []string
	for _, v := range a {
		_, exists := m[v]
		if !exists {
			ret = append(ret, v)
		}
	}
	return ret
}

func ContainsNumber(slice interface{}, element interface{}) bool {
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice && v.Kind() != reflect.Array {
		return false
	}
	numberElement := GetNumberValue(element)

	for i := 0; i < v.Len(); i++ {
		if GetNumberValue(v.Index(i).Interface()) == numberElement {
			return true
		}
	}

	return false
}

func ContainsFloat(slice interface{}, element float64) bool {
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice && v.Kind() != reflect.Array {
		return false
	}

	for i := 0; i < v.Len(); i++ {
		if v.Index(i).Interface() == element {
			return true
		}
	}

	return false
}

func SliceSubInt64(a []int64, b []int64) []int64 {
	m := make(map[int64]struct{}, len(b))
	for _, v := range b {
		m[v] = struct{}{}
	}

	var ret []int64
	for _, v := range a {
		_, exists := m[v]
		if !exists {
			ret = append(ret, v)
		}
	}
	return ret
}

func SliceInt64ToStrings(slices []int64) []string {
	newStrs := make([]string, 0, len(slices))
	for _, value := range slices {
		newStrs = append(newStrs, Int64ToString(value))
	}
	return newStrs
}
