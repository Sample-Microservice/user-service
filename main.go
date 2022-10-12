package main

import (
	"log"
	"net/http"
	"os"
)

func coffeeHandler(w http.ResponseWriter, r *http.Request) {
	servant, err := os.Hostname()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte("User service is running on - " + servant))
}

func main() {
	http.HandleFunc("/user", coffeeHandler)
	log.Fatal(http.ListenAndServe(":8090", nil))
}
