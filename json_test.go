package rut_test

import (
	"bytes"
	"encoding/json"
	"github.com/mnavarrocarter/go-rut"
	"testing"
)

type payload struct {
	Name string   `json:"name"`
	Rut  *rut.RUT `json:"rut"`
}

var unmarshalTests = []struct {
	name     string
	json     []byte
	expected string
}{
	{
		name:     "test one",
		json:     []byte(`{"name":"test one","rut":"16 894 365 2"}`),
		expected: "16.894.365-2",
	},
}

func TestUnmarshal(t *testing.T) {
	for _, test := range unmarshalTests {
		t.Run(test.name, func(t *testing.T) {
			p := &payload{}
			err := json.Unmarshal(test.json, p)
			if err != nil {
				t.Fatalf("unexpected error: %s", err.Error())
			}
			if test.expected != p.Rut.String() {
				t.Errorf(`expected rut "%s" is not the same as received "%s"`, test.expected, p.Rut.String())
			}
		})
	}
}

var marshalTests = []struct {
	name         string
	payload      payload
	expectedJson []byte
}{
	{
		name: "test one",
		payload: payload{
			Name: "test one",
			Rut:  rut.MustParse("16 894 365 2"),
		},
		expectedJson: []byte(`{"name":"test one","rut":"16.894.365-2"}`),
	},
}

func TestMarshal(t *testing.T) {
	for _, test := range marshalTests {
		t.Run(test.name, func(t *testing.T) {
			b, err := json.Marshal(&test.payload)
			if err != nil {
				t.Fatalf("unexpected error: %s", err.Error())
			}
			if bytes.Compare(test.expectedJson, b) != 0 {
				t.Fatalf("json payloads are not the same")
			}
		})
	}
}
