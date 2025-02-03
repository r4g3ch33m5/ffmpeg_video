package util

import (
	"os"
	"path/filepath"
)

func ApplyAllFileInDir(dir string, do func(path string) error) error {
	return filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		// Check if it's a file (not a directory)
		if !d.IsDir() {
			do(path)
		}

		return nil
	})
}
