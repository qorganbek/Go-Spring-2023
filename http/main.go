package main

import (
	"fmt"
	"io"
	"net/http"
)

func HelloView(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1> Hello World! </h1>")
}

func MainView(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>Main Page</h1>")
}


// admin panel
// sign up, sign in
// product


func main() {
	fmt.Println("Server started at this url http://localhost:8080/")
	http.HandleFunc("/", MainView)
	http.HandleFunc("/hello", HelloView)
	http.ListenAndServe(":8080", nil)
}
