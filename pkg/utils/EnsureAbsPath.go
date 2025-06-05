package utils

import (
	"path/filepath"
)

func EnsureAbsPath(fileName string, baseDir string) string {
	if filepath.IsAbs(fileName) {
		return fileName
	}
	return filepath.Join(baseDir, fileName)
}
