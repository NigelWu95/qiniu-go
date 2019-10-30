package main

import (
	"errors"
	"fmt"
	"github.com/qiniu/api.v7/cdn"
	"net/url"
	"strings"
	"time"
)

var cryptKey string
var cdnScheme string
var CDNHost string

func GenerateDownloadURL(keyURL string) (string, error) {

	if keyURL == "" {
		return "", errors.New("invalid source url")
	} else if !strings.HasPrefix(keyURL, "/") {
		keyURL = "/" + keyURL
	}

	duration := int64(time.Second * 3600 * 24 / time.Millisecond)
	accessURL, err := cdn.CreateTimestampAntileechURL(keyURL, cryptKey, duration)
	if err != nil {
		return "", err
	}

	u, err := url.Parse(accessURL)
	if err != nil {
		return "", err
	}

	u.Scheme = cdnScheme
	u.Host = CDNHost

	return u.String(), nil
}

func main() {

	cryptKey = ""
	cdnScheme = "http"
	CDNHost = ""
	// 签名时参数只传文件名部分
	singedUrl, err := GenerateDownloadURL("yd/e8CLsYzE5wk_720p.mp4");
	if err != nil {
		panic(err)
	}
	fmt.Println(singedUrl)
}
