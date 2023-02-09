package main

import (
	"fmt"
	"sort"
)

func swapPhoneNumbers(phoneNumbers map[string]string) map[string]string {
	// Create a new map to hold the swapped values
	swappedPhoneNumbers := make(map[string]string)
	for key, value := range phoneNumbers {
		swappedPhoneNumbers[value] = key
	}

	// Return the swapped map
	return swappedPhoneNumbers
}

func sortPhoneNumbers(phoneNumbers map[string]string) map[string]string {
	// Create a slice to hold the keys of the map
	keys := make([]string, 0, len(phoneNumbers))
	for key := range phoneNumbers {
		keys = append(keys, key)
	}

	// Sort the keys based on their corresponding values (names)
	sort.Slice(keys, func(i, j int) bool {
		return phoneNumbers[keys[i]] > phoneNumbers[keys[j]]
	})

	// Create a new map to hold the sorted values
	sortedPhoneNumbers := make(map[string]string)
	for _, key := range keys {
		sortedPhoneNumbers[key] = phoneNumbers[key]
	}

	// Return the sorted map
	return sortedPhoneNumbers
}

func main() {
	phoneNumbers := map[string]string{
		"John Doe":   "+1 (555) 555-5555",
		"Jane Doe":   "+1 (555) 555-5556",
		"Jim Brown":  "+1 (555) 555-5557",
		"Jim Green":  "+1 (555) 555-5558",
		"John Green": "+1 (555) 555-5559",
		"Jane Green": "+1 (555) 555-5560",
	}

	fmt.Println("Unsorted phone numbers:")
	for name, phoneNumber := range phoneNumbers {
		fmt.Printf("%s: %s\n", name, phoneNumber)
	}

	fmt.Println("\nSorted phone numbers:")
	sortedPhoneNumbers := swapPhoneNumbers(phoneNumbers)
	sortedPhoneNumbers = sortPhoneNumbers(sortedPhoneNumbers)
	for name, phoneNumber := range sortedPhoneNumbers {
		fmt.Printf("%s: %s\n", name, phoneNumber)
	}
}
