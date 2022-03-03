package strings_test

import (
	"testing"

	"githib.com/przmv/shiny-octo-palm-tree/errors"
	"githib.com/przmv/shiny-octo-palm-tree/strings"
)

const (
	validString   = `23-ab-48-caba-56-haha`
	invalidString = `aaa-bbb-ccc-ddd-eee-fff`
)

var validityTestCases = []struct {
	String  string
	IsValid bool
}{
	{
		String:  validString,
		IsValid: true,
	},
	{
		String:  invalidString,
		IsValid: false,
	},
}

func TestTestValidity(t *testing.T) {
	for _, tc := range validityTestCases {
		s := tc.String
		expected := tc.IsValid
		got := strings.TestValidity(s)
		if expected != got {
			t.Errorf("%q: expected %v, got %v", s, expected, got)
		}
	}
}

var averageNumberTestCases = []struct {
	String  string
	Error   error
	Average float64
}{
	{
		String:  validString,
		Error:   nil,
		Average: 42.333333333333336,
	},
	{
		String:  invalidString,
		Error:   errors.InvalidStringError,
		Average: 0,
	},
}

func TestAverageNumber(t *testing.T) {
	for _, tc := range averageNumberTestCases {
		s := tc.String
		expected := tc.Average
		expectedErr := tc.Error
		got, err := strings.AverageNumber(s)
		if expected != got {
			t.Errorf("%q: expected %v, got %v", s, expected, got)
		}
		if expectedErr != err {
			t.Errorf("%q: expected error %v, got %v", s, expectedErr, err)
		}
	}
}

var wholeStoryTestCases = []struct {
	String string
	Error  error
	Text   string
}{
	{
		String: validString,
		Error:  nil,
		Text:   "ab caba haha",
	},
	{
		String: invalidString,
		Error:  errors.InvalidStringError,
		Text:   "",
	},
}

func TestWholeStory(t *testing.T) {
	for _, tc := range wholeStoryTestCases {
		s := tc.String
		expected := tc.Text
		expectedErr := tc.Error
		got, err := strings.WholeStrory(s)
		if expected != got {
			t.Errorf("%q: expected %q, got %q", s, expected, got)
		}
		if expectedErr != err {
			t.Errorf("%q: expected error %v, got %v", s, expectedErr, err)
		}
	}
}
