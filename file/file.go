package file

import "os"

func FileExists(path string) (exist bool, err error) {
	return pathExists(path, false)
}

func RemoveFile(filePath string) (err error) {
	err = os.Remove(filePath)
	return
}

// WriteFile 写入内容到指定文件。
//
// 如果文件已存在，则覆盖原文件内容。
//
// 如果文件不存在，则创建新文件并写入内容。
//
// 参数：
//   - filepath: 文件的路径
//   - content: 需要写入的字节内容
//
// 返回值：
//   - err: 如果写入失败，返回错误信息；否则返回 nil。
func WriteFile(filepath string, content []byte) (err error) {
	err = os.WriteFile(filepath, content, 0644)
	return
}

// WriteFileString 写入内容到指定文件。
//
// 如果文件已存在，则覆盖原文件内容。
//
// 如果文件不存在，则创建新文件并写入内容。
func WriteFileString(filepath, content string) (err error) {
	return WriteFile(filepath, []byte(content))
}

// AppendToFile 追加内容到指定文件。
//
// 如果文件存在，则追加内容；如果文件不存在，则创建新文件并写入内容。
//
// 参数：
//   - filepath: 文件的路径
//   - content: 需要追加的字节内容
//
// 返回值：
//   - err: 如果写入失败，返回错误信息；否则返回 nil。
func AppendToFile(filepath string, content []byte) error {
	// 打开文件，若不存在则创建，使用追加模式
	f, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	// 写入内容到文件末尾
	_, err = f.Write(content)
	return err
}

// AppendToFileString 追加内容到指定文件。
//
// 如果文件存在，则追加内容；如果文件不存在，则创建新文件并写入内容。
func AppendToFileString(filepath, content string) (err error) {
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
