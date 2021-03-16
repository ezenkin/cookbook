package main

import (
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("Got", r.Method, "request from", r.RemoteAddr)
	switch r.Method {
	case "GET":
		w.Write([]byte("Response to GET request to the " + r.Host + "\n"))
	case "PUT":
		w.Write([]byte("Response to PUT to the " + r.Host + "\n"))
	case "POST":
		w.Write([]byte("Response to POST to the " + r.Host + "\n"))
	case "DELETE":
		w.Write([]byte("Response to DELETE to the " + r.Host + "\n"))
	}
}

func main() {
	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(":12345", nil); err != nil {
		log.Fatal(err)
	}
}
