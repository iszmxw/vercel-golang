package templates

import (
	"embed"
	"github.com/CloudyKit/jet/v6"
	"io/fs"
	"strings"
)

//go:embed statics
var Statics embed.FS

//go:embed views
var ViewsFile embed.FS

var Loader = jet.NewInMemLoader()
var Views = jet.NewSet(Loader, jet.InDevelopmentMode())

func LoadTemplates(root string) error {
	return fs.WalkDir(ViewsFile, root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			content, err := ViewsFile.ReadFile(path)
			if err != nil {
				return err
			}
			templatePath := strings.TrimPrefix(path, "views/")
			Loader.Set(templatePath, string(content))
		}
		return nil
	})
}
