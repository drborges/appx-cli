package appx_test

import (
	"testing"
	"github.com/drborges/go-ast/models"
)

func TestAPI(t *testing.T) {
	accounts, err := models.FromAccount(c).FindWhereTagsContains("Lol").Select("Email", "Token").Distinct().All()
	accounts[0].Save()
}
