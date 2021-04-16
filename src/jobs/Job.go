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
	s := m.ReadFromStatusTable(uid)
	if s != "" {
		RemoveWorker(uid)
		m.UpdateStatusTable(uid, "stopped")
		fmt.Println("Worker stopped:", uid)
	}
}

func GetJob(uid string) m.WorkerModel {
	w := m.ReadFromWorkerTable(uid)
	if w.Uuid == "" {
		w.Uuid = uid
		w.Status = "not found"
		w.Output = "not found"
	}
	return w
}

func GetAllJobs() map[string]string {
	return m.StatusTable()
}

func GetStatus(uid string) string {
	s := m.ReadFromStatusTable(uid)
	if s == "" {
		s = "Worker " + uid + " not found!"
	}
	return s
}
