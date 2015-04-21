package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var db *sql.DB

var cfg *Config

func randomHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("user_name")
	randValue := rand.Intn(100)
	msg := SlackMsg{
		Channel:   cfg.Channel,
		Username:  cfg.Username,
		Parse:     "full",
		Text:      fmt.Sprintf("%s's number : %d", username, randValue),
		IconEmoji: "",
	}

	err := msg.Post(cfg.WebhookUrl)
	if err != nil {
		log.Fatalf("Post failed: %v", err)
		fmt.Fprintf(w, "Post failed: %v", err)
	}

	InsertRandomToDB(db, username, randValue)
}

func main() {
	var err error
	cfg, err = ReadConfig()
	if err != nil {
		log.Fatalf("Could not read config: %v", err)
	}

	db = OpenDB()

	rand.Seed(time.Now().UTC().UnixNano())

	http.HandleFunc("/random", randomHandler)
	err = http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}
