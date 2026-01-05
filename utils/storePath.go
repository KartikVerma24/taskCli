package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func ResolveStorePath(dir string) (string, error) {
	const fileName = "taskStore.json"

	if dir == "" {
		cwd, err := os.Getwd()
		if err != nil {
			return "", err
		}
		return filepath.Join(cwd, fileName), nil
	}

	info, err := os.Stat(dir)
	if err != nil {
		return "", err
	}
	if !info.IsDir() {
		return "", fmt.Errorf("store path must be a directory")
	}

	return filepath.Join(dir, fileName), nil
}
