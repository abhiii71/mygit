package core

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func Add(files []string) error {
	if _, err := os.Stat(".mygit"); os.IsNotExist(err) {
		return fmt.Errorf("not a MyGit repository (or any of the parent directories): .mygit")
	}

	//	indexPath := ".mygit/index"

	//Read the existing index
	existingIndex, err := readIndex()
	if err != nil {
		return fmt.Errorf("error reading index: %v", err)
	}

	// Filter and add files

	for _, file := range files {
		if _, err := os.Stat(file); os.IsNotExist(err) {
			return fmt.Errorf(" file %s doesn't exists\n", file)
		}

		absPath, _ := filepath.Abs(file)
		if !isTracked(absPath, existingIndex) {
			existingIndex = append(existingIndex, absPath)
		}
	}
	if err := writeIndex(existingIndex); err != nil {
		return fmt.Errorf("error writing index: %v", err)
	}
	return nil
}

func isTracked(file string, trackedFiles []string) bool {
	for _, tracked := range trackedFiles {
		if strings.EqualFold(file, tracked) {
			return true
		}
	}
	return false
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
