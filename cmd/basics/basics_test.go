package basics

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/rs/zerolog/log"
)

func TestWidget(t *testing.T) {
	log.Info().Interface("id", NewWidget().ID()).Send()
}

func TestDeferCall(t *testing.T) {
	//deferCall()
	deferRecover()
}

func TestDeferAn(t *testing.T) {
	t.Log("increaseA", increaseA())
}

func increaseA() int {
	var i int
	defer func() {
		i++
	}()
	return i
}

func TestRange(t *testing.T) {
	var slice = []int{1, 6, 3, 8, 2}
	log.Info().Interface("map", Range(slice)).Send()
	RangAppend(slice)
}

func TestRangeLoop(t *testing.T) {
	RangeYy([]int{1})
}

func TestPointer(t *testing.T) {
	pointer()
}

func TestNew(t *testing.T) {
	gNew()
	gMake()
}

func TestSlice(t *testing.T) {
	gSlice()
	gSliceCut()
	sliceIsAttr()
}

func TestSliceAppend(t *testing.T) {
	gSliceAppend()
}

func TestSliceCopy(t *testing.T) {
	gSliceCopy()
}

func TestSliceCite(t *testing.T) {
	ss := []S{{Name: "mojo"}}
	gSliceCite(ss)
	log.Info().Interface("o", ss).Send()
}

func TestMap(t *testing.T) {
	m := map[string]string{"key": "mojo"}
	gMap(m)
	log.Info().Interface("m", m).Send()
}

func TestInterface(t *testing.T) {
	interfaceAssert()
}

func TestSelectorType(t *testing.T) {
	switchType(Circle{})
}

func TestInterfaceEmbed(t *testing.T) {
	interfaceEmbed()
}

func TestDefer(t *testing.T) {
	deferKnown(&Person{age: 28})
}

func TestIota(t *testing.T) {
	iotaDefined()
	iotaMulit()
	direction()
}

func TestChanSelect(t *testing.T) {
	selectChan()
	time.Sleep(time.Minute)
}

func TestChanBuffer(t *testing.T) {
	chanBuffer()
}

func TestChanClose(t *testing.T) {
	chanClose()
}

func TestCloseFunc(t *testing.T) {
	x := 1
	s := closeFunc(x)
	x = 2
	s()
	log.Info().Interface("out", s()).Interface("ss", closeInnerFunc()()).Send()
}

func TestCLose(t *testing.T) {
	bar := myFunc()
	val1 := bar()
	val2 := bar()
	f := extFunc(20)
	f()
	f()
	log.Info().Int("value 1", val1).Int("value 2", val2).Send()
}

func TestCloseFuncArr(t *testing.T) {
	ss := closeFuncArr([]int{1, 2, 3})
	for _, s := range ss {
		s()
	}
}

type User struct {
	IDs []int
}

func TestS(t *testing.T) {
	user := User{}
	user.IDs = append(user.IDs, 1)
	user.IDs = append(user.IDs, 2)
	user.IDs = append(user.IDs, 3)
	a := append(user.IDs, 4)
	log.Info().Interface("append 4 cap", cap(user.IDs)).Interface("slice", user.IDs).Interface("a", a).Send()
	b := append(user.IDs, 5)
	log.Info().Interface("append 5 cap", cap(user.IDs)).Interface("slice", user.IDs).Interface("b", b).Send()
	c := append(user.IDs, 6)
	log.Info().Interface("append 6 cap", cap(user.IDs)).Interface("slice", user.IDs).Interface("c", c).Send()
	d := append(user.IDs, 7)
	log.Info().Interface("append 7 cap", cap(user.IDs)).Interface("slice", user.IDs).Interface("d", d).Send()
	e := append(user.IDs, 8)
	log.Info().Interface("append 8 cap", cap(user.IDs)).Interface("slice", user.IDs).Interface("e", e).Send()
}

func TestM(t *testing.T) {
	s := `{"name": "mojo"}`
	m := make(map[string]string)
	json.Unmarshal([]byte(s), m)
	log.Info().Interface("info", m).Send()
}

func TestEqual(t *testing.T) {
	s1 := EqStruct{Name: "mojo", Sub: &SubStruct{Day: "monday"}}
	s2 := EqStruct{Name: "mojo", Sub: &SubStruct{Day: "monday"}}
	t.Log("is it equal?", StructEqual(s1, s2))
}
