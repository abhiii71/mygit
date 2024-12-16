package core

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fatih/color"
)

func Status() error {
	if _, err := os.Stat(".mygit"); os.IsNotExist(err) {
		return fmt.Errorf("not a MyGit repository (or any of the parent directories): .mygit")
	}

	patterns, err := ParseMyGitIgnore()
	if err != nil {
		return fmt.Errorf("error parsing .mygitignore: %v", err)
	}

	trackedFiles, err := GetTrackedFiles()
	if err != nil {
		return fmt.Errorf("error reading tracked files: %v", err)
	}

	untrackedFiles, err := getUntrackedFiles(trackedFiles, patterns)
	if err != nil {
		return fmt.Errorf("error reading untracked files: %v", err)
	}

	green := color.New(color.FgGreen).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()

	fmt.Println("Tracked Files:")
	for _, file := range trackedFiles {
		fmt.Println(" ", green(file))
	}

	fmt.Println("Untracked Files:")
	for _, file := range untrackedFiles {
		fmt.Println(" ", red(file))
	}

	return nil
}

func getUntrackedFiles(trackedFiles []string, patterns []string) ([]string, error) {
	var untracked []string
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() || path == ".mygit" || filepath.HasPrefix(path, ".mygit/") {
			return nil
		}

		if !isTracked(path, trackedFiles) && !IsIgnored(path, patterns) {
			untracked = append(untracked, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return untracked, nil
}

/*
func isTracked(path string, trackedFiles []string) bool {
	for _, tracked := range trackedFiles {
		if tracked == path {
			return true
		}
	}
	return false
}
*/
