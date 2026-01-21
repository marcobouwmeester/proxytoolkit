package utils

import (
	"strings"
	"unicode"
)

func Slugify(s string) string {
	s = strings.ToLower(s)

	var b strings.Builder
	prevDash := false

	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			b.WriteRune(r)
			prevDash = false
			continue
		}

		// If not alphanumeric, convert to dash (but avoid duplicates)
		if !prevDash {
			b.WriteByte('-')
			prevDash = true
		}
	}

	// Trim leading/trailing '-'
	result := strings.Trim(b.String(), "-")
	return result
}
