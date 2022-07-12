package collections

import (
	"fmt"
	"reflect"
)

func Contains(collection interface{}, target interface{}) bool {
	targetValue := reflect.ValueOf(collection)
	switch reflect.TypeOf(collection).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			fmt.Println(targetValue.Index(i).Interface(), "==?", target)
			if targetValue.Index(i).Interface() == target {
				return true
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(target)).IsValid() {
			return true
		}
	}
	return false
}
