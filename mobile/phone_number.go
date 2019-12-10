package mobile

import (
	"github.com/gopub/gox"
)

type PhoneNumber = gox.PhoneNumber

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

func NewPhoneNumberList() *PhoneNumberList {
	return &PhoneNumberList{}
}

func (l *PhoneNumberList) Add(phoneNumber *PhoneNumber) {
	l.List = append(l.List, phoneNumber)
}

func (l *PhoneNumberList) Contains(phoneNumber *PhoneNumber) bool {
	for _, pn := range l.List {
		if pn.String() == phoneNumber.String() {
			return true
		}
	}
	return false
}

func (l *PhoneNumberList) ContainsString(phoneNumber string) bool {
	for _, pn := range l.List {
		if pn.String() == phoneNumber {
			return true
		}
	}
	return false
}

func NewPhoneNumber(callingCode int, number int64) *PhoneNumber {
	return gox.NewPhoneNumber(callingCode, number)
}

func ParsePhoneNumber(s string) (*PhoneNumber, error) {
	return gox.ParsePhoneNumber(s)
}

func TidyPhoneNumber(s string, code int) *PhoneNumber {
	return gox.TidyPhoneNumber(s, code)
}
