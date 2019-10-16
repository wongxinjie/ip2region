/**
* @File: worker.go
* @Author: wongxinjie
* @Date: 2019/10/17
 */
package task

import (
	"fmt"
	"os"
)

var (
	MaxWorker = os.Getenv("MAX_WORKER")
	MaxQueue  = os.Getenv("MAX_QUEUE")
)

type Job struct {
	Payload Payload
}

var JobQueue chan Job

type Worker struct {
	WorkerPool chan chan Job
	JobChannel chan Job
	quit       chan bool
}

func NewWorker(workerPool chan chan Job) Worker {
	return Worker{
		WorkerPool: workerPool,
		JobChannel: make(chan Job),
		quit:       make(chan bool),
	}
}

func (w Worker) Start() {
	go func() {
		for {
			w.WorkerPool <- w.JobChannel

			select {
			case job := <-w.JobChannel:
				// process job here
				fmt.Printf("%v\n", job)
			case <-w.quit:
				return
			}
		}
	}()
}

func (w Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}
