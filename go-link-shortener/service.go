package main

import (
	"database/sql"
	"fmt"
	"log"
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
		defer rows.Close()
		defer database.Close()
		http.Error(w, "", 404)
	} else {
		defer rows.Close()
		defer database.Close()
		http.Redirect(w, r, originalURL, 301)
	}
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	url := r.Form.Get("url")
	if len(url) <= 8 || (url[0:7] != "http://" && url[0:8] != "https://") {
		url = "http://" + url
	}
	uid, _ := uuid.NewRandom()
	linkUID := uid.String()[0:config.URLLength]
	database, _ := sql.Open("sqlite3", config.DatabasePath)
	statement, _ := database.Prepare("INSERT INTO urls (original_url, shortened_url) VALUES (?, ?)")
	if _, err := statement.Exec(url, linkUID); err != nil {
		fmt.Println(err)
	}
	defer statement.Close()
	defer database.Close()
	w.Write([]byte(config.Domain + linkUID))
}

func startWebServer() {
	http.HandleFunc("/", linkHandler)
	http.HandleFunc("/api/v1/new/", apiHandler)
	log.Fatal(http.ListenAndServe(":"+config.Port, nil))
}
