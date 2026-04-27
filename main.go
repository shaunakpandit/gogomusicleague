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

	http.HandleFunc("/userSongs", userSongsHandler)

	log.Println("Server running on http://localhost:8080/search")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func userSongsHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	db, err := initDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	data := handlers.UserSongsHandler(name, db)

	tmpl := template.Must(template.ParseFiles("templates/userSongs.html"))
	tmplErr := tmpl.Execute(w, data)
	if tmplErr != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
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
