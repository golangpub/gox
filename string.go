package gox

import (
	"regexp"
	"strings"
)

func SmartLen(s string) int {
	n := 0
	for _, c := range s {
		if c <= 255 {
			n += 1
		} else {
			n += 2
		}
	}

	return n
}

var whitespaceRegexp = regexp.MustCompile(`\\s`)

func CombineSpaces(s string) string {
	s = strings.TrimSpace(s)
	s = whitespaceRegexp.ReplaceAllString(s, " ")
	return s
}
