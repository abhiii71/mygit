package core

import (
	"fmt"
	"os"
	"path/filepath"
)

func Status() error {
	if _, err := os.Stat(".mygit"); os.IsNotExist(err) {
		return fmt.Errorf("not a MyGit repository (or any of the parent directories): .mygit")
	}
	trackedFiles, err := getTrackedFiles()
	if err != nil {
		return fmt.Errorf("error reading tracked files: %v", err)
	}

	untrackedFiles, err := getUntrackedFiles(trackedFiles)
	if err != nil {
		return fmt.Errorf("error reading untracked files: %v", err)
	}

	fmt.Println("Tracked Files:")
	for _, file := range trackedFiles {
		fmt.Println(" ", file)
	}

	fmt.Println("Untracked Files:")
	for _, file := range untrackedFiles {
		fmt.Println(" ", file)
	}

	return nil
}

func getTrackedFiles() ([]string, error) {
	return []string{"example_tracked_file.txt"}, nil
}

func getUntrackedFiles(trackedFiles []string) ([]string, error) {
	var untracked []string
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() || path == ".mygit" || filepath.HasPrefix(path, ".mygit/") {
			return nil
		}

		if !isTracked(path, trackedFiles) {
			untracked = append(untracked, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return untracked, nil
}

func isTracked(path string, trackedFiles []string) bool {
	for _, tracked := range trackedFiles {
		if tracked == path {
			return true
		}
	}
	return false
}
