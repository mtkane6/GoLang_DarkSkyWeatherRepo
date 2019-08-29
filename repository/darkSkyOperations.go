package repository

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// DarkSky holds the api key for making api calls
type DarkSky struct {
	APIKey string
}

// DarkSkyResponse holds weather response information from DarkSky api
type DarkSkyResponse struct {
	ResortName string
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
	Timezone   string  `json:"timezone"`
	Currently  struct {
		Time                 int     `json:"time"` //  // t := time.Unix(1494505756, 0)
		Summary              string  `json:"summary"`
		Icon                 string  `json:"icon"`
		NearestStormDistance int     `json:"nearestStormDistance"`
		NearestStormBearing  int     `json:"nearestStormBearing"`
		PrecipIntensity      int     `json:"precipIntensity"` // inches of liquid per hour
		PrecipProbability    int     `json:"precipProbability"`
		Temperature          float64 `json:"temperature"`
		ApparentTemperature  float64 `json:"apparentTemperature"`
		DewPoint             float64 `json:"dewPoint"`
		Humidity             float64 `json:"humidity"`
		Pressure             float64 `json:"pressure"`
		WindSpeed            float64 `json:"windSpeed"`
		WindGust             float64 `json:"windGust"`
		WindBearing          int     `json:"windBearing"`
		CloudCover           float64 `json:"cloudCover"`
		UvIndex              int     `json:"uvIndex"`
		Visibility           int     `json:"visibility"`
		Ozone                int     `json:"ozone"`
	} `json:"currently"`
	Minutely struct {
		Summary string `json:"summary"`
		Icon    string `json:"icon"`
		Data    []struct {
			Time              int `json:"time"`
			PrecipIntensity   int `json:"precipIntensity"`
			PrecipProbability int `json:"precipProbability"`
		} `json:"data"`
	} `json:"minutely"`
	Hourly struct {
		Summary string `json:"summary"`
		Icon    string `json:"icon"`
		Data    []struct {
			Time                int     `json:"time"`
			Summary             string  `json:"summary"`
			Icon                string  `json:"icon"`
			PrecipIntensity     int     `json:"precipIntensity"`
			PrecipProbability   int     `json:"precipProbability"`
			Temperature         float64 `json:"temperature"`
			ApparentTemperature float64 `json:"apparentTemperature"`
			DewPoint            float64 `json:"dewPoint"`
			Humidity            float64 `json:"humidity"`
			Pressure            float64 `json:"pressure"`
			WindSpeed           float64 `json:"windSpeed"`
			WindGust            float64 `json:"windGust"`
			WindBearing         int     `json:"windBearing"`
			CloudCover          float64 `json:"cloudCover"`
			UvIndex             int     `json:"uvIndex"`
			Visibility          int     `json:"visibility"`
			Ozone               float64 `json:"ozone"`
			PrecipType          string  `json:"precipType,omitempty"`
		} `json:"data"`
	} `json:"hourly"`
	Daily struct {
		Summary string `json:"summary"`
		Icon    string `json:"icon"`
		Data    []struct {
			Time                        int     `json:"time"`
			Summary                     string  `json:"summary"`
			Icon                        string  `json:"icon"`
			SunriseTime                 int     `json:"sunriseTime"`
			SunsetTime                  int     `json:"sunsetTime"`
			MoonPhase                   float64 `json:"moonPhase"`
			PrecipIntensity             float64 `json:"precipIntensity"`
			PrecipIntensityMax          float64 `json:"precipIntensityMax"`
			PrecipIntensityMaxTime      int     `json:"precipIntensityMaxTime"`
			PrecipProbability           float64 `json:"precipProbability"`
			PrecipType                  string  `json:"precipType,omitempty"`
			TemperatureHigh             float64 `json:"temperatureHigh"`
			TemperatureHighTime         int     `json:"temperatureHighTime"`
			TemperatureLow              float64 `json:"temperatureLow"`
			TemperatureLowTime          int     `json:"temperatureLowTime"`
			ApparentTemperatureHigh     float64 `json:"apparentTemperatureHigh"`
			ApparentTemperatureHighTime int     `json:"apparentTemperatureHighTime"`
			ApparentTemperatureLow      float64 `json:"apparentTemperatureLow"`
			ApparentTemperatureLowTime  int     `json:"apparentTemperatureLowTime"`
			DewPoint                    float64 `json:"dewPoint"`
			Humidity                    float64 `json:"humidity"`
			Pressure                    float64 `json:"pressure"`
			WindSpeed                   float64 `json:"windSpeed"`
			WindGust                    float64 `json:"windGust"`
			WindGustTime                int     `json:"windGustTime"`
			WindBearing                 int     `json:"windBearing"`
			CloudCover                  float64 `json:"cloudCover"`
			UvIndex                     int     `json:"uvIndex"`
			UvIndexTime                 int     `json:"uvIndexTime"`
			Visibility                  float64 `json:"visibility"`
			Ozone                       float64 `json:"ozone"`
			TemperatureMin              float64 `json:"temperatureMin"`
			TemperatureMinTime          int     `json:"temperatureMinTime"`
			TemperatureMax              float64 `json:"temperatureMax"`
			TemperatureMaxTime          int     `json:"temperatureMaxTime"`
			ApparentTemperatureMin      float64 `json:"apparentTemperatureMin"`
			ApparentTemperatureMinTime  int     `json:"apparentTemperatureMinTime"`
			ApparentTemperatureMax      float64 `json:"apparentTemperatureMax"`
			ApparentTemperatureMaxTime  int     `json:"apparentTemperatureMaxTime"`
		} `json:"data"`
	} `json:"daily"`
	Flags struct {
		Sources        []string `json:"sources"`
		NearestStation float64  `json:"nearest-station"`
		Units          string   `json:"units"`
	} `json:"flags"`
	Offset int `json:"offset"`
}

// CallDarkSky uses the resort data to call DarkSky for forecast
func CallDarkSky(URLslice []URLinstance) []DarkSkyResponse {
	var DarkSkyResponseSlice []DarkSkyResponse

	for _, location := range URLslice {
		resp, err := GetForecastResponse(location.URL)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		var darkSkyResponse DarkSkyResponse
		b, err := ioutil.ReadAll(resp.Body)
		json.Unmarshal([]byte(b), &darkSkyResponse)
		darkSkyResponse.ResortName = location.ResortName
		DarkSkyResponseSlice = append(DarkSkyResponseSlice, darkSkyResponse)
	}
	return DarkSkyResponseSlice
}

// GetForecastResponse call DarkSky to retrieve resort forecast
func GetForecastResponse(URL string) (*http.Response, error) {
	resp, err := http.Get(URL)
	if err != nil {
		return resp, err
	}
	return resp, err
}
