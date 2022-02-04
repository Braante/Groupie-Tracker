package groupie

import (
	"strconv"
)

func filtered(creation_date string, creation_album string, nb_members []string, locations_choose string) []Artist {
	APIRequest()
	APIRequestLocation()
	var final_array []Artist

	crea_date_nb, _ := strconv.Atoi(creation_date)
	crea_album_nb, _ := strconv.Atoi(creation_album)

	for k := range ArtistsTab {
		good_places := false
		nb_good_members := false
		crea_album_temp := datetonum(ArtistsTab[k].FirstAlbum)
		crea_date_temp := ArtistsTab[k].CreationDate
		nb_members_temp := strconv.FormatInt(int64(len(ArtistsTab[k].Members)), 10)
		if locations_choose != "" {
			APIRequestLocationOne(ArtistsTab[k].Linkloca)
			for y := range LocationsGroupTab.OneLocationsGroup {
				if LocationsGroupTab.OneLocationsGroup[y] == locations_choose {
					good_places = true
				}
			}
		} else {
			good_places = true
		}
		if creation_date == "" {
			crea_date_temp = 0
			crea_date_nb = 0
		}
		if creation_album == "" {
			crea_album_temp = 0
			crea_album_nb = 0
		}
		if nb_members == nil {
			nb_good_members = true
		} else {
			for k := range nb_members {
				if nb_members_temp == nb_members[k] {
					nb_good_members = true
				}
			}
		}
		if crea_date_nb == crea_date_temp && crea_album_nb == crea_album_temp && nb_good_members && good_places {
			final_array = append(final_array, ArtistsTab[k])
		}
	}
	return final_array
}
