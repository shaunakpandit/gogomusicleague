package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"jssp.io/gogomusicleague/models"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}

	db, err := initDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	cmp, err := models.CompetitorByName("shaunakpandit", db)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Competitor found: %+v\n", cmp)
}

func initDB() (*sql.DB, error) {
	cfg := mysql.NewConfig()
	cfg.User = os.Getenv("DBUSER")
	cfg.Passwd = os.Getenv("DBPASS")
	cfg.Net = "tcp"
	cfg.Addr = "192.168.4.28:33061"
	cfg.DBName = "musicleague"

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
