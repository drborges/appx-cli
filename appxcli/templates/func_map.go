package templates

import (
	"strings"
	"text/template"
)

var FuncMap = template.FuncMap{
	"lower":        strings.ToLower,
	"startsWith":   strings.HasPrefix,
	"sliceElmType": func(elType string) string { return elType[2:] },
}
