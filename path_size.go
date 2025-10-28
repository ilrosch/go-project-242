package pathsize

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GetPathSize(path string, recursive, human, all bool) (string, error) {
	size, err := GetSize(path, recursive, all)

	if err != nil {
		return "", fmt.Errorf("failed get size - %w", err)
	}

	formatted := FormatSize(size, human)

	return fmt.Sprintf("%s\t%s", formatted, path), nil
}

func GetSize(path string, recursive, all bool) (int64, error) {
	var totalSize int64
	queue := []string{path}

	for len(queue) != 0 {
		curPath := queue[0]
		queue = queue[1:]

		info, err := os.Lstat(curPath)
		if err != nil {
			return 0, err
		}

		// file
		if !info.IsDir() {
			if !strings.HasPrefix(info.Name(), ".") || all {
				totalSize += info.Size()
			}
			continue
		}

		// directory
		children, err := os.ReadDir(curPath)
		if err != nil {
			return 0, err
		}

		for _, item := range children {
			newPath := filepath.Join(curPath, item.Name())
			
			if item.IsDir() {
				if recursive {
					queue = append(queue, newPath)
				}
			} else {
				if !strings.HasPrefix(item.Name(), ".") || all {
					fileInfo, err := item.Info()
					if err != nil {
						return 0, err
					}
					totalSize += fileInfo.Size()
				}
			}
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
