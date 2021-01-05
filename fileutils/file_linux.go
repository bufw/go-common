package fileutils

import (
	"os"
	"runtime"
	"syscall"
)

func GetFileCreateTime(path string) int64 {
	osType := runtime.GOOS
	fileInfo, _ := os.Stat(path)
	if osType == "linux" {
		stat_t := fileInfo.Sys().(*syscall.Stat_t)
		tCreate := int64(stat_t.Ctim.Sec)
		return tCreate
	}
	return 0
}
