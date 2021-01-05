package crypto

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
)

//Hash Type
const (
	HashMD5    = 0
	HashSHA1   = 1
	HashSHA256 = 3
	//384
	HashSHA512 = 6
)

//MD5 md5 hash
func MD5(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

//SHA1 SHA1 hash
func SHA1(data string) string {
	h := sha1.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

//SHA256  SHA256 hash
func SHA256(data string) string {
	h := sha256.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

//SHA512  SHA512 hash
func SHA512(data string) string {
	h := sha512.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

//Hash supports md5,sha1,sha256,sha512
func Hash(b []byte, hashType int) []byte {
	switch hashType {
	case HashMD5:
		h := md5.New()
		h.Write(b)
		return h.Sum(nil)
	case HashSHA1:
		h := sha1.New()
		h.Write(b)
		return h.Sum(nil)
	case HashSHA256:
		h := sha256.New()
		h.Write(b)
		return h.Sum(nil)
	case HashSHA512:
		h := sha512.New()
		h.Write(b)
		return h.Sum(nil)
	default:
		return nil
	}

}
