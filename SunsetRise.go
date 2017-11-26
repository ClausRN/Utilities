package Utilities

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// SunMoonInfo contains information about moon/sun rise/set and misc other info
type SunMoonInfo struct {
	Status              string  `json:"status"`
	Timezone            string  `json:"timezone"`
	DSTOffset           float32 `json:"dstOffset"`
	Date                string  `json:"date"`
	Time                string  `json:"currentTime"`
	MinutesSinceSunRise int16   `json:"MinutesSinceSunRise"`
	MoonRise            string  `json:"moonRise"`
	MoonSet             string  `json:"moonSet"`
	SunRise             string  `json:"sunRise"`
	SunSet              string  `json:"sunSet"`
	IsSunUp             string  `json:"isSunUp"`
	MoonPhase           string  `json:"moonPhase"`
	MoonIlluminated     float32 `json:"moonIlluminated"`
	MoonAge             string  `json:"moonAge"`
	IsWaxing            bool    `json:"isWaxing"`
	MextFullMoon        string  `json:"nextFullMoon"`
}

const (
	fullsuninfoProvider = "https://soltider.dk/api?lat=55.663426&lng=12.542953"
	suninfoProvider     = "https://soltider.dk/api?lat=%f&lng=%f"
)

// GetSunMoonInfo retrieves info about sun and moon rise/set
func GetSunMoonInfo(Latitude, Longitude float32) *SunMoonInfo {
	var data []*SunMoonInfo
	myURL := fmt.Sprintf(suninfoProvider, Latitude, Longitude)
	res, err := http.Get(myURL)

	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Printf("Couldn't read body: %s", err)
		panic(err.Error())
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Printf("Couldn't decode json: %s", err)
	}
	return data[0]
}
