package go_utils

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

// Read opens the file at the given path and returns its entire content as a string.
// Returns an error if the file cannot be opened or if reading fails.
func Read(path string) (string, error) {
	file, err := os.Open(path)

	if err != nil {
		return "", err
	}

	defer file.Close()
	var content string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content += scanner.Text() + "\n"
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return content, nil
}

// ReadIntoStrArr reads the file at the given path and returns its contents
// as a slice of strings, one element per line.
// The final line is trimmed if empty. Returns an error if reading the file fails.
func ReadIntoStrArr(path string) ([]string, error) {
	content, err := Read(path)

	if err != nil {
		return []string{}, err
	}

	lines := strings.Split(content, "\n")
	if len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	if len(lines) == 0 {
		return []string{}, errors.New("You forgot to add the input dummy")
	}

	return lines, err
}
