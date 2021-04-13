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
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodPost {
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

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(worker)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func QueryPool(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodGet {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(m.GetWorkerPool())
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func QueryJob(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodGet {
		uuid := r.Header.Get("uuid")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(m.GetJob(uuid))
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func StopJob(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodGet {
		uuid := r.Header.Get("uuid")
		w.WriteHeader(http.StatusOK)
		m.StopWorker(uuid)
		json.NewEncoder(w).Encode(m.GetJob(uuid))
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
