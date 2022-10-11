package grabjson

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"sync"

	"Abylkaiyr/groupie-tracker/internals/models"
	model "Abylkaiyr/groupie-tracker/internals/models"
)

func GetQuickArtistData(artist *[]model.Artist) error {
	// getting whole artists

	urlArtists := model.Url + "/" + "artists"

	if err := GetJson(urlArtists, artist); err != nil {
		return err
	}

	return nil
}

func GetDetailedData(i int, u *model.Artist) error {
	// Artist
	id := strconv.Itoa(i)
	var relations model.Relations
	var errArtist, errRelation error

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		urlArtist := model.Url + "/" + "artists" + "/" + id
		errArtist = GetJson(urlArtist, u)
	}()

	go func() {
		defer wg.Done()
		urlReleations := model.Url + "/" + "relation" + "/" + strconv.Itoa(i)
		errRelation = GetJson(urlReleations, &relations)
	}()

	wg.Wait()

	if errArtist != nil || errRelation != nil {
		return fmt.Errorf("error in getting detailed artist info")
	}

	u.DatesLocations = relations

	return nil
}

func GetLocation(id int) map[string]bool {
	tr, err := http.Get(model.Url + "/" + "locations" + "/" + strconv.Itoa(id))
	if err != nil {
		log.Println("can not get Locations url")
	}
	defer tr.Body.Close()
	dataLoc, err := io.ReadAll(tr.Body)
	if err != nil {
		log.Println("Error in reading detail of the artist", err)
	}
	var locations models.LocationData
	err = json.Unmarshal(dataLoc, &locations)
	if err != nil {
		log.Println("Error in unmarshalling json data in detail locatioons of the artist", err)
	}

	locationCheck := make(map[string]bool)

	for _, location := range locations.Locations {
		locationCheck[location] = true
	}
	return locationCheck
}
