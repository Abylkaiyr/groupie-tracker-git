package filter

import (
	"Abylkaiyr/groupie-tracker/internals/models"
	"fmt"
	"strconv"
	"sync"

	grabjson "Abylkaiyr/groupie-tracker/internals/grabJson"
)

func FilterOut(memberNum, creationDate, firstAlbum []string, location string) ([]models.Artist, error) {
	members, errMemberConv := StrArrToIntArr(memberNum)
	if errMemberConv != nil {
		return nil, fmt.Errorf("Invalid input for member filter: %w", errMemberConv)
	}

	createdFrom, errFrom := strconv.Atoi(creationDate[0])
	createdTo, errTo := strconv.Atoi(creationDate[1])
	if errFrom != nil || errTo != nil || createdTo < createdFrom {
		return nil, fmt.Errorf("Invalid input for creation date filter")
	}

	albumReleaseFrom, errFrom := strconv.Atoi(firstAlbum[0])
	albumReleaseTo, errTo := strconv.Atoi(firstAlbum[1])
	if errFrom != nil || errTo != nil || albumReleaseTo < albumReleaseFrom {
		return nil, fmt.Errorf("Invalid input for album release date filter")
	}

	var artists []models.Artist
	if err := grabjson.GetQuickArtistData(&artists); err != nil {
		return nil, fmt.Errorf("error in getting list of artists: %w", err)
	}

	var wg sync.WaitGroup
	var mu sync.Mutex
	sem := make(chan struct{}, 10)

	var result []models.Artist

	for i := 0; i < len(artists); i++ {
		wg.Add(1)
		sem <- struct{}{}

		go func(j int) {
			defer wg.Done()
			artist := artists[j]
			if (artist.CreationDate >= createdFrom && artist.CreationDate <= createdTo) &&
				CheckAlbum(artist.FirstAlbum, albumReleaseFrom, albumReleaseTo) &&
				CheckMembers(artist.Members, members) &&
				CheckLocation(artist.ID, location) {
				mu.Lock()
				result = append(result, artist)
				mu.Unlock()

			}
			<-sem
		}(i)
	}

	wg.Wait()

	return result, nil
}

func StrArrToIntArr(mem []string) ([]int, error) {
	res := []int{}
	for _, i := range mem {
		num, err := strconv.Atoi(i)
		if err != nil {
			return nil, fmt.Errorf("invalid input in members number: %w", err)
		}
		res = append(res, num)
	}
	return res, nil
}
