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
}

type Relation struct {
	DatesLocations map[string][]string `json:"datesLocations"`
}

type Location struct {
	LocationsPlaces []string `json:"locations"`
}

var ArtistsTab []Artist

var RelationsTab = Relation{}

var LocationTabs = Location{}

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
