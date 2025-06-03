package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Basic HTML content
	fmt.Fprint(w, `
		<!DOCTYPE html>
		<html>
		<head>
			<title>Golang Skeleton Code</title>
		</head>
		<body>
			<h1>Hello, Go Web!</h1>
			<p>This is a simple web page served using Golang.</p>
		</body>
		</html>
	`)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server starting at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
