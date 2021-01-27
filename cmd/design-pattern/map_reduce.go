package design_pattern

import "reflect"

func Map(data interface{}, fn interface{}) []interface{} {
	vfn := reflect.ValueOf(fn)
	vdata := reflect.ValueOf(data)
	result := make([]interface{}, vdata.Len())

	for i := 0; i < vdata.Len(); i++ {
		result[i] = vfn.Call([]reflect.Value{
			vdata.Index(i),
		})[0].Interface()
	}
	vfn.Type().NumIn()
	return result
}

func MapStrToStr(arr []string, fn func(s string) string) []string {
	var newArr []string
	for _, a := range arr {
		newArr = append(newArr, fn(a))
	}
	return newArr
}

func MapStrToInt(arr []string, fn func(s string) int) []int {
	var newArr []int
	for _, a := range arr {
		newArr = append(newArr, fn(a))
	}
	return newArr
}

func Filter(arr []int, fn func(s int) bool) []int {
	var newArr []int
	for _, a := range arr {
		if fn(a) {
			newArr = append(newArr, a)
		}
	}
	return newArr
}
