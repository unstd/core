package core

import (
	"fmt"
	"reflect"
	"strings"
)

// Contains check whether the target in collect, type requirements:
// when collect type is string, target mast be string
// when collect type is []T, target type mast be T
// when collect type is [n]T, target type mast be T
// when collect type is map[T]R, target type mast be T
//
// special:
// (nil, interface{})                => (false, -1)
// (map[string]string{"a": "a"}, "a) => (true, 0)
func Contains(collect interface{}, target interface{}) (bool, int) {
	collValue := reflect.ValueOf(collect)
	if !collValue.IsValid() {
		return false, -1
	}
	switch reflect.TypeOf(collect).Kind() {
	case reflect.Array, reflect.Slice:
			for i := 0; i < collValue.Len(); i++ {
				if target == collValue.Index(i).Interface() {
					return true, i
				}
			}
	case reflect.Map:
		mapIndex := collValue.MapIndex(reflect.ValueOf(target))
		return mapIndex.IsValid(), 0
	case reflect.String:
		switch reflect.TypeOf(target).Kind() {
		case reflect.String:
			index := strings.Index(collect.(string), target.(string))
			return index >= 0, index
		default:
			panic(fmt.Sprintf("unsupport target type of %s", reflect.TypeOf(target).Kind()))
		}
	default:
		panic(fmt.Sprintf("unsupport collect type of %s", reflect.TypeOf(target).Kind()))
	}
	return false, -1
}

// AllEqual check every element in the collect is equal to target
// when collect is nil or empty will return false
func AllEqual(collect interface{}, target interface{}) bool {
	collValue := reflect.ValueOf(collect)
	if !collValue.IsValid() {
		return false
	}
	switch reflect.TypeOf(collect).Kind() {
	case reflect.Array, reflect.Slice:
			for i := 0; i < collValue.Len(); i++ {
				if target == collValue.Index(i).Interface() {
					return true
				}
			}
	default:
		panic(fmt.Sprintf("unsupport collect type of %s", reflect.TypeOf(target).Kind()))
	}
	return false
}

// IsEmpty check whether the following types are empty
// string array slice nil
func IsEmpty(target interface{}) bool {
	value := reflect.ValueOf(target)
	if !value.IsValid() {
		return true
	}
	return value.IsZero() || value.Len() == 0
}

func IsNotEmpty(target interface{}) bool {
	return !IsEmpty(target)
}
