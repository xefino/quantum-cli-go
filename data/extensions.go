package data

import (
	"math/big"

	"github.com/shopspring/decimal"
)

// Size of the integer values we want to save (designed to fit inside an int64)
var offset = big.NewInt(1e18)

// NewFromDecimal creates a new representation of our Decimal from a decimal.Decimal
func NewFromDecimal(in decimal.Decimal) *Decimal {
	coefficient := in.Coefficient()
	sign := int64(coefficient.Sign())

	// First, get the absolute value of the coefficient
	rest := new(big.Int)
	rest.Abs(coefficient)

	// Next, iteratively div-mod the coefficient until there's no data remaining
	var ints []int64
	r := new(big.Int)
	for rest.BitLen() != 0 {

		// Divide the remaining value by the width of our integer value, saving the remainder
		// of the division to our temporary value, r and saving the quotient to the remainder
		rest.DivMod(rest, offset, r)

		// Append the remainder to our list (ensuring that we save the sign value)
		ints = append(ints, r.Int64()*sign)
	}

	// Inject the parts and the exponent into a Decimal value and return it
	return &Decimal{Value: ints, Exp: in.Exponent()}
}

// ToDecimal converts our internal representation of a Decimal to a decimal.Decimal
func (d *Decimal) ToDecimal() *decimal.Decimal {

	// First, create our decimal value with a zero-value
	resp := decimal.New(0, 0)

	// Next, iterate over all our sub-values and add them into the total
	for i, value := range d.Value {

		// Attempt to convert the value to its decimal equivalent based on where it is in the list
		temp := decimal.New(value, int32(i*18)+d.Exp)

		// Add the temporary value to the total
		resp = resp.Add(temp)
	}

	// Finally, return our total
	return &resp
}
