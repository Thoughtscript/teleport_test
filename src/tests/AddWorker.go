package tests

import (
	j "../job"
	m "../models"
	"fmt"
	"github.com/gofrs/uuid"
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

	uid := uuid.Must(uuid.NewV4()).String()
	fmt.Println("Adding worker:", uid)

	worker := m.WorkerModel{
		Uuid:    uid,
		Time:    time.Now(),
		Status:  "queued",
		Command: "ls",
	}

	j.AddWorker(worker)

	fmt.Println(m.StatusTable()[uid])
	fmt.Println(m.WorkerTable()[uid])
	fmt.Println(m.WorkerQueue()[uid])

	j.StopWorker(uid)
}

func AddWorkerAndExecuteWg(wg *sync.WaitGroup) {
	defer wg.Done()

	uid := uuid.Must(uuid.NewV4()).String()
	fmt.Println("Adding worker:", uid)

	worker := m.WorkerModel{
		Uuid:    uid,
		Time:    time.Now(),
		Status:  "queued",
		Command: "ls",
	}

	j.AddWorker(worker)

	fmt.Println(m.StatusTable()[uid])
	fmt.Println(m.WorkerTable()[uid])
	fmt.Println(m.WorkerQueue()[uid])

	worker.ExecuteCommand()
}

func AddWorkerAndPrintWg(wg *sync.WaitGroup) {
	defer wg.Done()

	uid := uuid.Must(uuid.NewV4()).String()
	fmt.Println("Adding worker:", uid)

	worker := m.WorkerModel{
		Uuid:    uid,
		Time:    time.Now(),
		Status:  "queued",
		Command: "ls",
	}

	j.AddWorker(worker)

	fmt.Println(m.StatusTable()[uid])
	fmt.Println(m.WorkerTable()[uid])
	fmt.Println(m.WorkerQueue()[uid])
}

// ----------------------------
// Private
// ----------------------------

func addWorkerAndExecute() {
	uid := uuid.Must(uuid.NewV4()).String()
	fmt.Println("Adding worker:", uid)

	worker := m.WorkerModel{
		Uuid:    uid,
		Time:    time.Now(),
		Status:  "queued",
		Command: "ls",
	}

	j.AddWorker(worker)

	fmt.Println(m.StatusTable()[uid])
	fmt.Println(m.WorkerTable()[uid])
	fmt.Println(m.WorkerQueue()[uid])

	worker.ExecuteCommand()
}

func addWorkerAndPrint() {
	uid := uuid.Must(uuid.NewV4()).String()
	fmt.Println("Adding worker:", uid)

	worker := m.WorkerModel{
		Uuid:    uid,
		Time:    time.Now(),
		Status:  "queued",
		Command: "ls",
	}

	j.AddWorker(worker)

	fmt.Println(m.StatusTable()[uid])
	fmt.Println(m.WorkerTable()[uid])
	fmt.Println(m.WorkerQueue()[uid])
}
