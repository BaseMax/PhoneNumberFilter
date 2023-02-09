package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Open the file
	file, err := os.Open("contacts.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a map to store the phone numbers and their corresponding names
	phoneNumbers := make(map[string]string)

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		columns := strings.Split(line, "\t")

		// Skip empty lines
		if len(columns) < 3 {
			continue
		}

		// Get the name and phone number
		name := columns[1]
		phoneNumber := columns[2]

		// Check if the phone number is already in the map
		if _, exists := phoneNumbers[phoneNumber]; exists {
			continue
		}

		// Trim name
		name = strings.TrimSpace(name)

		// Add the phone number and name to the map
		phoneNumbers[phoneNumber] = name
	}

	// Check for errors while reading the file
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Create a new file to store the filtered results
	resultFile, err := os.Create("filtered_contacts.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer resultFile.Close()

	// Write the header to the result file
	header := "ID\tName\tPhone Number\n"
	resultFile.WriteString(header)

	// Write the filtered results to the result file
	var id int
	for phoneNumber, name := range phoneNumbers {
		id++
		result := fmt.Sprintf("%d\t%s\t%s\n", id, name, phoneNumber)
		resultFile.WriteString(result)
	}

	// Print the filtered results in a table format
	fmt.Println(header)
	for phoneNumber, name := range phoneNumbers {
		id++
		fmt.Printf("%d\t%s\t%s\n", id, name, phoneNumber)
	}
}
