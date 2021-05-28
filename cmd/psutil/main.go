package main

import (
	"flag"
	"fmt"

	"github.com/shirou/gopsutil/process"
)

func main() {
	var pid int
	flag.IntVar(&pid, "pid", 0, "process pid")
	flag.Parse()
	proc, err := process.NewProcess(int32(pid))
	if err != nil {
		fmt.Println("can't new process", err.Error())
		return
	}
	percent, err := proc.CPUPercent()
	if err != nil {
		fmt.Println("can't get cpu percent", err)
		return
	}
	name, _ := proc.Name()
	//mem, _ := proc.MemoryInfo()
	fmt.Println("pid", pid, ",percent is", percent, "process name", name)
}
