package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB

type invitation struct {
	id             int    `json:"id,omitempty"`
	full_name      string `json:"full_name,omitempty"`
	emaill_address string `json:"emaill_address,omitempty"`
	referred_by    string `json:"referred_by,omitempty"`
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/student", getStudent).Methods("GET")
	r.HandleFunc("/student", storeStudent).Methods("POST")
	r.HandleFunc("/student", updateStudent).Methods("PUT")
	r.HandleFunc("/student", deleteStudent).Methods("DELETE")

	http.ListenAndServe(":3000", r)
}

func connectDatabase() *sql.DB {
	db, err := sql.Open("mysql", "root:password123!@(eightxmancity-instance.cbzeawye0kof.ap-southeast-1.rds.amazonaws.com:3306)/eightxmancitydb?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func getStudent(w http.ResponseWriter, r *http.Request) {
	db = connectDatabase()
	rows, err := db.Query("select * from invitations")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	inv := invitation{}
	response := []invitation{}

	for rows.Next() {
		rows.Scan(&inv.id, &inv.full_name, &inv.emaill_address, &inv.referred_by)
		response = append(response, inv)
		json.NewEncoder(w).Encode(inv)
		return
	}

}

func storeStudent(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "POST REQUEST")
}

func updateStudent(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "PUT REQUEST")
}

func deleteStudent(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "DELETE REQUEST")
}
