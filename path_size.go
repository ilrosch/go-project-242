package pathsize

import (
	"fmt"
	"os"
)

func GetPathSize(path string, humanReadable bool) (string, error) {
	size, err := GetSize(path)

	if err != nil {
		return "", fmt.Errorf("failed get size - %w", err)
	}

	formatted := FormatSize(size, humanReadable)

	return fmt.Sprintf("%s\t%s", formatted, path), nil
}

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

func FormatSize(size int64, humanReadable bool) string {
	const unit = 1024

	if !humanReadable || size < unit {
		return fmt.Sprint(size, "B")
	}

	units := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	sizeFloat := float64(size)

	var i int
	for sizeFloat >= unit && i < len(units) {
		sizeFloat /= unit
		i += 1
	}

	return fmt.Sprintf("%.1f%s", sizeFloat, units[i])
}
