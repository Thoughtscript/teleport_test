package models

import (
	"sync"
	"time"
)

// Adds workers to queue for processing
// Queue need not be an actual queue since these are triggered not by order but by timestamp
var workerQueue = make(map[string]time.Time)

// Define mutex
var wqLock = &sync.Mutex{}

// ----------------------------
// Helpers
// ----------------------------

func WorkerQueue() map[string]time.Time {
	wqLock.Lock()
	defer wqLock.Unlock()
	return workerQueue
}

func DeleteFromWorkerQueue(uid string) {
	wqLock.Lock()
	defer wqLock.Unlock()
	delete(workerQueue, uid)
}

func AddToWorkerQueue(worker WorkerModel) {
	wqLock.Lock()
	defer wqLock.Unlock()
	workerQueue[worker.Uuid] = worker.Time
}

func ReadFromWorkerQueue(uid string) time.Time {
	wqLock.Lock()
	defer wqLock.Unlock()
	return workerQueue[uid]
}