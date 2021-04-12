package main

import (
	h "./handlers"
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
	http.HandleFunc("/api/query", h.QueryJob)

	// TLS
	port := ":8888"
	http.ListenAndServeTLS(port, "cert.crt", "key.key", nil)
	log.Println("Listening on port {{port}} ...")
}