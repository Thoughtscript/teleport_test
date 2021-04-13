package handlers

import (
	m "../models"
	"encoding/json"
	"github.com/gofrs/uuid"
	"log"
	"net/http"
	"time"
)

func CreateJob(w http.ResponseWriter, r *http.Request) {
	scheduled, err := time.Parse(time.UnixDate, r.Header.Get("scheduled"))
	if err != nil {
		log.Fatal("Exception encountered - please supply a valid scheduled time!")
	}

	worker := m.WorkerModel{
		Uuid:    uuid.Must(uuid.NewV4()).String(),
		Time:    scheduled,
		Status:  "queued",
		Command: r.Header.Get("cmd"),
	}

	m.AddWorker(worker)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(worker)
}

func QueryJob(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(m.GetWorkerPool())
}

func QueryPool(w http.ResponseWriter, r *http.Request) {
	uuid := r.Header.Get("uuid")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(m.GetJob(uuid))
}
