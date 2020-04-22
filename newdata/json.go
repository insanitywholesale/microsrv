package newdata

import (
	"encoding/json"
	"io"
)

func (p *Products) ToJSON(w io.Writer) error {
	err := json.NewEncoder(w)
	return err.Encode(p)
}

func (p *Product) FromJSON(r io.Reader) error {
	err := json.NewDecoder(r)
	return err.Decode(p)
}
