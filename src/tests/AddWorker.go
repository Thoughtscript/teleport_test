package tests

import (
	j "../jobs"
	m "../models"
	"fmt"
	"sync"
	"time"
)

// ----------------------------
// Wait Group
// ----------------------------

func TestAddWorkerLoop(wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		var second time.Duration = 1000000000
		var seconds time.Duration = 3
		time.Sleep(seconds * second)

		addWorkerAndPrint()
	}
}

func TestExecuteWorkerLoop(wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		var second time.Duration = 1000000000
		var seconds time.Duration = 3
		time.Sleep(seconds * second)

		addWorkerAndExecute()
	}
}

func AddWorkerAndStop(wg *sync.WaitGroup) {
	defer wg.Done()

	worker := j.NewWorker(time.Now(), "ls")
	j.AddWorker(worker)
	uid := worker.Uuid

	fmt.Println(m.StatusTable()[uid])
	fmt.Println(m.WorkerTable()[uid])
	fmt.Println(m.WorkerQueue()[uid])

	j.StopWorker(uid)
}

func AddWorkerAndExecuteWg(wg *sync.WaitGroup) {
	defer wg.Done()

	worker := j.NewWorker(time.Now(), "ls")
	j.AddWorker(worker)
	uid := worker.Uuid

	fmt.Println(m.StatusTable()[uid])
	fmt.Println(m.WorkerTable()[uid])
	fmt.Println(m.WorkerQueue()[uid])

	worker.ExecuteCommand()
}

func AddWorkerAndPrintWg(wg *sync.WaitGroup) {
	defer wg.Done()

	worker := j.NewWorker(time.Now(), "ls")
	j.AddWorker(worker)
	uid := worker.Uuid

	fmt.Println(m.StatusTable()[uid])
	fmt.Println(m.WorkerTable()[uid])
	fmt.Println(m.WorkerQueue()[uid])
}

// ----------------------------
// Private
// ----------------------------

func addWorkerAndExecute() {
	worker := j.NewWorker(time.Now(), "ls")
	j.AddWorker(worker)
	uid := worker.Uuid

	fmt.Println(m.StatusTable()[uid])
	fmt.Println(m.WorkerTable()[uid])
	fmt.Println(m.WorkerQueue()[uid])

	worker.ExecuteCommand()
}

func addWorkerAndPrint() {
	worker := j.NewWorker(time.Now(), "ls")
	j.AddWorker(worker)
	uid := worker.Uuid

	fmt.Println(m.StatusTable()[uid])
	fmt.Println(m.WorkerTable()[uid])
	fmt.Println(m.WorkerQueue()[uid])
}
