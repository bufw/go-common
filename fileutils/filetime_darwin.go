// +build darwin

package fileutils

import (
	"os"
	"syscall"
	"time"
)

func GetFileTimeSpec(info os.FileInfo) FileTime {
	stat := info.Sys().(*syscall.Stat_t)
	times := timespec{
		Mtime: info.ModTime(),
		Atime: time.Unix(stat.Atimespec.Sec, stat.Atimespec.Nsec),
		Ctime: time.Unix(stat.Ctimespec.Sec, stat.Ctimespec.Nsec),
	}
	return times
}
