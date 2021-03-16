package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "Hello")
	})

	r.HandleFunc("/value/{value}", func(w http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		s := vars["value"]
		v, err := strconv.ParseFloat(s, 32)
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Fprintf(w, "%.1f\n", cToF(float32(v)))
	})

	http.ListenAndServe(":3000", r)
}

func cToF(v float32) float32 {
	return v*9/5 + 32
}

func fToC(v float32) float32 {
	return (v - 32) * 5 / 9
}
