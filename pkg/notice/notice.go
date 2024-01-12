package notice

import (
	"fmt"
	"goapi/pkg/config"
	"goapi/pkg/notice/bark"
	"goapi/pkg/notice/gotify"
)

func Notice(title, bodys string) {
	name := config.GetString("app.name")
	title = fmt.Sprintf("【%v】---- %v", name, title)
	go bark.Notice(title, bodys)
	go gotify.Notice(title, bodys)
}

func Error(title, bodys string) {
	name := config.GetString("app.name")
	title = fmt.Sprintf("【%v】---- %v", name, title)
	go bark.Notice(title, bodys)
	go gotify.Notice(title, bodys)
}
