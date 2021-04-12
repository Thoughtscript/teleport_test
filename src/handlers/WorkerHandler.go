package handlers

import (
	"log"
	"net/http"
	m "../models"
	"time"
)

func CreateJob(w http.ResponseWriter, r *http.Request) {
	scheduled, err := time.Parse(time.UnixDate, r.Header.Get("scheduled"))
	if err != nil {
		log.Fatal("Exeception encountered - please supply a valid scheduled time@")
	}

	worker := m.WorkerModel{
		Id: "",
		Time: scheduled,
	}


}

func QueryJob(w http.ResponseWriter, r *http.Request) {

}