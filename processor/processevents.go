package processor

import (
	"compclub/models"
	"fmt"
)

func ProcessEvents(club *models.Club) {
	for _, event := range club.Events {
		switch event.EventType {
		case 1:
			HandleClientArrival(club, event)
		case 2:
			HandleClientSit(club, event)
		case 3:
			HandleClientWait(club, event)
		case 4:
			HandleClientLeave(club, event)
		}
	}
	
	if checkFreeTables(club) {
		for clientName, client := range club.CurrentStatus {
			club.Logs = append(club.Logs, fmt.Sprintf("%s 11 %s\n", club.CloseTime.Format("15:04"), clientName))
			club.AddTimeAndRevenueForTable(client.Table, club.CloseTime.Sub(client.JoinTime))
		}

	}
}
