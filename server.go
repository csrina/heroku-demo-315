package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func createRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", serveIndex).Methods("GET")
	r.HandleFunc("/cat", serveCat).Methods("GET")
	r.HandleFunc("/dance", serveDance).Methods("GET")
	r.HandleFunc("/what", serveHack).Methods("GET")

	return r
}

func serveIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func serveCat(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "cat.html")
}

func serveDance(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "dance.html")
}

func serveHack(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "hack.html")
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
	//	port = "8080"
	 log.Fatal("$PORT must be set")
	}
	r := createRouter()
	http.Handle("/", r)
	log.Println("starting server.")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
