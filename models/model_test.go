package models_test

import "github.com/drborges/appx"

type Account struct {
	appx.Model
	Token string
	Name  string
	Email string
	Tags  []string
}
