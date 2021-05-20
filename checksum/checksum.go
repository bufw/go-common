package checksum

import (
	"bufio"
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"hash"
	"io"
	"os"
)

func MD5checksum(filename string) (string, error) {
	return fileChecksum(md5.New(), filename)
}

func SHA256checksum(filename string) (string, error) {
	return fileChecksum(sha256.New(), filename)
}

func fileChecksum(hashAlgorithm hash.Hash, filename string) (string, error) {
	if info, err := os.Stat(filename); err != nil || info.IsDir() {
		return "", err
	}

	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer func() { _ = file.Close() }()

	return sumReader(hashAlgorithm, bufio.NewReader(file))
}

func sumReader(hashAlgorithm hash.Hash, reader io.Reader) (string, error) {
	buf := make([]byte, 65536)
	for {
		switch n, err := reader.Read(buf); err {
		case nil:
			hashAlgorithm.Write(buf[:n])
		case io.EOF:
			return fmt.Sprintf("%x", hashAlgorithm.Sum(nil)), nil
		default:
			return "", err
		}
	}
}
