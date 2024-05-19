package models

import "time"

type Event struct {
	EventTime time.Time
	EventType int
	Details   []string
}
