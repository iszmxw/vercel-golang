package gotify

import (
	"goapi/pkg/logger"
	"io/ioutil"
	"net/http"
	"net/url"
)

func Notice(title, bodys string) {
	go send("AWxSCZ.uG4nV4zV", title, bodys)
}

func send(token, title, bodys string) {
	form, err := http.PostForm("https://gotify.ethanyep.cn/message?token="+token,
		url.Values{
			"title":   {title},
			"message": {bodys},
		})
	if err != nil {
		logger.Info(err.Error())
		return
	}
	defer form.Body.Close() // 关闭响应体
	body, err := ioutil.ReadAll(form.Body)
	if err != nil {
		logger.Info(err.Error())
		return
	}
	logger.Info(string(body))
}
