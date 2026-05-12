package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"text/template"

	"jssp.io/gogomusicleague/handlers"

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

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/votesForUser", func(w http.ResponseWriter, r *http.Request) {
		handlers.VotesForUserHandler(w, r, db)
	})
	http.HandleFunc("/userSongs", func(w http.ResponseWriter, r *http.Request) {
		handlers.UserSongsHandler(w, r, db)
	})
	http.HandleFunc("/songs", func(w http.ResponseWriter, r *http.Request) {
		handlers.SearchSongs(w, r, db)
	})

	log.Println("Server running on http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/home.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "something went wrong", http.StatusInternalServerError)
	}
}

func initDB() (*sql.DB, error) {
	cfg := mysql.NewConfig()
	cfg.User = os.Getenv("DBUSER")
	cfg.Passwd = os.Getenv("DBPASS")
	cfg.Net = "tcp"
	cfg.Addr = "192.168.4.28:33061"
	cfg.DBName = "musicleague"
	cfg.ParseTime = true

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
