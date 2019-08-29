package repository

import (
	"Url/config"
	"fmt"
)

// URLinstance is a URL request object
type URLinstance struct {
	URL        string
	ResortName string
}

// BuildURLslice returns slice of URL strings to make GET requests
func BuildURLslice(resorts *[]Resort) []URLinstance {
	var URLslice []URLinstance

	for _, resort := range *resorts {
		var currentResortURL URLinstance

		currentResortURL.URL = BuildForecastRequestURL(resort)
		currentResortURL.ResortName = resort.Name

		URLslice = append(URLslice, currentResortURL)
	}

	return URLslice
}

// BuildForecastRequestURL contructs the url for the api requests
func BuildForecastRequestURL(resort Resort) string {
	DarkSkyURL := fmt.Sprintf("%s%s%f,%f", config.GetBaseURL(), config.GetAPIkey(), resort.Latitude, resort.Longitude)
	return DarkSkyURL
}
