package handlers

import (
	j "../job"
	m "../models"
	"encoding/json"
	"github.com/gofrs/uuid"
	"log"
	"net/http"
	"time"
)

// CreateJob - POST - create a worker and it to the queue
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
		j.AddWorker(worker)

		w.WriteHeader(http.StatusCreated)
		err = json.NewEncoder(w).Encode(worker)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func QueryPool(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodGet {
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(j.GetAllJobs())
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func QueryJob(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodGet {
		uid := r.Header.Get("uuid")
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(j.GetJob(uid))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func StopJob(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodGet {
		uid := r.Header.Get("uuid")
		w.WriteHeader(http.StatusOK)
		j.StopWorker(uid)
		err := json.NewEncoder(w).Encode(j.GetJob(uid))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
