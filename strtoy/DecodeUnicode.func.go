package common

import (
	"strconv"
	"strings"
)

// DecodeUnicode 解析unicode, 把unicode转换为可解析成字符串的字节分片
func DecodeUnicode(raw []byte) ([]byte, error) {
	str, err := strconv.Unquote(strings.Replace(strconv.Quote(string(raw)), `\\u`, `\u`, -1))
	if err != nil {
		return nil, err
	}
	return []byte(str), nil
}
