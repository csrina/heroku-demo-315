package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	data, err := ioutil.ReadFile("index.html")
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Length", fmt.Sprint(len(data)))
	fmt.Fprint(w, string(data))
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	http.HandleFunc("/", rootHandler)
	log.Println("starting server.")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
