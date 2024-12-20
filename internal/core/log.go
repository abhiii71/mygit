package core

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ShowLog() error {
	if _, err := os.Stat(".mygit"); os.IsNotExist(err) {
		return fmt.Errorf("not a MyGit repository (or any of the parent directories): .mygit")
	}

	headFile := ".mygit/HEAD"
	headData, err := os.ReadFile(headFile)
	if err != nil {
		return fmt.Errorf("error reading HEAD file: %v", err)
	}

	currentCommitHash := strings.TrimSpace(string(headData))
	if currentCommitHash == "" {
		return fmt.Errorf("no commits found in the repository")
	}

	fmt.Println("Commit History:")
	for currentCommitHash != "" {
		commitFile := filepath.Join(".mygit", "objects", currentCommitHash)
		commitData, err := os.ReadFile(commitFile)
		if err != nil {
			return fmt.Errorf("error reading commit file %s: %v", commitFile, err)
		}

		// Parse the commit content
		scanner := bufio.NewScanner(strings.NewReader(string(commitData)))
		var commitMessage, timestamp, parentHash string
		for scanner.Scan() {
			line := scanner.Text()
			if strings.HasPrefix(line, "Commit Message:") {
				commitMessage = strings.TrimSpace(strings.TrimPrefix(line, "Commit Message:"))
			} else if strings.HasPrefix(line, "Timestamp:") {
				timestamp = strings.TrimSpace(strings.TrimPrefix(line, "Timestamp:"))
			} else if strings.HasPrefix(line, "Parent:") {
				parentHash = strings.TrimSpace(strings.TrimPrefix(line, "Parent:"))
			}
		}

		// Display  commit information
		fmt.Printf("\nCommit: %s\n", currentCommitHash)
		fmt.Printf("Date: %s\n", timestamp)
		fmt.Printf("Message: %s\n", commitMessage)

		// Move to the parent commit
		currentCommitHash = parentHash
	}
	return nil
}
