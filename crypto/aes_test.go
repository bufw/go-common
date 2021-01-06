package crypto

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestAes(t *testing.T) {
	rs := AESEncrypt([]byte("hello world"), "hiaes")
	fmt.Println(hex.EncodeToString(rs))
	rs1, _ := AESDecrypt(rs, "hiaes")
	fmt.Println(string(rs1))
}
