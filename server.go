package main

import (
	"encoding/json"
	"log"
	"net/http"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type TodoItem struct {
	Name string `json:"name"`
	Done bool   `json:"done"`
}

var db, err = gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})

func type_1(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		// Decode the JSON payload if a message was sent
		var newTodo TodoItem
		err := json.NewDecoder(r.Body).Decode(&newTodo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Add the new todo to the database
		db.Table("type_1").Create(&newTodo)
	}

	var todos []TodoItem
	//get from database
	db.Table("type_1").Find(&todos)

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(todos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func type_2(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		// Decode the JSON payload if a message was sent
		var newTodo TodoItem
		err := json.NewDecoder(r.Body).Decode(&newTodo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Add the new todo to the database
		db.Table("type_2").Create(&newTodo)
	}

	var todos []TodoItem
	//get from database
	db.Table("type_2").Find(&todos)

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(todos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func type_3(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		// Decode the JSON payload if a message was sent
		var newTodo TodoItem
		err := json.NewDecoder(r.Body).Decode(&newTodo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Add the new todo to the database
		db.Table("type_3").Create(&newTodo)
	}

	var todos []TodoItem
	//get from database
	db.Table("type_3").Find(&todos)

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(todos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func type_4(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		// Decode the JSON payload if a message was sent
		var newTodo TodoItem
		err := json.NewDecoder(r.Body).Decode(&newTodo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Add the new todo to the database
		db.Table("type_4").Create(&newTodo)
	}

	var todos []TodoItem
	//get from database
	db.Table("type_4").Find(&todos)

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(todos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func toggleDone(w http.ResponseWriter, r *http.Request) {
	type toggleItem struct {
		Table string `json:"table"`
		Name  string `json:"name"`
		Done  bool   `json:"done"`
	}
	var toToggle toggleItem

	err := json.NewDecoder(r.Body).Decode(&toToggle)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	db.Table(toToggle.Table).Where("name = ?", toToggle.Name).Update("done", !toToggle.Done)
}

func delete(w http.ResponseWriter, r *http.Request) {

	type Table struct {
		Table string
	}
	var addressedTable Table
	err := json.NewDecoder(r.Body).Decode(&addressedTable)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db.Table(addressedTable.Table).Where("done = ?", true).Delete(&TodoItem{})

}

func CreateTable(db *gorm.DB, tableName string) {
	table := db.Table(tableName)
	if !table.Migrator().HasTable(tableName) {
		table.AutoMigrate(&TodoItem{})
	}
}

func main() {
	/*----------------Database---------------------------*/
	CreateTable(db, "type_1")
	CreateTable(db, "type_2")
	CreateTable(db, "type_3")
	CreateTable(db, "type_4")

	/*----------------Endpoint Handler---------------------------*/

	http.HandleFunc("/type_1", type_1)
	http.HandleFunc("/type_2", type_2)
	http.HandleFunc("/type_3", type_3)
	http.HandleFunc("/type_4", type_4)
	http.HandleFunc("/toggle", toggleDone)
	http.HandleFunc("/delete", delete)
	/*----------------Server---------------------------*/
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.Handle("/", http.FileServer(http.Dir("static")))
	log.Fatal(http.ListenAndServe(":8300", nil))
}
