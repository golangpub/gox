package gox

import (
	"errors"

	"github.com/gopub/gox/protobuf/base"
	"github.com/nyaruka/phonenumbers"
)

func NewPhoneNumber(callingCode int, number int64) *PhoneNumber {
	pn := new(PhoneNumber)
	pn.CountryCode = callingCode
	pn.NationalNumber = number
	return pn
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

func FromPBPhoneNumber(pn *base.PhoneNumber) *PhoneNumber {
	return &PhoneNumber{
		CountryCode:    int(pn.CountryCode),
		NationalNumber: pn.NationalNumber,
		Extension:      pn.Extension,
	}
}

func FromPBPhoneNumbers(pns []*base.PhoneNumber) []*PhoneNumber {
	l := make([]*PhoneNumber, len(pns))
	for i, pn := range pns {
		l[i] = FromPBPhoneNumber(pn)
	}
	return l
}

func ToPBPhoneNumber(pn *PhoneNumber) *base.PhoneNumber {
	return &base.PhoneNumber{
		CountryCode:    int32(pn.CountryCode),
		NationalNumber: pn.NationalNumber,
		Extension:      pn.Extension,
	}
}

func ToPBPhoneNumbers(pns []*PhoneNumber) []*base.PhoneNumber {
	l := make([]*base.PhoneNumber, len(pns))
	for i, pn := range pns {
		l[i] = ToPBPhoneNumber(pn)
	}
	return l
}

func FromPBFullName(n *base.FullName) *FullName {
	return &FullName{
		FirstName:  n.FirstName,
		MiddleName: n.MiddleName,
		LastName:   n.LastName,
	}
}

func ToPBFullName(n *FullName) *base.FullName {
	return &base.FullName{
		FirstName:  n.FirstName,
		MiddleName: n.MiddleName,
		LastName:   n.LastName,
	}
}

func FromPBLocation(l *base.Location) *Location {
	return &Location{
		Name:      l.Name,
		FullName:  l.FullName,
		Latitude:  l.Latitude,
		Longitude: l.Longitude,
	}
}

func ToPBLocation(l *Location) *base.Location {
	return &base.Location{
		Name:      l.Name,
		FullName:  l.FullName,
		Latitude:  l.Latitude,
		Longitude: l.Longitude,
	}
}

func FromPBCoordinate(l *base.Coordinate) *Coordinate {
	return &Coordinate{
		Latitude:  l.Latitude,
		Longitude: l.Longitude,
	}
}

func ToPBCoordinate(l *Coordinate) *base.Coordinate {
	return &base.Coordinate{
		Latitude:  l.Latitude,
		Longitude: l.Longitude,
	}
}
