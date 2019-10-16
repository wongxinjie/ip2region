/**
* @File: main.go
* @Author: wongxinjie
* @Date: 2019/10/17
 */
package main

import (
	"context"
	"fmt"
	"time"

	"ip2region/task"
)

func producer(id int64) {
	work := task.Job{
		Payload: task.Payload{
			Id:    id,
			Token: "token",
			Data:  fmt.Sprintf("Payload-Data-%d", id),
		},
	}
	fmt.Printf("produce: %v\n", id)
	task.JobQueue <- work
}

func consumer(worker int) {
	dispatcher := task.NewDispatcher(worker)
	dispatcher.Run()
}

// 使用Golang 处理每分钟100万份请求
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	go consumer(5)

	var id int64
	select {
	case <-time.After(1 * time.Second):
		{
			id += 1
			producer(id)
		}
	case <-ctx.Done():
		break
	}
}
