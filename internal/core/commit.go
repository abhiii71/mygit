package core

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func Commit(message string) error {
	if _, err := os.Stat(".mygit"); os.IsNotExist(err) {
		return fmt.Errorf("not a MyGit repository (or any of the parent directories): .mygit")
	}

	trackedFils, err := readIndex()
	if err != nil {
		return fmt.Errorf("error reading index: %v", err)
	}

	if len(trackedFils) == 0 {
		return fmt.Errorf("nothing to commit, no tracked files")
	}

	// Create .mygit/objects directory if it doesn't exist
	objectsDir := ".mygit/objects"
	if _, err := os.Stat(objectsDir); os.IsNotExist(err) {
		if err := os.Mkdir(objectsDir, 0755); err != nil {
			return fmt.Errorf("error creating objects directory: %v", err)
		}
	}

	// Save the tracked files
	snapshot := make(map[string]string)
	for _, file := range trackedFils {
		content, err := ioutil.ReadFile(file)
		if err != nil {
			return fmt.Errorf("error reading file %s: %v", file, err)
		}
		snapshot[file] = string(content)
	}

	// generate commit hash
	commitHash := generateHash(message + time.Now().String())

	// save commit metadata
	commitFile := filepath.Join(objectsDir, commitHash)
	commitContent := buildCommitContent(message, snapshot)
	if err := ioutil.WriteFile(commitFile, []byte(commitContent), 0644); err != nil {
		return fmt.Errorf("error writing commit file: %v", err)
	}

	// update head
	headFile := ".mygit/HEAD"
	if err := ioutil.WriteFile(headFile, []byte(commitHash), 0644); err != nil {
		return fmt.Errorf("error updating HEAD: %v", err)
	}
	return nil
}

// Generates a SHA-1 hash
func generateHash(input string) string {
	h := sha1.New()
	h.Write([]byte(input))
	return hex.EncodeToString(h.Sum(nil))
}

// Builds the commit content
func buildCommitContent(message string, snapshot map[string]string) string {
	var builder strings.Builder
	builder.WriteString("Commit Message: " + message + "\n")
	builder.WriteString("Timestamp: " + time.Now().Format(time.RFC3339) + "\n")

	// Add parent reference
	headFile := ".mygit/HEAD"
	if headData, err := os.ReadFile(headFile); err == nil && len(headData) > 0 {
		parentHash := strings.TrimSpace(string(headData))
		builder.WriteString("Parent: " + parentHash + "\n")
	}

	builder.WriteString("Files:\n")
	for file, content := range snapshot {
		builder.WriteString(fmt.Sprintf("%s\n%s\n", file, content))
	}
	return builder.String()

}
