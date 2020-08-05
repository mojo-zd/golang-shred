package main

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"reflect"
)

type Padding struct {
	Name string `json:"name"`
}

func main() {
	paddings := []*Padding{}
	PaddingSlice(&paddings, Padding{}, 3)
	log.Info().Interface("paddings", paddings).Send()
	log.Info().Interface("newInstance", newInstance(Padding{})).Send()
}

// 思考: 给一个任意的slice 如何向这个slice中添加该类型的实例
// target 目标数组
// elem 数组元素类型
// num 为target添加num个实例
func PaddingSlice(target, elem interface{}, num int) {
	ty := reflect.TypeOf(target)
	if ty.Kind() == reflect.Ptr {
		ty = ty.Elem()
	}

	if ty.Kind() != reflect.Slice {
		log.Warn().Msg("sorry, only support slice")
		return
	}
	slice := reflect.Indirect(reflect.ValueOf(target))
	slice.Set(reflect.MakeSlice(slice.Type(), 0, 0))

	var elemTy reflect.Type
	if elemTy = reflect.TypeOf(elem); elemTy.Kind() == reflect.Ptr {
		elemTy = elemTy.Elem()
	}

	for i := 0; i < num; i++ {
		v := reflect.New(elemTy).Elem()
		for j := 0; j < elemTy.NumField(); j++ {
			switch elemTy.Field(j).Type.Name() {
			case "string":
				v.Field(j).SetString(fmt.Sprintf("rand%d", i))
			}
		}
		slice.Set(reflect.Append(slice, v.Addr()))
	}
}

// 解析elem struct字段 并实例化新的对象
func newInstance(elem interface{}) interface{} {
	var ty reflect.Type
	fields := []reflect.StructField{}
	if ty = reflect.TypeOf(elem); ty.Kind() == reflect.Ptr {
		ty = ty.Elem()
	}
	for i := 0; i < ty.NumField(); i++ {
		fieldName := ty.Field(i).Name
		fieldTyp := ty.Field(i).Type
		fields = append(fields, reflect.StructField{
			Name: fieldName,
			Type: fieldTyp,
		})
		log.Info().Str("field name", fieldName).Interface("type", fieldTyp.Name()).Send()
	}

	typ := reflect.StructOf(fields)
	v := reflect.New(typ).Elem()
	v.Field(0).SetString("mojo")
	return v.Interface()
}
