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

func FromPBGender(v base.Gender) Gender {
	switch v {
	case base.Gender_Male:
		return Male
	case base.Gender_Female:
		return Female
	default:
		return 0
	}
}

func ToPBGender(v Gender) base.Gender {
	switch v {
	case Male:
		return base.Gender_Male
	case Female:
		return base.Gender_Female
	default:
		return base.Gender_Unknown
	}
}

func ToPBFormItem(v *FormItem) *base.FormItem {
	return &base.FormItem{
		Type:        v.Type,
		Name:        v.Name,
		Options:     v.Options,
		Values:      v.Values,
		Optional:    v.Optional,
		Description: v.Description,
	}
}

func FromPBFormItem(v *base.FormItem) *FormItem {
	return &FormItem{
		Type:        v.Type,
		Name:        v.Name,
		Options:     v.Options,
		Values:      v.Values,
		Optional:    v.Optional,
		Description: v.Description,
	}
}

func ToPBFormItemList(items []*FormItem) []*base.FormItem {
	l := make([]*base.FormItem, len(items))
	for i, v := range items {
		l[i] = ToPBFormItem(v)
	}
	return l
}

func FromPBFormItems(items []*base.FormItem) []*FormItem {
	l := make([]*FormItem, len(items))
	for i, v := range items {
		l[i] = FromPBFormItem(v)
	}
	return l
}
