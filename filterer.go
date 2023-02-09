package main

import (
	"fmt"
)

// Return string or error
func file_reads(filepath string) 

func main() {
	// file path
	path := "data.txt"
	// open and read the file
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// create a new scanner and read the file line by line
	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		if i == 0 {

		fmt.Println(scanner.Text())
	}
}
