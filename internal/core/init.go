package core

import (
	"errors"
	"fmt"
	"os"
)

func InitializeRepo() error {
	const mygitDir = ".mygit"

	if _, err := os.Stat(mygitDir); !os.IsNotExist(err) {
		return errors.New("repository already initialized")
	}

	if err := os.Mkdir(mygitDir, 0755); err != nil {
		return fmt.Errorf("failed to create .mygit directory: %v", err)
	}

	// Create refs/heads directory
	refsPath := mygitDir + "/refs/heads"
	if err := os.MkdirAll(refsPath, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create refs/heads directory: %v", err)
	}

	// create object directory
	objectPath := mygitDir + "/objects"
	if err := os.MkdirAll(objectPath, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create objects directory: %v", err)
	}

	headFile := mygitDir + "/HEAD"
	headContent := []byte("ref: refs/heads/main\n")
	if err := os.WriteFile(headFile, headContent, 0644); err != nil {
		return fmt.Errorf("failed to create HEAD file: %v", err)
	}

	mainBranchFile := refsPath + "/main"
	if err := os.WriteFile(mainBranchFile, []byte{}, 0644); err != nil {
		return fmt.Errorf("failed to create default branch file: %v", err)
	}

	fmt.Println("Initialized empty MyGit repository in", mygitDir)

	return nil
}
