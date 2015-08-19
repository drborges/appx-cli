package appx_test

import (
	"testing"
	"github.com/drborges/go-ast/models"
)

func TestAPI(t *testing.T) {
	models.FromAccount(c).FindWhereTagsContains("Lol").All()
}
