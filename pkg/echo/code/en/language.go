package en

// 英文

type Code struct {
	SUCCESS             string `code:"1" msg:"operate successfully"`
	Failed              string `code:"0" msg:"operation failure"`
	LoginInvalid        string `code:"204" msg:"You are not logged in or the login fails"`
	ParamsLost          string `code:"104" msg:"Parameters will pass"`
	ParamsError         string `code:"104" msg:"Check the parameters"`
	SystemError         string `code:"944" msg:"system exception"`
	CountryLimitUse     string `code:"1006" msg:"State restrictions on use"`
	VerificationCodeErr string `code:"105" msg:"Check the parameters. The verification code is incorrect"`
	UserAlreadyExists   string `code:"901" msg:"User already exists"`
	RegistrationFailed  string `code:"902" msg:"Registration failed. Please try again later"`
}
