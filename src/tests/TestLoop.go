package tests

import (
	j "../job"
	"fmt"
	"sync"
	"time"
)

// ----------------------------
// Test Loop
// ----------------------------

func TestLoop() {
	wg := new(sync.WaitGroup)

	fmt.Println("============== (Some) Unit Regression Tests ============== ")

	wg.Add(1)
	go AddWorkerAndStop(wg)

	wg.Add(1)
	go AddWorkerAndExecuteWg(wg)

	wg.Add(1)
	go AddWorkerAndPrintWg(wg)

	var second time.Duration = 1000000000
	var seconds time.Duration = 10
	time.Sleep(second * seconds)

	fmt.Println("============== (Some) Integration and Loop Regression Tests ============== ")

	wg.Add(1)
	go TestAddWorkerLoop(wg)

	wg.Add(1)
	go TestExecuteWorkerLoop(wg)

	wg.Add(1)
	go j.ProcessQueue(wg)

	wg.Wait()
}
