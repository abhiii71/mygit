package core

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func AddFiles(files []string) error {
	ignoredFiles, err := getIgnoredFiles()
	if err != nil {
		return err
	}

	for _, file := range files {
		if contains(ignoredFiles, file) {
			fmt.Println("Warning: %s is ignored by .mygitignore\n", file)
			continue
		}

		err := index.AddToIndex(file)
		if err != nil {
			return fmt.Errorf("failed to add file %s: %v", file, err)
		}
	}
	return nil
}

func getIgnoredFiles() ([]string, error) {
	content, err := ioutil.ReadFile(".mygitignore")
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}
	ignoredFiles := strings.Split(string(content), "\n")
	return ignoredFiles, nil
}

func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}
