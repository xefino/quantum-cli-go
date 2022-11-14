package data

import (
	"math"
	"strconv"

	"github.com/shopspring/decimal"
)

// NewFromDecimal creates a new representation of our Decimal from a decimal.Decimal
func NewFromDecimal(in decimal.Decimal) *Decimal {
	const width = 18

	// First, convert the input decimal coefficient to a big.Int and from there to a string
	asStr := in.Coefficient().Text(10)
	strLen := len(asStr)

	// Next, if the original value was negative then we'll calculate an offset to use when
	// calculating the number of int64's and when propogating the negative sign
	offset := 0
	if in.IsNegative() {
		offset = 1
	}

	// Now, determine the number of int64's we want to split the value into
	length := int(math.Ceil(float64(strLen-offset) / width))

	// Finally, iterate over each of the int64's and extract a portion of the string to load into it
	ints := make([]int64, length)
	for i := 1; i <= length; i++ {

		// First, calculate the start position (we want this to be from the end of the string)
		start := strLen - (i * width)

		// Next, calculate the end of the string, which should have a constant value
		end := start + width

		// Now, if we're at risk of walking off the front of the string then set the start value to zero
		if start < 0 || (start == 1 && asStr[0] == '-') {
			start = 0
		}

		// Finally, parse the string back into an integer; since we came from an integer this cannot fail
		// If the value we're parsing from is negative but we're working with an intermediate value then
		// we'll need to be sure that the negative sign is propogated
		ints[i-1], _ = strconv.ParseInt(asStr[start:end], 10, 64)
		if offset == 1 && ints[i-1] > 0 {
			ints[i-1] = 0 - ints[i-1]
		}
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
