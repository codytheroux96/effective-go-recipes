/*
Opening a file and reading it line by line
*/

package reading

import (
	"bufio"
	"fmt"
	"os"
)

func readLines(filename string) error {
	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a new scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Iterate over each line
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line) // Process the line (here we just print it)
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func callReadLines() {
	    // Example usage
		if err := readLines("example.txt"); err != nil {
			fmt.Fprintln(os.Stderr, "Error reading file:", err)
			os.Exit(1)
		}
}