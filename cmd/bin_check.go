package cmd

func IsPython3Available() (version string, available bool) {
	return isBinAvailable("python3 --version")
}

func IsPip3Available() (version string, available bool) {
	return isBinAvailable("pip3 --version")
}

func isBinAvailable(command string) (version string, available bool) {
	version, err := ExecShell(command)
	available = err == nil
	return
}
