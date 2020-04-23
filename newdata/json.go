package newdata

import (
	"encoding/json"
	"io"
)

// ToJSON serializes the given interface into a string based JSON format
func ToJSON(i interface{}, w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(i)
}

// FromJSON deserializes the object from JSON string
//in an io.Reader to the given interface
func FromJSON(i interface{}, r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(i)
}

func (p *Products) ToJSON(w io.Writer) error {
	err := json.NewEncoder(w)
	return err.Encode(p)
}

func (p *Product) FromJSON(r io.Reader) error {
	err := json.NewDecoder(r)
	return err.Decode(p)
}
