package main

import (
	h "./handlers"
	j "./jobs"
	"log"
	"net/http"
)

func main() {
	// Static assets for client
	fileServer := http.FileServer(http.Dir("./assets"))
	http.Handle("/", http.StripPrefix("/public/", fileServer))
	http.Handle("/public/", http.StripPrefix("/public/", fileServer))

	// API endpoints
	http.HandleFunc("/api/create", h.CreateJob)
	http.HandleFunc("/api/jobs", h.QueryJob)
	http.HandleFunc("/api/pool", h.QueryPool)
	http.HandleFunc("/api/stop", h.StopJob)
	http.HandleFunc("/api/status", h.QueryStatus)

	go j.JobLoop()
	//go t.TestLoop()

	// TLS
	port := ":443"
	log.Println("Listening on port", port)
	go log.Fatal(http.ListenAndServeTLS(port, "cert.pem", "key.pem", nil))
}
