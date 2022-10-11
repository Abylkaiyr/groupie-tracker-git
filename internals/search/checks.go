package search

import (
	"strconv"
	"strings"

	grabjson "Abylkaiyr/groupie-tracker/internals/grabJson"
)

func SearchMembers(Members []string, searchTag string) bool {
	for _, member := range Members {
		if strings.ToLower(member) == searchTag {
			return true
		}
	}

	return false
}

func SearchCreationDate(createDate int, searchTag string) bool {
	strDate := strconv.Itoa(createDate)

	if strDate == searchTag {
		return true
	}

	return false
}

func SearchLocation(id int, location string) bool {
	locMap := grabjson.GetLocation(id)

	for key := range locMap {
		if strings.Contains(key, location) {
			return true
		}
	}

	return false
}
