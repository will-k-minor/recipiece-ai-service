package utils

import (
	"bufio"
	"os"
)

func ReadFileLines(filePath string) ([]string, error) {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Initialize a scanner
	scanner := bufio.NewScanner(file)
	var lines []string

	// Read each line and append to the slice
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
