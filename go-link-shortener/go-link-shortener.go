package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
)

func main() {
	configPath := flag.String("config", "", "The path to your config.json file.")
	flag.Parse()
	if *configPath == "" {
		log.Fatal("Please specify a config path")
	}
	fmt.Println("Loading config...")
	err := loadConfig(*configPath)
	if err != nil {
		fmt.Println("Unable to load config.json (does it exist?)")
		return
	}
	if config.URLLength > 16 {
		fmt.Println("URLLength too high, setting to 16.")
		fmt.Println("Don't worry, that is still 7 958 661 109 946 400 000 000 000 possible values")
		config.URLLength = 16
	}
	fmt.Println("Setting up database if not pre-existing")
	database, _ := sql.Open("sqlite3", config.DatabasePath)
	statement, err := database.Prepare("CREATE TABLE IF NOT EXISTS urls (id INTEGER PRIMARY  KEY, original_url TEXT, shortened_url TEXT)")
	if err != nil {
		fmt.Println(err)
		return
	}
	statement.Exec()
	fmt.Println("Starting service...")
	startWebServer()
}
