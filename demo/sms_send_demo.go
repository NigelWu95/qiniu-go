package main

import (
	"fmt"
	"github.com/qiniu/api.v7/auth"
	"github.com/qiniu/api.v7/sms"
)

func main() {

	accessKey := ""
	secretKey := ""
	mac := auth.New(accessKey, secretKey)
	manager := sms.NewManager(mac)

	// SendMessage
	args := sms.MessagesRequest{
		SignatureID: "",
		TemplateID:  "",
		Mobiles:     []string{""},
		Parameters: map[string]interface{}{
			"code": 111233,
		},
	}

	// 返回结果
	ret, err := manager.SendMessage(args)
	if err != nil {
		fmt.Println("发送短信错误：" + err.Error())
		return
	}
	if len(ret.JobID) == 0 {
		fmt.Println("发送短信错误： The job id cannot be empty")
		return
	}
}