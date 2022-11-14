package data

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/shopspring/decimal"
)

var _ = Describe("Decimal Extensions Tests", func() {

	// Tests that, if the value is greater than 0, then it will be encoded properly
	It("NewFromDecimal - Value greater than 0 - Encoded", func() {

		// First, create our decimal value
		dIn, err := decimal.NewFromString("1234512351234088800000.999")
		Expect(err).ShouldNot(HaveOccurred())

		// Next, attempt to convert it to a new Decimal object
		dOut := NewFromDecimal(dIn)

		// Finally, verify the data
		Expect(dOut).ShouldNot(BeNil())
		Expect(dOut.Exp).Should(Equal(int32(-3)))
		Expect(dOut.Value).Should(HaveLen(2))
		Expect(dOut.Value[0]).Should(Equal(int64(351234088800000999)))
		Expect(dOut.Value[1]).Should(Equal(int64(1234512)))
	})

	// Test that, if the value is less than 0, then it will be encoded properly
	It("NewFromDecimal - Value less than 0 - Encoded", func() {

		// First, create our decimal value
		dIn, err := decimal.NewFromString("-288341660781234512351234088800000.999")
		Expect(err).ShouldNot(HaveOccurred())

		// Next, attempt to convert it to a new Decimal object
		dOut := NewFromDecimal(dIn)

		// Finally, verify the data
		Expect(dOut).ShouldNot(BeNil())
		Expect(dOut.Exp).Should(Equal(int32(-3)))
		Expect(dOut.Value).Should(HaveLen(2))
		Expect(dOut.Value[0]).Should(Equal(int64(-351234088800000999)))
		Expect(dOut.Value[1]).Should(Equal(int64(-288341660781234512)))
	})

	// Tests that the Decimal value will be converted to a decimal.Decimal properly
	It("ToDecimal - Decoded", func() {

		// First, create a valid Decimal value
		dIn := Decimal{Value: []int64{351234088800000999, 1234512}, Exp: -3}

		// Next, attempt to convert the Decimal to a decimal
		dOut := dIn.ToDecimal()

		// Finally, verify the data
		Expect(dOut.String()).Should(Equal("1234512351234088800000.999"))
	})
})
