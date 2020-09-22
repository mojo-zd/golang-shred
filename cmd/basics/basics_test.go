package basics

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"testing"
	"time"
)

func TestWidget(t *testing.T) {
	log.Info().Interface("id", NewWidget().ID()).Send()
}

func TestDeferCall(t *testing.T) {
	deferCall()
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

func TestFunc(t *testing.T) {
	nil := 123
	fmt.Println(nil)
	var _ map[string]int = nil
}
