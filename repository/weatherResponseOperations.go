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
func PopulateResortWeather(Resorts []Resort, weatherResponseSlice []DarkSkyResponse) []Resort {

	for i := range Resorts {
		if Resorts[i].Name == weatherResponseSlice[i].ResortName {
			Resorts[i].Today.RainInchesPerHour = weatherResponseSlice[i].Daily.Data[0].PrecipIntensity
			Resorts[i].Today.ChanceOfPrecipitation = weatherResponseSlice[i].Daily.Data[0].PrecipProbability * 100
			Resorts[i].Today.TemperatureHigh = weatherResponseSlice[i].Daily.Data[0].TemperatureHigh
			Resorts[i].Today.TemperatureLow = weatherResponseSlice[i].Daily.Data[0].TemperatureLow

			Resorts[i].Tomorrow.ChanceOfPrecipitation = weatherResponseSlice[i].Daily.Data[1].PrecipProbability * 100
			Resorts[i].Tomorrow.RainInchesPerHour = weatherResponseSlice[i].Daily.Data[1].PrecipIntensity
			Resorts[i].Tomorrow.TemperatureHigh = weatherResponseSlice[i].Daily.Data[1].TemperatureHigh
			Resorts[i].Tomorrow.TemperatureLow = weatherResponseSlice[i].Daily.Data[1].TemperatureLow
		}
	}
	return Resorts
}
