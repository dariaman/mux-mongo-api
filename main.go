package main

import (
	"log"
	"mux-mongo-api/configs"
	"net/http"

	"github.com/gorilla/mux"
)

// == > https://dev.to/hackmamba/build-a-rest-api-with-golang-and-mongodb-gorillamux-version-57fh

func main() {
	router := mux.NewRouter()

	configs.ConnectDB()

	// router.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
	// 	rw.Header().Set("Content-Type", "application/json")

	// 	json.NewEncoder(rw).Encode(map[string]string{"data": "Hello from Mux & mongoDB"})
	// }).Methods("GET")

	log.Fatal(http.ListenAndServe(":6000", router))
}
