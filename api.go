package groupie

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ArtistStruct struct {
	Tab []Artist
}

type Artist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Link         string   `json:"relations"`
	Linkloca     string   `json:"locations"`
	Relation     []string
	LocationsLat []string
	LocationsLon []string
}

type Relation struct {
	DatesLocations map[string][]string `json:"datesLocations"`
}

type Location struct {
	Index []struct {
		LocationsPlaces []string `json:"locations"`
	} `json:"index"`
}

type LocationsGroup struct {
	OneLocationsGroup []string `json:"locations"`
}

type Coord struct {
	Lat string `json:"lat"`
	Lon string `json:"lon"`
}

var CoordTab []Coord

var ArtistsTab []Artist

var RelationsTab = Relation{}

var LocationTabs = Location{}

var LocationsGroupTab = LocationsGroup{}

func APIRequest() {

	req, errOne := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if errOne != nil {
		fmt.Println(errOne)
	}

	readOne, errOneA := ioutil.ReadAll(req.Body)
	if errOneA != nil {
		fmt.Println(errOneA)
	}
	json.Unmarshal(readOne, &ArtistsTab)

}

func APIRequestRelation(link string) {
	reqRela, errTwo := http.Get(link)
	if errTwo != nil {
		fmt.Println(errTwo)
	}

	readOneRela, errOneR := ioutil.ReadAll(reqRela.Body)
	if errOneR != nil {
		fmt.Println(errOneR)
	}
	json.Unmarshal(readOneRela, &RelationsTab)
}

func APIRequestLocation() {
	req, errOne := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	if errOne != nil {
		fmt.Println(errOne)
	}

	readOne, errOneA := ioutil.ReadAll(req.Body)
	if errOneA != nil {
		fmt.Println(errOneA)
	}
	json.Unmarshal(readOne, &LocationTabs)
}

func APIRequestLocationOne(link string) {
	req, errOne := http.Get(link)
	if errOne != nil {
		fmt.Println(errOne)
	}

	readOne, errOneA := ioutil.ReadAll(req.Body)
	if errOneA != nil {
		fmt.Println(errOneA)
	}
	json.Unmarshal(readOne, &LocationsGroupTab)
}

func APICoord(link string) {
	req, errOne := http.Get("https://nominatim.openstreetmap.org/search/" + link + "?format=json&limit=1")
	if errOne != nil {
		fmt.Println(errOne)
	}

	readOne, errOneA := ioutil.ReadAll(req.Body)
	if errOneA != nil {
		fmt.Println(errOneA)
	}
	json.Unmarshal(readOne, &CoordTab)
}
