package templates

import (
	"strings"
	"text/template"
)

var FuncMap = template.FuncMap{
	"startsWith":   strings.HasPrefix,
	"sliceElmType": func(elType string) string { return elType[2:] },
}

var AppxRepository = `// AUTO GENERATED BY appx-cli - DO NOT EDIT

package {{.Pkg}}

import (
	"appengine"
	"appengine/datastore"
	"github.com/drborges/appx"
	"github.com/drborges/rivers"
)

func From{{.Name}}(context appengine.Context) *{{.Name}}AppxRepository {
	return &{{.Name}}AppxRepository{
		db: appx.NewDatastore(context),
	}
}

type {{.Name}}AppxRepository struct {
	db *appx.Datastore
}

func (repo *{{.Name}}AppxRepository) New() *{{.Name}}AppxModel {
	return &{{.Name}}AppxModel{
		db: repo.db,
	}
}

func (repo *{{.Name}}AppxRepository) GetByKey(key *datastore.Key) (*{{.Name}}AppxModel, error) {
	item := &{{.Name}}{}
	item.SetKey(key)
	return item, repo.db.Load(item)
}

func (repo *{{.Name}}AppxRepository) GetByEncodedKey(key string) (*{{.Name}}AppxModel, error) {
	item := &{{.Name}}{}
	if err := item.SetEncodedKey(key); err != nil {
		return nil, err
	}
	return item, repo.db.Load(item)
}

{{ $model := . }}

{{range $field, $type := .Fields}}
func (repo *{{$model.Name}}AppxRepository) GetBy{{ $field }}(value {{ $type }}) (*{{$model.Name}}AppxModel, error) {
	item := &{{$model.Name}}{
		{{ $field }}: value,
	}
	return item, repo.db.Load(item)
}
{{end}}

{{range $field, $type := .Fields}}
func (repo *{{ $model.Name }}AppxRepository) FindWhere{{ $field }}(op string, value {{ $type }}) *{{$model.Name}}QueryRunner {
	q := datastore.NewQuery(new({{ $model.Name }}).KeySpec().Kind).Filter("{{ $field }}" + op, value)
	return &{{$model.Name}}QueryRunner{
		db: repo.db,
		q:  q,
	}
}
{{end}}

{{ range $field, $type := .Fields }}
{{ if startsWith $type "[]" }}
func (repo *{{ $model.Name }}AppxRepository) FindWhere{{ $field }}Contains(value {{ sliceElmType $type }}) *{{$model.Name}}QueryRunner {
	q := datastore.NewQuery(new({{ $model.Name }}).KeySpec().Kind).Filter("{{ $field }}=", value)
	return &{{$model.Name}}QueryRunner{
		db: repo.db,
		q:  q,
	}
}
{{ end }}
{{end}}

func (repo *{{.Name}}AppxRepository) FindWhere(filter string, value interface{}) *{{$model.Name}}QueryRunner {
	q := datastore.NewQuery(new({{.Name}}).KeySpec().Kind).Filter(filter, value)
	return &{{$model.Name}}QueryRunner{
		db: repo.db,
		q:  q,
	}
}

func (repo *{{.Name}}AppxRepository) FindWhereAncestorIs(ancestor appx.Entity) *{{$model.Name}}QueryRunner {
	q := datastore.NewQuery(new({{.Name}}).KeySpec().Kind).Ancestor(ancestor.Key())
	return &{{$model.Name}}QueryRunner{
		db: repo.db,
		q:  q,
	}
}

func (repo *{{ .Name }}AppxRepository) FindBy(q *datastore.Query) *{{.Name}}QueryRunner {
	return &{{ .Name }}QueryRunner{
		db: repo.db,
		q:  q,
	}
}

{{range $field, $type := .Fields}}
func (repo *{{ $model.Name }}AppxRepository) FindBy{{ $field }}(value {{ $type }}) *{{$model.Name}}QueryRunner {
	q := datastore.NewQuery(new({{ $model.Name }}).KeySpec().Kind).Filter("{{ $field }}=", value)
	return &{{$model.Name}}QueryRunner{
		db: repo.db,
		q:  q,
	}
}
{{end}}

type {{$model.Name}}QueryRunner struct {
	db *appx.Datastore
	q  *datastore.Query
}

func (runner *{{$model.Name}}QueryRunner) Select(fields ...string) *{{$model.Name}}QueryRunner {
	runner.q = runner.q.Project(fields...)
	return runner
}

func (runner *{{$model.Name}}QueryRunner) Distinct() *{{$model.Name}}QueryRunner {
	runner.q = runner.q.Distinct()
	return runner
}

func (runner *{{$model.Name}}QueryRunner) Limit(limit int) *{{$model.Name}}QueryRunner {
	runner.q = runner.q.Limit(limit)
	return runner
}

func (runner *{{$model.Name}}QueryRunner) KeysOnly() *{{$model.Name}}QueryRunner {
	runner.q = runner.q.KeysOnly()
	return runner
}

{{ range $field, $type := .Fields }}
func (runner *{{$model.Name}}QueryRunner) OrderBy{{ $field }}Asc() *{{$model.Name}}QueryRunner {
	runner.q = runner.q.Order("{{ $field }}")
	return runner
}
{{end}}

{{ range $field, $type := .Fields }}
func (runner *{{$model.Name}}QueryRunner) OrderBy{{ $field }}Desc() *{{$model.Name}}QueryRunner {
	runner.q = runner.q.Order("-{{ $field }}")
	return runner
}
{{end}}

func (runner *{{$model.Name}}QueryRunner) Stream() *rivers.Stage {
	return runner.db.Query(runner.q).StreamOf({{.Name}}{})
}

func (runner *{{$model.Name}}QueryRunner) Count() (int, error) {
	return runner.db.Query(runner.q).Count()
}

func (runner *{{$model.Name}}QueryRunner) All() ([]*{{.Name}}AppxModel, error) {
	items := []*{{.Name}}AppxModel{}
	return items, runner.db.Query(runner.q).Results(items)
}

func (runner *{{$model.Name}}QueryRunner) First() (*{{.Name}}AppxModel, error) {
	item := &{{.Name}}AppxModel{}
	return item, runner.db.Query(runner.q.Limit(1)).Result(item)
}

func (runner *{{$model.Name}}QueryRunner) PagesIterator() appx.Iterator {
	return runner.db.Query(runner.q).PagesIterator()
}

func (runner *{{$model.Name}}QueryRunner) ItemsIterator() appx.Iterator {
	return runner.db.Query(runner.q).ItemsIterator()
}

func (runner *{{$model.Name}}QueryRunner) StartFrom(cursor string) *{{$model.Name}}QueryRunner {
	q := runner.db.Query(runner.q).StartFrom(cursor)
	return &{{$model.Name}}QueryRunner{
		db: runner.db,
		q:  q,
	}
}`
