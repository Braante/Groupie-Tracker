package groupie

import (
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
	Res              string
}

func MainHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("static/index.html"))
	hide := ""
	no_res := ""
	switch r.Method {
	case "GET":
		APIRequest()
		filters()
		Test()

	case "POST":
		hide = "reset"
		creation_date := r.FormValue("creation")
		creation_album := r.FormValue("creationAlbum")
		nb_members := r.Form["nb_members[]"]
		locations_choose := r.FormValue("locations")
		//fmt.Println("creationdate:", creation_date, "creationalbum:", creation_album, "nbmembers:", nb_members, "locations:", locations_choose)
		ArtistsTab = filtered(creation_date, creation_album, nb_members, locations_choose)
		if len(ArtistsTab) == 0 {
			no_res = "Sorry, No results found :("
		}
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
		Res:              no_res,
	}

	tmpl.Execute(w, data)
}
