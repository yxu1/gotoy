package common

import (
	// "os"
	// "fmt"
	// "io/ioutil"
	"net/http"
)

// HttpHead 发送head请求
func HttpHead(
	url string,
) (
	int,
	error,
) {
	rsp, err := http.Head(url)
	if err != nil {
		return 0, err
	}
	return rsp.StatusCode, nil
}
