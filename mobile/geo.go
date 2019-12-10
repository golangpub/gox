package mobile

import (
	"github.com/gopub/gox/geo"
)

type Point = geo.Point
type Place = geo.Place
type Country = geo.Country

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
	return &CountryList{List: geo.GetCountries()}
}

func GetCountryByCallingCode(code int) *Country {
	return geo.GetCountryByCallingCode(code)
}

func GetCallingCodeByRegion(regionCode string) int {
	return geo.GetCallingCodeByRegion(regionCode)
}
