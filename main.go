package main

import (
	"net/http"

	"database/sql"

	"text/template"


	"github.com/gorilla/mux"
)

var tmpl = template.Must(template.ParseGlob("form/*"))

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", testServer)
	
	r.HandleFunc("/addItem", addItem).Methods("GET")
	r.HandleFunc("/addItem", insertItem).Methods("POST")
	
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
    ID    int
    Name  string
	Price int
	Stock int
	Category_id int
	Details string
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

func addItem(w http.ResponseWriter, r *http.Request) {
    tmpl.ExecuteTemplate(w, "addItem", "")
}

func insertItem(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    if r.Method == "POST" {
        name := r.FormValue("name")
        price := r.FormValue("price")
		stock := r.FormValue("stock")
		category_id := r.FormValue("category_id")
		details := r.FormValue("details")
        insForm, err := db.Prepare("INSERT INTO Items(name, price, stock, category_id, details) VALUES(?,?,?,?,?)")
        if err != nil {
            panic(err.Error())
        }
        insForm.Exec(name, price, stock,category_id,details)
    }
    defer db.Close()
    http.Redirect(w, r, "/", 301)
}
