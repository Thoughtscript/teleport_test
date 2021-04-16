package models

import (
	"sync"
)

// Tracks extant workers
var workerTable = make(map[string]WorkerModel)

// Define mutex
var wtLock = &sync.Mutex{}

// ----------------------------
// Helpers
// ----------------------------

func WorkerTable() map[string]WorkerModel {
	wtLock.Lock()
	defer wtLock.Unlock()
	return workerTable
}

func DeleteFromWorkerTable(uid string) {
	wtLock.Lock()
	defer wtLock.Unlock()
	delete(workerTable, uid)
}

// AddToWorkerTable - technically, an upsert
func AddToWorkerTable(worker WorkerModel) {
	wtLock.Lock()
	defer wtLock.Unlock()
	workerTable[worker.Uuid] = worker
}

func ReadFromWorkerTable(uid string) WorkerModel {
	wtLock.Lock()
	defer wtLock.Unlock()
	return workerTable[uid]
}