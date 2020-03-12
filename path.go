package gox

import (
	"path"
	"strings"
)

func JoinURLPath(segment ...string) string {
	p := path.Join(segment...)
	p = strings.Replace(p, ":/", "://", 1)
	return p
}
