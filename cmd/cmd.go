package cmd

import (
	"os/exec"
	"strings"
)

// ExecShell 阻塞式的执行外部shell命令的函数,等待执行完毕并返回标准输出
func ExecShell(command string) (string, error) {
	return ExecShellAt(command, "")
}

func ExecShellAt(command, dir string) (string, error) {
	// 函数返回一个*Cmd，用于使用给出的参数执行name指定的程序
	cmd := exec.Command("/bin/bash", "-c", command)
	cmd.Dir = dir
	output, err := cmd.CombinedOutput()

	return strings.TrimSpace(string(output)), err
}

func ExecShellFile(shFilePath string) (string, error) {
	cmd := exec.Command("/bin/bash", shFilePath)
	output, err := cmd.CombinedOutput()

	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), err
}
