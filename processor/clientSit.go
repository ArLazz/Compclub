package processor

import (
	"compclub/models"
	"fmt"
	"strconv"
)

func HandleClientSit(club *models.Club, event models.Event) {
	clientName := event.Details[0]
	table, _ := strconv.Atoi(event.Details[1])
	club.Logs = append(club.Logs, fmt.Sprintf("%s 2 %s %d\n", event.EventTime.Format("15:04"), clientName, table))

	client, exists := club.CurrentStatus[clientName]
	if !exists {
		club.Logs = append(club.Logs, fmt.Sprintf("%s 13 ClientUnknown\n", event.EventTime.Format("15:04")))
		return
	}

	if club.Tables[table].IsBusy {
		club.Logs = append(club.Logs, fmt.Sprintf("%s 13 PlaceIsBusy\n", event.EventTime.Format("15:04")))
		return
	}

	if client.Table != 0 {
		club.AddTimeAndRevenueForTable(client.Table, event.EventTime.Sub(client.JoinTime))
		club.Tables[client.Table].IsBusy = false
	}

	club.Tables[table].IsBusy = true
	client.Table = table
	client.JoinTime = event.EventTime
}
