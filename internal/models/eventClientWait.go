package models

import (
	"fmt"
)

type EventClientWait struct {
	Data EventData
}

func (e *EventClientWait) Procces(club *Club) {
	clientName := e.Data.Details[0]
	club.Logs = append(club.Logs, fmt.Sprintf("%s 3 %s\n", e.Data.EventTime.Format("15:04"), clientName))

	if CheckFreeTables(club) {
		club.Logs = append(club.Logs, fmt.Sprintf("%s 13 ICanWaitNoLonger!\n", e.Data.EventTime.Format("15:04")))
		return
	}

	client := club.CurrentStatus[clientName]
	if len(club.WaitingQueue) == len(club.Tables)-1 {
		club.Logs = append(club.Logs, fmt.Sprintf("%s 11 %s\n", e.Data.EventTime.Format("15:04"), clientName))
		club.AddTimeAndRevenueForTable(client.Table, e.Data.EventTime.Sub(client.JoinTime))
		client.Table = 0
		delete(club.CurrentStatus, clientName)
		return
	}

	club.WaitingQueue = append(club.WaitingQueue, client)
}

func CheckFreeTables(club *Club) bool {
	for i := 1; i < len(club.Tables); i++ {
		if !club.Tables[i].IsBusy {
			return true
		}
	}
	return false
}
