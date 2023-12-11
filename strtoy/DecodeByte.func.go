package common

import (
	"fmt"
	"unicode/utf8"
)

// DecodeByte 把byte解析为字符, 支持中文
func DecodeByte(bytes []byte) (string, error) {
	decoded := make([]byte, len(bytes)*4)
	decodedLen := 0

	for len(bytes) > 0 {
		r, size := utf8.DecodeRune(bytes)
		if r == utf8.RuneError {
			return "", fmt.Errorf("无效的UTF-8编码")
		}

		decodedLen += utf8.EncodeRune(decoded[decodedLen:], r)
		bytes = bytes[size:]
	}

	return string(decoded[:decodedLen]), nil
}
