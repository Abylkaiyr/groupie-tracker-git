package geolocalize

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	googleApiUri = "https://maps.googleapis.com/maps/api/geocode/json?key=AIzaSyDUaK5Jrw_3EFa0nS7q_btShDORnSf2EOs&address="
)

type googleApiResponse struct {
	Results Results `json:"results"`
}

type Results []Geometry

type Geometry struct {
	Geometry Location `json:"geometry"`
}

type Location struct {
	Location Coordinates `json:"location"`
}

type Coordinates struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lng"`
}

func GetCityCoordinates(city string) Coordinates {
	resp, err := http.Get(googleApiUri + city)
	if err != nil {
		log.Fatal("Fetching google api uri data error: ", err)
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal("Reading google api data error: ", err)
	}

	var d googleApiResponse
	json.Unmarshal(bytes, &d)
	return Coordinates{Latitude: d.Results[0].Geometry.Location.Latitude, Longitude: d.Results[0].Geometry.Location.Longitude}
}
