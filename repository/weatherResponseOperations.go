package repository

// AllWeatherResponse stores today and tomorrow forecast for html template display
type AllWeatherResponse struct {
	Today    []TodayWeatherResponse
	Tomorrow []TomorrowWeatherResponse
}

// TodayWeatherResponse aims to holds current weather conditions
type TodayWeatherResponse struct {
	RainInchesPerHour     float64
	ChanceOfPrecipitation float64
	TemperatureHigh       float64
	TemperatureLow        float64
}

// TomorrowWeatherResponse aims to holds current weather conditions
type TomorrowWeatherResponse struct {
	RainInchesPerHour     float64
	ChanceOfPrecipitation float64
	TemperatureHigh       float64
	TemperatureLow        float64
}

// PopulateResortWeather uses DarkSkyResponse data to populate desired resort data to display
func PopulateResortWeather(Resorts *[]Resort, weatherResponseSlice []DarkSkyResponse) {

	for i, resort := range *Resorts {
		if resort.Name == weatherResponseSlice[i].ResortName {
			resort.Today.RainInchesPerHour = weatherResponseSlice[i].Daily.Data[0].PrecipIntensity
			resort.Today.ChanceOfPrecipitation = weatherResponseSlice[i].Daily.Data[0].PrecipProbability * 100
			resort.Today.TemperatureHigh = weatherResponseSlice[i].Daily.Data[0].TemperatureHigh
			resort.Today.TemperatureHigh = weatherResponseSlice[i].Daily.Data[0].TemperatureLow

			resort.Tomorrow.ChanceOfPrecipitation = weatherResponseSlice[i].Daily.Data[1].PrecipProbability * 100
			resort.Tomorrow.RainInchesPerHour = weatherResponseSlice[i].Daily.Data[1].PrecipIntensity
			resort.Tomorrow.TemperatureHigh = weatherResponseSlice[i].Daily.Data[1].TemperatureHigh
			resort.Tomorrow.TemperatureHigh = weatherResponseSlice[i].Daily.Data[1].TemperatureLow
		}
	}

}
