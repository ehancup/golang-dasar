package database

import (
	"fmt"
)

var connection string

func init() {
	defer fmt.Println("connected to Mysql")
	connection = "MySQL"
}

func GetDatabase() string {
	return connection
}