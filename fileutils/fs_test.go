package fileutils

import (
	"os"
	"testing"
)

func TestPermmode(t *testing.T) {
	// io.Sta
	f, _ := os.Open("/etc/profile")
	finfo, _ := f.Stat()

	t.Error(FileModeToString(finfo.Mode()))
	t.Errorf("%o", finfo.Mode())
}
