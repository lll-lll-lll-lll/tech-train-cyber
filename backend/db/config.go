package db

import "fmt"

type DBConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	DBName   string
}

func (dc DBConfig) String() string {
	if dc.User == "" {
		dc.User = "user"
	}
	if dc.Password == "" {
		dc.Password = "password"
	}
	if dc.Port == "" {
		dc.Host = "3306"
	}
	if dc.Host == "" {
		dc.Host = "localhost"
	}
	if dc.DBName == "" {
		dc.DBName = "cyberdb"
	}
	s := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?", dc.User, dc.Password, dc.Host, dc.Port, dc.DBName)
	return s
}
