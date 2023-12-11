package common

// HttpPost 发送post请求
func HttpPost(
	url string,
	header map[string]string,
	data []byte,
) (
	[]byte,
	error,
) {
	body, err := HttpReq("POST", url, header, data)
	if err != nil {
		// fmt.Println(err)
		return nil, err
	}
	return body, nil
}

// func HttpPost(
// 	url string,
// 	header map[string]string,
// 	data []byte,
// ) (
// 	[]byte,
// 	error,
// ) {
// 	req, err := http.NewRequest("POST", url, bytes.NewReader(data))
// 	if err != nil {
// 		// fmt.Println(err)
// 		return nil, err
// 	}

// 	// default Connection: close
// 	_, ok := header["Connection"]
// 	if ok {
// 		//
// 	} else {
// 		req.Header.Set("Connection", header["Connection"])
// 	}

// 	for k, v := range header {
// 		req.Header.Set(k, v)
// 	}

// 	client := &http.Client{}
// 	rsp, err := client.Do(req)
// 	if err != nil {
// 		// fmt.Println(err)
// 		return nil, err
// 	}
// 	defer rsp.Body.Close()
// 	body, err := ioutil.ReadAll(rsp.Body)
// 	if err != nil {
// 		// fmt.Println(err)
// 		return nil, err
// 	}
// 	// fmt.Println(string(body))
// 	return body, nil
// }
