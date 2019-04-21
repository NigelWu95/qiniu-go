package main

import (
	"fmt"
	"github.com/qiniu/api.v7/auth/qbox"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	api := "http://ai.qiniuapi.com/v1/face/detect"
	// path := "/v1/face/detect?"
	data := "{\"data\": {\"uri\": \"http://xxx.com/xxx.jpg\"}}"

	req, reqErr := http.NewRequest("POST", api, strings.NewReader(data))
	if reqErr != nil {
		return
	}

	req.Header.Add("Content-Type", "application/json")

	mac := qbox.NewMac("", "")
	accessToken, signErr := mac.SignRequestV2(req)
	if signErr != nil {
		return
	}

	req.Header.Add("Authorization", "Qiniu "+accessToken)

	resp, respErr := http.DefaultClient.Do(req)
	if respErr != nil {
		return
	}
	defer resp.Body.Close()

	resData, ioErr := ioutil.ReadAll(resp.Body)
	if ioErr != nil {
		return
	}

	fmt.Println(string(resData))
}
