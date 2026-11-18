package hashs

import (
	"crypto/md5"
	"encoding/hex"
	"os"
)

func MD5(data []byte) string {
	sum := md5.Sum(data)
	return hex.EncodeToString(sum[:])
}

func FileMD5(file string) (h string, err error) {
	byteData, err := os.ReadFile(file)
	if err != nil {
		return
	}
	h = MD5(byteData)
	return
}
