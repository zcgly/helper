package helper

import (
	"strings"
	"unicode"
)

func FilterSpecChars(s string) string {
	buf := strings.Builder{}
	for _, r := range s {
		if unicode.IsSymbol(r) {
			continue
		}
		buf.WriteRune(r)
	}
	return buf.String()
}

func GetMaskName(s string) string {
	if s == "" {
		return ""
	}
	rs := []rune(s)
	count := len(rs)
	return strings.Repeat("*", count-1) + string(rs[count-1])
}
