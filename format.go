package gox

import (
	"errors"
	"github.com/nyaruka/phonenumbers"
	"regexp"
)

var emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func IsEmail(email string) bool {
	return emailRegexp.MatchString(email)
}

func IsPhoneNumber(phoneNumber string) bool {
	parsedNumber, err := phonenumbers.Parse(phoneNumber, "")
	if err != nil {
		return false
	}

	return phonenumbers.IsValidNumber(parsedNumber)
}

func ParsePhoneNumber(s string) (*PhoneNumber, error) {
	parsedNumber, err := phonenumbers.Parse(s, "")
	if err != nil {
		return nil, err
	}

	if phonenumbers.IsValidNumber(parsedNumber) {
		return &PhoneNumber{
			CountryCode:    int(parsedNumber.GetCountryCode()),
			NationalNumber: int64(parsedNumber.GetNationalNumber()),
			Extension:      parsedNumber.GetExtension(),
		}, nil
	}

	return nil, errors.New("invalid phone number")
}
