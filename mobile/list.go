package mobile

import "encoding/json"

/*
 * IntList, FloatList, StringList
 * As slice is not supported in gomobile bind
 */

type IntList struct {
	elements []int64
}

func NewIntList(elements []int64) *IntList {
	l := &IntList{}
	l.elements = append(l.elements, elements...)
	return l
}

func (l *IntList) Size() int {
	return len(l.elements)
}

func (l *IntList) Get(index int) int64 {
	return l.elements[index]
}

func (l *IntList) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &l.elements)
}

func (l *IntList) MarshalJSON() ([]byte, error) {
	return json.Marshal(l.elements)
}

type FloatList struct {
	elements []float64
}

func NewFloatList(elements []float64) *FloatList {
	l := &FloatList{}
	l.elements = append(l.elements, elements...)
	return l
}

func (l *FloatList) Size() int {
	return len(l.elements)
}

func (l *FloatList) Get(index int) float64 {
	return l.elements[index]
}

func (l *FloatList) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &l.elements)
}

func (l *FloatList) MarshalJSON() ([]byte, error) {
	return json.Marshal(l.elements)
}

type StringList struct {
	elements []string
}

func NewStringList(elements []string) *StringList {
	l := &StringList{}
	l.elements = append(l.elements, elements...)
	return l
}

func (l *StringList) Size() int {
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
