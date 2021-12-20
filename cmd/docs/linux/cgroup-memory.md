`memory`子系统

下面编写一个消耗内存的程序
`malloc.cpp`
```
#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <iostream>
#include <math.h>

using namespace std;
#define BLOCK_SIZE (100*1024*1024)

extern "C" {
    char* allocMemory() {
        char* out = (char*)malloc(BLOCK_SIZE);
        memset(out, 'A', BLOCK_SIZE);
        return out;
    }
}
```

`main.go`
```go
package main

//#cgo LDFLAGS:
//char* allocMemory();
import "C"
import (
	"fmt"
	"time"
)

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("Allocating %dMb memory, raw memory is %d\n")
		C.allocMemory()
		time.Sleep(time.Minute * 5)
	}
}

```

`Makefile`, 需要在linux下执行
```
build:
	CGO_ENABLED=1 GOOS=linux CGO_LDFLAGS="-static" go build -o malloc
```

通过`cgroup`控制cpu一样, 参考 <a href="./cgroup-cpu.md" title="标题">cgroup cpu</a>