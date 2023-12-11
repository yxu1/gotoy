package common

import "strings"

// ConcatString 合并字符串, 用于字符串拼接
func ConcatString(part ...string) (result string) {
	// concat string
	var strBuilder strings.Builder
	for _, x := range part {
		strBuilder.WriteString(x)
	}
	result = strBuilder.String()
	return result
}
