package job

import (
	m "../models"
	"fmt"
	"sync"
	"time"
)

// ----------------------------
// Main Process Loop
// ----------------------------

func ProcessQueue(wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		var second time.Duration = 1000000000
		var seconds time.Duration = 5

		time.Sleep(seconds * second)
		fmt.Println("=================== POLLING ===================")
		fmt.Println("Polling every:", seconds*second)
		c := time.Now()

		wks := m.WorkerQueue()

		for uid, v := range wks {

			// Execute only if hasn't run yet
			// Workers can have queued, executing, stopped, failed, completed
			// The latter three will be removed from queue and kept in status map
			w := m.ReadFromWorkerTable(uid)

			if (v.Before(c) || v.Equal(c)) && w.Status == "queued" {
				w.ExecuteCommand()
			}
		}

		fmt.Println(m.StatusTable())
		fmt.Println(m.WorkerTable())
		fmt.Println(m.WorkerQueue())
	}
}

