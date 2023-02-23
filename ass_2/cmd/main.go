package main

import (
	"ass2/pkg"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "<h1>Hello Go Http</h1>")
	})
	http.HandleFunc("/admin", func(w http.ResponseWriter, req *http.Request) {
		_, err := io.WriteString(w, "Admin Panel")
		if err != nil {
			return
		}
	})
	http.HandleFunc("/user", func(w http.ResponseWriter, req *http.Request) {
		user := req.URL.Query().Get("username")
		password := req.URL.Query().Get("password")
		fmt.Fprintf(w, "Username: %s, Password: %s", user, password)
	})

	http.HandleFunc("/products", func(w http.ResponseWriter, req *http.Request) {
		pkg.ProductDeserializer()
		json.NewEncoder(w).Encode(pkg.ProductDb)
	})

	http.ListenAndServe(":8080", nil)
}
