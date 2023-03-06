package main

import (
	"ass2/pkg"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

func MainPage(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("static/index.html")
	tmpl.Execute(w, nil)
}

func Registration(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("static/sign_up.html")
	username := r.FormValue("name")
	password := r.FormValue("password")
	u1 := pkg.User{Email: username, Password: password}
	pkg.Service{}.CreateUser(u1)
	tmpl.Execute(w, nil)
}

func Authorization(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("static/sign_in.html")
	tmpl.Execute(w, nil)
	username := r.FormValue("name")
	password := r.FormValue("password")
	fmt.Println(username, password)
	u1 := pkg.User{Email: username, Password: password}

	ok := pkg.Service{}.VerifyUser(u1)

	if ok {
		fmt.Println("Success")
		http.Redirect(w, r, "/products", http.StatusSeeOther)
	} else {
		fmt.Println("Error 404!")
	}
}

func ProductList(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("static/list.html")
	pkg.ProductDeserializer()
	tmpl.Execute(w, pkg.ProductDb)
}

func Handler() {
	router := mux.NewRouter()
	router.HandleFunc("/", MainPage)
	router.HandleFunc("/sign-up", Registration)
	router.HandleFunc("/sign-in", Authorization)
	router.HandleFunc("/products", ProductList)
	http.ListenAndServe(":8080", router)
}

func main() {
	fmt.Println("Server starting at server: http://127.0.0.1:8080")
	Handler()
}
