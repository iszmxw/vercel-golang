package v1

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"goapi/pkg/echo"
	"log"
	"net/http"
)

type SeoController struct {
	BaseController
}

func (h *SeoController) GetGoogleIndex(c *gin.Context) {
	text := Query()
	echo.Success(c, gin.H{"result": text}, "")
}

func Query() (text string) {
	url := "https://www.google.com/search?q=site%3Adianyingxs.cc&newwindow=1&sca_esv=596722195&sxsrf=AM9HkKlQgpfiKxCbCRr0Ro_LJY8AD3Pz4w%3A1704764009066&source=hp&ei=aaKcZfOPAoWjvr0Pkt2o4A8&iflsig=AO6bgOgAAAAAZZyweTAWI6Jq9nmNTjg3tQbMF8vTqSsg&oq=&gs_lp=Egdnd3Mtd2l6IgAqAggAMgcQIxjqAhgnMgcQIxjqAhgnMgcQIxjqAhgnMgcQIxjqAhgnMgcQIxjqAhgnMgcQIxjqAhgnMgcQIxjqAhgnMgcQIxjqAhgnMgcQIxjqAhgnMgcQIxjqAhgnSMsLUABYAHABeACQAQCYAQCgAQCqAQC4AQHIAQCoAgo&sclient=gws-wiz"
	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("authority", "www.google.com")
	req.Header.Add("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Add("accept-language", "en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7")
	req.Header.Add("cookie", "OTZ=7367255_24_24__24_; SEARCH_SAMESITE=CgQIkJoB; SID=fAiSALFKX5p5UMUSkqruzyPNCO8wCH4T1aiDYdpu5yIyEHVX5AmHNKIDpmeiSr_TJNaPAw.; __Secure-1PSID=fAiSALFKX5p5UMUSkqruzyPNCO8wCH4T1aiDYdpu5yIyEHVXCZ_VFiuadcr2WKN0FaPacg.; __Secure-3PSID=fAiSALFKX5p5UMUSkqruzyPNCO8wCH4T1aiDYdpu5yIyEHVXRZHz0ojxwvRCAhtwskdJ3Q.; HSID=AIVziGX-tV59dK9WQ; SSID=AK_2V8C2j02wx9NOV; APISID=-pp8CL-jPkr2wLrU/AFUyPyPqzooB1gxMb; SAPISID=-u3QRsDI_8i0qFUw/AdHQbKZcWdQCNWvvu; __Secure-1PAPISID=-u3QRsDI_8i0qFUw/AdHQbKZcWdQCNWvvu; __Secure-3PAPISID=-u3QRsDI_8i0qFUw/AdHQbKZcWdQCNWvvu; NID=511=IXYNiyewvNdeuaeRdc-guVWNNnZgc4f7vdsU1MK50WrwZgncipgXrGOQ8oCUw6RQh7qsaFRDpCc1A1yMvGJt6LisbtSkuJs2HpIp_V6xGcRseERfi6nQLmGsZeJgORwlIpSgmlzhzhv33k4wZ57eBMuXlL_dlQdTTbLac_bcLDosQH7iGi_E0vJAWRPPmQdgGWp0Ha2E3T0wXLJkhGb9GzsiHtRFfajnQQxVFhnAM4itIamE7SCrVppyjqy-nRuWAOmj0F1i40jzol0vVcb2VZOWT2TXhtPnfC6fUC2E61V-oVh9feVLyp9LTfjwEQAFms_Pm-Lih1iKCyUKVTW_DUTHmAG3z1M-; AEC=Ackid1TlbuYtgKIGoRLUXt-ITZVmn51CPl2WQUaQTBhdOvteK0Yr8c1KnQ; __Secure-1PSIDTS=sidts-CjABPVxjSpIB3nd9BcUrFFKsgXFV0_0XbeGiIGNxAS3qBP7RNk-cnDuWCzMjkpYggE8QAA; __Secure-3PSIDTS=sidts-CjABPVxjSpIB3nd9BcUrFFKsgXFV0_0XbeGiIGNxAS3qBP7RNk-cnDuWCzMjkpYggE8QAA; 1P_JAR=2024-01-09-02; DV=g8JNgSH2QH9YkAYfDddvFtewOjrAzliAuClUv9L1XQAAAICGSTB--aTyUAAAAOhvYGe9FtKiIAAAALVgeb96t-7vCQAAAA; UULE=a+cm9sZTogMQpwcm9kdWNlcjogMTIKdGltZXN0YW1wOiAxNzA0NzY2NzEyNTc3MDAwCmxhdGxuZyB7CiAgbGF0aXR1ZGVfZTc6IDIyNTg1NjUyMgogIGxvbmdpdHVkZV9lNzogMTEzOTU4NTI2Mwp9CnJhZGl1czogMTE4NTIuNTQKcHJvdmVuYW5jZTogNgo=; SIDCC=ABTWhQFF_fWtATVw2DWr3sEcvof_yKMBeMUVPtP9fvMgF1pC-0LKmCVvELli2piZQ8Zl-OgM1Gk; __Secure-1PSIDCC=ABTWhQGgmhhc5d28mzJJKTjHT2t9v1BYUZ2OEVx28oto1z_tRAUhOLBiYRWowNvBpRJqkIdRnuU; __Secure-3PSIDCC=ABTWhQEGRC6pyQ6NTDfGXagtXTVL0NhzolNhNJYew-mgQu-TTi-rKFu_FCiYLlEe9bZuYiFH4FU")
	req.Header.Add("referer", "https://www.google.com/")
	req.Header.Add("sec-ch-ua", "\"Not_A Brand\";v=\"8\", \"Chromium\";v=\"120\", \"Google Chrome\";v=\"120\"")
	req.Header.Add("sec-ch-ua-arch", "\"arm\"")
	req.Header.Add("sec-ch-ua-bitness", "\"64\"")
	req.Header.Add("sec-ch-ua-full-version", "\"120.0.6099.199\"")
	req.Header.Add("sec-ch-ua-full-version-list", "\"Not_A Brand\";v=\"8.0.0.0\", \"Chromium\";v=\"120.0.6099.199\", \"Google Chrome\";v=\"120.0.6099.199\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-model", "\"\"")
	req.Header.Add("sec-ch-ua-platform", "\"macOS\"")
	req.Header.Add("sec-ch-ua-platform-version", "\"14.0.0\"")
	req.Header.Add("sec-ch-ua-wow64", "?0")
	req.Header.Add("sec-fetch-dest", "document")
	req.Header.Add("sec-fetch-mode", "navigate")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("sec-fetch-user", "?1")
	req.Header.Add("upgrade-insecure-requests", "1")
	req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
	req.Header.Add("x-client-data", "CKq1yQEIkbbJAQiitskBCKmdygEIqJfLAQiVocsBCIagzQEI3L3NAQi6yM0BCI/hzQEIwu7NAQiD8M0BCIXwzQEIofLNARjt080BGKfqzQEYo+zNAQ==")
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	// 使用goquery解析HTML文档
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	// 替换为你要查找的ID
	idToFind := "result-stats"
	// 在文档中查找指定ID的元素
	doc.Find("#" + idToFind).Each(func(i int, s *goquery.Selection) {
		// 获取ID标签的文本内容或其他属性
		text = s.Text()
	})
	return
}
