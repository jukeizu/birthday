package birthday

import (
	mdb "github.com/shawntoffel/GoMongoDb"
)

type BirthdayStorage interface {
	mdb.Storage
}

type jobStorage struct {
	mdb.Store
}

func NewStorage(dbConfig mdb.DbConfig) (BirthdayStorage, error) {
	store, err := mdb.NewStorage(dbConfig)

	j := jobStorage{}
	j.Session = store.Session
	j.Collection = store.Collection

	return &j, err
}
