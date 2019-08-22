package api

import (
	domain "Url/domain"
	"log"
	"net/http"
)

// BeginServer route handles
func BeginServer() {
	http.HandleFunc("/", domain.HTMLRenderer)
	log.Fatal(http.ListenAndServe(":8080", nil))
	// domain.TestWeather()
}
