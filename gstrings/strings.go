package gstrings

import (
	"strings"
)

func IsUpperASCII(c byte) bool {
	return 'A' <= c && c <= 'Z'
}

func IsLowerASCII(c byte) bool {
	return 'a' <= c && c <= 'z'
}

func ToUpperASCII(c byte) byte {
	if IsLowerASCII(c) {
		return c - ('a' - 'A')
	}
	return c
}

func ToLowerASCII(c byte) byte {
	if IsUpperASCII(c) {
		return c + 'a' - 'A'
	}
	return c
}

func IsEmpty(str string) bool {
	return len(str) == 0
}

func Default(src string, defaultStr string) string {
	if IsEmpty(src) {
		return defaultStr
	}
	return src
}

func HasSuffix(src string, suffixArr ...string) bool {
	for _, suffix := range suffixArr {
		if strings.HasSuffix(src, suffix) {
			return true
		}
	}
	return false
}

func HasPrefix(src string, prefixArr ...string) bool {
	for _, prefix := range prefixArr {
		if strings.HasPrefix(src, prefix) {
			return true
		}
	}
	return false
}

func TrimSuffix(src string, suffixArr ...string) string {
	for _, suffix := range suffixArr {
		src = strings.TrimSuffix(src, suffix)
	}
	return src
}

//TrimPrefix TrimPrefix array
func TrimPrefix(src string, prefixArr ...string) string {
	for _, prefix := range prefixArr {
		src = strings.TrimPrefix(src, prefix)
	}
	return src
}

//ReplaceMap replace
func ReplaceMap(src string, kvMap map[string]interface{}) string {
	for key, val := range kvMap {
		src = strings.ReplaceAll(src, key, val.(string))
	}
	return src
}
