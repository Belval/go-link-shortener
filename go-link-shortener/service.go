package main

import (
	"database/sql"
	"net/http"

	"github.com/google/uuid"

	_ "github.com/mattn/go-sqlite3"
)

func linkHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path[1 : config.URLLength+1]
	database, _ := sql.Open("sqlite3", config.DatabasePath)
	rows, _ := database.Query("SELECT original_url FROM urls WHERE shortened_url = ? LIMIT 1", url)
	var originalURL string
	rows.Next()
	if err := rows.Scan(&originalURL); err != nil {
		http.Error(w, "", 404)
	} else {
		http.Redirect(w, r, originalURL, 301)
	}
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	url := r.Form.Get("url")
	if len(url) <= 8 || url[0:7] != "http://" || url[0:8] != "https://" {
		url = "http://" + url
	}
	uid, _ := uuid.NewRandom()
	linkUID := uid.String()[0:config.URLLength]
	database, _ := sql.Open("sqlite3", config.DatabasePath)
	statement, _ := database.Prepare("INSERT INTO urls (original_url, shortened_url) VALUES (?, ?)")
	statement.Exec(url, linkUID)
	w.Write([]byte(config.Domain + linkUID))
}

func startWebServer() {
	http.HandleFunc("/", linkHandler)
	http.HandleFunc("/api/v1/new/", apiHandler)
	http.ListenAndServe(":"+config.Port, nil)
}
