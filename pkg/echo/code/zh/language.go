package zh

// 中文

type Code struct {
	SUCCESS             string `code:"1" msg:"操作成功"`
	Failed              string `code:"0" msg:"操作失败"`
	LoginInvalid        string `code:"204" msg:"未登录或者登录失效"`
	ParamsLost          string `code:"104" msg:"参数必传"`
	ParamsError         string `code:"104" msg:"检查参数"`
	SystemError         string `code:"944" msg:"系统异常"`
	CountryLimitUse     string `code:"1006" msg:"国家限制使用"`
	VerificationCodeErr string `code:"105" msg:"检查参数,验证码不正确"`
	UserAlreadyExists   string `code:"901" msg:"用户已存在"`
	RegistrationFailed  string `code:"902" msg:"注册失败请稍后再试"`
}
