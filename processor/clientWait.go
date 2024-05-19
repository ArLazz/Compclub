package processor

import (
	"compclub/models"
	"fmt"
)

func HandleClientWait(club *models.Club, event models.Event) {
	clientName := event.Details[0]
	club.Logs = append(club.Logs, fmt.Sprintf("%s 3 %s\n", event.EventTime.Format("15:04"), clientName))

	if checkFreeTables(club) {
		club.Logs = append(club.Logs, fmt.Sprintf("%s 13 ICanWaitNoLonger!\n", event.EventTime.Format("15:04")))
		return
	}

	client := club.CurrentStatus[clientName]
	if len(club.WaitingQueue) == len(club.Tables)-1 {
		club.Logs = append(club.Logs, fmt.Sprintf("%s 11 %s\n", event.EventTime.Format("15:04"), clientName))
		club.AddTimeAndRevenueForTable(client.Table, event.EventTime.Sub(client.JoinTime))
		client.Table = 0
		delete(club.CurrentStatus, clientName)
		return
	}
	
	club.WaitingQueue = append(club.WaitingQueue, client)
}

func checkFreeTables(club *models.Club) bool {
	for i := 1; i < len(club.Tables); i++ {
		if !club.Tables[i].IsBusy {
			return true
		}
	}
	return false
}
