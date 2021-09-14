package rut

import (
	"errors"
	"strings"
)

var sanitizer = strings.NewReplacer("-", "", " ", "", ".", "")

var ErrInvalidVerifier = errors.New("invalid verifier digit")

// A RUT is a unique number to identify a chilean entity (person or organization) for taxing purposes
// Generally speaking, every resident of Chile and every registered company has a RUT.
//
// Company RUT are all those RUT bigger than ORGDelimiter
//
// A RUT is not as private as a SSN, but it is considered sensitive information nonetheless. You must
// be very careful where you store it or display it.
type RUT struct {
	number   NationalNumber
	verifier Verifier
}

// Parse parses rut into RUT.
//
// When a parsing error occurs, RUT == nil and error is an error explaining what happened, wrapping ErrParse.
//
// When the Verifier is invalid, then RUT != nil, but ErrInvalidVerifier is returned too.
func Parse(rut string) (*RUT, error) {
	rut = strings.ToUpper(sanitizer.Replace(rut))
	number, err := ParseNationalNumber(rut[0 : len(rut)-1])

	if err != nil {
		return nil, err
	}

	r := &RUT{
		number:   number,
		verifier: ParseVerifier(rut[len(rut)-1:]),
	}

	if !r.verifier.Valid() || !r.verifier.Equals(r.number.Verifier()) {
		return r, ErrInvalidVerifier
	}

	return r, nil
}

// MustParse is equivalent to Parse, but it panics if there is an error.
func MustParse(rut string) *RUT {
	r, err := Parse(rut)
	if err != nil {
		panic(err)
	}
	return r
}

// New creates a RUT out of a number.
//
// The Verifier is calculated automatically.
func New(number uint32) *RUT {
	num := NationalNumber(number)

	return &RUT{
		number:   num,
		verifier: num.Verifier(),
	}
}

// IsPerson returns true if this is a peron RUT.
func (r *RUT) IsPerson() bool {
	return r.number < ORGDelimiter
}

// IsOrg returns true if this is an organization RUT.
func (r *RUT) IsOrg() bool {
	return r.number >= ORGDelimiter
}

// Number returns the RUT number.
func (r *RUT) Number() uint32 {
	return uint32(r.number)
}

// Verifier returns the RUT verifier.
func (r *RUT) Verifier() byte {
	return byte(r.verifier)
}

// Equal returns true is both ruts are equal.
//
// RUT is considered equal when they have the same verifier.
func (r *RUT) Equal(s *RUT) bool {
	return r.Number() == s.Number() && r.Verifier() == s.Verifier()
}
