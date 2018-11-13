package main

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

var tmpl = template.Must(template.ParseGlob("form/*"))

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/test", testServer)

	r.HandleFunc("/addItem", testServer).Methods("GET")
	r.HandleFunc("/removeItem", testServer).Methods("GET")
	r.HandleFunc("/editItem", testServer).Methods("GET")
	r.HandleFunc("/deleteItem", testServer).Methods("GET")

	r.HandleFunc("/addCategory", testServer).Methods("GET")
	r.HandleFunc("/removeCategory", testServer).Methods("GET")
	r.HandleFunc("/editCategory", testServer).Methods("GET")
	r.HandleFunc("/deleteCategory", testServer).Methods("GET")

	http.ListenAndServe(":14045", r)
}

func testServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Server is Up")
}
