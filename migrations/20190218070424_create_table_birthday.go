package migrations

import "database/sql"

type CreateTableBirthday20190218070424 struct{}

func (m CreateTableBirthday20190218070424) Version() string {
	return "20190218070424_CreateTableBirthday"
}

func (m CreateTableBirthday20190218070424) Up(tx *sql.Tx) error {
	_, err := tx.Exec(`
		CREATE TABLE IF NOT EXISTS birthday (
			id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
			serverid STRING NOT NULL DEFAULT '',
			userid STRING NOT NULL DEFAULT '',
			month INT NOT NULL DEFAULT 0,
			day INT NOT NULL DEFAULT 0,
			created TIMESTAMPTZ NOT NULL DEFAULT NOW(),
			updated TIMESTAMPTZ,
			UNIQUE (serverid, userid)
		)`)

	return err
}

func (m CreateTableBirthday20190218070424) Down(tx *sql.Tx) error {
	_, err := tx.Exec(`DROP TABLE birthday`)
	return err
}
