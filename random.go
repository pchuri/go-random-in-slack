package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

func RandomHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%d", rand.Intn(100))
}

func main() {
	http.HandleFunc("/", RandomHandler)
	http.ListenAndServe(":8081", nil)
}
