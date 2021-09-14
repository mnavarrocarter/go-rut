package rut_test

import (
	"github.com/mnavarrocarter/go-rut"
	"testing"
)

var verifierTests = []struct {
	name     string
	input    rut.NationalNumber
	expected rut.Verifier
}{
	{
		name:     "verifier 0",
		input:    24_232_442,
		expected: '0',
	},
	{
		name:     "verifier 2",
		input:    16_894_365,
		expected: '2',
	},
	{
		name:     "verifier K",
		input:    22_457_309,
		expected: 'K',
	},
	{
		name:     "verifier K #2",
		input:    15_450_088,
		expected: 'K',
	},
}

func TestVerifier(t *testing.T) {
	for _, test := range verifierTests {
		t.Run(test.name, func(t *testing.T) {
			v := test.input.Verifier()
			if !test.expected.Equals(v) {
				t.Errorf("Expected verifier %c for number %d is not the same as received %c", test.expected, test.input, v)
			}
		})
	}
}
