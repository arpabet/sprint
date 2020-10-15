package app

import (
	"fmt"
	"arpabet.pkg.is/preferences"
	"os"
	"path/filepath"
)

func RunningExecutableName() string {
	executable, _ := os.Executable()
	fileName := filepath.Base(executable)
	return fileName
}

func ExecutableRotate(currExecutableName string) string {

	green := fmt.Sprintf(".%s.green", ExecutableName)
	blue := fmt.Sprintf(".%s.blue", ExecutableName)

	switch currExecutableName {
	case green:
		return blue
	case blue:
		return green
	default:
		return blue
	}

}

func ApplicationDir() string {
	if UserDir {
		dir := preferences.AppDataDir(ApplicationName)
		if _, err := os.Stat(dir); err != nil {
			os.MkdirAll(dir, UserDirPerm)
		}
		return dir
	} else {
		return "."
	}
}

func ExecutablePid() string {
	return filepath.Join(ApplicationDir(), ExecutableName + ".pid")
}

func ExecutableData() string {
	return filepath.Join(ApplicationDir(), ExecutableName + "_data")
}

func ExecutableLog() string {
	return filepath.Join(ApplicationDir(), ExecutableName + ".log")
}

