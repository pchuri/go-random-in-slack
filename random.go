package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
)

func RandomHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%d", rand.Intn(100))
}

func main() {
	http.HandleFunc("/random", RandomHandler)
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}
