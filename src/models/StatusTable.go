package models

import "sync"

// Tracks worker statuses even after workers are removed from WorkerTable
var statusTable = make(map[string]string)

// Define mutex
var stLock = &sync.Mutex{}

// ----------------------------
// Helpers
// ----------------------------

func StatusTable() map[string]string {
	stLock.Lock()
	defer stLock.Unlock()
	return statusTable
}

func DeleteFromStatusTable(uid string) {
	stLock.Lock()
	defer stLock.Unlock()
	delete(statusTable, uid)
}

func AddToStatusTable(worker WorkerModel) {
	stLock.Lock()
	defer stLock.Unlock()
	statusTable[worker.Uuid] = worker.Status
}

func ReadFromStatusTable(uid string) string {
	stLock.Lock()
	defer stLock.Unlock()
	return statusTable[uid]
}

func UpdateStatusTable(uid string, status string) {
	stLock.Lock()
	defer stLock.Unlock()
	statusTable[uid] = status
}