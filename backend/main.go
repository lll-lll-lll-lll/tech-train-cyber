package main

import (
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/tech-train-cyber/db"
)

func main() {
	d := db.DBConfig{
		User:     os.Getenv("MYSQL_USER"),
		Password: os.Getenv("MYSQL_PASSWORD"),
		Host:     os.Getenv("MYSQL_HOST"),
		Port:     os.Getenv("MYSQL_PORT"),
		DBName:   os.Getenv("MYSQL_DATABASE"),
	}.String()
	md := db.NewMySql(d)
	db, err := md.Open()
	defer db.Close()
	if err != nil {
		log.Fatal(err)
		return
	}

}
