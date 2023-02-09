package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sort"
)

// Sort by name (UTF-8)
func sortPhoneNumbers(phoneNumbers map[string]string) map[string]string {
	keyList := []string{}
	for key,_ := range phoneNumbers {
		keyList = append(keyList,key)
	}
	sort.Strings(keyList)
	ret := map[string]string{}
	for i,v := range keyList {
		ret[v] = phoneNumbers[keyList[i]]
	}
	return ret
}

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
		// Trim line
		line = strings.TrimSpace(line)

		// Skip empty lines
		if len(line) == 0 {
			continue
		}

		// Split the line by tabs
		columns := strings.Split(line, "\t")

		// Skip empty lines
		if len(columns) < 3 {
			continue
		}

		// Get the name and phone number
		name := columns[1]
		phoneNumber := columns[2]

		// Trim name
		name = strings.TrimSpace(name)

		// If the phone number not starts with 0, and it starts with 9 accept it otherwise skip it
		if len(phoneNumber) > 0 {
			if phoneNumber[0] != '0' && phoneNumber[0] != '9' {
				continue
			} else if phoneNumber[0] == '9' {
				// add 0 to the beginning of the phone number
				phoneNumber = "0" + phoneNumber
			}
		}

		// Skip phone numbers that are not 11 digits
		if len(phoneNumber) != 11 {
			continue
		}

		// Check if the phone number is already in the map
		if _, exists := phoneNumbers[phoneNumber]; exists {
			continue
		}

		// Add the phone number and name to the map
		phoneNumbers[phoneNumber] = name
	}

	// Check for errors while reading the file
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	sortedPhoneNumbers := sortPhoneNumbers(phoneNumbers)
	_ = sortedPhoneNumbers

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
