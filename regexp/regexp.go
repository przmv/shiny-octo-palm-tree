package regexp

import "regexp"

// Regular expression to match a sequence of numbers followed by dash followed
// by text.
var re = regexp.MustCompile(`^(?:[[:digit:]]+-[[:alpha:]]+-?)+[^-]$`)

// TestValidity checks if a string matches a regular expression pattern.
func TestValidity(s string) bool {
	return re.MatchString(s)
}
