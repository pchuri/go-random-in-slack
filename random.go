package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

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

	db, err := ConnectDB()
	defer db.Close()
	if err != nil {
		fmt.Fprintf(w, "Connect Database Failed : %v", err)
	}

	fmt.Fprintf(w, "INSERT INTO random_logs(username, random) VALUES (%v, %v)", username, randValue)

	err = InsertRandomToDB(db, username, randValue)
	if err != nil {
		fmt.Fprintf(w, "Insert Failed : %v", err)
	}

}

func main() {
	var err error
	cfg, err = ReadConfig()
	if err != nil {
		log.Fatalf("Could not read config: %v", err)
	}

	rand.Seed(time.Now().UTC().UnixNano())

	http.HandleFunc("/random", randomHandler)
	err = http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}
