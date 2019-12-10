package gox

import (
	"regexp"
	"strings"
)

var whitespaceRegexp = regexp.MustCompile(`\\s`)

// CombineSpaces combines multiple spaces to a single space
func CombineSpaces(s string) string {
	s = strings.TrimSpace(s)
	s = whitespaceRegexp.ReplaceAllString(s, " ")
	return s
}
