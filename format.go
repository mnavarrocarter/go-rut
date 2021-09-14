package rut

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// A Formatter is a function that takes num and v and returns a formatted rut string
type Formatter func(num NationalNumber, v Verifier) string

// The DefaultFormatter formats RUT in a standard way, with dots and hyphens. Example:
// 	16.894.365-2
var DefaultFormatter Formatter = func(num NationalNumber, v Verifier) string {
	p := message.NewPrinter(language.Spanish)
	return p.Sprintf("%d-%s", num, v.String())
}

// Converts r to a string
//
// This is equivalent to call Format with the DefaultFormatter
func (r *RUT) String() string {
	return r.Format(DefaultFormatter)
}

// Format formats r using f
// If f == nil, then DefaultFormatter is used
//
// You can implement a custom Formatter easily
func (r *RUT) Format(f Formatter) string {
	if f == nil {
		f = DefaultFormatter
	}
	return f(r.number, r.verifier)
}
