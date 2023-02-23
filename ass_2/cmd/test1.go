package main

import (
	"ass2/pkg"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"io/ioutil"
	"net/http"
	"strconv"
)

func mainpg(w http.ResponseWriter, req *http.Request) {
	tmpl, _ := template.ParseFiles("static/index.html")
	tmpl.Execute(w, nil)
}

func login(w http.ResponseWriter, req *http.Request) {
	tmpl, _ := template.ParseFiles("static/sign-in.html")
	tmpl.Execute(w, nil)
}

func List(w http.ResponseWriter, req *http.Request) {
	pkg.ProductDeserializer()
	json.NewEncoder(w).Encode(pkg.ProductDb)
}

func Retrieve(w http.ResponseWriter, req *http.Request) {
	pkg.ProductDeserializer()
	vars := mux.Vars(req)
	key, _ := strconv.Atoi(vars["id"])
	for _, i := range pkg.ProductDb {
		if i.ID == key {
			json.NewEncoder(w).Encode(i)
		}
	}
}

func DelProd(w http.ResponseWriter, req *http.Request) {
	body, _ := ioutil.ReadAll(req.Body)
	//fmt.Fprintf(w, string(body))
	var data pkg.Product
	json.Unmarshal([]byte(body), &data)
	//for key, _ := range data {
	//	fmt.Println(key)
	//}
	pkg.ProductDeserializer()
	pkg.Service{}.DeleteProduct(data)
}

func CreateProd(w http.ResponseWriter, req *http.Request) {
	body, _ := ioutil.ReadAll(req.Body)
	var data pkg.Product
	json.Unmarshal([]byte(body), &data)
	pkg.ProductDeserializer()
	pkg.Service{}.CreateProduct(data)
}

func MainHandler() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", mainpg)
	myRouter.HandleFunc("/sign-in", login)
	myRouter.HandleFunc("/products", List).Methods("GET")
	myRouter.HandleFunc("/products", DelProd).Methods("POST")
	myRouter.HandleFunc("/products/{id}", Retrieve)
	myRouter.HandleFunc("/create", CreateProd).Methods("POST")
	http.ListenAndServe(":8080", myRouter)
}

/*
EndPoints:

	/sign in,
	/sign up
	/ -> index.html
	/product -> ProductList, filters
	/admin -> User, Product
	/admin/user crud
	/admin/product crud
*/
func main() {
	fmt.Println("Server starting at server: http://127.0.0.1:8080")
	MainHandler()

}
