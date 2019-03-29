package gox

import (
	"encoding/json"
)

/*
 * Int64List, Float64List, StringList
 * As slice is not supported in gomobile bind
 */

type Int64List struct {
	List []int64
}

func NewIntList() *Int64List {
	l := &Int64List{}
	return l
}

func (l *Int64List) Len() int {
	if l == nil {
		return 0
	}
	return len(l.List)
}

func (l *Int64List) Get(index int) int64 {
	return l.List[index]
}

func (l *Int64List) Add(val int64) {
	l.List = append(l.List, val)
}

func (l *Int64List) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &l.List)
}

func (l *Int64List) MarshalJSON() ([]byte, error) {
	return json.Marshal(l.List)
}

type Float64List struct {
	elements []float64
}

func NewFloatList() *Float64List {
	l := &Float64List{}
	return l
}

func (l *Float64List) Len() int {
	if l == nil {
		return 0
	}
	return len(l.elements)
}

func (l *Float64List) Get(index int) float64 {
	return l.elements[index]
}

func (l *Float64List) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &l.elements)
}

func (l *Float64List) MarshalJSON() ([]byte, error) {
	return json.Marshal(l.elements)
}

type StringList struct {
	elements []string
}

func NewStringList() *StringList {
	l := &StringList{}
	return l
}

func (l *StringList) Len() int {
	if l == nil {
		return 0
	}
	return len(l.elements)
}

func (l *StringList) Get(index int) string {
	return l.elements[index]
}

func (l *StringList) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &l.elements)
}

func (l *StringList) MarshalJSON() ([]byte, error) {
	return json.Marshal(l.elements)
}

type ImageList struct {
	elements []*Image
}

func (l *ImageList) Len() int {
	if l == nil {
		return 0
	}
	return len(l.elements)
}

func (l *ImageList) Get(index int) *Image {
	return l.elements[index]
}
