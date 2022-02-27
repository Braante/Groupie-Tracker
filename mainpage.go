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
	Reset            string
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

	case "POST":
		hide = "Back to Main Page"
		search_temp := ""
		name := ""
		creation_date := r.FormValue("creation")
		creation_album := r.FormValue("creationAlbum")
		nb_members := r.Form["nb_members[]"]
		locations_choose := r.FormValue("locations")
		search_temp = r.FormValue("searchtemp")
		if search_temp != "" {
			if len(search_temp) == 4 && (search_temp[0] == '1' || search_temp[0] == '2') {
				creation_date = search_temp
			} else if len(search_temp) == 10 && search_temp[2] == '-' {
				creation_album = search_temp[6:]
			} else {
				for chr := range search_temp {
					if search_temp[chr] == '-' {
						locations_choose = search_temp
						break
					}
				}
				if locations_choose != search_temp {
					name = search_temp
				}
			}
		}
		//fmt.Println("creationdate:", creation_date, "creationalbum:", creation_album, "nbmembers:", nb_members, "locations:", locations_choose)
		ArtistsTab = filtered(creation_date, creation_album, nb_members, locations_choose, name)
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
		Reset:            hide,
		Res:              no_res,
	}

	tmpl.Execute(w, data)
}
