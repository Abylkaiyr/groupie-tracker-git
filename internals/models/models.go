package models

const (
	Url = "https://groupietrackers.herokuapp.com/api"
)

type TemplateData struct {
	Artists         []Artist
	UniqueLocations []string
	SearchArtists   []Artist
}

type LocationInfo struct {
	LocationInfo []LocationData `json:"index"`
}

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`

	DatesLocations Relations

	LongLat []LocCoordinates
}

type LocationData struct {
	ID        int `json:"id"`
	Locations []string
	Dates     string
}

type LocCoordinates struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lng"`
}

type Relations struct {
	DatesLocations map[string][]string `json:"datesLocations"`
}
