package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"log"
	"net/http"
	"goapi/api"
)
func main() {

	api.Init()
	router := mux.NewRouter()
	basePath:="api"
	apiPath:="people"
	route:=fmt.Sprintf("/%v/%v",basePath,apiPath)
	router.HandleFunc(route, api.ListAll).Methods(http.MethodGet)
	router.HandleFunc(route, api.CreateOne).Methods(http.MethodPost)
	router.HandleFunc(route+"/{id}", api.GetOne).Methods(http.MethodGet)
	router.HandleFunc(route+"/{id}", api.UpdateOne).Methods(http.MethodPut)
	router.HandleFunc(route+"/{id}", api.DeleteOne).Methods(http.MethodDelete)

	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})

	port:="3003"
	fmt.Println("Listening on ::"+port)
	log.Fatal(http.ListenAndServe(":"+port , handlers.CORS(allowedOrigins, allowedMethods)(router)))


}

