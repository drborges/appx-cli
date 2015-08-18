package models

import "github.com/drborges/appx"

// A comment
type Account struct {
	appx.Model
	Token string
	Name  string
	Email string
	Tags  []string
	Broadcast Broadcast
	Broadcasts []*Broadcast
}

func (model *Account) KeySpec() *appx.KeySpec {
	return &appx.KeySpec{
		Kind:     "Accounts",
		StringID: model.Token,
	}
}
