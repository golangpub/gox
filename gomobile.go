package types

import "encoding/json"

/*
 * IntArray, FloatArray, StringArray
 * As slice is not supported in gomobile bind
 */

type IntArray struct {
	elements []int64
}

func NewIntArray(elems []int64) *IntArray {
	a := &IntArray{}
	a.elements = append(a.elements, elems...)
	return a
}

func (a *IntArray) Size() int {
	return len(a.elements)
}

func (a *IntArray) Get(index int) int64 {
	return a.elements[index]
}

func (a *IntArray) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &a.elements)
}

func (a *IntArray) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.elements)
}

type FloatArray struct {
	elements []float64
}

func NewFloatArray(elems []float64) *FloatArray {
	a := &FloatArray{}
	a.elements = append(a.elements, elems...)
	return a
}

func (a *FloatArray) Size() int {
	return len(a.elements)
}

func (a *FloatArray) Get(index int) float64 {
	return a.elements[index]
}

func (a *FloatArray) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &a.elements)
}

func (a *FloatArray) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.elements)
}

type StringArray struct {
	elements []string
}

func NewStringArray(elems []string) *StringArray {
	a := &StringArray{}
	a.elements = append(a.elements, elems...)
	return a
}

func (a *StringArray) Size() int {
	return len(a.elements)
}

func (a *StringArray) Get(index int) string {
	return a.elements[index]
}

func (a *StringArray) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &a.elements)
}

func (a *StringArray) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.elements)
}
