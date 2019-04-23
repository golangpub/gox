package gox

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"math"
)

const (
	PI           = 3.14159265
	Earth_Radius = 6378.1 //km
	Earth_Circle = 2 * PI * Earth_Radius
	Degree       = Earth_Circle * 1000 / 360
)

type Coordinate struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lng"`
}

var _ driver.Valuer = (*Coordinate)(nil)
var _ sql.Scanner = (*Coordinate)(nil)

func (c *Coordinate) Scan(src interface{}) error {
	if src == nil {
		return nil
	}

	s, ok := src.(string)
	if !ok {
		var b []byte
		b, ok = src.([]byte)
		if ok {
			s = string(b)
		}
	}

	if !ok {
		return fmt.Errorf("failed to parse %v into geo.Coordinate", src)
	}

	k, err := fmt.Sscanf(s, "POINT(%f %f)", &c.Longitude, &c.Latitude)
	if k == 2 {
		return nil
	}

	return fmt.Errorf("failed to parse %v into geo.Coordinate: %v", s, err)
}

func (c *Coordinate) Value() (driver.Value, error) {
	if c == nil {
		return nil, nil
	}
	return fmt.Sprintf("POINT(%f %f)", c.Longitude, c.Latitude), nil
}

type Area struct {
	MinLat float64
	MaxLat float64
	MinLng float64
	MaxLng float64
}

func (a *Area) ContainsCoordinate(c *Coordinate) bool {
	if c.Latitude < a.MinLat || c.Latitude > a.MaxLat {
		return false
	}

	if a.MinLng > a.MaxLng {
		return c.Longitude >= a.MinLng || c.Longitude <= a.MaxLng
	}

	return c.Longitude >= a.MinLng && c.Longitude <= a.MaxLng
}

// @param radius is in km
func (c *Coordinate) GetArea(radius float64) Area {
	if radius > Earth_Circle {
		panic("radius is too large")
	}

	var a Area
	lat := c.Latitude * PI / 180
	lng := c.Longitude * PI / 180

	//弧面距离=半径*夹角，故，夹角=弧面距离/半径, 夹角永远为正
	angle := radius / Earth_Radius

	//纬度对应的半径均是地球的半径，故纬度的范围可以估算出来，但是不同经度对应的圆半径不一样，随着纬度的绝对值越大，经度半径越小
	a.MinLat = lat - angle
	a.MaxLat = lat + angle

	deltaLng := angle / math.Cos(lat)
	a.MinLng = lng - deltaLng
	a.MaxLng = lng + deltaLng

	//纬度的最大值和最小值是分开的，最大值大于90, 最小值小于-90，在计算范围的时候没有关系
	//经度的180和-180是临界点，因此最大值可能大于180，最小值可能小于-180，计算范围的时候需要进行转换

	a.MaxLat *= 180 / PI
	a.MinLat *= 180 / PI
	a.MaxLng *= 180 / PI
	a.MinLng *= 180 / PI
	if a.MaxLng > 180 {
		a.MaxLng -= 360
	}

	if a.MinLng < -180 {
		a.MinLng += 360
	}

	return a
}

func (c *Coordinate) DistanceTo(d Coordinate) float64 {
	lat1 := c.Latitude * PI / 180.0
	lat2 := d.Latitude * PI / 180
	deltaLat := lat2 - lat1
	deltaLng := (d.Longitude - c.Longitude) * PI / 180

	a := math.Sin(deltaLat/2)*math.Sin(deltaLat/2) + math.Cos(lat1)*math.Cos(lat2)*math.Sin(deltaLng/2)*math.Sin(deltaLng/2)
	k := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return k * Earth_Radius
}

type Location struct {
	Address
	Coordinate
}
