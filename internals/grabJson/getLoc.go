package grabjson

import (
	"sort"

	model "Abylkaiyr/groupie-tracker/internals/models"
)

func GetUniqueLocations() ([]string, error) {
	var locations model.LocationInfo
	urlLocations := model.Url + "/" + "locations"

	if err := GetJson(urlLocations, &locations); err != nil {
		return nil, err
	}
	return unique(locations.LocationInfo), nil
}

func unique(locationData []model.LocationData) []string {
	keys := make(map[string]bool)
	list := []string{}

	for _, locations := range locationData {
		for _, location := range locations.Locations {
			if _, value := keys[location]; !value {
				keys[location] = true
				list = append(list, location)
			}
		}
	}
	sort.Strings(list)
	return list
}
