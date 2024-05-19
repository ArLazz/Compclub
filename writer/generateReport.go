package writer

import (
	"compclub/models"
	"fmt"
)

func GenerateReport(club *models.Club) {
	fmt.Println(club.OpenTime.Format("15:04"))

	for _, log := range club.Logs {
		fmt.Print(log)
	}

	fmt.Println(club.CloseTime.Format("15:04"))

	for i := 1; i < len(club.Tables); i++ {
		table := club.Tables[i]
		fmt.Printf("%d %d %02d:%02d\n", i, table.Revenue*club.HourlyRate, int(table.TotalTime.Hours()), int(table.TotalTime.Minutes())%60)
	}

}
