package main

import (
	"log"
	"net/http"
	"portfolio/Eisenhower_Matrix/my_handlers"
)

func main() {

	/*--------------------Database-----------------------*/
	my_handlers.ToCreate("type_1", "type2", "type_3", "type_4")

	/*----------------Endpoint Handler-------------------*/
	http.HandleFunc("/type_1", my_handlers.Type_1)
	http.HandleFunc("/type_2", my_handlers.Type_2)
	http.HandleFunc("/type_3", my_handlers.Type_3)
	http.HandleFunc("/type_4", my_handlers.Type_4)
	http.HandleFunc("/toggle", my_handlers.ToggleDone)
	http.HandleFunc("/delete", my_handlers.Delete)

	/*-------------------Server---------------------------*/
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.Handle("/", http.FileServer(http.Dir("static")))
	log.Fatal(http.ListenAndServe(":8400", nil))
}
