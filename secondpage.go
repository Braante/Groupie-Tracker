package groupie

import (
	"html/template"
	"net/http"
	"strconv"
	"strings"
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
		TabA := ArtistsTab
		CardChoose := TabA[idcard-1]
		APIRequestRelation(CardChoose.Link)
		date := RelationsTab.DatesLocations
		RelationsTab.DatesLocations = nil
		var tableauConcerts []string
		var tableauLat []string
		var tableauLon []string
		for lieu := range date {
			h := date[lieu]
			listdate := strings.Join(h, " ")
			temp := lieu + ": " + listdate
			tableauConcerts = append(tableauConcerts, temp)
			APICoord(lieu)
			swap := CoordTab[0].Lat
			tableauLat = append(tableauLat, swap)
			swap = CoordTab[0].Lon
			tableauLon = append(tableauLon, swap)
		}
		data := Artist{
			Image:        CardChoose.Image,
			Name:         CardChoose.Name,
			Members:      CardChoose.Members,
			CreationDate: CardChoose.CreationDate,
			FirstAlbum:   CardChoose.FirstAlbum,
			Relation:     tableauConcerts,
			LocationsLat: tableauLat,
			LocationsLon: tableauLon,
		}
		tmplsec.Execute(w, data)
	}
}
