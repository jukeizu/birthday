package birthday

import (
	mdb "github.com/shawntoffel/GoMongoDb"
	"gopkg.in/mgo.v2/bson"
)

type Birthday struct {
	Id        string
	Month     string
	Day       int
	UserId    string
	ChannelId string
	ServerId  string
	JobId     string
	Enabled   bool
}

type BirthdayStorage interface {
	mdb.Storage

	Save(Birthday) error
	Birthday(userId string, channelId string) (Birthday, error)
	Disable(id string) error
}

type storage struct {
	mdb.Store
}

func NewStorage(dbConfig mdb.DbConfig) (BirthdayStorage, error) {
	store, err := mdb.NewStorage(dbConfig)

	j := storage{}
	j.Session = store.Session
	j.Collection = store.Collection

	return &j, err
}

func (s *storage) Save(b Birthday) error {
	return s.Collection.Insert(b)
}

func (s *storage) Birthday(userId string, serverId string) (Birthday, error) {
	b := Birthday{}

	err := s.Collection.Find(bson.M{"userid": userId, "serverid": serverId, "enabled": true}).One(&b)

	return b, err
}

func (s *storage) Disable(id string) error {
	_, err := s.Collection.Upsert(bson.M{"id": id}, bson.M{"$set": bson.M{"enabled": false}})

	return err
}
