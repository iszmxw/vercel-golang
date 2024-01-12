package email

type Params struct {
	From        string   `json:"From"`        // 发件人邮箱
	To          string   `json:"To"`          // 收件人邮箱
	Tos         []string `json:"Tos"`         // 收件人邮箱批量
	Subject     string   `json:"Subject"`     // 邮件标题
	Body        string   `json:"Body"`        // 邮件内容
	ContentType string   `json:"ContentType"` // 类型 text/html ，text/plain
	Host        string   `json:"host"`        // 邮件服务器
	Port        int      `json:"Port"`        // 邮件端口
	Username    string   `json:"Username"`    // 用户名
	Password    string   `json:"Password"`    // 密码
}
