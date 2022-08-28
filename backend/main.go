package main

import (
	"fmt"
	"net/http"

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
	defer dbcon.Close()
	if err != nil {
		return nil, fmt.Errorf("failed db init. %s", err)
	}
	if err := dbcon.Ping(); err != nil {
		return nil, err
	}
	return dbcon, nil
}

func main() {
	h1 := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("helloworld"))
	}
	http.HandleFunc("/", h1)
	http.ListenAndServe(":8080", nil)
}
