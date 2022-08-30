package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/tech-train-cyber/api"
)

func main() {
	server := api.NewServer()
	err := server.InitServer()
	if err != nil {
		log.Fatal(err)
		return
	}
	server.Run(80)
}
