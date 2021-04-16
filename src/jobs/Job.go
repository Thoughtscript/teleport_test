package jobs

import (
	m "../models"
	"fmt"
	"github.com/gofrs/uuid"
	"time"
)

// ----------------------------
// Worker Helpers
// ----------------------------

func NewWorker(scheduled time.Time, cmd string) m.WorkerModel {
	return m.WorkerModel{
		Uuid:    uuid.Must(uuid.NewV4()).String(),
		Time:    scheduled,
		Status:  "queued",
		Command: cmd,
	}
}

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
	if s != "" && s != "failed" && s != "completed"{
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

// ----------------------------
// Authentication Helpers
// ----------------------------

func VerifyPassword(user string, password string) bool {
	p := m.ReadFromAuthenticationTable(user)
	if p == "" {
		return false
	}
	return p == password
}

func SetPassword(user string, password string) {
	m.AddToAuthenticationTable(user, password)
}