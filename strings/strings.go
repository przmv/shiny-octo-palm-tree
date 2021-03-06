package strings

import (
	"math"
	sortpkg "sort"
	"strconv"
	"strings"
	"unicode"

	"githib.com/przmv/shiny-octo-palm-tree/errors"
	"githib.com/przmv/shiny-octo-palm-tree/sort"
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

// AverageNumber returns an average number from all the numbers from the string s.
func AverageNumber(s string) (float64, error) {
	if !TestValidity(s) {
		return 0, errors.InvalidStringError
	}
	a := strings.Split(s, "-")
	sum, n := 0, 0
	for i, s := range a {
		if i%2 == 0 {
			i, err := strconv.Atoi(s)
			if err != nil {
				return 0, err
			}
			sum += i
			n++
		}
	}
	avg := (float64(sum)) / (float64(n))
	return avg, nil
}

// WholeStory returns a text that is composed from all the text words
// separated by spaces.
func WholeStory(s string) (string, error) {
	if !TestValidity(s) {
		return "", errors.InvalidStringError
	}
	a := strings.Split(s, "-")
	words := make([]string, 0)
	for i, s := range a {
		if i%2 != 0 {
			words = append(words, s)
		}
	}
	return strings.Join(words, " "), nil
}

// StoryStats returns the shortest word, the longest word, the average word
// length, the list (or empty list) of all words from the story that have the
// length the same as the average length rounded up and down.
func StoryStats(s string) (shortest, longest string, avg float64, asAvg []string, err error) {
	s, err = WholeStory(s)
	if err != nil {
		return
	}
	words := strings.Fields(s)
	sortpkg.Sort(sort.ByLength(words))
	shortest = words[0]
	longest = words[len(words)-1]
	sum := 0.0
	m := make(map[float64][]string)
	for _, s := range words {
		n := float64(len(s))
		sum += n
		m[n] = append(m[n], s)
	}
	avg = sum / (float64(len(words)))
	ceil := math.Ceil(avg)
	floor := math.Floor(avg)
	asAvg = append(m[ceil], m[floor]...)
	return
}
