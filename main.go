package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func CreateDBConnection(dbName string) (*sql.DB, error){
	conn := mysql.NewConfig()
	conn.User = os.Getenv("DBUSER")
	conn.Passwd = os.Getenv("DBPASS")
	conn.Net = "tcp"
	conn.Addr = "127.0.0.1:3306"
	conn.DBName =dbName

	db, err := sql.Open("mysql", conn.FormatDSN())
	if err != nil {
		return nil, err
	}

	return db, nil
}

func main(){
	godotenv.Load()
	db, err := CreateDBConnection("heyshorty")
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Connected")

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	fileserver := http.FileServer(http.Dir("static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fileserver))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		Templates.ExecuteTemplate(w, "index.html", nil)
	})
	r.Get("/{id}", GetShortLinkByShortCode(db))
	r.Get("/shorten/", GetAllShortLinks(db))
	r.Get("/shorten/{id}", GetOneShortLinkByID(db))
	r.Post("/shorten", SaveShortLink(db))
	r.Delete("/shorten/{id}", DeleteShortLink(db))
	r.Put("/shorten/{id}", UpdateShortLink(db))

	var port string = fmt.Sprintf("%s", os.Getenv("PORT"))
	http.ListenAndServe(port, r)

}

// func redirectHander(w http.ResponseWriter, r *http.Request){
// 	http.Redirect(w, r, "http://localhost:3000/new-path", http.StatusPermanentRedirect)
// }