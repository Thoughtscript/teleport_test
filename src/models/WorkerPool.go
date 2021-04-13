package models

import (
	"fmt"
	"sync"
	"time"
)

type WorkerPool struct {
	Workers  []WorkerModel
	Statuses map[string]string
}

// ----------------------------
// Singleton Pool
// ----------------------------
// https://golangbyexample.com/singleton-design-pattern-go/

// Define one mutex
var lock = &sync.Mutex{}

// Define one top-level variable
var workerPool *WorkerPool

func GetWorkerPool() *WorkerPool {
	if workerPool == nil {
		lock.Lock()
		defer lock.Unlock()
		if workerPool == nil {
			fmt.Println("Creating singleton Worker Pool!")
			workerPool = &WorkerPool{}
			workerPool.Statuses = make(map[string]string)
		}
	}
	return workerPool
}

// ----------------------------
// Helpers
// ----------------------------

func AddWorker(worker WorkerModel) {
	GetWorkerPool().Workers = append(GetWorkerPool().Workers, worker)
	GetWorkerPool().Statuses[worker.Uuid] = worker.Status
}

func ProcessQueue() {
	for {
		var second time.Duration = 1000000000
		var seconds time.Duration =  5

		time.Sleep(seconds*second)
		fmt.Println("Polling: ", seconds*second)
		c := time.Now()
		w := GetWorkerPool().Workers
		for i := 0; i < len(w); i++ {
			t := w[i].Time
			r := t.Equal(c) || t.After(c)

			// Execute only if hasn't run yet
			if r && w[i].Status == "queued" {
				w[i].ExecuteCommand()
			}
		}
	}
}

func GetJob(uuid string) string {
	return GetWorkerPool().Statuses[uuid]
}
