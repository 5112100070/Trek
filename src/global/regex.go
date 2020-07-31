package global

import (
	"log"
	"regexp"
)

// restriction character on image name
var regexResImageNameChar *regexp.Regexp

// init regex compile. this is make we only once compiling the regex
func InitRegexCompileImageName() error {
	var err error
	regexResImageNameChar, err = regexp.Compile(`[^a-zA-Z\d\s:]`)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func IsValidEmail(email string) bool {
	pattern := `/^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/`

	reg, err := regexp.Compile(pattern)
	if err != nil {
		log.Println(err)
		return false
	}

	return !reg.MatchString(email)
}
