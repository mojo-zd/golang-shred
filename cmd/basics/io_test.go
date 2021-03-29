package basics

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"
	"unicode"
	"unsafe"
)

func TestIoReader(t *testing.T) {
	out, err := ReadFrom(os.Stdin, 10)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(out))
	time.Sleep(time.Minute)
}

type AA struct {
	Name string
}

func TestSeekEnd(t *testing.T) {
	aa := AA{Name: "mojo"}
	fmt.Printf("%+v\n", aa)
	fmt.Printf("%#v\n", aa)
	fmt.Printf("%t\n", false)
	fmt.Printf("%c\n", 11111111)
	strings.FieldsFunc("", unicode.IsSpace)
	//unicode.Is(unicode.Han) 判断汉字
	buff := bytes.NewBufferString("")
	buff.WriteString("sep")
	fmt.Println("mapping:", strings.Map(func(r rune) rune {
		if r != 'r' {
			return r
		}
		return 'R'
	}, "hello world"))
	r := strings.NewReplacer("<", "&lt;", ">", "&gt;")
	fmt.Println("replace:", r.Replace("This is <b>HTML</b>!"))
	b := strings.Builder{}
	b.WriteString("hello")
	b.WriteString(" world")
	fmt.Println("builder:", b.String())
	var i1 int = 1
	var i2 int8 = 2
	fmt.Println("i1 size:", unsafe.Sizeof(i1), ", i2 size:", unsafe.Sizeof(i2)) // out put bit of variable
	fmt.Println(strconv.Quote("go programe"))

	fmt.Println("format:", time.Now().Format("2006-01-02 15:04:05"),
		",time.round:", time.Now().Round(time.Minute), ", time.truncate:", time.Now().Truncate(time.Minute))
	out, err := strconv.ParseInt("128", 10, 8)
	if err != nil {
		t.Fatal("parse int:", err.Error(), ",out:", out)
	}
	fmt.Println("128 to int", out)
	//ll := list.New()
	//rr := ring.New(10)
	//el := list.Element{Value: ""}
	//sort.Search()
}

func TestReadSeek(t *testing.T) {
	fmt.Println("base:",
		filepath.Dir("aa/bb/cc/a.txt"),
		",ext:", filepath.Ext("aa/bb/cc/a.txt"))
	abs, _ := filepath.Abs("tester")
	fmt.Println("abs path:", abs)
	dir, file := filepath.Split("aa/bb/cc/dd")
	fmt.Println("dir:", dir, ",file:", file)
	f, err := os.Open("./README.md")
	if err != nil {
		t.Fatal("open file fault:", err.Error())
	}

	f.Seek(-1, io.SeekEnd)
	bytes := make([]byte, 1000)
	_, err = f.Read(bytes)
	if err != nil {
		t.Fatal("read fault", err.Error())
	}
	fmt.Println("out:", string(bytes))
	//pool := sync.Pool{New: func() interface{} {
	//	return nil
	//}}
	//pool.Put()
}

func TestCond(t *testing.T) {
	cond := sync.NewCond(new(sync.Mutex))
	cond.L.Lock()
	for i := 0; i < 10; i++ {
		go func(i int) {
			//cond.L.Lock()
			fmt.Println("sequence", i)
			//cond.L.Unlock()
			if i == 0 {
				//return
				cond.Signal()
			}
		}(i)
	}
	cond.Wait()
	t.Log("out print...")
}

func TestBrocast(t *testing.T) {
	cancel := make(chan bool, 1)
	go work(cancel, 1)
	go work(cancel, 2)

	time.Sleep(time.Second * 4)
	//cancel <- true
	close(cancel)
	time.Sleep(time.Second * 5)
}

func work(cancel chan bool, pid int) {
	for {
		select {
		case <-cancel:
			fmt.Println("exist", pid)
			return
		default:
			fmt.Printf("working[%d]...\n", pid)
			time.Sleep(time.Second)
		}
	}
}
