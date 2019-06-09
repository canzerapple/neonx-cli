package templates

import (
	"html/template"

	"github.com/huandu/xstrings"
)

var (
	tpl = template.New("")
)

func init() {
	tpl.Funcs(template.FuncMap{
		"Camel": xstrings.ToCamelCase,
		"Snake": xstrings.ToSnakeCase,
	})
}
