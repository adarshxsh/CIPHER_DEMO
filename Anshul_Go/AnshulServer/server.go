package main

import (
	"fmt"
	"net/http"
)

func name(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	html := `
	<!DOCTYPE html>
	<html>
	<head>
		<title>Name</title>
	</head>
	<body>
		<h1 style="color:red;"><b>Anshul</b></h1>
	</body>
	</html>
	`

	fmt.Fprint(w, html)
}

func Home(w http.ResponseWriter, req *http.Request){
	w.Header().Set("Content-Type", "text/html")

	myHome := `
	<!DOCTYPE html>
	<html>
	<head>
		<title>Home</title>
	</head>
	<body>
		<a href="http://localhost:8090/name">Name</a>
	</body>
	</html>
	`
	fmt.Fprint(w,myHome)
}

func main() {
	http.HandleFunc("/name", name)
	http.HandleFunc("/", Home)

	fmt.Println("Server running on http://localhost:8090")
	http.ListenAndServe(":8090", nil)
}