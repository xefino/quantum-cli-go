package data

import (
	"bytes"
	"encoding/gob"
	"fmt"

	"github.com/shopspring/decimal"
)

// NewFromDecimal creates a new representation of our Decimal from a decimal.Decimal
func NewFromDecimal(in decimal.Decimal) (*Decimal, error) {

	// Create our encoder from a new bytes buffer
	buf := new(bytes.Buffer)
	encoder := gob.NewEncoder(buf)

	// Attempt to gob-encode the data to a Decimal; if this fails then return an error
	if err := encoder.Encode(in); err != nil {
		return nil, fmt.Errorf("Failed to encode Decimal")
	}

	// The encoding succeeded; return the decimal
	return &Decimal{Data: buf.Bytes()}, nil
}

// ToDecimal converts our internal representation of a Decimal to a decimal.Decimal
func (d *Decimal) ToDecimal() (*decimal.Decimal, error) {
	var out decimal.Decimal

	// Create our decoder from a buffer of the data
	decoder := gob.NewDecoder(bytes.NewBuffer(d.Data))

	// Attempt to gob-decode the data to a Decimal; if this fails then return an error
	if err := decoder.Decode(&out); err != nil {
		return nil, fmt.Errorf("Failed to decode to Decimal")
	}

	// The decoding succeeded; return the decimal
	return &out, nil
}
