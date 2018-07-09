package funcs

import (
	"reflect"
)

//DeleteInts 删除[]int切片的某个键 并更新键值
func DeleteInts(s []int, key int) []int {
	if key == len(s)-1 {
		return s[:key]
	} else {
		return append(s[:key], s[key+1:]...)
	}
}

//DeleteStrings 删除[]string切片的某个键 并更新键值
func DeleteStrings(s []string, key int) []string {
	if key == len(s)-1 {
		return s[:key]
	} else {
		return append(s[:key], s[key+1:]...)
	}
}

//DeleteSlice 删除slice的某个键
func DeleteSlice(s interface{}, key int) interface{} {
	types := reflect.TypeOf(s)
	if types.Kind() == reflect.Slice {
		val := reflect.ValueOf(s)
		if key == val.Len()-1 {
			return val.Slice(0, key).Interface()
		} else {
			return reflect.AppendSlice(val.Slice(0, key), val.Slice(key+1, val.Len())).Interface()
		}
	}
	return nil
}
