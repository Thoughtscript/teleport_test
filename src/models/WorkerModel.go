package models

import (
	"bytes"
	"fmt"
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
	fmt.Println(w.Uuid, w.Status, w.Command)
	outputChannel := make(chan string)

	go func() {
		// Hardcoded for now - need client-injection protection
		// Possible real-world cases for this include common bash commands
		// Or, more likely, a pre-defined set of operations specific to the API or tool - like '$ teleport_test hello -> world'

		cmd := exec.Command("ls")
		var out bytes.Buffer
		cmd.Stdout = &out

		err := cmd.Run()
		if err != nil {
			w.Status = "Failed"
			GetWorkerPool().Statuses[w.Uuid] = w.Status
			outputChannel <- w.Uuid + " " + w.Status + " " + w.Command
		}

		w.Status = "Completed"
		GetWorkerPool().Statuses[w.Uuid] = w.Status
		outputChannel <- w.Uuid + " " + w.Status + " " + w.Command + " " + out.String()
	}()

	result := <- outputChannel
	fmt.Println(result)
}
