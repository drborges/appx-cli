// AUTO GENERATED BY appx-cli - DO NOT EDIT

package models

import (
	"appengine/datastore"
	"github.com/drborges/appx"
)

type AccountAppxModel struct {
	Account
	db *appx.Datastore
}

func (model *AccountAppxModel) Save() error {
	return model.db.Save(model)
}

func (model *AccountAppxModel) Delete() error {
	return model.db.Delete(model)
}
