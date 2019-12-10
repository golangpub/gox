package mobile

import (
	"github.com/gopub/gox/geo"
)

type Point = geo.Point
type Place = geo.Place

type Country geo.Country

type CountryList struct {
	List []*Country
}

func NewCountryList() *CountryList {
	return &CountryList{}
}

func (l *CountryList) Size() int {
	return len(l.List)
}

func (l *CountryList) Get(i int) *Country {
	return l.List[i]
}

func GetCountryList() *CountryList {
	l := geo.GetCountries()
	ll := NewCountryList()
	ll.List = make([]*Country, len(l))
	for i, c := range l {
		ll.List[i] = (*Country)(c)
	}
	return ll
}

func GetCountryByCallingCode(code int) *Country {
	return (*Country)(geo.GetCountryByCallingCode(code))
}

func GetCallingCodeByRegion(regionCode string) int {
	return geo.GetCallingCodeByRegion(regionCode)
}
