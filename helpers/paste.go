package helpers

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func PasteDirectory(source, destination string) error {
	sourceInfo, err := os.Stat(source)
	if err != nil {
		return err
	}

	if !sourceInfo.IsDir() {
		return fmt.Errorf("Source is not a directory")
	}

	if err := os.MkdirAll(destination, os.ModePerm); err != nil {
		return err
	}

	files, err := os.ReadDir(source)
	if err != nil {
		return err
	}

	for _, file := range files {
		srcPath := filepath.Join(source, file.Name())
		destPath := filepath.Join(destination, file.Name())

		if file.IsDir() {
			if err := PasteDirectory(srcPath, destPath); err != nil {
				return err
			}
		} else {
			if err := PasteFile(srcPath, destPath); err != nil {
				return err
			}
		}
	}

	return nil
}
func PasteFile(source, destination string) error {
	srcFile, err := os.Open(source)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	destFile, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)
	return err
}
