package utils

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// CreateFilePath will create the desired filePath if it does not exist
func CreateFilePath(filePath string) error {
	pathParts := strings.Split(filePath, "/")
	pathBuilder := ""
	for _, part := range pathParts {
		pathBuilder = pathBuilder + part + "/"
		if _, err := os.Stat(pathBuilder); os.IsNotExist(err) {
			if err := os.Mkdir(pathBuilder, fs.ModeAppend); err != nil {
				return err
			}
		}
	}
	return nil
}

// SaveStringToFile will create or update a file with the provided string
func SaveStringToFile(filePath, fileName, data string) error {
	err := CreateFilePath(filePath)
	if err != nil {
		return err
	}

	file, err := os.Create(filepath.Join(filePath, fileName)) // dir is directory where you want to save file.
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(data)

	return err

}
