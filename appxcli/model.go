package appxcli

type Model struct {
	Name       string
	Pkg        string
	Kind       string
	Incomplete bool
	IntID      bool
	StringID   bool
	ID         string
	HasParent  bool
	Cacheable  bool
	Queryable  bool
	CacheBy    string
	QueryBy    string
	Fields     map[string]string
}
