package main

import (
	"authDB/cmd/handleFunc"
	"authDB/pkg/application/repository"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const (
	POST   = "POST"
	DELETE = "DELETE"
	PUT    = "PUT"
)

func main() {

	log.Println("Server connected do DB")
	if err := repository.Connect(); err != nil {
		log.Fatal("Server can't connect do DB")
	}

	router := mux.NewRouter()
	port := ":1414"

	//AddUser add user
	router.HandleFunc("/addUser", handleFunc.AddUser).Methods(POST)
	//DeleteUser delete user
	router.HandleFunc("/user/{id}", handleFunc.DeleteUser).Methods(DELETE)
	//Update by id
	router.HandleFunc("/user/{id}", handleFunc.UpdateUser).Methods(PUT)

	fmt.Println("Server is run")
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatal("Server is not ready")
	}
}
