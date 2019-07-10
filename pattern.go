package gox

import (
	"github.com/gopub/log"
	"regexp"
	"time"

	"github.com/nyaruka/phonenumbers"
)

var nameRegexp = regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9\-._]*$`)
var emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func IsEmail(email string) bool {
	return emailRegexp.MatchString(email)
}

func IsName(name string) bool {
	return nameRegexp.MatchString(name)
}

func IsPhoneNumber(phoneNumber string) bool {
	parsedNumber, err := phonenumbers.Parse(phoneNumber, "")
	if err != nil {
		log.Error(err)
		return false
	}

	return phonenumbers.IsValidNumber(parsedNumber)
}

func IsBirthDate(s string) bool {
	t, err := time.Parse("2006-1-2", s)
	if err != nil {
		log.Error(err)
		return false
	}

	if t.After(time.Now()) {
		return false
	}

	return true
}
