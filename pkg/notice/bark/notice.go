package bark

import (
	"fmt"
	"goapi/pkg/config"
	"goapi/pkg/logger"
	"io/ioutil"
	"net/http"
	"net/url"
)

func icon() string {
	images := config.GetString("bark.logo")
	return images
}

func Notice(title, bodys string) {
	go send(config.GetString("bark.key.mac"), title, bodys)    // 发往 mac
	go send(config.GetString("bark.key.iphone"), title, bodys) // 发往 iPhone
}

func send(appKey, title, bodys string) {
	baseURL := config.GetString("bark.url") + url.QueryEscape(appKey)
	urls := fmt.Sprintf("%s/%s/%s?icon=%s", baseURL, url.QueryEscape(title), url.QueryEscape(bodys), icon())
	res, err := http.Get(urls)
	if err != nil {
		logger.Info(err.Error())
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logger.Info(err.Error())
		return
	}
	logger.Info(string(body))
}
