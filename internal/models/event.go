package models

import (
	"time"
)

type Event interface{
	Procces(club *Club)
} 

type EventData struct {
	EventTime time.Time
	Details   []string
}
