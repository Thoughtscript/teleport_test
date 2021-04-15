package main

import (
	h "./handlers"
	t "./tests"
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
	http.HandleFunc("/api/job", h.QueryJob)
	http.HandleFunc("/api/pool", h.QueryPool)
	http.HandleFunc("/api/stop", h.StopJob)

	// TLS
	port := ":8888"
	log.Println("Listening on port", port)

	//err := http.ListenAndServeTLS(port, "cert.pem", "key.pem", nil)
	//if err != nil {
	//	log.Fatal(err)
	//}

	//j.JobLoop()
	t.TestLoop()
}
