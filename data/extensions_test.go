package data

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/shopspring/decimal"
)

var _ = Describe("Decimal Extensions Tests", func() {

	// Tests that, if the gob encoder does not fail, then calling NewFromDecimal will encode the decimal
	// into a GOB and write it to the data in the Decimal object which will then be returned
	It("NewFromDecimal - No failures - Encoded", func() {

		// First, create our decimal value
		dIn, err := decimal.NewFromString("52.99")
		Expect(err).ShouldNot(HaveOccurred())

		// Next, attempt to convert it to a new Decimal object; this should not fail
		dOut, err := NewFromDecimal(dIn)
		Expect(err).ShouldNot(HaveOccurred())

		// Finally, verify the data
		Expect(dOut).ShouldNot(BeNil())
		Expect(string(dOut.Data)).Should(Equal("\x13\xff\x81\x05\x01\x01\aDecimal\x01\xff\x82\x00\x00\x00\v\xff\x82\x00\a\xff\xff\xff\xfe\x02\x14\xb3"))
	})

	// Tests that, if the gob decoder fails, then calling ToDecimal will return an error
	It("ToDecimal - Decode fails - Error", func() {

		// First, create an invalid Decimal value
		dIn := Decimal{Data: []byte("derp")}

		// Next, attempt to convert the Decimal to a decimal; this should fail
		dOut, err := dIn.ToDecimal()

		// Finally, verify the error
		Expect(dOut).Should(BeNil())
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("failed to decode to Decimal, error: unexpected EOF"))
	})

	// Tests that, if the gob decoder does not fail, then calling ToDecimal will not return an error,
	// but will decode the Decimal to a decimal value
	It("ToDecimal - No failures - Decoded", func() {

		// First, create a valid Decimal value
		dIn := Decimal{Data: []byte("\x13\xff\x81\x05\x01\x01\aDecimal\x01\xff\x82\x00\x00\x00\v\xff\x82\x00\a\xff\xff\xff\xfe\x02\x14\xb3")}

		// Next, attempt to convert the Decimal to a decima; this should not fail
		dOut, err := dIn.ToDecimal()
		Expect(err).ShouldNot(HaveOccurred())

		// Finally, verify the data
		Expect(dOut.String()).Should(Equal("52.99"))
	})
})
