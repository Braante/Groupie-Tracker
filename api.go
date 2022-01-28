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
	Relation     []string
}

type Relation struct {
	DatesLocations map[string][]string `json:"datesLocations"`
}

var ArtistsTab []Artist

var RelationsTab = Relation{}

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
	test := json.Unmarshal(readOneRela, &RelationsTab)
	fmt.Println(test)
	fmt.Println(RelationsTab)
}
