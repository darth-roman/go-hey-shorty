package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type ShortLink struct {
	ID string				`json:"id"`
	URL	string				`json:"url"`
	ShortCode string		`json:"shortcode"`
	CreatedAt string		`json:"created_at"`
}


// func ViewOneLinkHandler(w http.ResponseWriter, r *http.Request) {
// 	url := r.URL
// 	fmt.Println(url)

// 	err := templates.ExecuteTemplate(w, "index.html", url)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}
// }

func SaveShortLink(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		url := r.FormValue("url")
		shortCode := GenerateRandomLinkCode(6)
		
		shortUrl := &ShortLink{
			URL: url,
			ShortCode: shortCode,
		}
	
		_, err := db.Exec("Insert Into shortlink (url, shortcode) values (?, ?)", shortUrl.URL, shortUrl.ShortCode)
		if err != nil {
			log.Fatal(err)
		}
	
		// w.WriteHeader(http.StatusCreated)
		// RenderOneTemplate(w, "viewall", *shortUrl)
		http.Redirect(w, r, "/shorten", http.StatusSeeOther)
	}
}

func GetAllShortLinks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("Select * From shortlink Order By created_at desc")
		if err != nil {
			log.Fatal(err)
		}

		defer rows.Close()

		var links []ShortLink

		for rows.Next() {
			var sl ShortLink
			if err := rows.Scan(&sl.ID, &sl.URL, &sl.ShortCode, &sl.CreatedAt); err != nil {
				log.Fatal(err)
			}
			links = append(links, sl)
		}

		RenderTemplate(w, "viewall", &links)
	}
}

func GetOneShortLinkByID(db *sql.DB) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/shorten/")
		fmt.Println(id)
		rows, err := db.Query("Select * From shortlink Where id = ?", id)
		if err != nil {
			log.Fatal(err)
		}

		defer rows.Close()

		var links []ShortLink
		for rows.Next() {
			var sl ShortLink
			if err := rows.Scan(&sl.ID, &sl.URL, &sl.ShortCode, &sl.CreatedAt); err != nil {
				log.Fatal(err)
			}
			links = append(links, sl)
		}
		// w.WriteHeader(http.StatusOK)
		// w.Header().Set("Content-Type", "application/json")
		// json.NewEncoder(w).Encode(links)
		RenderOneTemplate(w, "view", links[0])
	}
}

func DeleteShortLink(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/shorten/")
		_, err := db.Exec("DELETE FROM shortlink WHERE id = ?", id)
		if err != nil {
			if err == sql.ErrNoRows{
				log.Fatal(err)
			}
		}
		w.WriteHeader(http.StatusNoContent)
	}
}

func UpdateShortLink(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/shorten/")
		url := r.FormValue("url")
		_, err := db.Exec("Update shortlink Set url = ? Where id = ?", url, id)
		if err != nil {
			log.Fatal(err)
		}
		w.WriteHeader(http.StatusOK)
	}
}

func GetShortLinkByShortCode(db *sql.DB) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		 shortCode := strings.TrimPrefix(r.URL.Path ,"/")
		 fmt.Println(shortCode)
		 rows, err := db.Query("Select * From shortlink Where shortcode like ?", shortCode)
		 if err != nil {
			log.Fatal(err)
		 }

		 defer rows.Close()

		 var links []ShortLink 
		 for rows.Next() {
			var sl ShortLink
			if err := rows.Scan(&sl.ID, &sl.URL, &sl.ShortCode, &sl.CreatedAt); err != nil {
				log.Fatal(err)
			}

			links = append(links, sl)
		 }

		 http.Redirect(w, r, links[0].URL, http.StatusTemporaryRedirect)
	}
}