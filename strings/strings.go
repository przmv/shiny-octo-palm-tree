package strings

import (
	"strings"
	"unicode"
)

// TestValidity checks if a string complies with the format.
func TestValidity(s string) bool {
	a := strings.Split(s, "-")
	if len(a) < 2 {
		return false
	}
	for i, s := range a {
		if i%2 == 0 {
			for _, r := range s {
				if !unicode.IsNumber(r) {
					return false
				}
			}
		} else {
			for _, r := range s {
				if !unicode.Is(unicode.Latin, r) {
					return false
				}
			}
		}
	}
	return true
}
