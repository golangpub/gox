package gox

import (
	"encoding/json"
	"github.com/gopub/log"
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

func JSONString(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		log.Errorf("Marshal: %v", err)
	}
	return string(b)
}
