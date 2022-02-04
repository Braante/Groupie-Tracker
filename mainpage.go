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
	Hide             string
}

func MainHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("static/index.html"))
	hide := ""
	switch r.Method {
	case "GET":
		APIRequest()
		filters()

	case "POST":
		hide = "reset"
		creation_date := r.FormValue("creation")
		creation_album := r.FormValue("creationAlbum")
		nb_members := r.FormValue("nb_members[]")
		locations_choose := r.FormValue("locations")
		fmt.Println("creationdate:", creation_date, "creationalbum:", creation_album, "nbmembers:", nb_members, "locations:", locations_choose)
		ArtistsTab = filtered(creation_date, creation_album, nb_members, locations_choose)
	}
	data := &Data{
		Tab:              ArtistsTab,
		MinCreation:      MinCreation,
		MaxCreation:      MaxCreation,
		MinCreationAlbum: MinCreationAlbum,
		MaxCreationAlbum: MaxCreationAlbum,
		NbMembers:        NbMembers,
		AllLocations:     AllLocations,
		Hide:             hide,
	}

	tmpl.Execute(w, data)
}
