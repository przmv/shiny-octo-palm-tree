package regexp

import (
	"errors"
	"regexp"
)

// Regular expression to match a sequence of numbers followed by dash followed
// by text.
var re = regexp.MustCompile(`^(?:[[:digit:]]+-[[:alpha:]]+-?)+[^-]$`)

var InvalidStringError = errors.New("invalid string")

// TestValidity checks if a string matches a regular expression pattern.
func TestValidity(s string) bool {
	return re.MatchString(s)
}

// AverageNumber returns an average number from all the numbers from the string s.
func AverageNumber(s string) (float64, error) {
	// TODO
	return 0, nil
}
