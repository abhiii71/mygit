package core

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

func ParseMyGitIgnore() ([]string, error) {
	var patterns []string

	file, err := os.Open(".mygitignore")
	if os.IsNotExist(err) {
		return patterns, nil
	} else if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" && !strings.HasPrefix(line, "#") {
			patterns = append(patterns, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return patterns, nil
}

func IsIgnored(path string, patterns []string) bool {
	for _, pattern := range patterns {
		matched, _ := filepath.Match(pattern, path)
		if matched {
			return true
		}
	}
	return false
}
