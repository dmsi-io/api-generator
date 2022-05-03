package utils

import "path/filepath"

func RemoveFileExtension(fileName string) string {
	return fileName[:len(fileName)-len(filepath.Ext(fileName))]
}
