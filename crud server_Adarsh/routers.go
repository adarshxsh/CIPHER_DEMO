package main

import (
	// "fmt"
	// "net/http"

	"github.com/gorilla/mux"
)

func Route() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", ServeHome).Methods("GET")
	router.HandleFunc("/tasks", ServeTasks).Methods("GET")
	router.HandleFunc("/task/{id}", ServeTask).Methods("GET")
	router.HandleFunc("/task", addTask).Methods("POST")
	router.HandleFunc("/task/{id}", updateTask).Methods("PUT")
	router.HandleFunc("/task/{id}", deleteTask).Methods("DELETE")
	return router

}
