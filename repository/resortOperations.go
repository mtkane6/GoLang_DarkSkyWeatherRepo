package repository

// Resort is the struct that keeps basic information of the resorts
type Resort struct {
	Name      string
	Latitude  float32
	Longitude float32
	Today     TodayWeatherResponse
	Tomorrow  TomorrowWeatherResponse
}

// BuildResortSlice builds up the slice of Resort structs
func BuildResortSlice() *[]Resort {
	var ResortSlice []Resort
	AddResort(&ResortSlice, "Stevens Pass", 47.7448, 121.0890)
	AddResort(&ResortSlice, "Crystal Mountin", 46.9282, 121.5045)
	AddResort(&ResortSlice, "Mt. Baker", 48.7767, 121.8144)
	AddResort(&ResortSlice, "Tahoe Heavenly", 38.9611, 119.8856)
	AddResort(&ResortSlice, "Jackson Hole", 43.5875, 110.8279)
	AddResort(&ResortSlice, "Alta/Snowbird", 40.5883, 111.6358)
	return &ResortSlice
}

// AddResort adds a resort struct to the slice of resort structs
func AddResort(resortSlice *[]Resort, name string, lat, long float32) {
	CurrentResort := Resort{
		Name:      name,
		Latitude:  lat,
		Longitude: long,
	}
	*resortSlice = append(*resortSlice, CurrentResort)
}
