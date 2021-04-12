package models

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"time"
)

type WorkerModel struct {
	Id   string
	Time time.Time
	Status string
	Command string
}

func (w WorkerModel) execute () {
	w.Status = "Executing"
	fmt.Printf(w.Command)

	cmd := exec.Command("ls")

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		w.Status = "Failed"
		log.Fatal(err)
	}

	w.Status = "Completed"
	fmt.Printf(out.String())
}

