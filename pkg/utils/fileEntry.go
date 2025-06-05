// file: pkg/utils/FileEntry.go
// (c) 2025 Asymmetric Effort, LLC. <scaldwell@asymmetric-effort.com>

package utils

// FileEntry is used to track each file to process.
type FileEntry struct {
	path    string
	baseDir string
	prefix  string
	isRoot  bool
}

func (fe *FileEntry) GetFile() string {
	return fe.path
}

func (fe *FileEntry) IsRoot() bool {
	return fe.isRoot
}

func (fe *FileEntry) BaseDir() string {
	return fe.baseDir
}

func (fe *FileEntry) Prefix() string {
	return fe.prefix
}

func (fe *FileEntry) SetPrefix(value string) {
	fe.prefix = value
}
