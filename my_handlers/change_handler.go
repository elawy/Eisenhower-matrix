package my_handlers

import (
	"encoding/json"
	"net/http"
)

func ToggleDone(w http.ResponseWriter, r *http.Request) {
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

func Delete(w http.ResponseWriter, r *http.Request) {

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
