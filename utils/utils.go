package utils

import (
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
