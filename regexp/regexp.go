package regexp

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

var (
	// Regular expression to match a sequence of numbers followed by dash
	// followed by text.
	re = regexp.MustCompile(`^(?:[[:digit:]]+-[[:alpha:]]+-?)+[^-]$`)

	// Regular expression to match numbers.
	numbers = regexp.MustCompile(`[[:digit:]]+`)

	// Regular expression to match words.
	words = regexp.MustCompile(`[[:alpha:]]+`)
)

var InvalidStringError = errors.New("invalid string")

// TestValidity checks if a string matches a regular expression pattern.
func TestValidity(s string) bool {
	return re.MatchString(s)
}

// AverageNumber returns an average number from all the numbers from the string s.
func AverageNumber(s string) (float64, error) {
	if !TestValidity(s) {
		return 0, InvalidStringError
	}
	a := numbers.FindAllString(s, -1)
	sum := 0
	for _, s := range a {
		i, err := strconv.Atoi(s)
		if err != nil {
			return 0, err
		}
		sum += i
	}
	avg := (float64(sum)) / (float64(len(a)))
	return avg, nil
}

// WholeStrory returns a text that is composed from all the text words
// separated by spaces.
func WholeStrory(s string) (string, error) {
	if !TestValidity(s) {
		return "", InvalidStringError
	}
	a := words.FindAllString(s, -1)
	return strings.Join(a, " "), nil
}
