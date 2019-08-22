package domain

import (
	"fmt"
)

// ValidateResort checks input data for resort
func ValidateResort(name string, lat, long float32) error {
	if name == "" {
		return fmt.Errorf("Empty resort name")
	}
	if lat > 360.0 || lat < -360.0 || long > 360.0 || long < -360.0 {
		return fmt.Errorf("Invalid coordinates: %f, %f", lat, long)
	}
	return nil
}

// ErrorCheck panics if error is discovered
func ErrorCheck(e error) {
	if e != nil {
		panic(e)
	}
}
