package main

import (
	"ass2/pkg"
	"encoding/json"
	"html/template"
	"net/http"
)

func MainPageHandler(w http.ResponseWriter, req *http.Request) {
	tmpl, _ := template.ParseFiles("static/index.html")
	tmpl.Execute(w, nil)
}

func signIn(w http.ResponseWriter, req *http.Request) {
	tmpl, _ := template.ParseFiles("static/sign-in.html")
	tmpl.Execute(w, nil)
}

func productList(w http.ResponseWriter, req *http.Request) {
	pkg.ProductDeserializer()
	json.NewEncoder(w).Encode(pkg.ProductDb)
}

func HandleRequest() {
	http.HandleFunc("/", MainPageHandler)
	http.HandleFunc("/sign-in", signIn)
	http.HandleFunc("/products", productList)
	http.ListenAndServe(":8080", nil)
}

func main() {
	HandleRequest()

}
