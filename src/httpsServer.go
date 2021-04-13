package main

import (
	h "./handlers"
	m "./models"
	"fmt"
	"github.com/gofrs/uuid"
	"log"
	"net/http"
	"time"
)

func main() {
	// Static assets for client
	fileServer := http.FileServer(http.Dir("./assets"))
	http.Handle("/", http.StripPrefix("/public/", fileServer))
	http.Handle("/public/", http.StripPrefix("/public/", fileServer))

	// API endpoints
	http.HandleFunc("/api/create", h.CreateJob)
	http.HandleFunc("/api/job", h.QueryJob)
	http.HandleFunc("/api/pool", h.QueryPool)

	// TLS
	port := ":8888"
	http.ListenAndServeTLS(port, "cert.crt", "key.key", nil)
	log.Println("Listening on port", port)

	Test()
	//m.ProcessQueue()
}

// Tests
func Test() {
	uuid1 := uuid.Must(uuid.NewV4()).String()
	worker1 := m.WorkerModel{
		Uuid:    uuid1,
		Time:    time.Now(),
		Status:  "queued",
		Command: "ls",
	}
	m.AddWorker(worker1)
	fmt.Println(uuid1, m.GetJob(uuid1))
	fmt.Println(m.GetWorkerPool().Workers)
	fmt.Println(m.GetWorkerPool().Statuses)
	worker1.ExecuteCommand()
	fmt.Println(m.GetAllJobs())
	fmt.Println(m.GetWorkerPool().Workers)

	uuid2 := uuid.Must(uuid.NewV4()).String()
	worker2 := m.WorkerModel{
		Uuid:    uuid2,
		Time:    time.Now(),
		Status:  "queued",
		Command: "ls",
	}
	m.AddWorker(worker2)
	fmt.Println(uuid1, m.GetJob(uuid1))
	fmt.Println(m.GetWorkerPool().Workers)
	fmt.Println(m.GetWorkerPool().Statuses)
	worker2.ExecuteCommand()
	fmt.Println(m.GetAllJobs())
	fmt.Println(m.GetWorkerPool().Workers)
}