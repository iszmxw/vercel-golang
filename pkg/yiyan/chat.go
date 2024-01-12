package yiyan

import (
	"fmt"
	"goapi/pkg/logger"

	"io/ioutil"
	"net/http"
	"strings"
)

func Chat() {
	clientId := "2M5Ke52SYo0F9Cf5Mgwd28KV"
	clientSecret := "cFBSwwPkP1n5L2UK5NYkn4XGryA1EKVe"
	accessToken := GetAccessToken(clientId, clientSecret)
	url := fmt.Sprintf("https://aip.baidubce.com/rpc/2.0/ai_custom/v1/wenxinworkshop/chat/completions?access_token=%s", accessToken)
	payload := strings.NewReader(`{
    "messages": [
        {
            "role": "user",
            "content": "什么是20日线呢"
        }
    ]
}`)
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, payload)

	if err != nil {
		logger.Error(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		logger.Error(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logger.Error(err)
		return
	}
	fmt.Println(string(body))
}
