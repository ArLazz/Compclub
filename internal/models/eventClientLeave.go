package models

import (
	"fmt"
)

type EventClientLeave struct {
	Data EventData
}

func (e *EventClientLeave) Procces(club *Club) {
	clientName := e.Data.Details[0]
	club.Logs = append(club.Logs, fmt.Sprintf("%s 4 %s\n", e.Data.EventTime.Format("15:04"), clientName))

	client, exists := club.CurrentStatus[clientName]
	if !exists {
		club.Logs = append(club.Logs, fmt.Sprintf("%s 13 ClientUnknown\n", e.Data.EventTime.Format("15:04")))
		return
	}

	club.AddTimeAndRevenueForTable(client.Table, e.Data.EventTime.Sub(client.JoinTime))
	club.Tables[client.Table].IsBusy = false
	table := client.Table

	if len(club.WaitingQueue) != 0 { //if queue is not empty
		clientFromQueue := club.WaitingQueue[0]   //take first client from queue
		club.WaitingQueue = club.WaitingQueue[1:] //delete first client from queue
		club.Logs = append(club.Logs, fmt.Sprintf("%s 12 %s %d\n", e.Data.EventTime.Format("15:04"), clientFromQueue.Name, table))

		if clientFromQueue.Table != 0 {
			club.AddTimeAndRevenueForTable(clientFromQueue.Table, e.Data.EventTime.Sub(clientFromQueue.JoinTime))
			club.Tables[clientFromQueue.Table].IsBusy = false
		}

		club.Tables[table].IsBusy = true
		clientFromQueue.Table = table
		clientFromQueue.JoinTime = e.Data.EventTime
	}
	client.Table = 0
	delete(club.CurrentStatus, clientName)
}
