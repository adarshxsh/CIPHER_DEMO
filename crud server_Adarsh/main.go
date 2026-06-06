package main

import (
	"fmt"
	"log"
	"net/http"
)


func main(){
	// Router.HandleFunc("/", ServeHome).Methods("GET")
	fmt.Println("Hello World")

	r := Route()
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening at port 4000 ...")
	

} 