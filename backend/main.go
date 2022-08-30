package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/tech-train-cyber/db"
)

func main() {
	cs := "user:password@tcp(localhost:3306)/cyberdb?charset=utf8mb4"
	md := db.NewMySql(cs)
	db, err := md.Open()
	defer db.Close()
	if err != nil {
		log.Fatal(err)
		return
	}

}
