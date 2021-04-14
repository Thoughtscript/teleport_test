package models

import (
	"time"
)

type WorkerModel struct {
	Uuid    string
	Time    time.Time
	Status  string
	Command string
	Output  string
}