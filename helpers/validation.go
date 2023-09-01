package helpers

import (
	"fmt"
	"os"
)

func ValidateDirectory(path string) (bool, error) {
	dirInfo, err := os.Stat(path)

	if err != nil {
		return false, err
	}

	if !dirInfo.IsDir() {
		return false, fmt.Errorf("The specified path is not a directory")
	}

	return true, nil
}

func ValidateFile(path string) (bool, error) {
	fileInfo, err := os.Stat(path)

	if err != nil {
		return false, err
	}

	if fileInfo.IsDir() {
		return false, fmt.Errorf("The specified path is a directory, not a file")
	}

	return true, nil

}
