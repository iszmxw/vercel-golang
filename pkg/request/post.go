package request

import (
	"net/http"
	"net/url"
	"strings"
)

func Post(urlAddress string, Header http.Header, body string, proxy *url.URL) (map[string]interface{}, http.Header, error) {
	req, err := http.NewRequest("POST", urlAddress, strings.NewReader(body))
	if err != nil {
		return nil, nil, err
	}
	req.Header = Header
	client := &http.Client{
		Transport: &http.Transport{
			// 设置代理
			Proxy: http.ProxyURL(proxy),
		},
	}
	response, err := client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	// 解析Response
	return ParseResponse(response)

}
