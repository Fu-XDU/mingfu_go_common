package file

import "os"

func FileExists(path string) (exist bool, err error) {
	return pathExists(path, false)
}

func RemoveFile(filePath string) (err error) {
	err = os.Remove(filePath)
	return
}

func WriteFile(filepath string, content []byte) (err error) {
	err = os.WriteFile(filepath, content, 0644)
	return
}

func WriteFileString(filepath, content string) (err error) {
	return WriteFile(filepath, []byte(content))
}

func fileExists(path string) (exist bool, isDir bool) {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return
		}
		return
	}
	exist = true
	isDir = info.IsDir()
	return
}
