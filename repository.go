package main

import (
	"database/sql"
	"fmt"

	"github.com/jukeizu/birthday/api/protobuf-spec/birthdaypb"
	"github.com/jukeizu/birthday/migrations"
	_ "github.com/lib/pq"
	"github.com/shawntoffel/gossage"
)

const (
	DatabaseName = "birthday"
)

type Repository interface {
	SetBirthday(*birthdaypb.SetBirthdayRequest) (*birthdaypb.UserBirthday, error)
	Migrate() error
}

type repository struct {
	Db *sql.DB
}

func NewRepository(url string) (Repository, error) {
	conn := fmt.Sprintf("postgresql://%s/%s?sslmode=disable", url, DatabaseName)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}

	r := repository{
		Db: db,
	}

	return &r, nil
}

func (r *repository) Migrate() error {
	_, err := r.Db.Exec(`CREATE DATABASE IF NOT EXISTS ` + DatabaseName)
	if err != nil {
		return err
	}

	g, err := gossage.New(r.Db)
	if err != nil {
		return err
	}

	err = g.RegisterMigrations(migrations.CreateTableBirthday20190218070424{})
	if err != nil {
		return err
	}

	return g.Up()
}

func (r *repository) SetBirthday(req *birthdaypb.SetBirthdayRequest) (*birthdaypb.UserBirthday, error) {
	q := `INSERT INTO birthday (serverid, userId, month, day) 
		VALUES($1, $2, $3, $4) 
		ON CONFLICT (serverid, userid) DO UPDATE SET month = excluded.month, day = excluded.day, updated = NOW()
		RETURNING id, serverid, userid, month, day`

	userBirthday := birthdaypb.UserBirthday{}

	err := r.Db.QueryRow(q,
		req.ServerId,
		req.UserId,
		req.Month,
		req.Day,
	).Scan(
		&userBirthday.Id,
		&userBirthday.ServerId,
		&userBirthday.UserId,
		&userBirthday.Month,
		&userBirthday.Day,
	)

	return &userBirthday, err
}
