package main

// https://golangbyexample.com/singleton-design-pattern-go/

import (
	m "./models"
	"fmt"
	"sync"
	"time"
)

// Define one mutex
var lock = &sync.Mutex{}
// Define one top-level variable
var workerPool *m.WorkerPool

func WorkerPool() *m.WorkerPool {
	if workerPool == nil {
		lock.Lock()
		defer lock.Unlock()
		if workerPool == nil {
			fmt.Println("Creting Single Instance Now")
			workerPool = &m.WorkerPool{}
		} else {
			fmt.Println("Single Instance already created-1")
		}
	} else {
		fmt.Println("Single Instance already created-2")
	}
	return workerPool
}

func AddWorker(worker m.WorkerModel) {
	WorkerPool().Workers = append(WorkerPool().Workers, worker)
}

func processQueue() {
	for {
		c := time.Now()
		w := WorkerPool().Workers
		for i := 0; i < len(w); i++ {
			t := w[i].Time
			if t.Equal(c) || t.After(c) {

			}
		}
	}
}
