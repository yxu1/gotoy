package common

// HttpGet 发送get请求
func HttpGet(
	url string,
	header map[string]string,
) (
	[]byte,
	error,
) {
	body, err := HttpReq("GET", url, header, nil)
	if err != nil {
		// fmt.Println(err)
		return nil, err
	}
	return body, nil
}

// func HttpGet(
// 	url string,
// 	header map[string]string,
// ) (
// 	[]byte,
// 	error,
// ) {
// 	req, err := http.NewRequest("GET", url, nil)
// 	if err != nil {
// 		// fmt.Println(err)
// 		return nil, err
// 	}

// 	if _, ok := header["Connection"]; ok {
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
