package config

import (
	"database/sql"
	"fmt"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
)

type Database struct {
	Server string `json:"-" envconfig:"BD_SERVER" default:"192.168.39.192"`
	Port string `json:"-" envconfig:"BD_PORT" default:"32216"`
	User string `json:"-" envconfig:"BD_USER" default:"admin"`
	Password string `json:"-" envconfig:"BD_PASWWORD" default:"adminadmin1"`
	DatabaseName string `json:"-" envconfig:"BD_NAME" default:"pokemon_project"`
}

func NewDatabase() (*sql.DB, error) {
	database := &Database{}
	err := envconfig.Process("", database)
	if err != nil {
		return nil, err
	}
	psqlInfo := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		database.User,database.Password,database.Server, database.Port, database.DatabaseName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
