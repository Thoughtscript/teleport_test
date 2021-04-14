package main

import (
	h "./handlers"
	m "./models"
	"fmt"
	"github.com/gofrs/uuid"
	"log"
	"net/http"
	"sync"
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
	http.HandleFunc("/api/stop", h.StopJob)

	// TLS
	port := ":8888"
	//err := http.ListenAndServe(":8888", nil)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//http.ListenAndServeTLS(port, "cert.pem", "key.pem", nil)
	log.Println("Listening on port", port)

	handleQueue()
}

// ----------------------------
// Queue
// ----------------------------

func handleQueue() {
	wg := new(sync.WaitGroup)

	wg.Add(1)
	go m.ProcessQueue(wg)

	wg.Add(1)
	go Test1(wg)
	wg.Add(1)
	go Test2(wg)

	wg.Wait()
}

// ----------------------------
// Debugging and tests
// ----------------------------

func Test1(wg *sync.WaitGroup) {
	defer wg.Done()

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

	uuid3 := uuid.Must(uuid.NewV4()).String()
	worker3 := m.WorkerModel{
		Uuid:    uuid3,
		Time:    time.Now(),
		Status:  "queued",
		Command: "ls",
	}
	m.AddWorker(worker3)
	fmt.Println(uuid3, m.GetJob(uuid3))
	fmt.Println(m.GetWorkerPool().Workers)
	fmt.Println(m.GetWorkerPool().Statuses)
	m.StopWorker(uuid3)
	fmt.Println(m.GetAllJobs())
	fmt.Println(m.GetWorkerPool().Workers)
}

func Test2(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		var second time.Duration = 1000000000
		var seconds time.Duration = 5
		time.Sleep(seconds * second)

		uud := uuid.Must(uuid.NewV4()).String()
		fmt.Println("Adding worker", uud)

		worker := m.WorkerModel{
			Uuid:    uud,
			Time:    time.Now(),
			Status:  "queued",
			Command: "ls",
		}
		m.AddWorker(worker)
	}
}
