package yiyan

import (
	"encoding/json"
	"fmt"
	"goapi/pkg/logger"
	"goapi/pkg/redis"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetAccessToken(clientId, clientSecret string) string {
	key := "yiyan:accessToken"
	accessToken, _ := redis.Client.Get(key)
	payload := strings.NewReader(``)
	client := &http.Client{}
	url := fmt.Sprintf("https://aip.baidubce.com/oauth/2.0/token?client_id=%s&client_secret=%s&grant_type=client_credentials", clientId, clientSecret)
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		logger.Error(err)
		return accessToken
	}
	res, err := client.Do(req)
	if err != nil {
		logger.Error(err)
		return accessToken
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logger.Error(err)
		return accessToken
	}
	var resp AccessToken
	err = json.Unmarshal(body, &resp)
	if err != nil {
		logger.Error(err)
		return accessToken
	}
	_, _ = redis.Client.Add(key, resp.AccessToken, resp.ExpiresIn)
	return resp.AccessToken
}
