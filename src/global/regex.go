package global

import (
	"log"
	"regexp"
)

func IsValidEmail(email string) bool {
	pattern := `/^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/`

	reg, err := regexp.Compile(pattern)
	if err != nil {
		log.Println(err)
		return false
	}

	return !reg.MatchString(email)
}
