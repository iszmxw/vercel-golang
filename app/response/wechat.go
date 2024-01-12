package response

type Wechat struct {
	Account     string `json:"account"`
	Avatar      string `json:"avatar"`
	City        string `json:"city"`
	Country     string `json:"country"`
	LabelIdList string `json:"labelid_list"`
	Nickname    string `json:"nickname"`
	Province    string `json:"province"`
	Remark      string `json:"remark"`
	Sex         int    `json:"sex"`
	Wxid        string `json:"wxid"`
}

type RespWx struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}
