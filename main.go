package main

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

func main(){
	godotenv.Load()
	cfg := mysql.NewConfig()
}