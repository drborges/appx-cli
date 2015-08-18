package models

import "github.com/drborges/appx"

type Broadcast struct {
	appx.Model
	URL    string
	Length int
}

func (model *Broadcast) KeySpec() *appx.KeySpec {
	return &appx.KeySpec{
		Kind:       "Broadcasts",
		HasParent:  true,
		Incomplete: true,
	}
}
