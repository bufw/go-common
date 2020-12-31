package crypto

import "encoding/base64"

func Encode(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

func Decode(src string) string {
	data, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		return ""
	}
	return string(data)
}
