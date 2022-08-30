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

func (md *MySql) InitDB() (*sqlx.DB, error) {
	db, err := md.Open()
	defer db.Close()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (md *MySql) Open() (*sqlx.DB, error) {
	dbcon, err := sqlx.Open("mysql", md.datasource)
	if err != nil {
		return nil, fmt.Errorf("failed db init. %s", err)
	}
	if err := dbcon.Ping(); err != nil {
		log.Fatal("Failed to connect to MySQL database")
		return nil, err
	}
	log.Println("DB Connect Successfully ")
	return dbcon, nil
}
