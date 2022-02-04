package groupie

import "strconv"

func filtered(creation_date string, creation_album string, nb_members string, locations_choose string) []Artist {
	APIRequest()
	var test []Artist
	crea_date, _ := strconv.Atoi(creation_date)
	for k := range ArtistsTab {
		if crea_date == ArtistsTab[k].CreationDate {
			test = append(test, ArtistsTab[k])
		}
	}

	return test
}
