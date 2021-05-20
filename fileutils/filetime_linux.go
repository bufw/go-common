package fileutils

import (
	"os"
	"syscall"
	"time"
)

func GetFileTimeSpec(info os.FileInfo) FileTime {
	stat := info.Sys().(*syscall.Stat_t)
	times := FileTime{
		Mtime: info.ModTime(),
		Atime: time.Unix(int64(stat.Atim.Sec), int64(stat.Atim.Nsec)),
		Ctime: time.Unix(int64(stat.Ctim.Sec), int64(stat.Ctim.Nsec)),
	}
	return times
}
