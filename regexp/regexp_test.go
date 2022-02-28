package regexp_test

import (
	"testing"

	"githib.com/przmv/shiny-octo-palm-tree/regexp"
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
		got := regexp.TestValidity(s)
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
		Error:   regexp.InvalidStringError,
		Average: 0,
	},
}

func TestAverageNumber(t *testing.T) {
	for _, tc := range averageNumberTestCases {
		s := tc.String
		expected := tc.Average
		expectedErr := tc.Error
		got, err := regexp.AverageNumber(s)
		if expected != got {
			t.Errorf("%q: expected %v, got %v", s, expected, got)
		}
		if expectedErr != err {
			t.Errorf("%q: expected error %v, got %v", s, expectedErr, err)
		}
	}
}
