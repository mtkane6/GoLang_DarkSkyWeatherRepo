package domain

// ErrorCheck panics if error is discovered
func ErrorCheck(e error) {
	if e != nil {
		panic(e)
	}
}
