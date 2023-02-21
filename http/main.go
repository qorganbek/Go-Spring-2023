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
	
	fmt.Println("Server is listening on localhost:8080, open your browser on http://localhost:8080/")
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	http.HandleFunc("/about", func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "About Page")
	})

	http.HandleFunc("/contact", func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "<h1>Contact Page</h1>")
	})

	http.HandleFunc("/person", func(w http.ResponseWriter, req *http.Request) {
		name := req.URL.Query().Get("name")
		age := req.URL.Query().Get("age")
		fmt.Fprintf(w, "Name: %s Age: %s", name, age)
	})

	http.ListenAndServe(":8080", nil)
}
