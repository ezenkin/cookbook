package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
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
	mx := mux.NewRouter()
	mx.HandleFunc("/", handleGet).Methods("GET")
	mx.HandleFunc("/", handlePost).Methods("POST")
	mx.HandleFunc("/", handlePut).Methods("PUT")
	mx.HandleFunc("/", handleDelete).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":12345", mx))
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	log.Println("Got", r.Method, "request from", r.RemoteAddr)
	w.Write([]byte("got some data from " + r.Host + "\n"))
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	log.Println("Got", r.Method, "request from", r.RemoteAddr)
	w.Write([]byte("add some data on " + r.Host + "\n"))
}

func handlePut(w http.ResponseWriter, r *http.Request) {
	log.Println("Got", r.Method, "request from", r.RemoteAddr)
	w.Write([]byte("update some data on " + r.Host + "\n"))
}

func handleDelete(w http.ResponseWriter, r *http.Request) {
	log.Println("Got", r.Method, "request from", r.RemoteAddr)
	w.Write([]byte("delete some data on " + r.Host + "\n"))
}
