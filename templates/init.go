package templates

import (
	"embed"
)

//go:embed default
var Default embed.FS

//go:embed statics
var Statics embed.FS
