package fileutils

import (
	"time"
)

type FileTime struct {
	Mtime time.Time
	Atime time.Time
	Ctime time.Time
}
