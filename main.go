package main

import (
	api "./api"
)

func main() {
	api.BeginServer()

	// instantiate DarkSky request model
	// Build DarkSky URL
	// Call Darksky With URL, receive JSON.
	// Unmarshall JSON into DarkSky response model
	// Repackage that info into my display model

}
