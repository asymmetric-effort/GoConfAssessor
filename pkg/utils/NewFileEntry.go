// file: pkg/utils/NewFileEntry.go
// (c) 2025 Asymmetric Effort, LLC. <scaldwell@asymmetric-effort.com>

package utils

import (
	"fmt"
	"path/filepath"
)

func NewFileEntry(filename *string, isRoot bool) (file FileEntry, err error) {
	if file.path, err = filepath.Abs(*filename); err != nil {
		return file, fmt.Errorf("failed to get absolute path: %w", err)
	}
	file.baseDir = filepath.Dir(file.path)
	file.prefix = ""
	file.isRoot = isRoot
	return file, err
}
