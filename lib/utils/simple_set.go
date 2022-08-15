package util

import (
	"reflect"
	"sync"
)

type SimpleSet map[interface{}]struct{}

func NewSimpleSet() SimpleSet {
	return make(SimpleSet)
}

func NewSetWithArray(values interface{}) SimpleSet {
	set := make(SimpleSet)
	set.AddList(values)
	return set
}

func (set SimpleSet) Add(i interface{}) bool {
	_, ok := set[i]
	if ok {
		// False if it existed already
		return false
	}

	set[i] = struct{}{}
	return true
}

func (set SimpleSet) AddList(values interface{}) {
	for i := 0; i < reflect.ValueOf(values).Len(); i++ {
		val := reflect.ValueOf(values).Index(i).Interface()
		set[val] = struct{}{}
	}
}

func (set SimpleSet) Contains(i ...interface{}) bool {
	for _, val := range i {
		if _, ok := set[val]; !ok {
			return false
		}
	}
	return true
}

func (set SimpleSet) Remove(i interface{}) bool {
	if set.Contains(i) {
		delete(set, i)
		return true
	}
	return false
}

func (set SimpleSet) Cardinality() int {
	return len(set)
}

func (set SimpleSet) Values() []interface{} {
	keys := make([]interface{}, 0, set.Cardinality())
	for elem := range set {
		keys = append(keys, elem)
	}
	return keys
}

func (set SimpleSet) ToStrings() []string {
	keys := make([]string, 0, set.Cardinality())
	for elem := range set {
		keys = append(keys, elem.(string))
	}
	return keys
}

func (set SimpleSet) ToFloat64() []float64 {
	keys := make([]float64, 0, set.Cardinality())
	for elem := range set {
		keys = append(keys, elem.(float64))
	}
	return keys
}

func (set SimpleSet) ToInt64() []int64 {
	keys := make([]int64, 0, set.Cardinality())
	for elem := range set {
		num := GetNumberValue(elem)
		if num == MinNumber {
			return nil
		}
		keys = append(keys, num)
	}
	return keys
}

type SafeSet struct {
	safeMap sync.Map
}

func NewSafeSet() *SafeSet {
	return &SafeSet{}
}

func (set *SafeSet) Add(i interface{}) bool {
	if _, ok := set.safeMap.Load(i); ok {
		return false
	}
	set.safeMap.Store(i, struct{}{})
	return true
}

func (set *SafeSet) Contains(i ...interface{}) bool {
	for _, val := range i {
		if _, ok := set.safeMap.Load(val); !ok {
			return false
		}
	}
	return true
}

func (set *SafeSet) Remove(i interface{}) bool {
	if set.Contains(i) {
		set.safeMap.Delete(i)
		return true
	}
	return false
}

func (set *SafeSet) Cardinality() int {
	count := 0
	set.safeMap.Range(func(key, value interface{}) bool {
		count++
		return true
	})
	return count
}

func (set *SafeSet) ToInt64() []int64 {
	keys := make([]int64, 0)
	set.safeMap.Range(func(key, value interface{}) bool {
		num := GetNumberValue(key)
		if num == MinNumber {
			return false
		}
		keys = append(keys, num)
		return true
	})
	return keys
}
