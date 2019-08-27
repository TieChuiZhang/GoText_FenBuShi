//File  : channellee.go
//Author: 燕人Lee&骚气又迷人的反派
//Date  : 2019-08-21

package main

import (
	"fmt"
	"sync"
)

func chanDemo() {
	var wg sync.WaitGroup

	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}
	//wg.Add(20)
	for i, worker := range workers {
		worker.in <- 'a' + i
		wg.Add(1)
	}
	for i, worker := range workers {
		worker.in <- 'A' + i
		wg.Add(1)
	}
	wg.Wait()
}

func doWork(id int, w worker) {
	for n := range w.in {
		fmt.Printf("Worker %d receiver %d \n", id, n)
		// go func() {}()
		w.done()
	}
}

type worker struct {
	in   chan int
	done func()
}

func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		make(chan int),
		func() {
			wg.Done()
		},
	}
	go doWork(id, w)
	return w
}

func main() {
	chanDemo()

}
