package core

import (
	"fmt"
	"reflect"
	"strings"
)

func Contains(collect interface{}, target interface{}) (bool, int) {
	switch reflect.TypeOf(collect).Kind() {
	case reflect.Array, reflect.Slice:
			collValue := reflect.ValueOf(collect)
			for i := 0; i < collValue.Len(); i++ {
				if target == collValue.Index(i).Interface() {
					return true, i
				}
			}
	case reflect.Map:
		mapIndex := reflect.ValueOf(collect).MapIndex(reflect.ValueOf(target))
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

func en() {
	a := []string{"3"}
	Contains(a, "")
}
