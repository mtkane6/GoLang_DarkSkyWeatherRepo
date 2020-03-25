package repository

import "math"

// AllWeatherResponse stores today and tomorrow forecast for html template display
type AllWeatherResponse struct {
	Today    []TodayWeatherResponse
	Tomorrow []TomorrowWeatherResponse
}

// TodayWeatherResponse aims to holds current weather conditions
type TodayWeatherResponse struct {
	SnowAccumulation      float64
	ChanceOfPrecipitation float64
	TemperatureHigh       float64
	TemperatureLow        float64
}

// TomorrowWeatherResponse aims to holds current weather conditions
type TomorrowWeatherResponse struct {
	SnowAccumulation      float64
	ChanceOfPrecipitation float64
	TemperatureHigh       float64
	TemperatureLow        float64
}

// PopulateResortWeather uses DarkSkyResponse data to populate desired resort data to display
func PopulateResortWeather(Resorts []Resort, weatherResponseSlice []DarkSkyResponse) []Resort {
	// fmt.Printf("%v+", weatherResponseSlice[0])
	for i := range Resorts {
		if Resorts[i].Name == weatherResponseSlice[i].ResortName {
			Resorts[i].Today.SnowAccumulation = weatherResponseSlice[i].Daily.Data[0].PrecipAccumulation
			Resorts[i].Today.ChanceOfPrecipitation = math.Round(weatherResponseSlice[i].Daily.Data[0].PrecipProbability * 100)
			Resorts[i].Today.TemperatureHigh = weatherResponseSlice[i].Daily.Data[0].TemperatureHigh
			Resorts[i].Today.TemperatureLow = weatherResponseSlice[i].Daily.Data[0].TemperatureLow

			Resorts[i].Tomorrow.SnowAccumulation = weatherResponseSlice[i].Daily.Data[1].PrecipAccumulation
			Resorts[i].Tomorrow.ChanceOfPrecipitation = math.Round(weatherResponseSlice[i].Daily.Data[1].PrecipProbability * 100)
			Resorts[i].Tomorrow.TemperatureHigh = weatherResponseSlice[i].Daily.Data[1].TemperatureHigh
			Resorts[i].Tomorrow.TemperatureLow = weatherResponseSlice[i].Daily.Data[1].TemperatureLow
		}
	}
	return Resorts
}
