// AUTO GENERATED BY appx-cli - DO NOT EDIT

package models

import (
	"appengine"
	"appengine/datastore"
	"github.com/drborges/appx"
	"github.com/drborges/rivers"
)

func FromAccount(context appengine.Context) *AccountAppxRepository {
	return &AccountAppxRepository{
		db: appx.NewDatastore(context),
	}
}

type AccountAppxRepository struct {
	db *appx.Datastore
}

func (repo *AccountAppxRepository) New() *AccountAppxModel {
	return &AccountAppxModel{
		db: repo.db,
	}
}

func (repo *AccountAppxRepository) GetByKey(key *datastore.Key) (*AccountAppxModel, error) {
	item := &Account{}
	item.SetKey(key)
	return item, repo.db.Load(item)
}

func (repo *AccountAppxRepository) GetByEncodedKey(key string) (*AccountAppxModel, error) {
	item := &Account{}
	if err := item.SetEncodedKey(key); err != nil {
		return nil, err
	}
	return item, repo.db.Load(item)
}




func (repo *AccountAppxRepository) GetByBroadcast(value Broadcast) (*AccountAppxModel, error) {
	item := &Account{
		Broadcast: value,
	}
	return item, repo.db.Load(item)
}

func (repo *AccountAppxRepository) GetByEmail(value string) (*AccountAppxModel, error) {
	item := &Account{
		Email: value,
	}
	return item, repo.db.Load(item)
}

func (repo *AccountAppxRepository) GetByName(value string) (*AccountAppxModel, error) {
	item := &Account{
		Name: value,
	}
	return item, repo.db.Load(item)
}

func (repo *AccountAppxRepository) GetByToken(value string) (*AccountAppxModel, error) {
	item := &Account{
		Token: value,
	}
	return item, repo.db.Load(item)
}



func (repo *AccountAppxRepository) FindWhereBroadcast(op string, value interface{}) *AccountQueryRunner {
	q := datastore.NewQuery(new(Account).KeySpec().Kind).Filter("Broadcast" + op, value)
	return &AccountQueryRunner{
		db: repo.db,
		q:  q,
	}
}

func (repo *AccountAppxRepository) FindWhereEmail(op string, value interface{}) *AccountQueryRunner {
	q := datastore.NewQuery(new(Account).KeySpec().Kind).Filter("Email" + op, value)
	return &AccountQueryRunner{
		db: repo.db,
		q:  q,
	}
}

func (repo *AccountAppxRepository) FindWhereName(op string, value interface{}) *AccountQueryRunner {
	q := datastore.NewQuery(new(Account).KeySpec().Kind).Filter("Name" + op, value)
	return &AccountQueryRunner{
		db: repo.db,
		q:  q,
	}
}

func (repo *AccountAppxRepository) FindWhereToken(op string, value interface{}) *AccountQueryRunner {
	q := datastore.NewQuery(new(Account).KeySpec().Kind).Filter("Token" + op, value)
	return &AccountQueryRunner{
		db: repo.db,
		q:  q,
	}
}



func (repo *AccountAppxRepository) FindByBroadcast(value interface{}) *AccountQueryRunner {
	q := datastore.NewQuery(new(Account).KeySpec().Kind).Filter("Broadcast=", value)
	return &AccountQueryRunner{
		db: repo.db,
		q:  q,
	}
}

func (repo *AccountAppxRepository) FindByEmail(value interface{}) *AccountQueryRunner {
	q := datastore.NewQuery(new(Account).KeySpec().Kind).Filter("Email=", value)
	return &AccountQueryRunner{
		db: repo.db,
		q:  q,
	}
}

func (repo *AccountAppxRepository) FindByName(value interface{}) *AccountQueryRunner {
	q := datastore.NewQuery(new(Account).KeySpec().Kind).Filter("Name=", value)
	return &AccountQueryRunner{
		db: repo.db,
		q:  q,
	}
}

