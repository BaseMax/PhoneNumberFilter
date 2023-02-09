package main

import (
	"fmt"
	"sort"
)

func main() {
	// Create a slice of Persian strings
	list := []string{"علی", "اسمان", "ابرو", "", "مهدی", "محمد"}

	// Sort the slice of strings
	sort.Strings(list)

	// Print the sorted slice of strings
	for _, v := range list {
		fmt.Println(v)
	}
}
