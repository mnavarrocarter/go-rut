package rut

import "strings"

func (r *RUT) UnmarshalJSON(bytes []byte) error {
	str := strings.Trim(string(bytes), `"`)
	rb, err := Parse(str)
	if err != nil {
		return err
	}
	r.number = rb.number
	r.verifier = rb.verifier
	return nil
}

func (r *RUT) MarshalJSON() ([]byte, error) {
	return []byte(`"` + r.String() + `"`), nil
}
