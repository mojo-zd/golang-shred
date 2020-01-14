package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	bytes := str2Byte("test")
	fmt.Println("bytes",bytes)
	fmt.Println("string", byte2Str(bytes))
}

func byte2Str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func str2Byte(s string) []byte {
	sHeader := (*reflect.StringHeader)(unsafe.Pointer(&s))
	sliceHeader := reflect.SliceHeader{
		Data: sHeader.Data,
		Len:  sHeader.Len,
		Cap:  sHeader.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&sliceHeader))
}
