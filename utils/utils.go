package utils

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
)

// RemoveSpecific recursively removes directories or files with a specified name
func RemoveSpecific(root, nameToRemove string) error {
	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.Name() == nameToRemove {
			if err := os.RemoveAll(path); err != nil {
				return fmt.Errorf("failed to remove %s: %v", path, err)
			}
			fmt.Printf("Removed %s\n", path)
		}

		return nil
	})
}

func DecodeBase64ToImage(textPath, imagePath string) error {
	textData, err := os.ReadFile(textPath)
	if err != nil {
		return err
	}
	base64Str := string(textData)

	imageData, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		return err
	}

	return os.WriteFile(imagePath, imageData, 0644)
}

func EncodeImageToBase64(imagePath, textPath string) error {
	imageData, err := os.ReadFile(imagePath)
	if err != nil {
		return err
	}

	encodedStr := base64.StdEncoding.EncodeToString(imageData)
	return os.WriteFile(textPath, []byte(encodedStr), 0644)
}

// ReadFileToLines reads a file from the given path and returns its lines as a slice of strings.
func ReadFileToLines(filePath string) (lines []string, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err = file.Close()
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
