package main

import (
	"fmt"
	"net/http"
)

func bio(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, `
	<html>
	<body>
		<h1 style="color:blue;">Hamza</h1>
		<p>Go Developer</p>
		<a href="/">Back</a>
	</body>
	</html>
	`)
}

func home(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, `
	<html>
	<body>
		<h2>Welcome!</h2>
		<a href="/bio">Visit Bio</a>
	</body>
	</html>
	`)
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/bio", bio)
	fmt.Println("Server running on http://localhost:8090")
	http.ListenAndServe(":8090", nil)
}