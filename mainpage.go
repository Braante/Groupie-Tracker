package groupie

import (
	"fmt"
	"html/template"
	"net/http"
)

type Data struct {
	Tab              []Artist
	MinCreation      int
	MaxCreation      int
	MinCreationAlbum int
	MaxCreationAlbum int
	NbMembers        []int
	AllLocations     []string
}

func MainHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("static/index.html"))

	switch r.Method {
	case "GET":
		APIRequest()
		filters()

	case "POST":
		creation_date := r.FormValue("creation")
		creation_album := r.FormValue("creationAlbum")
		nb_members := r.FormValue("nb_members")
		locations_choose := r.FormValue("locations")
		fmt.Println("creationdate:", creation_date, "creationalbum:", creation_album, "nbmembers:", nb_members, "locations:", locations_choose)

	}
	data := &Data{
		Tab:              ArtistsTab,
		MinCreation:      MinCreation,
		MaxCreation:      MaxCreation,
		MinCreationAlbum: MinCreationAlbum,
		MaxCreationAlbum: MaxCreationAlbum,
		NbMembers:        NbMembers,
		AllLocations:     AllLocations,
	}

	tmpl.Execute(w, data)
}
