package bruno

import (
	"os"
	"path/filepath"
)

func GetFilePath(slug string, fileName string) string {
	dir := filepath.Join("dist", slug)
	path := filepath.Join(dir, fileName)

	return path
}

func CheckIfFileExists(path string) (bool, error) {
	// check if the filename already exists
	if _, err := os.Stat(path); err == nil {
		return true, nil
	}
	return false, nil
}

func CreateDirIfFileNotExists(path string) error {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	return nil
}
