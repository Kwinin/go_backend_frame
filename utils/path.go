package utils

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func GetFilePath(filePath string) (path, dir string) {
	path, _ = filepath.Abs(filePath)
	index := strings.LastIndex(path, string(os.PathSeparator))

	return path, path[:index]
}

func GetAppPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))
	return path[:index]
}
