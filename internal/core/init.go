package core

import (
	"errors"
	"os"
)

func InitializeRepo() error {
	const mygitDir = ".mygit"

	if _, err := os.Stat(mygitDir); !os.IsNotExist(err) {
		return errors.New("repository already initialized")
	}

	err := os.Mkdir(mygitDir, 0755)
	if err != nil {
		return err
	}

	headFile := mygitDir + "/HEAD"
	err = os.WriteFile(headFile, []byte("ref: refs/head/main\n"), 0644)
	if err != nil {
		return err
	}
	return nil
}
