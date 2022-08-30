package api

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/tech-train-cyber/db"
)

type Server struct {
	db     *sqlx.DB
	router *gin.Engine
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) InitServer() error {
	r := InitRouter()
	d := db.DBConfig{
		User:     os.Getenv("MYSQL_USER"),
		Password: os.Getenv("MYSQL_PASSWORD"),
		Host:     os.Getenv("MYSQL_HOST"),
		Port:     os.Getenv("MYSQL_PORT"),
		DBName:   os.Getenv("MYSQL_DATABASE"),
	}.String()
	md := db.NewMySql(d)
	db, err := md.InitDB()
	if err != nil {
		log.Fatal(err)
		return err
	}
	s.db = db
	s.router = r
	return nil
}

func (s *Server) Run(port int) {
	log.Printf("Listening on port %d", port)
	err := http.ListenAndServe(
		fmt.Sprintf(":%d", port), nil,
	)
	if err != nil {
		panic(err)
	}
}
