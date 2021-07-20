package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// Employee : class with id, name and city
type Employee struct {
	ID   int
	Name string
	City string
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "lingga.putra"
	dbPassword := "a5k4CooL"
	dbName := "GO_TEST"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPassword+"@/"+dbName)

	if err != nil {
		panic(err.Error())
	}

	return db
}

var tmpl = template.Must(template.ParseGlob("form/*"))

// Index : function to handle request and response
func Index(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM Employee ORDER BY id ASC")

	if err != nil {
		panic(err.Error())
	}

	emp := Employee{}
	res := []Employee{}

	for selDB.Next() {
		var id int
		var name, city string
		err = selDB.Scan(&id, &name, &city)

		if err != nil {
			panic(err.Error())
		}

		emp.ID = id
		emp.Name = name
		emp.City = city
		res = append(res, emp)
	}

	tmpl.ExecuteTemplate(w, "Index", res)
	defer db.Close()
}

// Show : function to handle request and response
func Show(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nID := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM Employee WHERE id=?", nID)

	if err != nil {
		panic(err.Error())
	}
	emp := Employee{}

	for selDB.Next() {
		var id int
		var name, city string
		err = selDB.Scan(&id, &name, &city)
		if err != nil {
			panic(err.Error())
		}
		emp.ID = id
		emp.Name = name
		emp.City = city
	}

	tmpl.ExecuteTemplate(w, "Show", emp)
	defer db.Close()
}

// New : funtion to create new data
func New(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}

// Edit : funtion to edit data
func Edit(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nID := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM Employee WHERE id=?", nID)

	if err != nil {
		panic(err.Error())
	}

	emp := Employee{}

	for selDB.Next() {
		var id int
		var name, city string
		err = selDB.Scan(&id, &name, &city)

		if err != nil {
			panic(err.Error())
		}

		emp.ID = id
		emp.Name = name
		emp.City = city
	}

	tmpl.ExecuteTemplate(w, "Edit", emp)
	defer db.Close()
}

// Insert : function to add new data
func Insert(w http.ResponseWriter, r *http.Request) {
	db := dbConn()

	if r.Method == "POST" {
		name := r.FormValue("name")
		city := r.FormValue("city")
		insForm, err := db.Prepare("INSERT INTO Employee(name, city) VALUES(?,?)")

		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(name, city)
		log.Println("INSERT: Name: " + name + " | City: " + city)
	}

	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

// Update : function to edit data
func Update(w http.ResponseWriter, r *http.Request) {
	db := dbConn()

	if r.Method == "POST" {
		name := r.FormValue("name")
		city := r.FormValue("city")
		id := r.FormValue("uid")
		insForm, err := db.Prepare("UPDATE Employee SET name=?, city=? WHERE id=?")

		if err != nil {
			panic(err.Error())
		}

		insForm.Exec(name, city, id)
		log.Println("UPDATE: Name: " + name + " | City: " + city)
	}

	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

// Delete : function to delete data
func Delete(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	emp := r.URL.Query().Get("id")
	delForm, err := db.Prepare("DELETE FROM Employee WHERE id=?")

	if err != nil {
		panic(err.Error())
	}

	delForm.Exec(emp)
	log.Println("DELETE")
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func main() {
	log.Println("Server started on: http://localhost:8080")
	http.HandleFunc("/", Index)
	http.HandleFunc("/show", Show)
	http.HandleFunc("/new", New)
	http.HandleFunc("/edit", Edit)
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/delete", Delete)
	http.ListenAndServe(":8080", nil)
}
