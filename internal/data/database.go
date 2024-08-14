package data

import (
	"github.com/lauro-ss/goe"
	"github.com/lauro-ss/postgres"
)

type User struct {
	Id       uint
	Name     uint
	Password string `goe:"type:varchar(64)"`
}

type Database struct {
	User *User
	*goe.DB
}

func OpenAndMigrate(dns string) (*Database, error) {
	db := &Database{DB: &goe.DB{}}
	err := goe.Open(db, postgres.Open(dns))
	if err != nil {
		return nil, err
	}
	err = db.Migrate(goe.MigrateFrom(db))
	if err != nil {
		return nil, err
	}
	return db, nil
}
