package pathsize

import (
	"fmt"
	"os"
)

func GetSize(path string) (int64, error) {
	info, err := os.Lstat(path)

	if err != nil {
		return 0, err
	}

	// file
	if !info.IsDir() {
		return info.Size(), nil
	}

	// directory
	children, err := os.ReadDir(path)

	if err != nil {
		return 0, err
	}

	var totalSize int64
	for _, item := range children {
		infoItem, err := item.Info()
		
		if err != nil {
			return 0, fmt.Errorf("failed to get info for %s: %w", item.Name(), err)
		}
		
		if !infoItem.IsDir() {
			totalSize += infoItem.Size()
		}
	}

	return totalSize, nil
}

func GetPathSize(path string) (string, error) {
	size, err := GetSize(path)

	if err != nil {
		return "", fmt.Errorf("failed get size - %s", err)
	}

	return fmt.Sprintf("%dB\t%s", size, path), nil
}