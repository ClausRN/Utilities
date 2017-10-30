package Utilities

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// LocationInfo contains different kinds of infirmation about current location
type LocationInfo struct {
	IP            string  `json:"ip"`
	City          string  `json:"city"`
	Region        string  `json:"region"`
	Country       string  `json:"country_name"`
	CountryCode   string  `json:"country_code"`
	Continent     string  `json:"continent_name"`
	ContinentCode string  `json:"continent_code"`
	Latitude      float32 `json:"latitude"`
	Longitude     float32 `json:"longitude"`
	ASN           string  `json:"asn"`
	Organisation  string  `json:"organisation"`
	Postal        string  `json:"postal"`
	Currency      string  `json:"currency"`
	CallingCode   string  `json:"calling_code"`
	FlagURL       string  `json:"flag"`
	Timezone      string  `json:"time_zone"`
}

const (
	locationProvider = "https://api.ipdata.co"
)

// GetLocation returns the location data for IP address
func GetLocation() (data *LocationInfo) {
	res, err := http.Get(locationProvider)

	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err.Error())
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Printf("Couldn't decode json: %s", err)
	}
	return
}
