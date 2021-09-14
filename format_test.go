package rut_test

import (
	"github.com/mnavarrocarter/go-rut"
	"testing"
)

var stringTests = []struct {
	name     string
	input    uint32
	expected string
}{
	{
		name:     "#1",
		input:    24_232_442,
		expected: "24.232.442-0",
	},
	{
		name:     "#2",
		input:    16_894_365,
		expected: "16.894.365-2",
	},
	{
		name:     "#3",
		input:    15_450_088,
		expected: "15.450.088-K",
	},
}

func TestString(t *testing.T) {
	for _, test := range stringTests {
		t.Run(test.name, func(t *testing.T) {
			r := rut.Make(test.input)
			str := r.String()
			if str != test.expected {
				t.Errorf("Unexpected string \"%s\" received. Expected is \"%s\"", str, test.expected)
			}
		})
	}
}
