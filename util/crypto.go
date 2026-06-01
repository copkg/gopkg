package util

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(p string) string {
	hash := md5.New()
	hash.Write([]byte(p))
	hashInBytes := hash.Sum(nil)
	return hex.EncodeToString(hashInBytes)
}
func Md5Byte(p string) []byte {
	hash := md5.New()
	hash.Write([]byte(p))
	hashInBytes := hash.Sum(nil)
	return hashInBytes
}
