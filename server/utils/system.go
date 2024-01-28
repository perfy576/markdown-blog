package utils

import (
	"path/filepath"
	"runtime"
)

func convertPathToWindows(path string) string {
	if runtime.GOOS != "windows" {
		return path
	}
	windowsPath := filepath.FromSlash(path)

	return windowsPath
}

func ConvertPathToLinux(path string) string {
	// 将路径中的反斜杠转换为正斜杠
	linuxPath := filepath.ToSlash(path)

	return linuxPath
}

func ConverPath(path string) string {
	if runtime.GOOS == "linux" {
		return ConvertPathToLinux(path)
	} else if runtime.GOOS == "windows" {
		return convertPathToWindows(path)
	} else {
		return path
	}
}
