package insa

// insa

type Code struct {
	SUCCESS             string `code:"1" msg:"操作成功"`
	Failed              string `code:"0" msg:"操作失败"`
	LoginInvalid        string `code:"204" msg:"未登录或者登录失效"`
	ParamsLost          string `code:"104" msg:"参数必传"`
	ParamsError         string `code:"104" msg:"检查参数"`
	SystemError         string `code:"944" msg:"系统异常"`
	CountryLimitUse     string `code:"1006" msg:"Pembatasan daerah Anda"`
	VerificationCodeErr string `code:"105" msg:"Memeriksa parameter, kode verifikasi salah"`
	UserAlreadyExists   string `code:"901" msg:"Pengguna sudah ada"`
	RegistrationFailed  string `code:"902" msg:"Pendaftaran gagal silahkan coba lagi nanti"`
}
