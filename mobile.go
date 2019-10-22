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

func NewInt64List() *Int64List {
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

type IntList struct {
	List []int
}

func NewIntList() *IntList {
	l := &IntList{}
	return l
}

func (l *IntList) Len() int {
	if l == nil {
		return 0
	}
	return len(l.List)
}

func (l *IntList) Get(index int) int {
	return l.List[index]
}

func (l *IntList) Add(val int) {
	l.List = append(l.List, val)
}

func (l *IntList) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &l.List)
}

func (l *IntList) MarshalJSON() ([]byte, error) {
	return json.Marshal(l.List)
}

type Float64List struct {
	List []float64
}

func NewFloatList() *Float64List {
	l := &Float64List{}
	return l
}

func (l *Float64List) Len() int {
	if l == nil {
		return 0
	}
	return len(l.List)
}

func (l *Float64List) Get(index int) float64 {
	return l.List[index]
}

func (l *Float64List) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &l.List)
}

func (l *Float64List) MarshalJSON() ([]byte, error) {
	return json.Marshal(l.List)
}

type StringList struct {
	List []string
}

func NewStringList() *StringList {
	l := &StringList{}
	return l
}

func (l *StringList) Len() int {
	if l == nil {
		return 0
	}
	return len(l.List)
}

func (l *StringList) Get(index int) string {
	return l.List[index]
}

func (l *StringList) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &l.List)
}

func (l *StringList) MarshalJSON() ([]byte, error) {
	return json.Marshal(l.List)
}

type ImageList struct {
	List []*Image
}

func (l *ImageList) Len() int {
	if l == nil {
		return 0
	}
	return len(l.List)
}

func (l *ImageList) Get(index int) *Image {
	return l.List[index]
}

type PhoneNumberList struct {
	List []*PhoneNumber
}

func (l *PhoneNumberList) Len() int {
	if l == nil {
		return 0
	}
	return len(l.List)
}

func (l *PhoneNumberList) Get(index int) *PhoneNumber {
	return l.List[index]
}

// For iOS/Android

func NewPhoneNumberList() *PhoneNumberList {
	return &PhoneNumberList{}
}

func (l *PhoneNumberList) Add(phoneNumber *PhoneNumber) {
	l.List = append(l.List, phoneNumber)
}
