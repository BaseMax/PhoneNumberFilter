package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func partition(phoneNumbers []string, left int, right int) int {
	// Get the pivot
	pivot := phoneNumbers[right]

	// Create a variable to store the index of the smaller element
	i := left - 1

	// Loop through the array
	for j := left; j < right; j++ {
		// If the current element is smaller than the pivot
		if phoneNumbers[j] < pivot {
			// Increment the index of the smaller element
			i++

			// Swap the current element with the element at the index of the smaller element
			phoneNumbers[i], phoneNumbers[j] = phoneNumbers[j], phoneNumbers[i]
		}
	}

	// Swap the element at the index of the smaller element with the pivot
	phoneNumbers[i+1], phoneNumbers[right] = phoneNumbers[right], phoneNumbers[i+1]

	// Return the index of the smaller element
	return i + 1
}

func quickSort(phoneNumbers []string, left int, right int) {
	if left < right {
		// Partition the array
		pivot := partition(phoneNumbers, left, right)

		// Sort the left side
		quickSort(phoneNumbers, left, pivot-1)

		// Sort the right side
		quickSort(phoneNumbers, pivot+1, right)
	}
}

// Sort by name (UTF-8)
func sortPhoneNumbers(phoneNumbers map[string]string) map[string]string {
	// Create a slice to store the phone numbers
	var phoneNumbersSlice []string

	// Add the phone numbers to the slice
	for phoneNumber := range phoneNumbers {
		phoneNumbersSlice = append(phoneNumbersSlice, phoneNumber)
	}

	// Sort the phone numbers
	quickSort(phoneNumbersSlice, 0, len(phoneNumbersSlice)-1)

	// Create a map to store the sorted phone numbers
	sortedPhoneNumbers := make(map[string]string)

	// Add the phone numbers to the map
	for _, phoneNumber := range phoneNumbersSlice {
		sortedPhoneNumbers[phoneNumber] = phoneNumbers[phoneNumber]
	}

	// Return the sorted phone numbers
	return sortedPhoneNumbers
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

	// Sort the phone numbers by name
	phoneNumbers = sortPhoneNumbers(phoneNumbers)

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
