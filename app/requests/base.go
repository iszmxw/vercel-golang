package requests

type Base struct {
	ItemCode     string `json:"itemCode" validate:"required"`     // 项目编码
	Nation       string `json:"nation" validate:"required"`       // 国家编码
	LanguageCode string `json:"languageCode" validate:"required"` // 语言编码
}
