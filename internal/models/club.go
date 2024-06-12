package models

import (
	"math"
	"time"
)

type Club struct {
	Tables        []Table
	OpenTime      time.Time
	CloseTime     time.Time
	HourlyRate    int
	Events        []Event
	CurrentStatus map[string]*Client
	WaitingQueue  []*Client
	Logs          []string
}

type Table struct {
	Revenue   int
	TotalTime time.Duration
	IsBusy    bool
}

func (c *Club) AddTimeAndRevenueForTable(table int, time time.Duration) {
	c.Tables[table].TotalTime += time
	c.Tables[table].Revenue += int(math.Ceil(time.Hours()))
}
