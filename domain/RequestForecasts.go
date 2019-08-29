package domain

import (
	repo "URL/repository"
	"net/http"
	"text/template"
)

//HTMLRenderer builds the webpage
func HTMLRenderer(w http.ResponseWriter, r *http.Request) {
	// Resorts is the array of desired ski resorts
	var Resorts *[]repo.Resort
	Resorts = repo.BuildResortSlice()

	URLslice := repo.BuildURLslice(Resorts)            // returns slice of URLinstances{URL, resortName}
	responseSlice := repo.CallDarkSky(URLslice)        // returns slice of DarkSkyResponses
	repo.PopulateResortWeather(Resorts, responseSlice) // populates resort weather info for html

	tmpl, err := template.ParseFiles("./repository/WeatherDisplayTemplate.html")
	ErrorCheck(err)
	tmpl.Execute(w, &Resorts)
}
