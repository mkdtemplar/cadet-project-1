package utils

import (
	"html"
	"strings"
)

func Clean(str string) string {
	return strings.TrimSpace(html.EscapeString(str))
}
