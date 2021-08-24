package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// 实现双哈希算法
func sha256Double(text string, isHex bool) []byte {
	hashInstance := sha256.New()
	if isHex {
		arr, _ := hex.DecodeString(text)
		hashInstance.Write(arr)
	} else {
		hashInstance.Write([]byte(text))
	}
	bytes := hashInstance.Sum(nil)
	hashInstance.Reset()
	hashInstance.Write(bytes)
	bytes = hashInstance.Sum(nil)
	return bytes
}
func Sha256DoubleString(text string, isHex bool) string {
	bytes := sha256Double(text, isHex)
	return fmt.Sprintf("%x", bytes)
}

func sha256Hash(text string, isHex bool) []byte {
	hashInstance := sha256.New()
	if isHex {
		arr, _ := hex.DecodeString(text)
		hashInstance.Write(arr)
	} else {
		hashInstance.Write([]byte(text))
	}
	bytes := hashInstance.Sum(nil)
	return bytes
}

func Sha256HashString(text string, isHex bool) string {
	bytes := sha256Hash(text, isHex)
	return fmt.Sprintf("%x", bytes)
}
