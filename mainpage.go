package groupie

import (
	"html/template"
	"net/http"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("static/index.html"))

	APIRequest()

	data := ArtistStruct{
		Tab: ArtistsTab,
	}

	tmpl.Execute(w, data)
}
