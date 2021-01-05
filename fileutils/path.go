package fileutils

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

//GetAppExecutablePath 获取App所在目录
func GetAppExecutablePath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return dir
}

func FileNameWithoutExtension(fileName string) string {
	ext := filepath.Ext(fileName)
	if ext == "" {
		return fileName
	}
	return strings.TrimSuffix(fileName, ext)
}
