/**
* @File: dispatcher.go
* @Author: wongxinjie
* @Date: 2019/10/17
 */
package task

type Dispatcher struct {
	maxWorkers int
	pool       chan chan Job
}

func NewDispatcher(maxWorkers int) *Dispatcher {
	pool := make(chan chan Job, maxWorkers)
	return &Dispatcher{
		maxWorkers: maxWorkers,
		pool:       pool,
	}
}

func (d *Dispatcher) Run() {
	for i := 0; i < d.maxWorkers; i++ {
		worker := NewWorker(d.pool)
		worker.Start()
	}

	go d.dispatch()
}

func (d *Dispatcher) dispatch() {
	for {
		select {
		case job := <-JobQueue:
			go func(job Job) {
				jobChannel := <-d.pool

				jobChannel <- job

			}(job)

		}
	}
}
