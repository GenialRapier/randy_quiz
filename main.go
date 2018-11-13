package main

import (
	"database/sql"
	"net/http"
	"text/template"
	"github.com/gorilla/mux"
)

var tmpl = template.Must(template.ParseGlob("form/*"))

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", testServer)

	r.HandleFunc("/addItem", testServer).Methods("GET")
	r.HandleFunc("/removeItem", testServer).Methods("GET")
	r.HandleFunc("/editItem", testServer).Methods("GET")
	r.HandleFunc("/deleteItem", testServer).Methods("GET")

	r.HandleFunc("/addCategory", testServer).Methods("GET")
	r.HandleFunc("/removeCategory", testServer).Methods("GET")
	r.HandleFunc("/editCategory", testServer).Methods("GET")
	r.HandleFunc("/deleteCategory", testServer).Methods("GET")

	r.HandleFunc("/addIncoming", getAddIncoming).Methods("GET")
	r.HandleFunc("/addOutgoing", testServer).Methods("GET")
	r.HandleFunc("/editHistory", testServer).Methods("GET")

	http.ListenAndServe(":14045", r)
}

func testServer(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "home", "")
}

func getAddIncoming(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "IncomingStock", "")
}

type Items struct {
    id    int
    name  string
	price int
	stock int
	category_id int
	details string
}

func dbConn() (db *sql.DB) {
    dbDriver := "mysql"
    dbUser := "root"
    dbPass := ""
    dbName := "quiz3"
    db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
    if err != nil {
        panic(err.Error())
    }
    return db
}
