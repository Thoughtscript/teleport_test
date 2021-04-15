package job

import (
	m "../models"
)

// ----------------------------
// Worker-Queue Helpers
// ----------------------------

func AddWorker(worker m.WorkerModel) {
	m.AddToStatusTable(worker)
	m.AddToWorkerTable(worker)
	m.AddToWorkerQueue(worker)
}

func RemoveWorker(uid string) {
	m.DeleteFromWorkerQueue(uid)
	m.DeleteFromWorkerTable(uid)
}

func StopWorker(uid string) {
	RemoveWorker(uid)
	m.UpdateStatusTable(uid, "stopped")
}

func GetJob(uid string) m.WorkerModel {
	return m.ReadFromWorkerTable(uid)
}

func GetAllJobs() map[string]string {
	return m.StatusTable()
}
