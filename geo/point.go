package geo

import (
	"database/sql/driver"
	"fmt"
	"strings"

	"github.com/gopub/gox/sql"
)

const (
	PI          = 3.14159265
	EarthRadius = 6378.1 //km
	EarthCircle = 2 * PI * EarthRadius
	Degree      = EarthCircle * 1000 / 360
)

type Point struct {
	X float64 `json:"x"` // X is longitude for geodetic coordinate
	Y float64 `json:"y"` // Y is latitude for geodetic coordinate
}

func NewPoint(x, y float64) *Point {
	return &Point{}
}

var _ driver.Valuer = (*Point)(nil)
var _ sql.Scanner = (*Point)(nil)

func (p *Point) Scan(src interface{}) error {
	if src == nil {
		return nil
	}

	var s string
	switch v := src.(type) {
	case string:
		s = v
	case []byte:
		s = string(v)
	default:
		return fmt.Errorf("cannot parse %v into string", src)
	}
	if s == "" {
		return nil
	}
	fields, err := sql.ParseCompositeFields(s)
	if err != nil {
		return fmt.Errorf("parse composite fields %s: %w", s, err)
	}
	if len(fields) == 1 {
		fields = strings.Split(fields[0], " ")
	}
	if len(fields) != 2 {
		return fmt.Errorf("parse composite fields %s", s)
	}
	_, err = fmt.Sscanf(fields[0], "%f", &p.X)
	if err != nil {
		return fmt.Errorf("parse x %s: %w", fields[0], err)
	}
	_, err = fmt.Sscanf(fields[1], "%f", &p.Y)
	if err != nil {
		return fmt.Errorf("parse y %s: %w", fields[1], err)
	}
	return nil
}

func (p *Point) Value() (driver.Value, error) {
	if p == nil {
		return nil, nil
	}
	return fmt.Sprintf("POINT(%f %f)", p.X, p.Y), nil
}
