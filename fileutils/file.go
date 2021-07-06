package fileutils

import "os"

func IsSymlink(mode os.FileMode) bool {
	return mode&os.ModeSymlink != 0
}
