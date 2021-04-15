package job

import (
	"sync"
)

// ----------------------------
// Main Job Loop
// ----------------------------

func JobLoop() {
	wg := new(sync.WaitGroup)

	wg.Add(1)
	go ProcessQueue(wg)

	wg.Wait()
}
