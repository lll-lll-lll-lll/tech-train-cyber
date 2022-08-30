package db

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type MySql struct {
	datasource string
}

func NewMySql(datasource string) *MySql {
	return &MySql{
		datasource: datasource,
	}
}

func (db *MySql) Open() (*sqlx.DB, error) {
	dbcon, err := sqlx.Open("mysql", db.datasource)
	if err != nil {
		return nil, fmt.Errorf("failed db init. %s", err)
	}
	if err := dbcon.Ping(); err != nil {
		return nil, err
	}
	log.Println("DB Connect Successfully ")
	return dbcon, nil
}
