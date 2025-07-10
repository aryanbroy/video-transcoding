package files

import (
	"fmt"
	"os"
)

func FileExists(uploadPath string) (bool, error) {
	info, err := os.Stat(uploadPath)

	if os.IsNotExist(err) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	if info.IsDir() {
		return false, fmt.Errorf("%v is a directory", uploadPath)
	}

	return true, nil
}

func DeleteFile(filePath string) error {
	err := os.RemoveAll(filePath)
	if err != nil {
		return fmt.Errorf("Error deleting file/folder at the specified path: %w", err)
	}
	return nil
}
