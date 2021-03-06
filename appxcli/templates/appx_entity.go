package templates

var AppxEntity = `// AUTO GENERATED BY appx-cli - DO NOT EDIT

package {{.Pkg}}

import (
	{{ if .Queryable }}
	"appengine/datastore"
	{{ end }}
	"github.com/drborges/appx"
)

type {{.Name}} struct {
	appx.Model
	{{ range $field, $type := .Fields }}
	{{ $field }} {{ $type }}
	{{ end }}
}

func ({{lower .Name}} *{{.Name}}) KeySpec() *appx.KeySpec {
	return &appx.KeySpec{
		Kind: "{{ .Kind }}",
		{{ if .Incomplete }}
		Incomplete: true,
		{{ end }}
		{{ if .IntID }}
		IntID: {{lower .Name}}.{{ .ID }},
		{{ end }}
		{{ if .StringID }}
		StringID: {{lower .Name}}.{{ .ID }},
		{{ end }}
		{{ if .HasParent }}
		HasParent: true,
		{{ end }}
	}
}

{{ if .Cacheable }}
func ({{lower .Name}} *{{.Name}}) CacheID() string {
	return {{lower .Name}}.{{ .CacheBy }}
}
{{ end }}

{{ if .Queryable }}
func ({{ lower .Name }} *{{ .Name }}) Query() *datastore.Query {
	return datastore.NewQuery({{ lower .Name }}.KeySpec().Kind).Filter("{{ .QueryBy }}=", {{ lower .Name}}.{{ .QueryBy }})
}
{{ end }}`
