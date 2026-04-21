package main

import (
	"html/template"
	"math/rand/v2"
	"net/http"
	"strings"
)

var Templates = template.Must(template.ParseFiles("index.html", "view.html", "viewall.html"))

func GenerateRandomLinkCode(size uint) string {
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	var shortCode strings.Builder
	for i := 0; i <= int(size); i++ {
		shortCode.WriteString(string(alphabet[rand.IntN(len(alphabet))]))
	}

	return shortCode.String()
}

func RenderTemplate(w http.ResponseWriter, template string, sl *[]ShortLink){
	err := Templates.ExecuteTemplate(w, template+".html", sl)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func RenderOneTemplate(w http.ResponseWriter, template string, sl ShortLink){
	err := Templates.ExecuteTemplate(w, template+".html", sl)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