func (repo *AccountAppxRepository) FindByToken(value interface{}) *AccountQueryRunner {
	q := datastore.NewQuery(new(Account).KeySpec().Kind).Filter("Token=", value)
	return &AccountQueryRunner{
		db: repo.db,
		q:  q,
	}
}


func (repo *AccountAppxRepository) FindWhere(filter string, value interface{}) *AccountQueryRunner {
	q := datastore.NewQuery(new(Account).KeySpec().Kind).Filter(filter, value)
	return &AccountQueryRunner{
		db: repo.db,
		q:  q,
	}
}

func (repo *AccountAppxRepository) FindWhereAncestorIs(ancestor appx.Entity) *AccountQueryRunner {
	q := datastore.NewQuery(new(Account).KeySpec().Kind).Ancestor(ancestor.Key())
	return &AccountQueryRunner{
		db: repo.db,
		q:  q,
	}
}

type AccountQueryRunner struct {
	db *appx.Datastore
	q  *datastore.Query
}

func (runner *AccountQueryRunner) Project(fields ...string) *AccountQueryRunner {
	runner.q = runner.q.Project(fields...)
	return runner
}

func (runner *AccountQueryRunner) Distinct() *AccountQueryRunner {
	runner.q = runner.q.Distinct()
	return runner
}

func (runner *AccountQueryRunner) Limit(limit int) *AccountQueryRunner {
	runner.q = runner.q.Limit(limit)
	return runner
}

func (runner *AccountQueryRunner) KeysOnly() *AccountQueryRunner {
	runner.q = runner.q.KeysOnly()
	return runner
}


func (runner *AccountQueryRunner) OrderByBroadcastAsc() *AccountQueryRunner {
	runner.q = runner.q.Order("Broadcast")
	return runner
}

func (runner *AccountQueryRunner) OrderByEmailAsc() *AccountQueryRunner {
	runner.q = runner.q.Order("Email")
	return runner
}

func (runner *AccountQueryRunner) OrderByNameAsc() *AccountQueryRunner {
	runner.q = runner.q.Order("Name")
	return runner
}

func (runner *AccountQueryRunner) OrderByTokenAsc() *AccountQueryRunner {
	runner.q = runner.q.Order("Token")
	return runner
}



func (runner *AccountQueryRunner) OrderByBroadcastDesc() *AccountQueryRunner {
	runner.q = runner.q.Order("-Broadcast")
	return runner
}

func (runner *AccountQueryRunner) OrderByEmailDesc() *AccountQueryRunner {
	runner.q = runner.q.Order("-Email")
	return runner
}

func (runner *AccountQueryRunner) OrderByNameDesc() *AccountQueryRunner {
	runner.q = runner.q.Order("-Name")
	return runner
}

func (runner *AccountQueryRunner) OrderByTokenDesc() *AccountQueryRunner {
	runner.q = runner.q.Order("-Token")
	return runner
}


func (runner *AccountQueryRunner) Stream() *rivers.Stage {
	return runner.db.Query(runner.q).StreamOf(Account{})
}

func (runner *AccountQueryRunner) Count() (int, error) {
	return runner.db.Query(runner.q).Count()
}

func (runner *AccountQueryRunner) All() ([]*Account, error) {
	items := []*Account{}
	return items, runner.db.Query(runner.q).Results(items)
}

func (runner *AccountQueryRunner) First() (*Account, error) {
	item := &Account{}
	return item, runner.db.Query(runner.q.Limit(1)).Result(item)
}

func (runner *AccountQueryRunner) PagesIterator() appx.Iterator {
	return runner.db.Query(runner.q).PagesIterator()
}

func (runner *AccountQueryRunner) ItemsIterator() appx.Iterator {
	return runner.db.Query(runner.q).ItemsIterator()
}

func (runner *AccountQueryRunner) StartFrom(cursor string) *AccountQueryRunner {
	q := runner.db.Query(runner.q).StartFrom(cursor)
	return &AccountQueryRunner{
		db: runner.db,
		q:  q,
	}
}