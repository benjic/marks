package cmd

import (
	"os"
	"os/user"
)

const (
	defaultPermissions = 0770
)

func getDataPath() (string, error) {
	u, err := user.Current()

	if err != nil {
		return "", err
	}

	path := u.HomeDir + "/.marks"

	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.MkdirAll(path, defaultPermissions); err != nil {
			return "", err
		}
	}

	return path, nil
}

func getBookmarksFile() (*os.File, error) {
	dataPath, err := getDataPath()
	if err != nil {
		return nil, err
	}

	return os.OpenFile(dataPath+"/bookmarks", os.O_RDWR|os.O_CREATE, defaultPermissions)
}
