package core

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

const refsDir = ".mygit/refs/heads"

// ListBranches lists all branches in the repository
func ListBranches() {
	files, err := ioutil.ReadDir(refsDir)
	if err != nil {
		fmt.Println("Error reading branches: ", err)
		return
	}

	fmt.Println("Branches: ")
	for _, file := range files {
		fmt.Printf(" %s\n", file.Name())
	}
}

// CreateBranch creates a new branch
func CreateBranch(name string) error {
	branchPath := filepath.Join(refsDir, name)
	if _, err := os.Stat(branchPath); err == nil {
		return fmt.Errorf("branch %s already exists", name)
	}

	// Get the current HEAD commit hash
	headPath := ".mygit/HEAD"
	headBytes, err := ioutil.ReadFile(headPath)
	if err != nil {
		return fmt.Errorf("failed to read HEAD: %v", err)
	}

	// Write the branch reference pointing to the current commit
	err = ioutil.WriteFile(branchPath, headBytes, 0644)
	if err != nil {
		return fmt.Errorf("failed to create branch '%s': %v", name, err)
	}
	return nil
}
