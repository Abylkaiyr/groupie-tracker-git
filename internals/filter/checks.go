package filter

import (
	"strconv"

	grabjson "Abylkaiyr/groupie-tracker/internals/grabJson"
)

func CheckAlbum(artistAlbum string, filterDateFrom, filterDateTo int) bool {
	yearDate := artistAlbum[6:]
	releaseDate, err := strconv.Atoi(yearDate)
	if err != nil {
		return false
	}

	if releaseDate >= filterDateFrom && releaseDate <= filterDateTo {
		return true
	}

	return false
}

func CheckMembers(members []string, membersArr []int) bool {
	memNum := len(members)

	if len(membersArr) == 0 {
		return true
	}

	for _, v := range membersArr {
		if v == memNum {
			return true
		}
	}
	return false
}

func CheckLocation(id int, location string) bool {
	if location == "" {
		return true
	}
	return grabjson.GetLocation(id)[location]
}
