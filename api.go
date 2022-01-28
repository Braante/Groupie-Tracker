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
}

var ArtistsTab []Artist

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
