package common

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

// HttpReq 发送请求
func HttpReq(
	method string,
	url string,
	header map[string]string,
	data []byte,
) (
	body []byte,
	err error,
) {
	var req *http.Request
	// var payload *bytes.Reader
	if method == "GET" {
		// payload = nil
		req, err = http.NewRequest(method, url, nil)
	} else {
		// payload = bytes.NewReader(data)
		req, err = http.NewRequest(method, url, bytes.NewReader(data))
	}
	// req, err = http.NewRequest(method, url, payload)
	if err != nil {
		// fmt.Println(err)
		return nil, err
	}

	// default Connection: close
	_, ok := header["Connection"]
	if ok {
		//
	} else {
		req.Header.Set("Connection", header["Connection"])
	}

	for k, v := range header {
		req.Header.Set(k, v)
	}

	client := &http.Client{}
	rsp, err := client.Do(req)
	if err != nil {
		// fmt.Println(err)
		return nil, err
	}
	defer rsp.Body.Close()
	body, err = ioutil.ReadAll(rsp.Body)
	if err != nil {
		// fmt.Println(err)
		return nil, err
	}
	// fmt.Println(string(body))
	return body, nil
}
