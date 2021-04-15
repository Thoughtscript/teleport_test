package jobs

import (
	m "../models"
	"fmt"
)

// ----------------------------
// Worker-Queue Helpers
// ----------------------------

func AddWorker(worker m.WorkerModel) {
	m.AddToStatusTable(worker)
	m.AddToWorkerTable(worker)
	m.AddToWorkerQueue(worker)
	fmt.Println("Worker added:", worker.Uuid)
}

func RemoveWorker(uid string) {
	m.DeleteFromWorkerQueue(uid)
	m.DeleteFromWorkerTable(uid)
	fmt.Println("Worker removed:", uid)
}

func StopWorker(uid string) {
	RemoveWorker(uid)
	m.UpdateStatusTable(uid, "stopped")
	fmt.Println("Worker stopped:", uid)
}

func GetJob(uid string) m.WorkerModel {
	return m.ReadFromWorkerTable(uid)
}

func GetAllJobs() map[string]string {
	return m.StatusTable()
}

func GetStatus(uid string) string {
	return m.ReadFromStatusTable(uid)
}