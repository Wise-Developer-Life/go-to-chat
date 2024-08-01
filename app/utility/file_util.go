package utility

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

func SaveFileLocally(file *multipart.FileHeader, destinationPath string) (string, error) {
	// Open the uploaded file
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	// Ensure the destination directory exists
	if err := os.MkdirAll(filepath.Dir(destinationPath), os.ModePerm); err != nil {
		return "", err
	}

	// Create the destination file
	dst, err := os.Create(destinationPath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	// Copy the file content to the destination file
	if _, err := io.Copy(dst, src); err != nil {
		return "", err
	}

	absDestinationPath, err := filepath.Abs(destinationPath)

	if err != nil {
		return "", err
	}

	return absDestinationPath, nil
}
