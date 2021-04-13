package models

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"time"
)

type WorkerModel struct {
	Uuid    string
	Time    time.Time
	Status  string
	Command string
}

// Receiver-type function
func (w WorkerModel) ExecuteCommand() {
	w.Status = "Executing"
	GetWorkerPool().Statuses[w.Uuid] = w.Status
	fmt.Printf(w.Command)

	cmd := exec.Command("ls")

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		w.Status = "Failed"
		GetWorkerPool().Statuses[w.Uuid] = w.Status
		log.Fatal(err)
	}

	w.Status = "Completed"
	GetWorkerPool().Statuses[w.Uuid] = w.Status
	fmt.Printf(out.String())
}
