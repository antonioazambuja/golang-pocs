package main

import (
	"log"
	"net/http"

	handler_v1 "github.com/antonioazambuja/golang-pocs/api-design/app/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/{name}", handler_v1.GetEmployeeByName).Methods("GET")
	router.HandleFunc("/{name}/{accountID}", handler_v1.SaveEmployee).Methods("POST")
	router.HandleFunc("/health", handler_v1.HealthCheck).Methods("GET")
	log.Print("POC Golang API Desing!")
	log.Print("Serving at port 5000!")
	log.Fatal(http.ListenAndServe(":5000", router))
}
