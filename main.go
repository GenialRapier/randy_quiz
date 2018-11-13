package main

import (
	"database/sql"
	// "fmt"
	"net/http"
	"text/template"
	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"
)

type Items struct {
    Id    int
    Name  string
	Price int
	Stock int
	Category_id int
	Details string
}

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
	r.HandleFunc("/addIncoming", postAddIncoming).Methods("POST")
	r.HandleFunc("/addOutgoing", testServer).Methods("GET")
	r.HandleFunc("/editHistory", testServer).Methods("GET")

	http.ListenAndServe(":14045", r)
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

func testServer(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "home", "")
}

func getAddIncoming(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	result, err := db.Query("SELECT * FROM Items")
	if err != nil {panic(err.Error())}

	item := Items{}
	itemList := []Items{}

	for result.Next() {
		var id, price, stock ,category_id int
		var name, details string

		err = result.Scan(&id, &name, &price, &stock, &category_id, &details)
        if err != nil { panic(err.Error()) }

        item.Id = id
        item.Name = name
        item.Price = price
        item.Stock = stock
        item.Category_id = category_id
        item.Details = details

        itemList = append(itemList, item)
	}

	tmpl.ExecuteTemplate(w, "IncomingStock", itemList)
	defer db.Close()
}

func postAddIncoming(w http.ResponseWriter, r * http.Request) {
	db := dbConn()

	itemId := r.FormValue("item_id")
    quantity := r.FormValue("quantity")
    details := r.FormValue("details")

    query, err := db.Prepare("INSERT INTO histories(items_id, date, movement, quantity, details) VALUES(?,NOW(),'incoming',?,?)")
    if err != nil { panic(err.Error()) }
    query.Exec(itemId, quantity, details)

    defer db.Close()
}
