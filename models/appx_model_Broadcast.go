// AUTO GENERATED BY appx-cli - DO NOT EDIT

package models

import (
	"appengine/datastore"
	"github.com/drborges/appx"
)

type BroadcastAppxModel struct {
	Broadcast
	db *appx.Datastore
}

func (model *BroadcastAppxModel) Save() error {
	return model.db.Save(model)
}

func (model *BroadcastAppxModel) Delete() error {
	return model.db.Delete(model)
}
