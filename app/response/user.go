package response

/** 用户登录数据 **/

type ClientUserBeans struct {
	AccountEmail            string `json:"accountEmail"`
	AppUserId               string `json:"appUserId"`
	BankCredit              int    `json:"bankCredit"`
	BonusBalance            string `json:"bonusBalance"`
	CurrencyType            string `json:"currencyType"`
	EmailCredit             int    `json:"emailCredit"`
	HadRechargeAmounts      string `json:"hadRechargeAmounts"`
	HadWithdrawAmounts      string `json:"hadWithdrawAmounts"`
	IdentityCredit          int    `json:"identityCredit"`
	IsFinishCredit          int    `json:"isFinishCredit"`
	IsNew                   int    `json:"isNew"`
	MobileCredit            int    `json:"mobileCredit"`
	Nation                  string `json:"nation"`
	RealBalance             string `json:"realBalance"`
	RechargeAmounts         string `json:"rechargeAmounts"`
	TotalHisBonusAmounts    string `json:"totalHisBonusAmounts"`
	TotalHisRechargeAmounts string `json:"totalHisRechargeAmounts"`
	TotalHisWithdrawAmounts string `json:"totalHisWithdrawAmounts"`
	TradingAmounts          string `json:"tradingAmounts"`
	UserLevel               int    `json:"userLevel"`
	VirtualBalance          string `json:"virtualBalance"`
	WithdrawAmounts         string `json:"withdrawAmounts"`
}
