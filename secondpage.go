package groupie

import (
	"html/template"
	"net/http"
	"strconv"
)

func MainHandlersec(w http.ResponseWriter, r *http.Request) {
	tmplsec := template.Must(template.ParseFiles("static/second.html"))

	APIRequest()
	var idcard int
	switch r.Method {
	case "GET":
		http.Redirect(w, r, "/", 301)
	case "POST":

		idcardstr := r.FormValue("id")
		idcard, _ = strconv.Atoi(idcardstr)
	}
	Tab := ArtistsTab
	CardChoose := Tab[idcard-1]
	data := Artist{
		Image:        CardChoose.Image,
		Name:         CardChoose.Name,
		Members:      CardChoose.Members,
		CreationDate: CardChoose.CreationDate,
		FirstAlbum:   CardChoose.FirstAlbum,
	}

	tmplsec.Execute(w, data)
}
