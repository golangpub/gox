package gox

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"strings"
)

const (
	PI          = 3.14159265
	EarthRadius = 6378.1 //km
	EarthCircle = 2 * PI * EarthRadius
	Degree      = EarthCircle * 1000 / 360
)

type Point struct {
	X float64 `json:"x"` // X
	Y float64 `json:"y"` // Y
}

func NewPoint() *Point {
	return &Point{}
}

var _ driver.Valuer = (*Point)(nil)
var _ sql.Scanner = (*Point)(nil)

func (c *Point) Scan(src interface{}) error {
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

	if !ok || len(s) < 2 {
		return fmt.Errorf("failed to parse %v into geo.Point", src)
	}

	s = s[1 : len(s)-1]
	s = strings.Replace(s, ",", " ", -1)
	k, err := fmt.Sscanf(s, "%f %f", &c.X, &c.Y)
	if k == 2 {
		return nil
	}

	return fmt.Errorf("parse %v into geo.Location: %w", s, err)
}

func (c *Point) Value() (driver.Value, error) {
	if c == nil {
		return nil, nil
	}
	return fmt.Sprintf("(%f,%f)", c.X, c.Y), nil
}

// Place
type Place struct {
	Code     string `json:"code,omitempty"`
	Name     string `json:"name,omitempty"`
	Location *Point `json:"point,omitempty"`
}

func NewPlace() *Place {
	return &Place{}
}

var _ driver.Valuer = (*Place)(nil)
var _ sql.Scanner = (*Place)(nil)

func (p *Place) Scan(src interface{}) error {
	if src == nil {
		return nil
	}

	b, ok := src.([]byte)
	if !ok || len(b) < 2 || b[0] != '(' || b[len(b)-1] != ')' {
		return fmt.Errorf("failed to parse %v into gox.Place", src)
	}
	b = b[1 : len(b)-1]
	strs := strings.Split(string(b), ",")
	if len(strs) != 4 {
		return fmt.Errorf("failed to parse %v into gox.Place", src)
	}
	p.Code = strs[0]
	p.Name = strs[1]
	p.Location = new(Point)
	p.Location.Scan(strs[3])
	if len(strs[2]) > 0 || len(strs[3]) > 0 {
		p.Location = &Point{}
		var err error
		p.Location.X, err = ParseFloat(strs[2])
		if err != nil {
			return fmt.Errorf("parse point.x: %w", err)
		}
		p.Location.Y, err = ParseFloat(strs[3])
		if err != nil {
			return fmt.Errorf("parse point.y: %w", err)
		}
	}
	return nil
}

func (p *Place) Value() (driver.Value, error) {
	if p == nil {
		return nil, nil
	}
	name := strings.Replace(p.Name, ",", "\\,", -1)
	x, y := "", ""
	if p.Location != nil {
		x, y = fmt.Sprint(p.Location.X), fmt.Sprint(p.Location.Y)
	}
	s := fmt.Sprintf("(%s,%s,'(%s,%s)')", p.Code, name, x, y)
	return s, nil
}
