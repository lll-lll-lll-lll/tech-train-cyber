package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/tech-train-cyber/api"
)

func main() {
	server := api.NewServer()
	server.Run(80)
}
