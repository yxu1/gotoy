package common

import (
	"fmt"
	"unicode/utf8"
)

// DecodeString 把字符串解析为字符串, 支持中文
func DecodeString(str string) (string, error) {
	decoded := make([]byte, len(str)*4)
	decodedLen := 0

	for i := 0; i < len(str); {
		r, size := utf8.DecodeRuneInString(str[i:])
		if r == utf8.RuneError {
			return "", fmt.Errorf("Invalid UTF-8 encoding")
		}

		decodedLen += utf8.EncodeRune(decoded[decodedLen:], r)
		i += size
	}

	return string(decoded[:decodedLen]), nil
}
