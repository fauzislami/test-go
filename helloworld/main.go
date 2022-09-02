package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

func Health(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, `{"alive": true}`)
}

func main() {
	fmt.Println("Starting hello world")
	router := mux.NewRouter()

	router.HandleFunc("/health", Health).Methods("GET")
	http.ListenAndServe(":8000", router)
}
