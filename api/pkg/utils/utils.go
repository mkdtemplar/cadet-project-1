package utils

import (
	"html"
	"strings"
)

func CleanUserData(str string) string {
	return strings.TrimSpace(html.EscapeString(str))
}
