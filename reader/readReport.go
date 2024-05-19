package reader

import (
	"bufio"
	"compclub/models"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func ParseTime(s string) (time.Time, error) {
	return time.Parse("15:04", s)
}

func ReadInput(fileName string) (*models.Club, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	club := &models.Club{}
	club.CurrentStatus = make(map[string]*models.Client)

	// Read the number of tables
	scanner.Scan()
	line := scanner.Text()
	numOfTables, err := strconv.Atoi(line)
	if err != nil {
		fmt.Println(line)
		return nil, err
	}
	club.Tables = make([]models.Table, numOfTables+1)

	// Read the opening and closing times
	scanner.Scan()
	line = scanner.Text()
	times := strings.Split(line, " ")
	if len(times) != 2 {
		fmt.Println(line)
		return nil, fmt.Errorf("error to parse opening and closing times")
	}

	club.OpenTime, err = ParseTime(times[0])
	if err != nil {
		fmt.Println(line)
		return nil, err
	}

	club.CloseTime, err = ParseTime(times[1])
	if err != nil {
		fmt.Println(line)
		return nil, err
	}

	// Read the hourly rate
	scanner.Scan()
	line = scanner.Text()
	club.HourlyRate, err = strconv.Atoi(line)
	if err != nil {
		fmt.Println(line)
		return nil, err
	}

	// Read the events
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")

		eventTime, err := ParseTime(parts[0])
		if err != nil {
			fmt.Println(line)
			return nil, err
		}

		eventType, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Println(line)
			return nil, err
		}

		event := models.Event{
			EventTime: eventTime,
			EventType: eventType,
			Details:   parts[2:],
		}
		club.Events = append(club.Events, event)
	}
	
	return club, nil
}
