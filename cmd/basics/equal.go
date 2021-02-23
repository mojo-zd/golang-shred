package basics

import "reflect"

type EqStruct struct {
	Name string
	Sub  *SubStruct
}

type SubStruct struct {
	Day string
}

func StructEqual(s1, s2 EqStruct) bool {
	return reflect.DeepEqual(s1, s2)
}
