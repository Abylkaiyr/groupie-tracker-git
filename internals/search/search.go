package search

import (
	"fmt"
	"strings"
	"sync"

	grabjson "Abylkaiyr/groupie-tracker/internals/grabJson"
	"Abylkaiyr/groupie-tracker/internals/models"
)

func Search(searchTag string) ([]models.Artist, error) {
	var artists []models.Artist
	if err := grabjson.GetQuickArtistData(&artists); err != nil {
		return nil, fmt.Errorf("error in getting list of artists in search: %w", err)
	}

	var searchArtist []models.Artist
	searchTag = strings.ToLower(strings.TrimSpace(searchTag))

	var wg sync.WaitGroup
	var mu sync.Mutex
	sem := make(chan struct{}, 10)

	for i := range artists {
		wg.Add(1)
		sem <- struct{}{}

		go func(j int) {
			defer wg.Done()
			artist := artists[j]
			if strings.ToLower(strings.TrimSpace(artist.Name)) == searchTag || artist.FirstAlbum == searchTag ||
				SearchCreationDate(artist.CreationDate, searchTag) ||
				SearchLocation(artist.ID, searchTag) ||
				SearchMembers(artist.Members, searchTag) {
				mu.Lock()
				searchArtist = append(searchArtist, artist)
				mu.Unlock()
			}
			<-sem
		}(i)
	}

	wg.Wait()

	return searchArtist, nil
}
