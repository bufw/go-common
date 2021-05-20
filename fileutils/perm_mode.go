package fileutils

import (
	"fmt"
	"io/fs"
)

func FileModeToString(fMode fs.FileMode) string {
	// return fMode.String()
	return fmt.Sprintf("%v", fMode)
}

func FileModeToInt(fMode fs.FileMode) string {
	return fmt.Sprintf("%o", fMode)
}
