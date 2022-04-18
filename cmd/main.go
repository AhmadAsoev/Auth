package main

import (
	"authDB/cmd/handleFunc"
	"authDB/pkg/application/repository"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

const (
	POST   = "POST"
	DELETE = "DELETE"
	PUT    = "PUT"
)

func main() {
	log.Println("Server start read configuration")
	v := viper.New()
	v.SetConfigName("auth")
	v.SetConfigType("json")
	v.AddConfigPath("./config")
	if err := v.ReadInConfig(); err != nil {
		log.Fatal("Server can't read configuration: ", err)
	}

	port := v.GetString("server.port")
	dsn := v.GetString("server.db.dsn")

	log.Println("Server connected do DB")
	if err := repository.Connect(dsn); err != nil {
		log.Fatal("Server can't connect do DB")
	}

	router := mux.NewRouter()
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
