package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hello", handler).Methods("GET")

	fileDirectory := http.Dir("./assets/")
	fileHandler := http.StripPrefix("/assets/", http.FileServer(fileDirectory))
	r.PathPrefix("/assets/").Handler(fileHandler).Methods("GET")

	r.HandleFunc("/bird", getBirdHandler).Methods("GET")
	r.HandleFunc("/bird", createBirdHandler).Methods("POST")

	return r
}
func main() {

	r := newRouter()
	//	r := mux.NewRouter()

	//	r.HandleFunc("/hello", handler).Methods("GET")

	//http.HandleFunc("/", handler)
	fmt.Println("Listening at Port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
