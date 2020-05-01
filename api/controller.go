package api

import (
	"log"
	"net/http"

	domain "Url/domain"
	// "log"
	// "net/http"
)

// BeginServer route handles
func BeginServer() {
	http.HandleFunc("/", domain.HTMLRenderer)
	http.Handle("/repository/", http.StripPrefix("/repository/", http.FileServer(http.Dir("repository"))))
	log.Fatal(http.ListenAndServe(":8080", nil))
	// domain.HTMLRenderer()
}
