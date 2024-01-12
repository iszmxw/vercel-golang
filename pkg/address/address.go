package address

import (
	"encoding/json"
	"errors"
	"fmt"
	cmap "github.com/orcaman/concurrent-map"
	"goapi/app/models"
	"goapi/pkg/config"
	"goapi/pkg/helpers"
	"goapi/pkg/logger"
	"goapi/pkg/mysql"
	"goapi/pkg/redis"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type Address struct {
	Status      string  `json:"status"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	Zip         string  `json:"zip"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Timezone    string  `json:"timezone"`
	Isp         string  `json:"isp"`
	Org         string  `json:"org"`
	As          string  `json:"as"`
	Query       string  `json:"query"`
}

func dataBase() map[string]interface{} {
	data := cmap.New().Items()
	data["countrys"] = "China|Israel"           // 限制国家
	data["timezones"] = "Europe"                // 限制时区
	data["whitelist"] = GetPassLimitHostList(1) // ip白名单
	return data
}

func GetPassLimitHostList(passLevel int) []string {
	var result []string
	mysql.DB.Debug().Model(models.GoAppPassLimitHost{}).Where(map[string]interface{}{
		"is_show":    1,
		"is_delete":  0,
		"pass_level": passLevel,
	}).Select("pass_host").Find(&result)
	return result
}

/**
 * . 提取请求IP的波段
 * @param ip
 * @return
 */

func IpDomainByIp(ip string, num int) string {
	str := ""
	if 0 == len(ip) {
		return str
	}
	nums := strings.Split(ip, ".")
	for index, string_ := range nums {
		if index < num {
			if index == num-1 {
				str = str + string_
			} else {
				str = str + string_ + "."
			}
		}
	}
	return str
}

// 缓存 ip

func IpSet(ip string, isLimit bool) {
	if isLimit {
		_, _ = redis.Client.SelectDbAdd(config.GetInt("redis.db"), fmt.Sprintf("LimitAddress:ip:%v", ip), "true", 60*60*24)
		_, _ = redis.Client.SelectDbAdd(config.GetInt("redis.db"), fmt.Sprintf("LimitAddress:noallow:%v", ip), time.Now().Format("2006-01-02 15:04:05"), 60*60*24)
	} else {
		_, _ = redis.Client.SelectDbAdd(config.GetInt("redis.db"), fmt.Sprintf("LimitAddress:ip:%v", ip), "false", 60*60*24)
		_, _ = redis.Client.SelectDbAdd(config.GetInt("redis.db"), fmt.Sprintf("LimitAddress:allow:%v", ip), time.Now().Format("2006-01-02 15:04:05"), 60*60*24)
	}
}

// LimitAddress 地址限制
func LimitAddress(ip string) bool {
	if len(ip) == 0 {
		logger.Error(errors.New("未读取到ip"))
		return false
	}
	if ip == "127.0.0.1" {
		return false
	}
	data := dataBase()
	if helpers.InArray(ip, data["whitelist"].([]string)) {
		return false
	}
	if helpers.InArray(IpDomainByIp(ip, 2), data["whitelist"].([]string)) {
		return false
	}
	ok, _ := redis.Client.SelectDbGet(config.GetInt("redis.db"), fmt.Sprintf("LimitAddress:ip:%v", ip))
	if ok == "true" {
		return true
	}
	if ok == "false" {
		return false
	}
	res := GetMessageByIp(ip)
	/* 异常请求，但是为了不影响正常操作，让它过 */
	if len(res.Country) == 0 {
		logger.Error(errors.New("异常请求，但是为了不影响正常操作，让它过"))
		return false
	}
	/* 限制国家 */
	if strings.Contains(data["countrys"].(string), res.Country) { // 缓存24小时
		IpSet(ip, true)
		return true
	} else {
		IpSet(ip, false)
	}
	/* 限制时区是欧洲且不是英国的地区 */
	if strings.Contains(data["timezones"].(string), res.Timezone) && !strings.EqualFold("United Kingdom", res.Country) { // 缓存24小时
		IpSet(ip, true)
		return true
	}
	return false
}

/*
  - {
    "status": "success",
    "country": "Indonesia",
    "countryCode": "ID",
    "region": "SU",
    "regionName": "North Sumatra",
    "city": "Medan",
    "zip": "",
    "lat": 3.5847,
    "lon": 98.6629,
    "timezone": "Asia/Jakarta",
    "isp": "SMART-TELECOM",
    "org": "",
    "as": "AS18004 PT WIRELESS INDONESIA ( WIN )",
    "query": "114.79.56.80"
    }
    *
*/
func GetMessageByIp(ip string) Address {
	var address Address
	url := fmt.Sprintf("http://ip-api.com/json/%v", ip)
	Client := new(http.Client)
	response, err := Client.Get(url)
	if err != nil {
		logger.Error(err)
		return address
	}
	if response.StatusCode != 200 {
		return address
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logger.Error(err)
		return address
	}
	err = json.Unmarshal(body, &address)
	if err != nil {
		logger.Error(err)
	}
	return address
}
