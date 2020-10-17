package utils

import "log"

// CheckError checks for errors
func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
