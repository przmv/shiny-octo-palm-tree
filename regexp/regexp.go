package regexp

import (
	"math"
	"regexp"
	sortpkg "sort"
	"strconv"
	"strings"

	"githib.com/przmv/shiny-octo-palm-tree/errors"
	"githib.com/przmv/shiny-octo-palm-tree/sort"
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

// TestValidity checks if a string matches a regular expression pattern.
func TestValidity(s string) bool {
	return re.MatchString(s)
}

// AverageNumber returns an average number from all the numbers from the string s.
func AverageNumber(s string) (float64, error) {
	if !TestValidity(s) {
		return 0, errors.InvalidStringError
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
		return "", errors.InvalidStringError
	}
	a := words.FindAllString(s, -1)
	return strings.Join(a, " "), nil
}

// StoryStats returns the shortest word, the longest word, the average word
// length, the list (or empty list) of all words from the story that have the
// length the same as the average length rounded up and down.
func StoryStats(s string) (shortest, longest string, avg float64, asAvg []string, err error) {
	if !TestValidity(s) {
		err = errors.InvalidStringError
		return
	}
	a := words.FindAllString(s, -1)
	sortpkg.Sort(sort.ByLength(a))
	shortest = a[0]
	longest = a[len(a)-1]
	sum := 0.0
	m := make(map[float64][]string)
	for _, s := range a {
		n := float64(len(s))
		sum += n
		m[n] = append(m[n], s)
	}
	avg = sum / (float64(len(a)))
	ceil := math.Ceil(avg)
	floor := math.Floor(avg)
	asAvg = append(m[ceil], m[floor]...)
	return
}
