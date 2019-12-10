package mobile

import (
	"encoding/json"
	"regexp"
	"strings"
)

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

func SmartLen(s string) int {
	n := 0
	for _, c := range s {
		if c <= 255 {
			n++
		} else {
			n += 2
		}
	}

	return n
}

var whitespaceRegexp = regexp.MustCompile(`\\s`)

// CombineSpaces cobines multiple spaces to a single space
func CombineSpaces(s string) string {
	s = strings.TrimSpace(s)
	s = whitespaceRegexp.ReplaceAllString(s, " ")
	return s
}
