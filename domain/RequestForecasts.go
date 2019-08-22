package domain

import (
	repo "URL/repository"
	"Url/config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"
)

//HTMLRenderer builds the webpage
func HTMLRenderer(w http.ResponseWriter, r *http.Request) {
	// Resorts is the array of desired ski resorts
	var Resorts *[]repo.Resort
	Resorts = BuildResortSlice()

	CallDarkSky(Resorts)

	tmpl, err := template.ParseFiles("./repository/WeatherDisplayTemplate.html")
	ErrorCheck(err)
	tmpl.Execute(w, &Resorts)
}

// CallDarkSky uses the resort data to call DarkSky for forecast
func CallDarkSky(resorts *[]repo.Resort) {

	for _, resort := range *resorts {
		URL := BuildForecastRequestURL(resort)
		resp, err := GetForecastResponse(URL)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		var WeatherResponse repo.DarkSkyResponse
		b, err := ioutil.ReadAll(resp.Body)
		json.Unmarshal([]byte(b), &WeatherResponse)

		PopulateResortWeather(&resort, &WeatherResponse)
	}
}

// PopulateResortWeather uses DarkSkyResponse data to populate desired resort data to display
func PopulateResortWeather(resort *repo.Resort, weatherResponse *repo.DarkSkyResponse) {
	resort.Today.RainInchesPerHour = weatherResponse.Daily.Data[0].PrecipIntensity
	resort.Today.ChanceOfPrecipitation = weatherResponse.Daily.Data[0].PrecipProbability * 100
	resort.Today.TemperatureHigh = weatherResponse.Daily.Data[0].TemperatureHigh
	resort.Today.TemperatureHigh = weatherResponse.Daily.Data[0].TemperatureLow

	resort.Tomorrow.ChanceOfPrecipitation = weatherResponse.Daily.Data[1].PrecipProbability * 100
	resort.Tomorrow.RainInchesPerHour = weatherResponse.Daily.Data[1].PrecipIntensity
	resort.Tomorrow.TemperatureHigh = weatherResponse.Daily.Data[1].TemperatureHigh
	resort.Tomorrow.TemperatureHigh = weatherResponse.Daily.Data[1].TemperatureLow
}

// BuildResortSlice builds up the slice of Resort structs
func BuildResortSlice() *[]repo.Resort {
	var ResortSlice []repo.Resort
	AddResort(&ResortSlice, "Stevens Pass", 47.7448, 121.0890)
	AddResort(&ResortSlice, "Crystal Mountin", 46.9282, 121.5045)
	AddResort(&ResortSlice, "Mt. Baker", 48.7767, 121.8144)
	AddResort(&ResortSlice, "Tahoe Heavenly", 38.9611, 119.8856)
	AddResort(&ResortSlice, "Jackson Hole", 43.5875, 110.8279)
	AddResort(&ResortSlice, "Alta/Snowbird", 40.5883, 111.6358)
	return &ResortSlice
}

// AddResort adds a resort struct to the slice of resort structs
func AddResort(resortSlice *[]repo.Resort, name string, lat, long float32) {
	err := ValidateResort(name, lat, long)
	if err != nil {
		log.Fatal(err)
	}
	CurrentResort := repo.Resort{
		Name:      name,
		Latitude:  lat,
		Longitude: long,
	}
	*resortSlice = append(*resortSlice, CurrentResort)
}

// BuildForecastRequestURL contructs the url for the api requests
func BuildForecastRequestURL(resort repo.Resort) string {
	DarkSkyURL := fmt.Sprintf("%s%s%f,%f", config.GetBaseURL(), config.GetAPIkey(), resort.Latitude, resort.Longitude)
	return DarkSkyURL
}

// GetForecastResponse call DarkSky to retrieve resort forecast
func GetForecastResponse(URL string) (*http.Response, error) {
	resp, err := http.Get(URL)
	if err != nil {
		return resp, err
	}
	return resp, err
}
