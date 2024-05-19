package processor

import (
	"compclub/models"
	"fmt"
)

func HandleClientArrival(club *models.Club, event models.Event) {
	clientName := event.Details[0]
	club.Logs = append(club.Logs, fmt.Sprintf("%s 1 %s\n", event.EventTime.Format("15:04"), clientName))

	if _, exists := club.CurrentStatus[clientName]; exists {
		club.Logs = append(club.Logs, fmt.Sprintf("%s 13 YouShallNotPass\n", event.EventTime.Format("15:04")))
		return
	}

	if event.EventTime.Before(club.OpenTime) || event.EventTime.After(club.CloseTime) {
		club.Logs = append(club.Logs, fmt.Sprintf("%s 13 NotOpenYet\n", event.EventTime.Format("15:04")))
		return
	}

	club.CurrentStatus[clientName] = &models.Client{Name: clientName}
}
