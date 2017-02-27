package main

import (
	"encoding/json"
	"fmt"
	lib "golangmicroservice/lib"
	"net/http"
)

//Message of the jason response
type Message struct {
	Text string
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome, %s", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/about", about)
	http.ListenAndServe(":8080", nil)
}

func about(w http.ResponseWriter, r *http.Request) {
	m := Message{"Welcome to the about route"}
	lib.PostRequest()
	b, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	w.Write(b)
}
