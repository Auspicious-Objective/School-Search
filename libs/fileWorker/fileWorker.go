package fileWorker

import (
	"encoding/csv"
	"fmt"
	"os"
)

// Generates the preliminary file with the given name
func makeFile(fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("There was an error in generating your file. Please try again.")
		os.Exit(1)
	}
	defer file.Close()
}

// adds the header e.g. school name etc.
func addFileHeaders(fileName string) {
	headers := []string{"School Name", "Location", "Acceptance Rate", "Cost of Attendance", "Students:Teachers", "Undergrad Population", "Testing Requirements"}
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("There was an error in opening the file to write the initial headers. Please try again.")
		os.Exit(1)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write(headers)
}

// used to generate the starting file that all of the other data will get passed into
func GenerateFile(fileName string) {
	makeFile(fileName)
	addFileHeaders(fileName)
}

// Primary writing function that adds each schools info to the file.
func AddDataToFile(fileName string, data []string) {
	// 7 headers needs 7 bits of info
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("There was an error in opening the file to write college data. Please try again.")
		os.Exit(1)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write(data)
}
