package rut

import (
	"errors"
	"strconv"
)

var ErrParse = errors.New("parse error")

const ORGDelimiter = NationalNumber(50_000_000)

// A NationalNumber represents a number to uniquely identify any Chilean person or organization
//
// The biggest this number can be is 4_294_967_295, but most ruts are under 100_000_000
type NationalNumber uint32

func (n NationalNumber) toSeq() []uint8 {
	var ret []uint8

	for n != 0 {
		ret = append(ret, uint8(n%10))
		n /= 10
	}

	return ret
}

// ParseNationalNumber parses a string to a NationalNumber.
//
// Returns ErrParse if there was an error parsing the number
func ParseNationalNumber(num string) (NationalNumber, error) {
	parsed, err := strconv.ParseUint(num, 10, 32)
	if err != nil {
		return NationalNumber(0), ErrParse
	}
	return NationalNumber(parsed), nil
}

// Verifier calculates the Verifier digit of the NationalNumber.
//
// It uses the standard module 11 algorithm.
func (n NationalNumber) Verifier() Verifier {
	seq := n.toSeq()
	x := uint8(2)
	s := uint8(0)
	for _, d := range seq {
		if x > 7 {
			x = 2
		}
		s += d * x
		x += 1
	}
	dv := 11 - (s % 11)
	v, _ := GetVerifierByMod(dv)
	return v
}

// Rut creates a RUT for this NationalNumber.
func (n NationalNumber) Rut() *RUT {
	return &RUT{
		number:   n,
		verifier: n.Verifier(),
	}
}
