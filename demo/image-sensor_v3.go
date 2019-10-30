package main

import (
	"fmt"
	"github.com/qiniu/api.v7/auth/qbox"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	data := "{\"data\": {\"uri\": \"http://p3l1d5mx4.bkt.clouddn.com/0000021\"}, \"params\":{\"scenes\":[\"pulp\"]}}"

	api := "http://ai.qiniuapi.com/v3/image/censor"

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

	req.Header.Add("Authorization", "Qiniu " + accessToken)

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
