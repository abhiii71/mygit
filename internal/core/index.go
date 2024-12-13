package core

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetTrackedFiles() ([]string, error) {
	indexPath := ".mygit/index"

	if _, err := os.Stat(indexPath); os.IsNotExist(err) {
		return []string{}, nil
	}

	file, err := os.Open(indexPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var trackedFiles []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			trackedFiles = append(trackedFiles, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return trackedFiles, nil
}

func AddToTrackedFiles(filepath string) error {
	indexPath := ".mygit/index"

	file, err := os.OpenFile(indexPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(filepath + "\n")
	return err
}

func IsTracked(filepath string, trackedFiles []string) bool {
	for _, file := range trackedFiles {
		if file == filepath {
			return true
		}
	}
	return false
}

func AddToIndex(file string) error {
	_, err := os.Stat(file)
	if err != nil {
		return fmt.Errorf("file not found: %v", file)
	}

	fmt.Printf("File %s added to the index\n", file)
	return nil
}
