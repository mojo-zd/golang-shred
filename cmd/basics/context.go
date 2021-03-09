package basics

import (
	"context"
	"fmt"
	"time"
)

func TimeoutContext(ctx context.Context, timeout time.Duration) {
	con, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	go task1(con)
	go task2(con)
	select {
	case <-con.Done():
		fmt.Println("main", con.Err())
	}
}

func task1(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("cancel task1", ctx.Err())
	case <-time.After(time.Second):
		fmt.Println("task1 after 1 second")
	}
}

func task2(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("cancel task2", ctx.Err())
	case <-time.After(time.Second):
		fmt.Println("task2 after 1 second")
	}
}
