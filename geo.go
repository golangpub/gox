package gox

import (
	"database/sql/driver"
	"fmt"

	"github.com/gopub/gox/sql"
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
	s, err := ParseString(src)
	if err != nil {
		return fmt.Errorf("parse string: %w", err)
	}
	if s == "" {
		return nil
	}
	fields, err := sql.ParseCompositeFields(s)
	if err != nil {
		return fmt.Errorf("parse composite fields %s: %w", s, err)
	}
	if len(fields) != 2 {
		return fmt.Errorf("parse composite fields %s", s)
	}
	c.X, err = ParseFloat(fields[0])
	if err != nil {
		return fmt.Errorf("parse x %s: %w", fields[0], err)
	}
	c.Y, err = ParseFloat(fields[1])
	if err != nil {
		return fmt.Errorf("parse y %s: %w", fields[1], err)
	}
	return nil
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
	s, err := ParseString(src)
	if err != nil {
		return fmt.Errorf("parse string: %w", err)
	}
	if s == "" {
		return nil
	}
	fields, err := sql.ParseCompositeFields(s)
	if err != nil {
		return fmt.Errorf("parse composite fields %s: %w", s, err)
	}
	if len(fields) != 3 {
		return fmt.Errorf("parse composite fields %s", s)
	}
	p.Code = fields[0]
	p.Name = fields[1]
	if len(fields[2]) > 0 {
		p.Location = new(Point)
		if err := p.Location.Scan(fields[2]); err != nil {
			return fmt.Errorf("scan place.location: %w", err)
		}
	}
	return nil
}

func (p *Place) Value() (driver.Value, error) {
	if p == nil {
		return nil, nil
	}
	loc, err := p.Location.Value()
	if err != nil {
		return nil, fmt.Errorf("get location value: %w", err)
	}
	if locStr, ok := loc.(string); ok {
		loc = sql.Escape(locStr)
	}
	s := fmt.Sprintf("(%s,%s,%s)", sql.Escape(p.Code), sql.Escape(p.Name), loc)
	fmt.Println(s)
	return s, nil
}
