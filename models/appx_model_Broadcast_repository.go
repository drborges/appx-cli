// AUTO GENERATED BY appx-cli - DO NOT EDIT

package models

import (
	"appengine"
	"appengine/datastore"
	"github.com/drborges/appx"
	"github.com/drborges/rivers"
)

func FromBroadcast(context appengine.Context) *BroadcastAppxRepository {
	return &BroadcastAppxRepository{
		db: appx.NewDatastore(context),
	}
}

type BroadcastAppxRepository struct {
	db *appx.Datastore
}

func (repo *BroadcastAppxRepository) New() *BroadcastAppxModel {
	return &BroadcastAppxModel{
		db: repo.db,
	}
}

func (repo *BroadcastAppxRepository) GetByKey(key *datastore.Key) (*BroadcastAppxModel, error) {
	item := &Broadcast{}
	item.SetKey(key)
	return item, repo.db.Load(item)
}

func (repo *BroadcastAppxRepository) GetByEncodedKey(key string) (*BroadcastAppxModel, error) {
	item := &Broadcast{}
	if err := item.SetEncodedKey(key); err != nil {
		return nil, err
	}
	return item, repo.db.Load(item)
}




func (repo *BroadcastAppxRepository) GetByLength(value int) (*BroadcastAppxModel, error) {
	item := &Broadcast{
		Length: value,
	}
	return item, repo.db.Load(item)
}

func (repo *BroadcastAppxRepository) GetByURL(value string) (*BroadcastAppxModel, error) {
	item := &Broadcast{
		URL: value,
	}
	return item, repo.db.Load(item)
}



func (repo *BroadcastAppxRepository) FindWhereLength(op string, value int) *BroadcastQueryRunner {
	q := datastore.NewQuery(new(Broadcast).KeySpec().Kind).Filter("Length" + op, value)
	return &BroadcastQueryRunner{
		db: repo.db,
		q:  q,
	}
}

func (repo *BroadcastAppxRepository) FindWhereURL(op string, value string) *BroadcastQueryRunner {
	q := datastore.NewQuery(new(Broadcast).KeySpec().Kind).Filter("URL" + op, value)
	return &BroadcastQueryRunner{
		db: repo.db,
		q:  q,
	}
}









func (repo *BroadcastAppxRepository) FindByLength(value int) *BroadcastQueryRunner {
	q := datastore.NewQuery(new(Broadcast).KeySpec().Kind).Filter("Length=", value)
	return &BroadcastQueryRunner{
		db: repo.db,
		q:  q,
	}
}

func (repo *BroadcastAppxRepository) FindByURL(value string) *BroadcastQueryRunner {
	q := datastore.NewQuery(new(Broadcast).KeySpec().Kind).Filter("URL=", value)
	return &BroadcastQueryRunner{
		db: repo.db,
		q:  q,
	}
}


func (repo *BroadcastAppxRepository) FindWhere(filter string, value interface{}) *BroadcastQueryRunner {
	q := datastore.NewQuery(new(Broadcast).KeySpec().Kind).Filter(filter, value)
	return &BroadcastQueryRunner{
		db: repo.db,
		q:  q,
	}
}

func (repo *BroadcastAppxRepository) FindWhereAncestorIs(ancestor appx.Entity) *BroadcastQueryRunner {
	q := datastore.NewQuery(new(Broadcast).KeySpec().Kind).Ancestor(ancestor.Key())
	return &BroadcastQueryRunner{
		db: repo.db,
		q:  q,
	}
}

type BroadcastQueryRunner struct {
	db *appx.Datastore
	q  *datastore.Query
}

func (runner *BroadcastQueryRunner) Select(fields ...string) *BroadcastQueryRunner {
	runner.q = runner.q.Project(fields...)
	return runner
}

func (runner *BroadcastQueryRunner) Distinct() *BroadcastQueryRunner {
	runner.q = runner.q.Distinct()
	return runner
}

func (runner *BroadcastQueryRunner) Limit(limit int) *BroadcastQueryRunner {
	runner.q = runner.q.Limit(limit)
	return runner
}

func (runner *BroadcastQueryRunner) KeysOnly() *BroadcastQueryRunner {
	runner.q = runner.q.KeysOnly()
	return runner
}


func (runner *BroadcastQueryRunner) OrderByLengthAsc() *BroadcastQueryRunner {
	runner.q = runner.q.Order("Length")
	return runner
}

func (runner *BroadcastQueryRunner) OrderByURLAsc() *BroadcastQueryRunner {
	runner.q = runner.q.Order("URL")
	return runner
}



func (runner *BroadcastQueryRunner) OrderByLengthDesc() *BroadcastQueryRunner {
	runner.q = runner.q.Order("-Length")
	return runner
}

func (runner *BroadcastQueryRunner) OrderByURLDesc() *BroadcastQueryRunner {
	runner.q = runner.q.Order("-URL")
	return runner
}


func (runner *BroadcastQueryRunner) Stream() *rivers.Stage {
	return runner.db.Query(runner.q).StreamOf(Broadcast{})
}

func (runner *BroadcastQueryRunner) Count() (int, error) {
	return runner.db.Query(runner.q).Count()
}

func (runner *BroadcastQueryRunner) All() ([]*BroadcastAppxModel, error) {
	items := []*BroadcastAppxModel{}
	return items, runner.db.Query(runner.q).Results(items)
}

func (runner *BroadcastQueryRunner) First() (*BroadcastAppxModel, error) {
	item := &BroadcastAppxModel{}
	return item, runner.db.Query(runner.q.Limit(1)).Result(item)
}

func (runner *BroadcastQueryRunner) PagesIterator() appx.Iterator {
	return runner.db.Query(runner.q).PagesIterator()
}

func (runner *BroadcastQueryRunner) ItemsIterator() appx.Iterator {
	return runner.db.Query(runner.q).ItemsIterator()
}

func (runner *BroadcastQueryRunner) StartFrom(cursor string) *BroadcastQueryRunner {
	q := runner.db.Query(runner.q).StartFrom(cursor)
	return &BroadcastQueryRunner{
		db: runner.db,
		q:  q,
	}
}