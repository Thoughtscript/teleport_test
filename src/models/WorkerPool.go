package models

import (
	"bytes"
	"fmt"
	"os/exec"
	"sync"
	"time"
)

type WorkerPool struct {
	Workers  map[string]WorkerModel
	Queue 	 map[string]time.Time
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
			workerPool.Queue = make(map[string]time.Time)
			workerPool.Statuses = make(map[string]string)
			workerPool.Workers = make(map[string]WorkerModel)
		}
	}
	return workerPool
}

// ----------------------------
// Worker-Specific Helpers
// ----------------------------

// Receiver-type function
func (w WorkerModel) ExecuteCommand() {
	w.Status = "executing"
	GetWorkerPool().Statuses[w.Uuid] = w.Status
	fmt.Println(w.Uuid, w.Status, w.Command)
	outputChannel := make(chan string)

	go func() {
		// Hardcoded for now - need client-injection protection
		// Possible real-world cases for this include common bash commands
		// Or, more likely, a pre-defined set of operations specific to the API or tool - like '$ teleport_test hello -> world'

		cmd := exec.Command("ls")
		var out bytes.Buffer
		cmd.Stdout = &out

		err := cmd.Run()
		if err != nil {
			w.Status = "failed"
			GetWorkerPool().Statuses[w.Uuid] = w.Status
			outputChannel <- w.Uuid + " " + w.Status + " " + w.Command
			RemoveWorker(w.Uuid)
		}

		w.Status = "completed"
		GetWorkerPool().Statuses[w.Uuid] = w.Status
		outputChannel <- w.Uuid + " " + w.Status + " " + w.Command + " " + out.String()
		RemoveWorker(w.Uuid)
	}()

	result := <-outputChannel
	fmt.Println(result)
}

// ----------------------------
// Worker-Queue Helpers
// ----------------------------

func AddWorker(worker WorkerModel) {
	GetWorkerPool().Queue[worker.Uuid] = worker.Time
	GetWorkerPool().Workers[worker.Uuid] = worker
	GetWorkerPool().Statuses[worker.Uuid] = worker.Status
}

func RemoveWorker(uuid string) {
	delete(GetWorkerPool().Workers, uuid)
	delete(GetWorkerPool().Queue, uuid)
}

func StopWorker(uuid string) {
	RemoveWorker(uuid)
	GetWorkerPool().Statuses[uuid] = "stopped"
}

func ProcessQueue(wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		var second time.Duration = 1000000000
		var seconds time.Duration = 5

		time.Sleep(seconds * second)
		fmt.Println("Polling: ", seconds*second)
		c := time.Now()

		for k, v := range GetWorkerPool().Queue {
			r := v.Equal(c) || v.After(c)

			// Execute only if hasn't run yet
			// Workers can have queued, executing, failed, completed
			// The latter two will be removed from queue and kept in status map
			w := GetWorkerPool().Workers[k]
			if r && w.Status == "queued" {
				w.ExecuteCommand()
			}
		}

		fmt.Println(GetWorkerPool().Queue)
		fmt.Println(GetWorkerPool().Workers)
		fmt.Println(GetWorkerPool().Statuses)
	}
}

// Query to get status/history of worker
func GetJob(uuid string) string {
	return GetWorkerPool().Statuses[uuid]
}

// Query to get all worker histories
func GetAllJobs() map[string]string {
	return GetWorkerPool().Statuses
}
