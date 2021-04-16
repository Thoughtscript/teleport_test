package handlers

import (
	j "../jobs"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// CreateJob - POST - create a worker and it to the queue
func CreateJob(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodPost {
		scheduled, err := time.Parse(time.RFC3339, r.Header.Get("scheduled"))
		if err != nil {
			log.Println("Exception encountered - please supply a valid scheduled time!")
		}

		user := r.Header.Get("user")
		password := r.Header.Get("password")
		cmd := r.Header.Get("cmd")

		if j.VerifyPassword(user, password) {
			worker := j.NewWorker(scheduled, cmd)
			j.AddWorker(worker)
			w.WriteHeader(http.StatusCreated)
			err = json.NewEncoder(w).Encode(worker)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
			}
		} else {
			w.WriteHeader(http.StatusForbidden)
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func QueryPool(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodGet {
		user := r.Header.Get("user")
		password := r.Header.Get("password")

		if j.VerifyPassword(user, password) {
			w.WriteHeader(http.StatusOK)
			err := json.NewEncoder(w).Encode(j.GetAllJobs())
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
			}
		} else {
			w.WriteHeader(http.StatusForbidden)
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func QueryJob(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodGet {
		user := r.Header.Get("user")
		password := r.Header.Get("password")

		if j.VerifyPassword(user, password) {
			uid := r.Header.Get("uuid")
			w.WriteHeader(http.StatusOK)
			err := json.NewEncoder(w).Encode(j.GetJob(uid))
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
			}
		} else {
			w.WriteHeader(http.StatusForbidden)
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func StopJob(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodPost {
		user := r.Header.Get("user")
		password := r.Header.Get("password")

		if j.VerifyPassword(user, password) {
			uid := r.Header.Get("uuid")
			w.WriteHeader(http.StatusOK)
			j.StopWorker(uid)
			err := json.NewEncoder(w).Encode(j.GetStatus(uid))
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
			}
		} else {
			w.WriteHeader(http.StatusForbidden)
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func QueryStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodGet {
		user := r.Header.Get("user")
		password := r.Header.Get("password")

		if j.VerifyPassword(user, password) {
			uid := r.Header.Get("uuid")
			w.WriteHeader(http.StatusOK)
			err := json.NewEncoder(w).Encode(j.GetStatus(uid))
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
			}
		} else {
			w.WriteHeader(http.StatusForbidden)
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}