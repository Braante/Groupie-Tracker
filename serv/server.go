package main

import (
	"fmt"
	"groupie"
	"net/http"
	"text/template"
)

type ToDoPage struct {
	PageTitle string
}

func main() {
	tmpl := template.Must(template.ParseFiles("static/index.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		groupie.MainHandler(w, r)
		data := ToDoPage{}
		tmpl.Execute(w, data)
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	err := http.ListenAndServe("localhost:8080", nil)

	if err != nil {
		fmt.Println(err)
	}
}
