package mobile

import "github.com/gopub/gox"

type PhoneNumberList struct {
	List []*gox.PhoneNumber
}

func (l *PhoneNumberList) Len() int {
	if l == nil {
		return 0
	}
	return len(l.List)
}

func (l *PhoneNumberList) Get(index int) *gox.PhoneNumber {
	return l.List[index]
}

func NewPhoneNumberList() *PhoneNumberList {
	return &PhoneNumberList{}
}

func (l *PhoneNumberList) Add(phoneNumber *gox.PhoneNumber) {
	l.List = append(l.List, phoneNumber)
}

func (l *PhoneNumberList) Contains(phoneNumber *gox.PhoneNumber) bool {
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
