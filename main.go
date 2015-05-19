package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", redirect)
	err := http.ListenAndServe(":"+os.Args[1], nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, os.Args[2], 301)
}
