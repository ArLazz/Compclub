package main

import (
	"compclub/processor"
	"compclub/reader"
	"compclub/writer"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No input file")
		return
	}
	fileName := os.Args[1]
	club, err := reader.ReadInput(fileName)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	processor.ProcessEvents(club)
	writer.GenerateReport(club)
}
