package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Database struct {

}

func NewDatabase() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		"admin","adminadmin1","192.168.39.192", 32216, "pokemon_project")
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	return db, nil
}
