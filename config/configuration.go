package config

// GetBaseURL returns the beginning of the API call to DarkSky
func GetBaseURL() string {
	BaseURL := "https://api.darksky.net/forecast/"
	return BaseURL
}

// GetAPIkey returns the unique API key to authenticate calls to DarkSky
func GetAPIkey() string {
	APIKey := "secret :)"
	return APIKey
}
