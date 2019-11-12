package gox

import (
	"errors"
	"fmt"

	"github.com/gopub/gox/protobuf/base"
	"github.com/nyaruka/phonenumbers"
)

func NewPhoneNumber(callingCode int, number int64) *PhoneNumber {
	pn := new(PhoneNumber)
	pn.Code = callingCode
	pn.Number = number
	return pn
}

func ParsePhoneNumber(s string) (*PhoneNumber, error) {
	parsedNumber, err := phonenumbers.Parse(s, "")
	if err != nil {
		return nil, err
	}

	if phonenumbers.IsValidNumber(parsedNumber) {
		return &PhoneNumber{
			Code:      int(parsedNumber.GetCountryCode()),
			Number:    int64(parsedNumber.GetNationalNumber()),
			Extension: parsedNumber.GetExtension(),
		}, nil
	}

	return nil, errors.New("invalid phone number")
}

func FromPBPhoneNumber(pn *base.PhoneNumber) *PhoneNumber {
	if pn == nil {
		return nil
	}
	return &PhoneNumber{
		Code:      int(pn.CountryCode),
		Number:    pn.NationalNumber,
		Extension: pn.Extension,
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
	if pn == nil {
		return nil
	}
	return &base.PhoneNumber{
		CountryCode:    int32(pn.Code),
		NationalNumber: pn.Number,
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
	if n == nil {
		return nil
	}
	return &FullName{
		FirstName:  n.FirstName,
		MiddleName: n.MiddleName,
		LastName:   n.LastName,
	}
}

func ToPBFullName(n *FullName) *base.FullName {
	if n == nil {
		return nil
	}
	return &base.FullName{
		FirstName:  n.FirstName,
		MiddleName: n.MiddleName,
		LastName:   n.LastName,
	}
}

func FromPBPoint(l *base.Point) *Point {
	if l == nil {
		return nil
	}
	return &Point{
		Y: l.Y,
		X: l.X,
	}
}

func ToPBPoint(l *Point) *base.Point {
	if l == nil {
		return nil
	}
	return &base.Point{
		Y: l.Y,
		X: l.X,
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

type ByteUnit int64

const (
	_           = iota
	KB ByteUnit = 1 << (10 * iota) // 1 << (10*1)
	MB                             // 1 << (10*2)
	GB                             // 1 << (10*3)
	TB                             // 1 << (10*4)
	PB                             // 1 << (10*5)
)

func (b ByteUnit) HumanReadable() string {
	return ByteCountToDisplaySize(int64(b))
}

func ByteCountToDisplaySize(count int64) string {
	n := ByteUnit(count)
	if n < KB {
		return fmt.Sprintf("%d B", count)
	} else if n < MB {
		return fmt.Sprintf("%.2f KB", float64(count)/float64(KB))
	} else if n < GB {
		return fmt.Sprintf("%.2f MB", float64(count)/float64(MB))
	} else if n < TB {
		return fmt.Sprintf("%.2f GB", float64(count)/float64(GB))
	} else if n < PB {
		return fmt.Sprintf("%.2f TB", float64(count)/float64(TB))
	} else {
		return fmt.Sprintf("%.2f PB", float64(count)/float64(PB))
	}
}
