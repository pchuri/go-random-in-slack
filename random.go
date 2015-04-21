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

var cfg *Config
var db *sql.DB

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

	err = InsertRandomToDB(db, username, randValue)
	if err != nil {
		fmt.Fprintf(w, "Insert Failed : %v", err)
	}
}

func avgRandomHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("user_name")

	avgRand, err := SelectAvgRandomFromDB(db, username)
	if err != nil {
		fmt.Fprintf(w, "Select Avg Random : %v", err)
	}

	msg := SlackMsg{
		Channel:   cfg.Channel,
		Username:  cfg.Username,
		Parse:     "full",
		Text:      fmt.Sprintf("%s's avg number : %d", username, avgRand),
		IconEmoji: "",
	}

	err = msg.Post(cfg.WebhookUrl)
	if err != nil {
		log.Fatalf("Post failed: %v", err)
		fmt.Fprintf(w, "Post failed: %v", err)
	}
}

func main() {
	var err error
	cfg, err = ReadConfig()
	if err != nil {
		log.Fatalf("Could not read config: %v", err)
	}

	db, err = ConnectDB()
	defer db.Close()
	if err != nil {
		log.Fatalf("Connect Database Failed : %v", err)
	}

	rand.Seed(time.Now().UTC().UnixNano())

	http.HandleFunc("/random", randomHandler)
	http.HandleFunc("/avgRandom", avgRandomHandler)

	err = http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}
