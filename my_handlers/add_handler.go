package my_handlers

import (
	"encoding/json"
	"net/http"
)

func all_types(w http.ResponseWriter, r *http.Request, used_type string) {
	if r.Method == http.MethodPost {

		// Decode the JSON payload if a message was sent
		var newTodo TodoItem
		err := json.NewDecoder(r.Body).Decode(&newTodo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Add the new todo to the database
		db.Table(used_type).Create(&newTodo)
	}

	var todos []TodoItem
	//get from database
	db.Table(used_type).Find(&todos)

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(todos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Type_1(w http.ResponseWriter, r *http.Request) {
	all_types(w, r, "type_1")
}

func Type_2(w http.ResponseWriter, r *http.Request) {
	all_types(w, r, "type_2")
}

func Type_3(w http.ResponseWriter, r *http.Request) {
	all_types(w, r, "type_3")
}

func Type_4(w http.ResponseWriter, r *http.Request) {
	all_types(w, r, "type_4")
}
