package appxcli_test

import (
	"github.com/drborges/go-ast/models"
	"github.com/drborges/rivers/rx"
	"testing"
)

func TestAPI(t *testing.T) {
	accounts, err := models.FromAccount(c).FindWhereTagsContains("Lol").Select("Name", "Tags", "Token").Distinct().OrderByEmailDesc().All()
	accounts[0].Name = "Borges"
	accounts[0].Save()

	iter := models.FromAccount(c).FindWhereTagsContains("Lol").Select("Name", "Tags", "Token").Distinct().OrderByEmailDesc().PagesIterator()
	for iter.HasNext() {
		var page []*models.Account
		iter.LoadNext(&page)
	}

	account, err := models.FromAccount(c).FindWhereTagsContains("Lol").OrderByEmailDesc().Stream().Each(func(data rx.T) {
		acc := data.(models.AccountAppxModel)
		acc.Email = "LOL"
		acc.Save()
	})

}
