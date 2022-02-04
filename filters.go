package groupie

import (
	"strconv"
)

type Filters struct {
	MinCreation      int
	MaxCreation      int
	MinCreationAlbum int
	MaxCreationAlbum int
	NbMembers        []int
	AllLocations     []string
}

var MinCreation int
var MaxCreation int
var MinCreationAlbum int
var MaxCreationAlbum int
var NbMembers []int
var AllLocations []string

func filters() {

	APIRequest()
	APIRequestLocation()

	MinCreation = 3000
	MaxCreation = 1000
	MinCreationAlbum = 3000
	MaxCreationAlbum = 1000
	var tableauplaces []string
	tableauplaces = append(tableauplaces, "")
	location_id := LocationTabs.Index
	for k := range location_id {
		for y := 0; y < len(location_id[k].LocationsPlaces); y++ {
			tableauplaces = append(tableauplaces, location_id[k].LocationsPlaces[y])
		}
	}

	for k := range ArtistsTab {
		if MinCreation > ArtistsTab[k].CreationDate {
			MinCreation = ArtistsTab[k].CreationDate
		}
		if MaxCreation < ArtistsTab[k].CreationDate {
			MaxCreation = ArtistsTab[k].CreationDate
		}
		if MinCreationAlbum > datetonum(ArtistsTab[k].FirstAlbum) {
			MinCreationAlbum = datetonum(ArtistsTab[k].FirstAlbum)
		}
		if MaxCreationAlbum < datetonum(ArtistsTab[k].FirstAlbum) {
			MaxCreationAlbum = datetonum(ArtistsTab[k].FirstAlbum)
		}
	}
	AllLocations = removeDuplicateStr(tableauplaces)
	NbMembers = []int{1, 2, 3, 4, 5, 6, 7, 8}
}

func datetonum(str string) int {
	temp := str[len(str)-4:]
	nb, _ := strconv.Atoi(temp)
	return nb
}

func removeDuplicateStr(strDupl []string) []string {
	all := make(map[string]bool)
	list := []string{}
	for _, k := range strDupl {
		if _, value := all[k]; !value {
			all[k] = true
			list = append(list, k)
		}
	}
	return list
}
