package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("START")

	http.Handle("/", http.FileServer(http.Dir("./html/")))

	http.ListenAndServe(":80", nil)
}
