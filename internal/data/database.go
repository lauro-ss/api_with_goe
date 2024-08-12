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

func InitDatabase(dns string) *Database {
	db := &Database{DB: &goe.DB{}}
	goe.Open(db, postgres.Open(dns))
	db.Migrate(goe.MigrateFrom(db))
	return db
}
