package tests

import (
	j "../job"
	"sync"
)

// ----------------------------
// Test Loop
// ----------------------------

func TestLoop() {
	wg := new(sync.WaitGroup)

	wg.Add(1)
	go TestAddWorkerLoop(wg)

	wg.Add(1)
	go j.ProcessQueue(wg)

	wg.Wait()
}
