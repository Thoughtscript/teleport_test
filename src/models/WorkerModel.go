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
	Output  string
}

// ----------------------------
// Worker-Specific Helpers
// ----------------------------

func (worker WorkerModel) ExecuteCommand() {
	msg := "executing"
	worker.updateStatus(msg)
	outputChannel := make(chan string)

	go func() {
		// Hardcoded for now - need client-injection protection
		// Possible real-world cases for this include common bash commands
		// Or, more likely, a pre-defined set of operations specific to the API or tool - like '$ teleport_test hello -> world'
		cmd := exec.Command("ls")
		var out bytes.Buffer
		cmd.Stdout = &out

		var second time.Duration = 1000000000
		var seconds time.Duration = 3
		time.Sleep(second * seconds)

		// Check if worker stopped
		if worker.Status != "stopped" {
			err := cmd.Run()
			if err != nil {
				msg = "failed"
			} else {
				msg = "completed"
			}
			worker.log(msg, outputChannel, out.String())
		}
	}()

	result := <-outputChannel
	fmt.Println(result)
	close(outputChannel)
}

func (worker WorkerModel) log(msg string, outputChannel chan string, stdoutStr string) {
	worker.Output = stdoutStr
	worker.updateStatus(msg)
	outputChannel <- worker.Uuid + " " + worker.Status + " " + worker.Command + " " + stdoutStr
	DeleteFromWorkerQueue(worker.Uuid)
}

func (worker WorkerModel) updateStatus(msg string) {
	worker.Status = msg
	UpdateStatusTable(worker.Uuid, msg)
	// Update worker
	AddToWorkerTable(worker)
}
