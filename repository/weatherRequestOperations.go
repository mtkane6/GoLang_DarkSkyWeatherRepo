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
func BuildURLslice(resorts []Resort) []URLinstance {
	var URLslice []URLinstance

	for _, resort := range resorts {
		URLslice = append(URLslice, URLinstance{
			URL:        BuildForecastRequestURL(resort),
			ResortName: resort.Name,
		})
	}

	return URLslice
}

// BuildForecastRequestURL contructs the url for the api requests
func BuildForecastRequestURL(resort Resort) string {
	return fmt.Sprintf("%s%s%f,%f", config.GetBaseURL(), config.GetAPIkey(), resort.Latitude, resort.Longitude)
}
