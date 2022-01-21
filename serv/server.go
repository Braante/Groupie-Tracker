package main

import (
	"net/http"
	"text/template"
)

type ToDoPage struct {
	PageTitle string
}

func main() {
	tmpl := template.Must(template.ParseFiles("static/index.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := ToDoPage{}
		tmpl.Execute(w, data)
	})
	http.ListenAndServe("localhost:8080", nil)
}
