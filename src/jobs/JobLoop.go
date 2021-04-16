package jobs

import (
	"sync"
)

// ----------------------------
// Main Job Loop
// ----------------------------

func JobLoop() {
	wg := new(sync.WaitGroup)

	SetPassword("test", "test")

	wg.Add(1)
	go ProcessQueue(wg)

	wg.Wait()
}
