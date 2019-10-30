package main

import (
	"errors"
	"fmt"
	"github.com/qiniu/api.v7/cdn"
	"net/url"
	"strings"
)

var cryptKey string
var cdnScheme string
var videoCDNHost string

func GenerateDownloadURL(keyURL string) (string, error) {

	if keyURL == "" {
		return "", errors.New("invalid source url")
	} else if !strings.HasPrefix(keyURL, "/") {
		keyURL = "/" + keyURL
	}

	duration := int64(3600 * 24)
	accessURL, err := cdn.CreateTimestampAntileechURL(keyURL, cryptKey, duration)
	if err != nil {
		return "", err
	}

	fmt.Println(accessURL)

	u, err := url.Parse(accessURL)
	if err != nil {
		return "", err
	}

	u.Scheme = cdnScheme
	u.Host = videoCDNHost

	return u.String(), nil
}

func main() {
	cryptKey = ""
	cdnScheme = "http"
	videoCDNHost = ""
	singedUrl, err := GenerateDownloadURL("yd/--0HU52H2RU_720p.mp4");
	if err != nil {
		panic(err)
	}
	fmt.Println(singedUrl)
}
