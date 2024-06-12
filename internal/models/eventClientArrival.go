package models

import (
	"fmt"
)

type EvemtClientArrival struct {
	Data EventData
}

func (e *EvemtClientArrival) Procces(club *Club) {
	clientName := e.Data.Details[0]
	club.Logs = append(club.Logs, fmt.Sprintf("%s 1 %s\n", e.Data.EventTime.Format("15:04"), clientName))
	if _, exists := club.CurrentStatus[clientName]; exists {
		club.Logs = append(club.Logs, fmt.Sprintf("%s 13 YouShallNotPass\n", e.Data.EventTime.Format("15:04")))
		return
	}

	if e.Data.EventTime.Before(club.OpenTime) || e.Data.EventTime.After(club.CloseTime) {
		club.Logs = append(club.Logs, fmt.Sprintf("%s 13 NotOpenYet\n", e.Data.EventTime.Format("15:04")))
		return
	}

	club.CurrentStatus[clientName] = &Client{Name: clientName}
}
