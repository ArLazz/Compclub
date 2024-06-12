package models

import (
	"fmt"
	"strconv"
)

type EventCLientSit struct {
	Data EventData
}

func (e *EventCLientSit) Procces(club *Club) {
	clientName := e.Data.Details[0]
	table, _ := strconv.Atoi(e.Data.Details[1])
	club.Logs = append(club.Logs, fmt.Sprintf("%s 2 %s %d\n", e.Data.EventTime.Format("15:04"), clientName, table))

	client, exists := club.CurrentStatus[clientName]
	if !exists {
		club.Logs = append(club.Logs, fmt.Sprintf("%s 13 ClientUnknown\n", e.Data.EventTime.Format("15:04")))
		return
	}

	if club.Tables[table].IsBusy {
		club.Logs = append(club.Logs, fmt.Sprintf("%s 13 PlaceIsBusy\n", e.Data.EventTime.Format("15:04")))
		return
	}

	if client.Table != 0 {
		club.AddTimeAndRevenueForTable(client.Table, e.Data.EventTime.Sub(client.JoinTime))
		club.Tables[client.Table].IsBusy = false
	}

	club.Tables[table].IsBusy = true
	client.Table = table
	client.JoinTime = e.Data.EventTime
}
