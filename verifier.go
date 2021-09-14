package rut

var verifiersArr = [11]Verifier{'1', '2', '3', '4', '5', '6', '7', '8', '9', 'K', '0'}

type Verifier byte

func (v Verifier) Valid() bool {
	for k := range verifiersArr {
		if verifiersArr[k] == v {
			return true
		}
	}
	return false
}

func (v Verifier) Equals(b Verifier) bool {
	return v == b
}

// GetVerifierByMod obtains a verifier from the modulo 11
func GetVerifierByMod(dv uint8) (Verifier, error) {
	if len(verifiersArr) < int(dv) {
		return verifiersArr[0], ErrInvalidVerifier
	}
	k := dv - 1
	return verifiersArr[k], nil
}

// ParseVerifier parses a verifier from a string
func ParseVerifier(str string) Verifier {
	return Verifier(str[0])
}

func (v Verifier) String() string {
	return string(v)
}
