package regexp_test

import (
	"testing"

	"githib.com/przmv/shiny-octo-palm-tree/regexp"
)

const (
	validString   = `23-ab-48-caba-56-haha`
	invalidString = `aaa-bbb-ccc-ddd-eee-fff`
)

var testCases = []struct {
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
	for _, tc := range testCases {
		s := tc.String
		expected := tc.IsValid
		got := regexp.TestValidity(s)
		if expected != got {
			t.Errorf("%q: expected %v, got %v", s, expected, got)
		}
	}
}
