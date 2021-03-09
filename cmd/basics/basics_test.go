package basics

import (
	"context"
	"encoding/json"
	"reflect"
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
	t.Log("increaseB", increaseB())
	t.Log("increaseC", increaseC())
	t.Log("increaseD", increaseD())
	a := []int{1, 2, 3}
	t.Log("cap:", cap(a))
	var std = Student{}
	std.Speak("")
}

type People interface {
	Speak(talk string) string
}

type Student struct {
}

func (s *Student) Speak(talk string) string {
	return "hi"
}

func increaseA() int {
	var i int
	defer func() {
		i++
	}()
	return i
}

func increaseB() (r int) {
	defer func() {
		r++
	}()
	return 0
}

func increaseC() (r int) {
	t := 5
	defer func() {
		t += 5
	}()
	return t
}

func increaseD() (r int) {
	defer func(r int) {
		r += 5
	}(r)
	return 1
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

func TestSelect(t *testing.T) {
	ch := make(chan int)
	close(ch)
	t.Log("running...", <-ch)
}

func TestChain(t *testing.T) {
	//one := 0
	//one := 0
	var s []int
	s = append(s, 1)
	ch := make(chan int)
	close(ch)
	if _, ok := <-ch; ok {
		close(ch)
	}

	//t.Log("start write...")
	//ch <- 4
	//ch <- 4
	t.Log("only write...")
}

func TestInterfaceEmbed(t *testing.T) {
	interfaceEmbed()
}

func TestDefer(t *testing.T) {
	deferKnown(&Person{age: 28})
}

func TestIota(t *testing.T) {
	//aa := []string{"kk", "mm", "zz"}
	//for i, v := range aa {
	//	go func() {
	//		t.Log("i:", i, ",v:", v)
	//	}()
	//}

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

func TestValPoint(t *testing.T) {
	d := &Dog{}
	sendNotify(d)
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

func TestTimeoutContext(t *testing.T) {
	TimeoutContext(context.Background(), time.Millisecond*100)
}

func TestUpdateValue(t *testing.T) {
	var x = 4.0
	v := reflect.ValueOf(&x)
	v.Elem().SetFloat(5.3)
	t.Log("can address", v.CanAddr())
}

func TestTicker(t *testing.T) {
	c := time.NewTicker(time.Second)
	for range c.C {
		t.Log(".")
	}
}
