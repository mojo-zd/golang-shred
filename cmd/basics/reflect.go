package basics

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
)

func makeFunc(fnptr interface{}) {
	fn := reflect.ValueOf(fnptr).Elem()
	v := reflect.MakeFunc(fn.Type(), func(args []reflect.Value) (results []reflect.Value) {
		return []reflect.Value{args[1], args[0]}
	})
	v.Addr()
	fn.Set(v)
}

func StructOf() {
	value := reflect.New(reflect.StructOf([]reflect.StructField{
		{
			Name: "Height",
			Type: reflect.TypeOf(float64(0)),
			Tag:  "height",
		},
		{
			Name: "Width",
			Type: reflect.TypeOf(int(0)),
			Tag:  "width",
		},
	})).Elem()
	value.Field(0).SetFloat(2.3)
	value.Field(1).SetInt(5)
	v := value.Addr().Interface()
	w := new(bytes.Buffer)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		panic(err)
	}
	fmt.Printf("value: %+v\n", v)
	fmt.Printf("json: %s\n", w.Bytes())
}
