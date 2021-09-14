package rut_test

import (
	"github.com/mnavarrocarter/go-rut"
	"testing"
)

var parseTests = []struct {
	name     string
	input    string
	err      string
	isPerson bool
	isOrg    bool
}{
	{
		name:     "nicely formatted",
		input:    "24.232.442-0",
		isPerson: true,
		isOrg:    false,
	},
	{
		name:     "no dots",
		input:    "16894365-2",
		isPerson: true,
		isOrg:    false,
	},
	{
		name:  "good format but invalid",
		input: "16.894.365-K",
		err:   "invalid verifier digit",
	},
	{
		name:  "bad rut",
		input: "1K.24a.45b-5",
		err:   "parse error",
	},
	{
		name:  "really bad rut",
		input: "224&$£^$£42342",
		err:   "parse error",
	},
	{
		name:     "lower case k",
		input:    "15.450.088-k",
		isPerson: true,
		isOrg:    false,
	},
	{
		name:  "invalid verifier",
		input: "15.450.088-Z",
		err:   "invalid verifier digit",
	},
	{
		name:     "badly formatted",
		input:    "24. 736  732 - 2",
		isPerson: true,
		isOrg:    false,
	},
	{
		name:     "org rut with no spaces",
		input:    "760194530",
		isPerson: false,
		isOrg:    true,
	},
}

func TestParse(t *testing.T) {
	for _, test := range parseTests {
		t.Run(test.name, func(t *testing.T) {
			r, err := rut.Parse(test.input)
			if test.err == "" {
				if err != nil {
					t.Errorf("Unexpected error received: %s", err.Error())
				}
				if r.IsPerson() != test.isPerson {
					t.Errorf("Rut type mismatch")
				}
				if r.IsOrg() != test.isOrg {
					t.Errorf("Rut type mismatch")
				}
			}
			if test.err != "" {
				if err == nil {
					t.Fatalf("Expected error was not received")
				}
				if err.Error() != test.err {
					t.Errorf("Unexpected error message \"%s\" received. Expecting \"%s\"", err.Error(), test.err)
				}
			}
		})
	}
}
