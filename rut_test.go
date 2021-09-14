package rut_test

import (
	"github.com/mnavarrocarter/go-rut"
	"testing"
)

var parseTests = []struct {
	name  string
	input string
	err   string
}{
	{
		name:  "nicely formatted",
		input: "24.232.442-0",
	},
	{
		name:  "no dots",
		input: "16894365-2",
	},
	{
		name:  "good format but invalid",
		input: "16.894.365-K",
		err:   "invalid verifier digit",
	},
	{
		name:  "bad rut",
		input: "1K.24a.45b-5",
		err:   "parse error: strconv.ParseUint: parsing \"1K24A45B\": invalid syntax",
	},
	{
		name:  "really bad rut",
		input: "224&$£^$£42342",
		err:   "parse error: strconv.ParseUint: parsing \"224&$£^$£4234\": invalid syntax",
	},
	{
		name:  "lower case k",
		input: "15.450.088-k",
	},
	{
		name:  "invalid verifier",
		input: "15.450.088-Z",
		err:   "invalid verifier digit",
	},
	{
		name:  "badly formatted",
		input: "24. 736  732 - 2",
	},
}

func TestParse(t *testing.T) {
	for _, test := range parseTests {
		t.Run(test.name, func(t *testing.T) {
			_, err := rut.Parse(test.input)
			if test.err == "" {
				if err != nil {
					t.Errorf("Unexpected error received: %s", err.Error())
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
