package processor

import (
	"compclub/internal/models"
	"fmt"
)

func ProcessEvents(club *models.Club) {
	for _, event := range club.Events {
		event.Procces(club)
	}

	if models.CheckFreeTables(club) {
		for clientName, client := range club.CurrentStatus {
			club.Logs = append(club.Logs, fmt.Sprintf("%s 11 %s\n", club.CloseTime.Format("15:04"), clientName))
			club.AddTimeAndRevenueForTable(client.Table, club.CloseTime.Sub(client.JoinTime))
		}

	}
}
