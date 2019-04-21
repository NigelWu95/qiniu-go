package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io"
	"net/url"
)

func main() {
	var URL = "http://p3l1d5mx4.bkt.clouddn.com/0000021"
	// var URL = "http://"
	var accessKey = ""
	var secretKey = ""
	var saveBucket = "temp"
	var saveKey = "test-2.jpg"

	encodedEntryURI := base64.URLEncoding.EncodeToString([]byte(saveBucket + ":" + saveKey))
	URL += "|saveas/" + encodedEntryURI
	h := hmac.New(sha1.New, []byte(secretKey))
	// 签名内容不包括Scheme，仅含如下部分：
	// <Domain>/<Path>?<Query>
	u, _ := url.Parse(URL)
	_, _ = io.WriteString(h, u.Host+u.RequestURI())
	d := h.Sum(nil)
	sign := accessKey + ":" + base64.URLEncoding.EncodeToString(d)
	fmt.Println(URL + "/sign/" + sign)
}
