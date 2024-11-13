package file

import (
	"errors"
	"io"
	"os"
	"path/filepath"
)

var (
	ErrWantDirGotFile = errors.New("expected directory, but found file")
	ErrWantFileGotDir = errors.New("expected file, but found directory")
)

func DirExists(path string) (exist bool, err error) {
	return pathExists(path, true)
}

func MkdirAll(dir string) (err error) {
	err = os.MkdirAll(filepath.Join(dir), 0755)
	return
}

func RemoveDir(dirPath string) (err error) {
	err = os.RemoveAll(dirPath)
	return
}

func ReadFile(filepath string) (content string, err error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return
	}
	content = string(data)
	return
}

func pathExists(path string, expectDir bool) (exist bool, err error) {
	exist, isDir := fileExists(path)
	if !exist {
		return
	}
	if expectDir && !isDir {
		return false, ErrWantDirGotFile
	}
	if !expectDir && isDir {
		return false, ErrWantFileGotDir
	}
	return
}

// IsDirEmpty 检查目录是否为空
func IsDirEmpty(dir string) (empty bool, err error) {
	empty = false
	f, err := os.Open(dir)
	if err != nil {
		return
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	names, err := f.Readdirnames(1)
	if err != nil {
		if err == io.EOF {
			empty = true
			err = nil
		}
		return
	}
	for _, name := range names {
		if name != ".DS_Store" {
			return
		}
	}
	empty = true
	return
}
