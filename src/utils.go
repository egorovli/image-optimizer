package main

import (
	"os"
	"path/filepath"
)

func ensureTempDir() (string, error) {
	osTempDir := os.TempDir()
	tempDir := filepath.Join(osTempDir, ".image-optimizer")

	stats, err := os.Stat(tempDir)

	if err != nil && !os.IsNotExist(err) {
		return "", err
	}

	if stats != nil && stats.IsDir() {
		return tempDir, nil
	}

	if err := os.Mkdir(tempDir, os.ModePerm); err != nil {
		return "", err
	}

	return tempDir, nil
}
