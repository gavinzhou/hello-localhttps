package main

import (
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Hi, This is an example of https server in golang!\n"))
}

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServeTLS(":4433", "localhost+3.pem", "localhost+3-key.pem", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
