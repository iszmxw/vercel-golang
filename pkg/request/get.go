package request

import (
	"net/http"
	"net/url"
)

func Get(url string, Header http.Header, proxy *url.URL) (map[string]interface{}, http.Header, error) {
	req, err := http.NewRequest("GET", url, nil)
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
