package data

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"gopkg.in/yaml.v3"
)

var _ = Describe("data.FillPolicy Marshal/Unmarshal Tests", func() {

	// Test that converting the data.FillPolicy enum to JSON works for all values
	DescribeTable("MarshalJSON Tests",
		func(enum FillPolicy, value string) {
			data, err := json.Marshal(enum)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(string(data)).Should(Equal(value))
		},
		Entry("FillOrKill - Works", FillPolicy_FillOrKill, "\"FOK\""),
		Entry("ImmediateOrCancel - Works", FillPolicy_ImmediateOrCancel, "\"IOC\""),
		Entry("Return - Works", FillPolicy_Return, "\"Return\""))

	// Test that converting the data.FillPolicy enum to a CSV column works for all values
	DescribeTable("MarshalCSV Tests",
		func(enum FillPolicy, value string) {
			data, err := enum.MarshalCSV()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(string(data)).Should(Equal(value))
		},
		Entry("FillOrKill - Works", FillPolicy_FillOrKill, "FOK"),
		Entry("ImmediateOrCancel - Works", FillPolicy_ImmediateOrCancel, "IOC"),
		Entry("Return - Works", FillPolicy_Return, "Return"))

	// Test that converting the data.FillPolicy enum to a YAML node works for all values
	DescribeTable("MarshalYAML - Works",
		func(enum FillPolicy, value string) {
			data, err := enum.MarshalYAML()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(data).Should(Equal(value))
		},
		Entry("FillOrKill - Works", FillPolicy_FillOrKill, "FOK"),
		Entry("ImmediateOrCancel - Works", FillPolicy_ImmediateOrCancel, "IOC"),
		Entry("Return - Works", FillPolicy_Return, "Return"))

	// Test that converting the data.FillPolicy enum to a DynamoDB AttributeVAlue works for all values
	DescribeTable("MarshalDynamoDBAttributeValue - Works",
		func(enum FillPolicy, value string) {
			data, err := attributevalue.Marshal(enum)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(data.(*types.AttributeValueMemberS).Value).Should(Equal(value))
		},
		Entry("FillOrKill - Works", FillPolicy_FillOrKill, "FOK"),
		Entry("ImmediateOrCancel - Works", FillPolicy_ImmediateOrCancel, "IOC"),
		Entry("Return - Works", FillPolicy_Return, "Return"))

	// Test that converting the data.FillPolicy enum to an SQL value for all values
	DescribeTable("Value Tests",
		func(enum FillPolicy, value string) {
			data, err := enum.Value()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(data).Should(Equal(value))
		},
		Entry("FillOrKill - Works", FillPolicy_FillOrKill, "FOK"),
		Entry("ImmediateOrCancel - Works", FillPolicy_ImmediateOrCancel, "IOC"),
		Entry("Return - Works", FillPolicy_Return, "Return"))

	// Test that attempting to deserialize a data.FillPolicy will fail and return an error if the value
	// cannot be deserialized from a JSON value to a string
	It("UnmarshalJSON fails - Error", func() {

		// Attempt to convert a non-parseable string value into a data.FillPolicy; this should return an error
		enum := new(FillPolicy)
		err := enum.UnmarshalJSON([]byte("derp"))

		// Verify the error
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of \"derp\" cannot be mapped to a data.FillPolicy"))
	})

	// Test that attempting to deserialize a data.FillPolicy will fail and return an error if the value
	// cannot be converted to either the name value or integer value of the enum option
	It("UnmarshalJSON - Value is invalid - Error", func() {

		// Attempt to convert a fake string value into a data.FillPolicy; this should return an error
		enum := new(FillPolicy)
		err := enum.UnmarshalJSON([]byte("\"derp\""))

		// Verify the error
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of \"derp\" cannot be mapped to a data.FillPolicy"))
	})

	// Test the conditions under which values should be convertible to a data.FillPolicy
	DescribeTable("UnmarshalJSON Tests",
		func(value string, shouldBe FillPolicy) {

			// Attempt to convert the string value into a data.FillPolicy; this should not fail
			var enum FillPolicy
			err := enum.UnmarshalJSON([]byte(value))

			// Verify that the deserialization was successful
			Expect(err).ShouldNot(HaveOccurred())
			Expect(enum).Should(Equal(shouldBe))
		},
		Entry("FOK - Works", "\"FOK\"", FillPolicy_FillOrKill),
		Entry("IOC - Works", "\"IOC\"", FillPolicy_ImmediateOrCancel),
		Entry("FillOrKill - Works", "\"FillOrKill\"", FillPolicy_FillOrKill),
		Entry("ImmediateOrCancel - Works", "\"ImmediateOrCancel\"", FillPolicy_ImmediateOrCancel),
		Entry("Return - Works", "\"Return\"", FillPolicy_Return),
		Entry("0 - Works", "\"0\"", FillPolicy_FillOrKill),
		Entry("1 - Works", "\"1\"", FillPolicy_ImmediateOrCancel),
		Entry("2 - Works", "\"2\"", FillPolicy_Return))

	// Test that attempting to deserialize a data.FillPolicy will fail and return an error if the value
	// cannot be converted to either the name value or integer value of the enum option
	It("UnmarshalCSV - Value is empty - Error", func() {

		// Attempt to convert a fake string value into a data.FillPolicy; this should return an error
		enum := new(FillPolicy)
		err := enum.UnmarshalCSV("")

		// Verify the error
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of \"\" cannot be mapped to a data.FillPolicy"))
	})

	// Test the conditions under which values should be convertible to a data.FillPolicy
	DescribeTable("UnmarshalCSV Tests",
		func(value string, shouldBe FillPolicy) {

			// Attempt to convert the value into a data.FillPolicy; this should not fail
			var enum FillPolicy
			err := enum.UnmarshalCSV(value)

			// Verify that the deserialization was successful
			Expect(err).ShouldNot(HaveOccurred())
			Expect(enum).Should(Equal(shouldBe))
		},
		Entry("FOK - Works", "FOK", FillPolicy_FillOrKill),
		Entry("IOC - Works", "IOC", FillPolicy_ImmediateOrCancel),
		Entry("FillOrKill - Works", "FillOrKill", FillPolicy_FillOrKill),
		Entry("ImmediateOrCancel - Works", "ImmediateOrCancel", FillPolicy_ImmediateOrCancel),
		Entry("Return - Works", "Return", FillPolicy_Return),
		Entry("0 - Works", "0", FillPolicy_FillOrKill),
		Entry("1 - Works", "1", FillPolicy_ImmediateOrCancel),
		Entry("2 - Works", "2", FillPolicy_Return))

	// Test that attempting to deserialize a data.FillPolicy will fail and return an error if the YAML
	// node does not represent a scalar value
	It("UnmarshalYAML - Node type is not scalar - Error", func() {
		enum := new(FillPolicy)
		err := enum.UnmarshalYAML(&yaml.Node{Kind: yaml.AliasNode})
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("YAML node had an invalid kind (expected scalar value)"))
	})

	// Test that attempting to deserialize a data.FillPolicy will fail and return an error if the YAML
	// node value cannot be converted to either the name value or integer value of the enum option
	It("UnmarshalYAML - Parse fails - Error", func() {
		enum := new(FillPolicy)
		err := enum.UnmarshalYAML(&yaml.Node{Kind: yaml.ScalarNode, Value: "derp"})
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of \"derp\" cannot be mapped to a data.FillPolicy"))
	})

	// Test the conditions under which YAML node values should be convertible to a data.FillPolicy
	DescribeTable("UnmarshalYAML Tests",
		func(value string, shouldBe FillPolicy) {
			var enum FillPolicy
			err := enum.UnmarshalYAML(&yaml.Node{Kind: yaml.ScalarNode, Value: value})
			Expect(err).ShouldNot(HaveOccurred())
			Expect(enum).Should(Equal(shouldBe))
		},
		Entry("FOK - Works", "FOK", FillPolicy_FillOrKill),
		Entry("IOC - Works", "IOC", FillPolicy_ImmediateOrCancel),
		Entry("FillOrKill - Works", "FillOrKill", FillPolicy_FillOrKill),
		Entry("ImmediateOrCancel - Works", "ImmediateOrCancel", FillPolicy_ImmediateOrCancel),
		Entry("Return - Works", "Return", FillPolicy_Return),
		Entry("0 - Works", "0", FillPolicy_FillOrKill),
		Entry("1 - Works", "1", FillPolicy_ImmediateOrCancel),
		Entry("2 - Works", "2", FillPolicy_Return))

	// Tests that, if the attribute type submitted to UnmarshalDynamoDBAttributeValue is not one we
	// recognize, then the function will return an error
	It("UnmarshalDynamoDBAttributeValue - AttributeValue type invalid - Error", func() {
		enum := new(FillPolicy)
		err := attributevalue.Unmarshal(&types.AttributeValueMemberBOOL{Value: true}, &enum)
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("Attribute value of *types.AttributeValueMemberBOOL could not be converted to a data.FillPolicy"))
	})

	// Tests that, if time parsing fails, then calling UnmarshalDynamoDBAttributeValue will return an error
	It("UnmarshalDynamoDBAttributeValue - Parse fails - Error", func() {
		enum := new(FillPolicy)
		err := attributevalue.Unmarshal(&types.AttributeValueMemberS{Value: "derp"}, &enum)
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of \"derp\" cannot be mapped to a data.FillPolicy"))
	})

	// Tests the conditions under which UnmarshalDynamoDBAttributeValue is called and no error is generated
	DescribeTable("UnmarshalDynamoDBAttributeValue - AttributeValue Conditions",
		func(value types.AttributeValue, expected FillPolicy) {
			var enum FillPolicy
			err := attributevalue.Unmarshal(value, &enum)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(enum).Should(Equal(expected))
		},
		Entry("Value is []bytes, FOK - Works",
			&types.AttributeValueMemberB{Value: []byte("FOK")}, FillPolicy_FillOrKill),
		Entry("Value is []bytes, IOC - Works",
			&types.AttributeValueMemberB{Value: []byte("IOC")}, FillPolicy_ImmediateOrCancel),
		Entry("Value is []bytes, FillOrKill - Works",
			&types.AttributeValueMemberB{Value: []byte("FillOrKill")}, FillPolicy_FillOrKill),
		Entry("Value is []bytes, ImmediateOrCancel - Works",
			&types.AttributeValueMemberB{Value: []byte("ImmediateOrCancel")}, FillPolicy_ImmediateOrCancel),
		Entry("Value is []bytes, Return - Works",
			&types.AttributeValueMemberB{Value: []byte("Return")}, FillPolicy_Return),
		Entry("Value is []bytes, 0 - Works",
			&types.AttributeValueMemberB{Value: []byte("0")}, FillPolicy_FillOrKill),
		Entry("Value is []bytes, 1 - Works",
			&types.AttributeValueMemberB{Value: []byte("1")}, FillPolicy_ImmediateOrCancel),
		Entry("Value is []bytes, 2 - Works",
			&types.AttributeValueMemberB{Value: []byte("2")}, FillPolicy_Return),
		Entry("Value is int, 0 - Works",
			&types.AttributeValueMemberN{Value: "0"}, FillPolicy_FillOrKill),
		Entry("Value is int, 1 - Works",
			&types.AttributeValueMemberN{Value: "1"}, FillPolicy_ImmediateOrCancel),
		Entry("Value is int, 2 - Works",
			&types.AttributeValueMemberN{Value: "2"}, FillPolicy_Return),
		Entry("Value is NULL - Works", new(types.AttributeValueMemberNULL), FillPolicy(0)),
		Entry("Value is string, FOK - Works",
			&types.AttributeValueMemberS{Value: "FOK"}, FillPolicy_FillOrKill),
		Entry("Value is string, IOC - Works",
			&types.AttributeValueMemberS{Value: "IOC"}, FillPolicy_ImmediateOrCancel),
		Entry("Value is string, FillOrKill - Works",
			&types.AttributeValueMemberS{Value: "FillOrKill"}, FillPolicy_FillOrKill),
		Entry("Value is string, ImmediateOrCancel - Works",
			&types.AttributeValueMemberS{Value: "ImmediateOrCancel"}, FillPolicy_ImmediateOrCancel),
		Entry("Value is string, Return - Works",
			&types.AttributeValueMemberS{Value: "Return"}, FillPolicy_Return),
		Entry("Value is string, 0 - Works",
			&types.AttributeValueMemberS{Value: "0"}, FillPolicy_FillOrKill),
		Entry("Value is string, 1 - Works",
			&types.AttributeValueMemberS{Value: "1"}, FillPolicy_ImmediateOrCancel),
		Entry("Value is string, 2 - Works",
			&types.AttributeValueMemberS{Value: "2"}, FillPolicy_Return))

	// Test that attempting to deserialize a data.FillPolicy will fial and return an error if the value
	// cannot be converted to either the name value or integer value of the enum option
	It("Scan - Value is nil - Error", func() {

		// Attempt to convert a fake string value into a data.FillPolicy; this should return an error
		var enum *FillPolicy
		err := enum.Scan(nil)

		// Verify the error
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of %!q(<nil>) had an invalid type of <nil>"))
		Expect(enum).Should(BeNil())
	})

	// Test the conditions under which values should be convertible to a data.FillPolicy
	DescribeTable("Scan Tests",
		func(value interface{}, shouldBe FillPolicy) {

			// Attempt to convert the value into a data.FillPolicy; this should not fail
			var enum FillPolicy
			err := enum.Scan(value)

			// Verify that the deserialization was successful
			Expect(err).ShouldNot(HaveOccurred())
			Expect(enum).Should(Equal(shouldBe))
		},
		Entry("FOK - Works", "FOK", FillPolicy_FillOrKill),
		Entry("IOC - Works", "IOC", FillPolicy_ImmediateOrCancel),
		Entry("FillOrKill - Works", "FillOrKill", FillPolicy_FillOrKill),
		Entry("ImmediateOrCancel - Works", "ImmediateOrCancel", FillPolicy_ImmediateOrCancel),
		Entry("Return - Works", "Return", FillPolicy_Return),
		Entry("0 - Works", 0, FillPolicy_FillOrKill),
		Entry("1 - Works", 1, FillPolicy_ImmediateOrCancel),
		Entry("2 - Works", 2, FillPolicy_Return))
})

var _ = Describe("data.ExpirationPolicy Marshal/Unmarshal Tests", func() {

	// Test that converting the data.ExpirationPolicy enum to JSON works for all values
	DescribeTable("MarshalJSON Tests",
		func(enum ExpirationPolicy, value string) {
			data, err := json.Marshal(enum)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(string(data)).Should(Equal(value))
		},
		Entry("UntilCancelled - Works", ExpirationPolicy_UntilCancelled, "\"GTC\""),
		Entry("Today - Works", ExpirationPolicy_Today, "\"Today\""),
		Entry("TimeSpecified - Works", ExpirationPolicy_TimeSpecified, "\"TimeSpecified\""),
		Entry("TimeSpecifiedDay - Works", ExpirationPolicy_TimeSpecifiedDay, "\"TimeSpecifiedDay\""))

	// Test that converting the data.ExpirationPolicy enum to a CSV column works for all values
	DescribeTable("MarshalCSV Tests",
		func(enum ExpirationPolicy, value string) {
			data, err := enum.MarshalCSV()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(string(data)).Should(Equal(value))
		},
		Entry("UntilCancelled - Works", ExpirationPolicy_UntilCancelled, "GTC"),
		Entry("Today - Works", ExpirationPolicy_Today, "Today"),
		Entry("TimeSpecified - Works", ExpirationPolicy_TimeSpecified, "TimeSpecified"),
		Entry("TimeSpecifiedDay - Works", ExpirationPolicy_TimeSpecifiedDay, "TimeSpecifiedDay"))

	// Test that converting the data.ExpirationPolicy enum to a YAML node works for all values
	DescribeTable("MarshalYAML - Works",
		func(enum ExpirationPolicy, value string) {
			data, err := enum.MarshalYAML()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(data).Should(Equal(value))
		},
		Entry("UntilCancelled - Works", ExpirationPolicy_UntilCancelled, "GTC"),
		Entry("Today - Works", ExpirationPolicy_Today, "Today"),
		Entry("TimeSpecified - Works", ExpirationPolicy_TimeSpecified, "TimeSpecified"),
		Entry("TimeSpecifiedDay - Works", ExpirationPolicy_TimeSpecifiedDay, "TimeSpecifiedDay"))

	// Test that converting the data.ExpirationPolicy enum to a DynamoDB AttributeVAlue works for all values
	DescribeTable("MarshalDynamoDBAttributeValue - Works",
		func(enum ExpirationPolicy, value string) {
			data, err := attributevalue.Marshal(enum)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(data.(*types.AttributeValueMemberS).Value).Should(Equal(value))
		},
		Entry("UntilCancelled - Works", ExpirationPolicy_UntilCancelled, "GTC"),
		Entry("Today - Works", ExpirationPolicy_Today, "Today"),
		Entry("TimeSpecified - Works", ExpirationPolicy_TimeSpecified, "TimeSpecified"),
		Entry("TimeSpecifiedDay - Works", ExpirationPolicy_TimeSpecifiedDay, "TimeSpecifiedDay"))

	// Test that converting the data.ExpirationPolicy enum to an SQL value for all values
	DescribeTable("Value Tests",
		func(enum ExpirationPolicy, value string) {
			data, err := enum.Value()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(data).Should(Equal(value))
		},
		Entry("UntilCancelled - Works", ExpirationPolicy_UntilCancelled, "GTC"),
		Entry("Today - Works", ExpirationPolicy_Today, "Today"),
		Entry("TimeSpecified - Works", ExpirationPolicy_TimeSpecified, "TimeSpecified"),
		Entry("TimeSpecifiedDay - Works", ExpirationPolicy_TimeSpecifiedDay, "TimeSpecifiedDay"))

	// Test that attempting to deserialize a data.ExpirationPolicy will fail and return an error if
	// the value cannot be deserialized from a JSON value to a string
	It("UnmarshalJSON fails - Error", func() {

		// Attempt to convert a non-parseable string value into a data.ExpirationPolicy; this should return an error
		enum := new(ExpirationPolicy)
		err := enum.UnmarshalJSON([]byte("derp"))

		// Verify the error
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of \"derp\" cannot be mapped to a data.ExpirationPolicy"))
	})

	// Test that attempting to deserialize a data.ExpirationPolicy will fail and return an error if
	// the value cannot be converted to either the name value or integer value of the enum option
	It("UnmarshalJSON - Value is invalid - Error", func() {

		// Attempt to convert a fake string value into a data.ExpirationPolicy; this should return an error
		enum := new(ExpirationPolicy)
		err := enum.UnmarshalJSON([]byte("\"derp\""))

		// Verify the error
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of \"derp\" cannot be mapped to a data.ExpirationPolicy"))
	})

	// Test the conditions under which values should be convertible to a data.ExpirationPolicy
	DescribeTable("UnmarshalJSON Tests",
		func(value string, shouldBe ExpirationPolicy) {

			// Attempt to convert the string value into a data.ExpirationPolicy; this should not fail
			var enum ExpirationPolicy
			err := enum.UnmarshalJSON([]byte(value))

			// Verify that the deserialization was successful
			Expect(err).ShouldNot(HaveOccurred())
			Expect(enum).Should(Equal(shouldBe))
		},
		Entry("GTC - Works", "\"GTC\"", ExpirationPolicy_UntilCancelled),
		Entry("UntilCancelled - Works", "\"UntilCancelled\"", ExpirationPolicy_UntilCancelled),
		Entry("Today - Works", "\"Today\"", ExpirationPolicy_Today),
		Entry("TimeSpecified - Works", "\"TimeSpecified\"", ExpirationPolicy_TimeSpecified),
		Entry("TimeSpecifiedDay - Works", "\"TimeSpecifiedDay\"", ExpirationPolicy_TimeSpecifiedDay),
		Entry("0 - Works", "\"0\"", ExpirationPolicy_UntilCancelled),
		Entry("1 - Works", "\"1\"", ExpirationPolicy_Today),
		Entry("2 - Works", "\"2\"", ExpirationPolicy_TimeSpecified),
		Entry("3 - Works", "\"3\"", ExpirationPolicy_TimeSpecifiedDay))

	// Test that attempting to deserialize a data.ExpirationPolicy will fail and return an error if
	// the value cannot be converted to either the name value or integer value of the enum option
	It("UnmarshalCSV - Value is empty - Error", func() {

		// Attempt to convert a fake string value into a data.ExpirationPolicy; this should return an error
		enum := new(ExpirationPolicy)
		err := enum.UnmarshalCSV("")

		// Verify the error
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of \"\" cannot be mapped to a data.ExpirationPolicy"))
	})

	// Test the conditions under which values should be convertible to a data.ExpirationPolicy
	DescribeTable("UnmarshalCSV Tests",
		func(value string, shouldBe ExpirationPolicy) {

			// Attempt to convert the value into a data.ExpirationPolicy; this should not fail
			var enum ExpirationPolicy
			err := enum.UnmarshalCSV(value)

			// Verify that the deserialization was successful
			Expect(err).ShouldNot(HaveOccurred())
			Expect(enum).Should(Equal(shouldBe))
		},
		Entry("GTC - Works", "GTC", ExpirationPolicy_UntilCancelled),
		Entry("UntilCancelled - Works", "UntilCancelled", ExpirationPolicy_UntilCancelled),
		Entry("Today - Works", "Today", ExpirationPolicy_Today),
		Entry("TimeSpecified - Works", "TimeSpecified", ExpirationPolicy_TimeSpecified),
		Entry("TimeSpecifiedDay - Works", "TimeSpecifiedDay", ExpirationPolicy_TimeSpecifiedDay),
		Entry("0 - Works", "0", ExpirationPolicy_UntilCancelled),
		Entry("1 - Works", "1", ExpirationPolicy_Today),
		Entry("2 - Works", "2", ExpirationPolicy_TimeSpecified),
		Entry("3 - Works", "3", ExpirationPolicy_TimeSpecifiedDay))

	// Test that attempting to deserialize a data.ExpirationPolicy will fail and return an error if
	// the YAML node does not represent a scalar value
	It("UnmarshalYAML - Node type is not scalar - Error", func() {
		enum := new(ExpirationPolicy)
		err := enum.UnmarshalYAML(&yaml.Node{Kind: yaml.AliasNode})
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("YAML node had an invalid kind (expected scalar value)"))
	})

	// Test that attempting to deserialize a data.ExpirationPolicy will fail and return an error if the
	// YAML node value cannot be converted to either the name value or integer value of the enum option
	It("UnmarshalYAML - Parse fails - Error", func() {
		enum := new(ExpirationPolicy)
		err := enum.UnmarshalYAML(&yaml.Node{Kind: yaml.ScalarNode, Value: "derp"})
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of \"derp\" cannot be mapped to a data.ExpirationPolicy"))
	})

	// Test the conditions under which YAML node values should be convertible to a data.ExpirationPolicy
	DescribeTable("UnmarshalYAML Tests",
		func(value string, shouldBe ExpirationPolicy) {
			var enum ExpirationPolicy
			err := enum.UnmarshalYAML(&yaml.Node{Kind: yaml.ScalarNode, Value: value})
			Expect(err).ShouldNot(HaveOccurred())
			Expect(enum).Should(Equal(shouldBe))
		},
		Entry("GTC - Works", "GTC", ExpirationPolicy_UntilCancelled),
		Entry("UntilCancelled - Works", "UntilCancelled", ExpirationPolicy_UntilCancelled),
		Entry("Today - Works", "Today", ExpirationPolicy_Today),
		Entry("TimeSpecified - Works", "TimeSpecified", ExpirationPolicy_TimeSpecified),
		Entry("TimeSpecifiedDay - Works", "TimeSpecifiedDay", ExpirationPolicy_TimeSpecifiedDay),
		Entry("0 - Works", "0", ExpirationPolicy_UntilCancelled),
		Entry("1 - Works", "1", ExpirationPolicy_Today),
		Entry("2 - Works", "2", ExpirationPolicy_TimeSpecified),
		Entry("3 - Works", "3", ExpirationPolicy_TimeSpecifiedDay))

	// Tests that, if the attribute type submitted to UnmarshalDynamoDBAttributeValue is not one we
	// recognize, then the function will return an error
	It("UnmarshalDynamoDBAttributeValue - AttributeValue type invalid - Error", func() {
		enum := new(ExpirationPolicy)
		err := attributevalue.Unmarshal(&types.AttributeValueMemberBOOL{Value: true}, &enum)
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("Attribute value of *types.AttributeValueMemberBOOL could not be converted to a data.ExpirationPolicy"))
	})

	// Tests that, if time parsing fails, then calling UnmarshalDynamoDBAttributeValue will return an error
	It("UnmarshalDynamoDBAttributeValue - Parse fails - Error", func() {
		enum := new(ExpirationPolicy)
		err := attributevalue.Unmarshal(&types.AttributeValueMemberS{Value: "derp"}, &enum)
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of \"derp\" cannot be mapped to a data.ExpirationPolicy"))
	})

	// Tests the conditions under which UnmarshalDynamoDBAttributeValue is called and no error is generated
	DescribeTable("UnmarshalDynamoDBAttributeValue - AttributeValue Conditions",
		func(value types.AttributeValue, expected ExpirationPolicy) {
			var enum ExpirationPolicy
			err := attributevalue.Unmarshal(value, &enum)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(enum).Should(Equal(expected))
		},
		Entry("Value is []bytes, GTC - Works",
			&types.AttributeValueMemberB{Value: []byte("GTC")}, ExpirationPolicy_UntilCancelled),
		Entry("Value is []bytes, UntilCancelled - Works",
			&types.AttributeValueMemberB{Value: []byte("UntilCancelled")}, ExpirationPolicy_UntilCancelled),
		Entry("Value is []bytes, Today - Works",
			&types.AttributeValueMemberB{Value: []byte("Today")}, ExpirationPolicy_Today),
		Entry("Value is []bytes, TimeSpecified - Works",
			&types.AttributeValueMemberB{Value: []byte("TimeSpecified")}, ExpirationPolicy_TimeSpecified),
		Entry("Value is []bytes, TimeSpecifiedDay - Works",
			&types.AttributeValueMemberB{Value: []byte("TimeSpecifiedDay")}, ExpirationPolicy_TimeSpecifiedDay),
		Entry("Value is []bytes, 0 - Works",
			&types.AttributeValueMemberB{Value: []byte("0")}, ExpirationPolicy_UntilCancelled),
		Entry("Value is []bytes, 1 - Works",
			&types.AttributeValueMemberB{Value: []byte("1")}, ExpirationPolicy_Today),
		Entry("Value is []bytes, 2 - Works",
			&types.AttributeValueMemberB{Value: []byte("2")}, ExpirationPolicy_TimeSpecified),
		Entry("Value is []bytes, 3 - Works",
			&types.AttributeValueMemberB{Value: []byte("3")}, ExpirationPolicy_TimeSpecifiedDay),
		Entry("Value is int, 0 - Works",
			&types.AttributeValueMemberN{Value: "0"}, ExpirationPolicy_UntilCancelled),
		Entry("Value is int, 1 - Works",
			&types.AttributeValueMemberN{Value: "1"}, ExpirationPolicy_Today),
		Entry("Value is int, 2 - Works",
			&types.AttributeValueMemberN{Value: "2"}, ExpirationPolicy_TimeSpecified),
		Entry("Value is int, 3 - Works",
			&types.AttributeValueMemberN{Value: "3"}, ExpirationPolicy_TimeSpecifiedDay),
		Entry("Value is NULL - Works", new(types.AttributeValueMemberNULL), ExpirationPolicy(0)),
		Entry("Value is string, GTC - Works",
			&types.AttributeValueMemberS{Value: "GTC"}, ExpirationPolicy_UntilCancelled),
		Entry("Value is string, UntilCancelled - Works",
			&types.AttributeValueMemberS{Value: "UntilCancelled"}, ExpirationPolicy_UntilCancelled),
		Entry("Value is string, Today - Works",
			&types.AttributeValueMemberS{Value: "Today"}, ExpirationPolicy_Today),
		Entry("Value is string, TimeSpecified - Works",
			&types.AttributeValueMemberS{Value: "TimeSpecified"}, ExpirationPolicy_TimeSpecified),
		Entry("Value is string, TimeSpecifiedDay - Works",
			&types.AttributeValueMemberS{Value: "TimeSpecifiedDay"}, ExpirationPolicy_TimeSpecifiedDay),
		Entry("Value is string, 0 - Works",
			&types.AttributeValueMemberS{Value: "0"}, ExpirationPolicy_UntilCancelled),
		Entry("Value is string, 1 - Works",
			&types.AttributeValueMemberS{Value: "1"}, ExpirationPolicy_Today),
		Entry("Value is string, 2 - Works",
			&types.AttributeValueMemberS{Value: "2"}, ExpirationPolicy_TimeSpecified),
		Entry("Value is string, 3 - Works",
			&types.AttributeValueMemberS{Value: "3"}, ExpirationPolicy_TimeSpecifiedDay))

	// Test that attempting to deserialize a data.ExpirationPolicy will fial and return an error if
	// the value cannot be converted to either the name value or integer value of the enum option
	It("Scan - Value is nil - Error", func() {

		// Attempt to convert a fake string value into a data.ExpirationPolicy; this should return an error
		var enum *ExpirationPolicy
		err := enum.Scan(nil)

		// Verify the error
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of %!q(<nil>) had an invalid type of <nil>"))
		Expect(enum).Should(BeNil())
	})

	// Test the conditions under which values should be convertible to a data.ExpirationPolicy
	DescribeTable("Scan Tests",
		func(value interface{}, shouldBe ExpirationPolicy) {

			// Attempt to convert the value into a data.ExpirationPolicy; this should not fail
			var enum ExpirationPolicy
			err := enum.Scan(value)

			// Verify that the deserialization was successful
			Expect(err).ShouldNot(HaveOccurred())
			Expect(enum).Should(Equal(shouldBe))
		},
		Entry("GTC - Works", "GTC", ExpirationPolicy_UntilCancelled),
		Entry("UntilCancelled - Works", "UntilCancelled", ExpirationPolicy_UntilCancelled),
		Entry("Today - Works", "Today", ExpirationPolicy_Today),
		Entry("TimeSpecified - Works", "TimeSpecified", ExpirationPolicy_TimeSpecified),
		Entry("TimeSpecifiedDay - Works", "TimeSpecifiedDay", ExpirationPolicy_TimeSpecifiedDay),
		Entry("0 - Works", 0, ExpirationPolicy_UntilCancelled),
		Entry("1 - Works", 1, ExpirationPolicy_Today),
		Entry("2 - Works", 2, ExpirationPolicy_TimeSpecified),
		Entry("3 - Works", 3, ExpirationPolicy_TimeSpecifiedDay))
})

var _ = Describe("data.Deal.Type Marshal/Unmarshal Tests", func() {

	// Test that converting the data.Deal.Type enum to JSON works for all values
	DescribeTable("MarshalJSON Tests",
		func(enum Deal_Type, value string) {
			data, err := json.Marshal(enum)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(string(data)).Should(Equal(value))
		},
		Entry("Buy - Works", Deal_Buy, "\"Buy\""),
		Entry("Sell - Works", Deal_Sell, "\"Sell\""),
		Entry("Balance - Works", Deal_Balance, "\"Balance\""),
		Entry("Credit - Works", Deal_Credit, "\"Credit\""),
		Entry("Charge - Works", Deal_Charge, "\"Charge\""),
		Entry("Correction - Works", Deal_Correction, "\"Correction\""),
		Entry("Bonus - Works", Deal_Bonus, "\"Bonus\""),
		Entry("Commission - Works", Deal_Commission, "\"Commission\""),
		Entry("CommissionDaily - Works", Deal_CommissionDaily, "\"CommissionDaily\""),
		Entry("CommissionMonthly - Works", Deal_CommissionMonthly, "\"CommissionMonthly\""),
		Entry("CommissionAgentDaily - Works", Deal_CommissionAgentDaily, "\"CommissionAgentDaily\""),
		Entry("CommissionAgentMonthly - Works", Deal_CommissionAgentMonthly, "\"CommissionAgentMonthly\""),
		Entry("Interest - Works", Deal_Interest, "\"Interest\""),
		Entry("BuyCancelled - Works", Deal_BuyCancelled, "\"BuyCancelled\""),
		Entry("SellCancelled - Works", Deal_SellCancelled, "\"SellCancelled\""),
		Entry("Dividend - Works", Deal_Dividend, "\"Dividend\""),
		Entry("Franked - Works", Deal_Franked, "\"Franked\""),
		Entry("Tax - Works", Deal_Tax, "\"Tax\""))

	// Test that converting the data.Deal.Type enum to a CSV column works for all values
	DescribeTable("MarshalCSV Tests",
		func(enum Deal_Type, value string) {
			data, err := enum.MarshalCSV()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(string(data)).Should(Equal(value))
		},
		Entry("Buy - Works", Deal_Buy, "Buy"),
		Entry("Sell - Works", Deal_Sell, "Sell"),
		Entry("Balance - Works", Deal_Balance, "Balance"),
		Entry("Credit - Works", Deal_Credit, "Credit"),
		Entry("Charge - Works", Deal_Charge, "Charge"),
		Entry("Correction - Works", Deal_Correction, "Correction"),
		Entry("Bonus - Works", Deal_Bonus, "Bonus"),
		Entry("Commission - Works", Deal_Commission, "Commission"),
		Entry("CommissionDaily - Works", Deal_CommissionDaily, "CommissionDaily"),
		Entry("CommissionMonthly - Works", Deal_CommissionMonthly, "CommissionMonthly"),
		Entry("CommissionAgentDaily - Works", Deal_CommissionAgentDaily, "CommissionAgentDaily"),
		Entry("CommissionAgentMonthly - Works", Deal_CommissionAgentMonthly, "CommissionAgentMonthly"),
		Entry("Interest - Works", Deal_Interest, "Interest"),
		Entry("BuyCancelled - Works", Deal_BuyCancelled, "BuyCancelled"),
		Entry("SellCancelled - Works", Deal_SellCancelled, "SellCancelled"),
		Entry("Dividend - Works", Deal_Dividend, "Dividend"),
		Entry("Franked - Works", Deal_Franked, "Franked"),
		Entry("Tax - Works", Deal_Tax, "Tax"))

	// Test that converting the data.Deal.Type enum to a YAML node works for all values
	DescribeTable("MarshalYAML - Works",
		func(enum Deal_Type, value string) {
			data, err := enum.MarshalYAML()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(data).Should(Equal(value))
		},
		Entry("Buy - Works", Deal_Buy, "Buy"),
		Entry("Sell - Works", Deal_Sell, "Sell"),
		Entry("Balance - Works", Deal_Balance, "Balance"),
		Entry("Credit - Works", Deal_Credit, "Credit"),
		Entry("Charge - Works", Deal_Charge, "Charge"),
		Entry("Correction - Works", Deal_Correction, "Correction"),
		Entry("Bonus - Works", Deal_Bonus, "Bonus"),
		Entry("Commission - Works", Deal_Commission, "Commission"),
		Entry("CommissionDaily - Works", Deal_CommissionDaily, "CommissionDaily"),
		Entry("CommissionMonthly - Works", Deal_CommissionMonthly, "CommissionMonthly"),
		Entry("CommissionAgentDaily - Works", Deal_CommissionAgentDaily, "CommissionAgentDaily"),
		Entry("CommissionAgentMonthly - Works", Deal_CommissionAgentMonthly, "CommissionAgentMonthly"),
		Entry("Interest - Works", Deal_Interest, "Interest"),
		Entry("BuyCancelled - Works", Deal_BuyCancelled, "BuyCancelled"),
		Entry("SellCancelled - Works", Deal_SellCancelled, "SellCancelled"),
		Entry("Dividend - Works", Deal_Dividend, "Dividend"),
		Entry("Franked - Works", Deal_Franked, "Franked"),
		Entry("Tax - Works", Deal_Tax, "Tax"))

	// Test that converting the data.Deal.Type enum to a DynamoDB AttributeVAlue works for all values
	DescribeTable("MarshalDynamoDBAttributeValue - Works",
		func(enum Deal_Type, value string) {
			data, err := attributevalue.Marshal(enum)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(data.(*types.AttributeValueMemberS).Value).Should(Equal(value))
		},
		Entry("Buy - Works", Deal_Buy, "Buy"),
		Entry("Sell - Works", Deal_Sell, "Sell"),
		Entry("Balance - Works", Deal_Balance, "Balance"),
		Entry("Credit - Works", Deal_Credit, "Credit"),
		Entry("Charge - Works", Deal_Charge, "Charge"),
		Entry("Correction - Works", Deal_Correction, "Correction"),
		Entry("Bonus - Works", Deal_Bonus, "Bonus"),
		Entry("Commission - Works", Deal_Commission, "Commission"),
		Entry("CommissionDaily - Works", Deal_CommissionDaily, "CommissionDaily"),
		Entry("CommissionMonthly - Works", Deal_CommissionMonthly, "CommissionMonthly"),
		Entry("CommissionAgentDaily - Works", Deal_CommissionAgentDaily, "CommissionAgentDaily"),
		Entry("CommissionAgentMonthly - Works", Deal_CommissionAgentMonthly, "CommissionAgentMonthly"),
		Entry("Interest - Works", Deal_Interest, "Interest"),
		Entry("BuyCancelled - Works", Deal_BuyCancelled, "BuyCancelled"),
		Entry("SellCancelled - Works", Deal_SellCancelled, "SellCancelled"),
		Entry("Dividend - Works", Deal_Dividend, "Dividend"),
		Entry("Franked - Works", Deal_Franked, "Franked"),
		Entry("Tax - Works", Deal_Tax, "Tax"))

	// Test that converting the data.Deal.Type enum to an SQL value for all values
	DescribeTable("Value Tests",
		func(enum Deal_Type, value string) {
			data, err := enum.Value()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(data).Should(Equal(value))
		},
		Entry("Buy - Works", Deal_Buy, "Buy"),
		Entry("Sell - Works", Deal_Sell, "Sell"),
		Entry("Balance - Works", Deal_Balance, "Balance"),
		Entry("Credit - Works", Deal_Credit, "Credit"),
		Entry("Charge - Works", Deal_Charge, "Charge"),
		Entry("Correction - Works", Deal_Correction, "Correction"),
		Entry("Bonus - Works", Deal_Bonus, "Bonus"),
		Entry("Commission - Works", Deal_Commission, "Commission"),
		Entry("CommissionDaily - Works", Deal_CommissionDaily, "CommissionDaily"),
		Entry("CommissionMonthly - Works", Deal_CommissionMonthly, "CommissionMonthly"),
		Entry("CommissionAgentDaily - Works", Deal_CommissionAgentDaily, "CommissionAgentDaily"),
		Entry("CommissionAgentMonthly - Works", Deal_CommissionAgentMonthly, "CommissionAgentMonthly"),
		Entry("Interest - Works", Deal_Interest, "Interest"),
		Entry("BuyCancelled - Works", Deal_BuyCancelled, "BuyCancelled"),
		Entry("SellCancelled - Works", Deal_SellCancelled, "SellCancelled"),
		Entry("Dividend - Works", Deal_Dividend, "Dividend"),
		Entry("Franked - Works", Deal_Franked, "Franked"),
		Entry("Tax - Works", Deal_Tax, "Tax"))

	// Test that attempting to deserialize a data.Deal.Type will fail and return an error if the value
	// cannot be deserialized from a JSON value to a string
	It("UnmarshalJSON fails - Error", func() {

		// Attempt to convert a non-parseable string value into a data.Deal.Type; this should return an error
		enum := new(Deal_Type)
		err := enum.UnmarshalJSON([]byte("derp"))

		// Verify the error
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of \"derp\" cannot be mapped to a data.Deal_Type"))
	})

	// Test that attempting to deserialize a data.Deal.Type will fail and return an error if the value
	// cannot be converted to either the name value or integer value of the enum option
	It("UnmarshalJSON - Value is invalid - Error", func() {

		// Attempt to convert a fake string value into a data.Deal.Type; this should return an error
		enum := new(Deal_Type)
		err := enum.UnmarshalJSON([]byte("\"derp\""))

		// Verify the error
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of \"derp\" cannot be mapped to a data.Deal_Type"))
	})

	// Test the conditions under which values should be convertible to a data.Deal.Type
	DescribeTable("UnmarshalJSON Tests",
		func(value string, shouldBe Deal_Type) {

			// Attempt to convert the string value into a data.Deal.Type; this should not fail
			var enum Deal_Type
			err := enum.UnmarshalJSON([]byte(value))

			// Verify that the deserialization was successful
			Expect(err).ShouldNot(HaveOccurred())
			Expect(enum).Should(Equal(shouldBe))
		},
		Entry("Buy - Works", "\"Buy\"", Deal_Buy),
		Entry("Sell - Works", "\"Sell\"", Deal_Sell),
		Entry("Balance - Works", "\"Balance\"", Deal_Balance),
		Entry("Credit - Works", "\"Credit\"", Deal_Credit),
		Entry("Charge - Works", "\"Charge\"", Deal_Charge),
		Entry("Correction - Works", "\"Correction\"", Deal_Correction),
		Entry("Bonus - Works", "\"Bonus\"", Deal_Bonus),
		Entry("Commission", "\"Commission\"", Deal_Commission),
		Entry("CommissionDaily", "\"CommissionDaily\"", Deal_CommissionDaily),
		Entry("CommissionMonthly", "\"CommissionMonthly\"", Deal_CommissionMonthly),
		Entry("CommissionAgentDaily", "\"CommissionAgentDaily\"", Deal_CommissionAgentDaily),
		Entry("CommissionAgentMonthly", "\"CommissionAgentMonthly\"", Deal_CommissionAgentMonthly),
		Entry("Interest", "\"Interest\"", Deal_Interest),
		Entry("BuyCancelled", "\"BuyCancelled\"", Deal_BuyCancelled),
		Entry("SellCancelled", "\"SellCancelled\"", Deal_SellCancelled),
		Entry("Dividend", "\"Dividend\"", Deal_Dividend),
		Entry("Franked", "\"Franked\"", Deal_Franked),
		Entry("Tax", "\"Tax\"", Deal_Tax),
		Entry("0 - Works", "\"0\"", Deal_Buy),
		Entry("1 - Works", "\"1\"", Deal_Sell),
		Entry("2 - Works", "\"2\"", Deal_Balance),
		Entry("3 - Works", "\"3\"", Deal_Credit),
		Entry("4 - Works", "\"4\"", Deal_Charge),
		Entry("5 - Works", "\"5\"", Deal_Correction),
		Entry("6 - Works", "\"6\"", Deal_Bonus),
		Entry("7 - Works", "\"7\"", Deal_Commission),
		Entry("8 - Works", "\"8\"", Deal_CommissionDaily),
		Entry("9 - Works", "\"9\"", Deal_CommissionMonthly),
		Entry("10 - Works", "\"10\"", Deal_CommissionAgentDaily),
		Entry("11 - Works", "\"11\"", Deal_CommissionAgentMonthly),
		Entry("12 - Works", "\"12\"", Deal_Interest),
		Entry("13 - Works", "\"13\"", Deal_BuyCancelled),
		Entry("14 - Works", "\"14\"", Deal_SellCancelled),
		Entry("15 - Works", "\"15\"", Deal_Dividend),
		Entry("16 - Works", "\"16\"", Deal_Franked),
		Entry("17 - Works", "\"17\"", Deal_Tax))

	// Test that attempting to deserialize a data.Deal.Type will fail and return an error if the value
	// cannot be converted to either the name value or integer value of the enum option
	It("UnmarshalCSV - Value is empty - Error", func() {

		// Attempt to convert a fake string value into a data.Deal.Type; this should return an error
		enum := new(Deal_Type)
		err := enum.UnmarshalCSV("")

		// Verify the error
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of \"\" cannot be mapped to a data.Deal_Type"))
	})

	// Test the conditions under which values should be convertible to a data.Deal.Type
	DescribeTable("UnmarshalCSV Tests",
		func(value string, shouldBe Deal_Type) {

			// Attempt to convert the value into a data.Deal.Type; this should not fail
			var enum Deal_Type
			err := enum.UnmarshalCSV(value)

			// Verify that the deserialization was successful
			Expect(err).ShouldNot(HaveOccurred())
			Expect(enum).Should(Equal(shouldBe))
		},
		Entry("Buy - Works", "Buy", Deal_Buy),
		Entry("Sell - Works", "Sell", Deal_Sell),
		Entry("Balance - Works", "Balance", Deal_Balance),
		Entry("Credit - Works", "Credit", Deal_Credit),
		Entry("Charge - Works", "Charge", Deal_Charge),
		Entry("Correction - Works", "Correction", Deal_Correction),
		Entry("Bonus - Works", "Bonus", Deal_Bonus),
		Entry("Commission", "Commission", Deal_Commission),
		Entry("CommissionDaily", "CommissionDaily", Deal_CommissionDaily),
		Entry("CommissionMonthly", "CommissionMonthly", Deal_CommissionMonthly),
		Entry("CommissionAgentDaily", "CommissionAgentDaily", Deal_CommissionAgentDaily),
		Entry("CommissionAgentMonthly", "CommissionAgentMonthly", Deal_CommissionAgentMonthly),
		Entry("Interest", "Interest", Deal_Interest),
		Entry("BuyCancelled", "BuyCancelled", Deal_BuyCancelled),
		Entry("SellCancelled", "SellCancelled", Deal_SellCancelled),
		Entry("Dividend", "Dividend", Deal_Dividend),
		Entry("Franked", "Franked", Deal_Franked),
		Entry("Tax", "Tax", Deal_Tax),
		Entry("0 - Works", "0", Deal_Buy),
		Entry("1 - Works", "1", Deal_Sell),
		Entry("2 - Works", "2", Deal_Balance),
		Entry("3 - Works", "3", Deal_Credit),
		Entry("4 - Works", "4", Deal_Charge),
		Entry("5 - Works", "5", Deal_Correction),
		Entry("6 - Works", "6", Deal_Bonus),
		Entry("7 - Works", "7", Deal_Commission),
		Entry("8 - Works", "8", Deal_CommissionDaily),
		Entry("9 - Works", "9", Deal_CommissionMonthly),
		Entry("10 - Works", "10", Deal_CommissionAgentDaily),
		Entry("11 - Works", "11", Deal_CommissionAgentMonthly),
		Entry("12 - Works", "12", Deal_Interest),
		Entry("13 - Works", "13", Deal_BuyCancelled),
		Entry("14 - Works", "14", Deal_SellCancelled),
		Entry("15 - Works", "15", Deal_Dividend),
		Entry("16 - Works", "16", Deal_Franked),
		Entry("17 - Works", "17", Deal_Tax))

	// Test that attempting to deserialize a data.Deal.Type will fail and return an error if the YAML
	// node does not represent a scalar value
	It("UnmarshalYAML - Node type is not scalar - Error", func() {
		enum := new(Deal_Type)
		err := enum.UnmarshalYAML(&yaml.Node{Kind: yaml.AliasNode})
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("YAML node had an invalid kind (expected scalar value)"))
	})

	// Test that attempting to deserialize a data.Deal.Type will fail and return an error if the YAML
	// node value cannot be converted to either the name value or integer value of the enum option
	It("UnmarshalYAML - Parse fails - Error", func() {
		enum := new(Deal_Type)
		err := enum.UnmarshalYAML(&yaml.Node{Kind: yaml.ScalarNode, Value: "derp"})
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of \"derp\" cannot be mapped to a data.Deal_Type"))
	})

	// Test the conditions under which YAML node values should be convertible to a data.Deal.Type
	DescribeTable("UnmarshalYAML Tests",
		func(value string, shouldBe Deal_Type) {
			var enum Deal_Type
			err := enum.UnmarshalYAML(&yaml.Node{Kind: yaml.ScalarNode, Value: value})
			Expect(err).ShouldNot(HaveOccurred())
			Expect(enum).Should(Equal(shouldBe))
		},
		Entry("Buy - Works", "Buy", Deal_Buy),
		Entry("Sell - Works", "Sell", Deal_Sell),
		Entry("Balance - Works", "Balance", Deal_Balance),
		Entry("Credit - Works", "Credit", Deal_Credit),
		Entry("Charge - Works", "Charge", Deal_Charge),
		Entry("Correction - Works", "Correction", Deal_Correction),
		Entry("Bonus - Works", "Bonus", Deal_Bonus),
		Entry("Commission", "Commission", Deal_Commission),
		Entry("CommissionDaily", "CommissionDaily", Deal_CommissionDaily),
		Entry("CommissionMonthly", "CommissionMonthly", Deal_CommissionMonthly),
		Entry("CommissionAgentDaily", "CommissionAgentDaily", Deal_CommissionAgentDaily),
		Entry("CommissionAgentMonthly", "CommissionAgentMonthly", Deal_CommissionAgentMonthly),
		Entry("Interest", "Interest", Deal_Interest),
		Entry("BuyCancelled", "BuyCancelled", Deal_BuyCancelled),
		Entry("SellCancelled", "SellCancelled", Deal_SellCancelled),
		Entry("Dividend", "Dividend", Deal_Dividend),
		Entry("Franked", "Franked", Deal_Franked),
		Entry("Tax", "Tax", Deal_Tax),
		Entry("0 - Works", "0", Deal_Buy),
		Entry("1 - Works", "1", Deal_Sell),
		Entry("2 - Works", "2", Deal_Balance),
		Entry("3 - Works", "3", Deal_Credit),
		Entry("4 - Works", "4", Deal_Charge),
		Entry("5 - Works", "5", Deal_Correction),
		Entry("6 - Works", "6", Deal_Bonus),
		Entry("7 - Works", "7", Deal_Commission),
		Entry("8 - Works", "8", Deal_CommissionDaily),
		Entry("9 - Works", "9", Deal_CommissionMonthly),
		Entry("10 - Works", "10", Deal_CommissionAgentDaily),
		Entry("11 - Works", "11", Deal_CommissionAgentMonthly),
		Entry("12 - Works", "12", Deal_Interest),
		Entry("13 - Works", "13", Deal_BuyCancelled),
		Entry("14 - Works", "14", Deal_SellCancelled),
		Entry("15 - Works", "15", Deal_Dividend),
		Entry("16 - Works", "16", Deal_Franked),
		Entry("17 - Works", "17", Deal_Tax))

	// Tests that, if the attribute type submitted to UnmarshalDynamoDBAttributeValue is not one we
	// recognize, then the function will return an error
	It("UnmarshalDynamoDBAttributeValue - AttributeValue type invalid - Error", func() {
		enum := new(Deal_Type)
		err := attributevalue.Unmarshal(&types.AttributeValueMemberBOOL{Value: true}, &enum)
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("Attribute value of *types.AttributeValueMemberBOOL could not be converted to a data.Deal.Type"))
	})

	// Tests that, if time parsing fails, then calling UnmarshalDynamoDBAttributeValue will return an error
	It("UnmarshalDynamoDBAttributeValue - Parse fails - Error", func() {
		enum := new(Deal_Type)
		err := attributevalue.Unmarshal(&types.AttributeValueMemberS{Value: "derp"}, &enum)
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of \"derp\" cannot be mapped to a data.Deal_Type"))
	})

	// Tests the conditions under which UnmarshalDynamoDBAttributeValue is called and no error is generated
	DescribeTable("UnmarshalDynamoDBAttributeValue - AttributeValue Conditions",
		func(value types.AttributeValue, expected Deal_Type) {
			var enum Deal_Type
			err := attributevalue.Unmarshal(value, &enum)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(enum).Should(Equal(expected))
		},
		Entry("Value is []bytes, Buy - Works",
			&types.AttributeValueMemberB{Value: []byte("Buy")}, Deal_Buy),
		Entry("Value is []bytes, Sell - Works",
			&types.AttributeValueMemberB{Value: []byte("Sell")}, Deal_Sell),
		Entry("Value is []bytes, Balance - Works",
			&types.AttributeValueMemberB{Value: []byte("Balance")}, Deal_Balance),
		Entry("Value is []bytes, Credit - Works",
			&types.AttributeValueMemberB{Value: []byte("Credit")}, Deal_Credit),
		Entry("Value is []bytes, Charge - Works",
			&types.AttributeValueMemberB{Value: []byte("Charge")}, Deal_Charge),
		Entry("Value is []bytes, Correction - Works",
			&types.AttributeValueMemberB{Value: []byte("Correction")}, Deal_Correction),
		Entry("Value is []bytes, Bonus - Works",
			&types.AttributeValueMemberB{Value: []byte("Bonus")}, Deal_Bonus),
		Entry("Value is []bytes, Commission - Works",
			&types.AttributeValueMemberB{Value: []byte("Commission")}, Deal_Commission),
		Entry("Value is []bytes, CommissionDaily - Works",
			&types.AttributeValueMemberB{Value: []byte("CommissionDaily")}, Deal_CommissionDaily),
		Entry("Value is []bytes, CommissionMonthly - Works",
			&types.AttributeValueMemberB{Value: []byte("CommissionMonthly")}, Deal_CommissionMonthly),
		Entry("Value is []bytes, CommissionAgentDaily - Works",
			&types.AttributeValueMemberB{Value: []byte("CommissionAgentDaily")}, Deal_CommissionAgentDaily),
		Entry("Value is []bytes, CommissionAgentMonthly - Works",
			&types.AttributeValueMemberB{Value: []byte("CommissionAgentMonthly")}, Deal_CommissionAgentMonthly),
		Entry("Value is []bytes, Interest - Works",
			&types.AttributeValueMemberB{Value: []byte("Interest")}, Deal_Interest),
		Entry("Value is []bytes, BuyCancelled - Works",
			&types.AttributeValueMemberB{Value: []byte("BuyCancelled")}, Deal_BuyCancelled),
		Entry("Value is []bytes, SellCancelled - Works",
			&types.AttributeValueMemberB{Value: []byte("SellCancelled")}, Deal_SellCancelled),
		Entry("Value is []bytes, Dividend - Works",
			&types.AttributeValueMemberB{Value: []byte("Dividend")}, Deal_Dividend),
		Entry("Value is []bytes, Franked - Works",
			&types.AttributeValueMemberB{Value: []byte("Franked")}, Deal_Franked),
		Entry("Value is []bytes, Tax - Works",
			&types.AttributeValueMemberB{Value: []byte("Tax")}, Deal_Tax),
		Entry("Value is []bytes, 0 - Works",
			&types.AttributeValueMemberB{Value: []byte("0")}, Deal_Buy),
		Entry("Value is []bytes, 1 - Works",
			&types.AttributeValueMemberB{Value: []byte("1")}, Deal_Sell),
		Entry("Value is []bytes, 2 - Works",
			&types.AttributeValueMemberB{Value: []byte("2")}, Deal_Balance),
		Entry("Value is []bytes, 3 - Works",
			&types.AttributeValueMemberB{Value: []byte("3")}, Deal_Credit),
		Entry("Value is []bytes, 4 - Works",
			&types.AttributeValueMemberB{Value: []byte("4")}, Deal_Charge),
		Entry("Value is []bytes, 5 - Works",
			&types.AttributeValueMemberB{Value: []byte("5")}, Deal_Correction),
		Entry("Value is []bytes, 6 - Works",
			&types.AttributeValueMemberB{Value: []byte("6")}, Deal_Bonus),
		Entry("Value is []bytes, 7 - Works",
			&types.AttributeValueMemberB{Value: []byte("7")}, Deal_Commission),
		Entry("Value is []bytes, 8 - Works",
			&types.AttributeValueMemberB{Value: []byte("8")}, Deal_CommissionDaily),
		Entry("Value is []bytes, 9 - Works",
			&types.AttributeValueMemberB{Value: []byte("9")}, Deal_CommissionMonthly),
		Entry("Value is []bytes, 10 - Works",
			&types.AttributeValueMemberB{Value: []byte("10")}, Deal_CommissionAgentDaily),
		Entry("Value is []bytes, 11 - Works",
			&types.AttributeValueMemberB{Value: []byte("11")}, Deal_CommissionAgentMonthly),
		Entry("Value is []bytes, 12 - Works",
			&types.AttributeValueMemberB{Value: []byte("12")}, Deal_Interest),
		Entry("Value is []bytes, 13 - Works",
			&types.AttributeValueMemberB{Value: []byte("13")}, Deal_BuyCancelled),
		Entry("Value is []bytes, 14 - Works",
			&types.AttributeValueMemberB{Value: []byte("14")}, Deal_SellCancelled),
		Entry("Value is []bytes, 15 - Works",
			&types.AttributeValueMemberB{Value: []byte("15")}, Deal_Dividend),
		Entry("Value is []bytes, 16 - Works",
			&types.AttributeValueMemberB{Value: []byte("16")}, Deal_Franked),
		Entry("Value is []bytes, 17 - Works",
			&types.AttributeValueMemberB{Value: []byte("17")}, Deal_Tax),
		Entry("Value is int, 0 - Works",
			&types.AttributeValueMemberN{Value: "0"}, Deal_Buy),
		Entry("Value is int, 1 - Works",
			&types.AttributeValueMemberN{Value: "1"}, Deal_Sell),
		Entry("Value is int, 2 - Works",
			&types.AttributeValueMemberN{Value: "2"}, Deal_Balance),
		Entry("Value is int, 3 - Works",
			&types.AttributeValueMemberN{Value: "3"}, Deal_Credit),
		Entry("Value is int, 4 - Works",
			&types.AttributeValueMemberN{Value: "4"}, Deal_Charge),
		Entry("Value is int, 5 - Works",
			&types.AttributeValueMemberN{Value: "5"}, Deal_Correction),
		Entry("Value is int, 6 - Works",
			&types.AttributeValueMemberN{Value: "6"}, Deal_Bonus),
		Entry("Value is int, 7 - Works",
			&types.AttributeValueMemberN{Value: "7"}, Deal_Commission),
		Entry("Value is int, 8 - Works",
			&types.AttributeValueMemberN{Value: "8"}, Deal_CommissionDaily),
		Entry("Value is int, 9 - Works",
			&types.AttributeValueMemberN{Value: "9"}, Deal_CommissionMonthly),
		Entry("Value is int, 10 - Works",
			&types.AttributeValueMemberN{Value: "10"}, Deal_CommissionAgentDaily),
		Entry("Value is int, 11 - Works",
			&types.AttributeValueMemberN{Value: "11"}, Deal_CommissionAgentMonthly),
		Entry("Value is int, 12 - Works",
			&types.AttributeValueMemberN{Value: "12"}, Deal_Interest),
		Entry("Value is int, 13 - Works",
			&types.AttributeValueMemberN{Value: "13"}, Deal_BuyCancelled),
		Entry("Value is int, 14 - Works",
			&types.AttributeValueMemberN{Value: "14"}, Deal_SellCancelled),
		Entry("Value is int, 15 - Works",
			&types.AttributeValueMemberN{Value: "15"}, Deal_Dividend),
		Entry("Value is int, 16 - Works",
			&types.AttributeValueMemberN{Value: "16"}, Deal_Franked),
		Entry("Value is int, 17 - Works",
			&types.AttributeValueMemberN{Value: "17"}, Deal_Tax),
		Entry("Value is NULL - Works", new(types.AttributeValueMemberNULL), Deal_Type(0)),
		Entry("Value is string, Buy - Works",
			&types.AttributeValueMemberN{Value: "Buy"}, Deal_Buy),
		Entry("Value is string, Sell - Works",
			&types.AttributeValueMemberN{Value: "Sell"}, Deal_Sell),
		Entry("Value is string, Balance - Works",
			&types.AttributeValueMemberN{Value: "Balance"}, Deal_Balance),
		Entry("Value is string, Credit - Works",
			&types.AttributeValueMemberN{Value: "Credit"}, Deal_Credit),
		Entry("Value is string, Charge - Works",
			&types.AttributeValueMemberN{Value: "Charge"}, Deal_Charge),
		Entry("Value is string, Correction - Works",
			&types.AttributeValueMemberN{Value: "Correction"}, Deal_Correction),
		Entry("Value is string, Bonus - Works",
			&types.AttributeValueMemberN{Value: "Bonus"}, Deal_Bonus),
		Entry("Value is string, Commission - Works",
			&types.AttributeValueMemberN{Value: "Commission"}, Deal_Commission),
		Entry("Value is string, CommissionDaily - Works",
			&types.AttributeValueMemberN{Value: "CommissionDaily"}, Deal_CommissionDaily),
		Entry("Value is string, CommissionMonthly - Works",
			&types.AttributeValueMemberN{Value: "CommissionMonthly"}, Deal_CommissionMonthly),
		Entry("Value is string, CommissionAgentDaily - Works",
			&types.AttributeValueMemberN{Value: "CommissionAgentDaily"}, Deal_CommissionAgentDaily),
		Entry("Value is string, CommissionAgentMonthly - Works",
			&types.AttributeValueMemberN{Value: "CommissionAgentMonthly"}, Deal_CommissionAgentMonthly),
		Entry("Value is string, Interest - Works",
			&types.AttributeValueMemberN{Value: "Interest"}, Deal_Interest),
		Entry("Value is string, BuyCancelled - Works",
			&types.AttributeValueMemberN{Value: "BuyCancelled"}, Deal_BuyCancelled),
		Entry("Value is string, SellCancelled - Works",
			&types.AttributeValueMemberN{Value: "SellCancelled"}, Deal_SellCancelled),
		Entry("Value is string, Dividend - Works",
			&types.AttributeValueMemberN{Value: "Dividend"}, Deal_Dividend),
		Entry("Value is string, Franked - Works",
			&types.AttributeValueMemberN{Value: "Franked"}, Deal_Franked),
		Entry("Value is string, Tax - Works",
			&types.AttributeValueMemberN{Value: "Tax"}, Deal_Tax),
		Entry("Value is string, 0 - Works",
			&types.AttributeValueMemberN{Value: "0"}, Deal_Buy),
		Entry("Value is string, 1 - Works",
			&types.AttributeValueMemberN{Value: "1"}, Deal_Sell),
		Entry("Value is string, 2 - Works",
			&types.AttributeValueMemberN{Value: "2"}, Deal_Balance),
		Entry("Value is string, 3 - Works",
			&types.AttributeValueMemberN{Value: "3"}, Deal_Credit),
		Entry("Value is string, 4 - Works",
			&types.AttributeValueMemberN{Value: "4"}, Deal_Charge),
		Entry("Value is string, 5 - Works",
			&types.AttributeValueMemberN{Value: "5"}, Deal_Correction),
		Entry("Value is string, 6 - Works",
			&types.AttributeValueMemberN{Value: "6"}, Deal_Bonus),
		Entry("Value is string, 7 - Works",
			&types.AttributeValueMemberN{Value: "7"}, Deal_Commission),
		Entry("Value is string, 8 - Works",
			&types.AttributeValueMemberN{Value: "8"}, Deal_CommissionDaily),
		Entry("Value is string, 9 - Works",
			&types.AttributeValueMemberN{Value: "9"}, Deal_CommissionMonthly),
		Entry("Value is string, 10 - Works",
			&types.AttributeValueMemberN{Value: "10"}, Deal_CommissionAgentDaily),
		Entry("Value is string, 11 - Works",
			&types.AttributeValueMemberN{Value: "11"}, Deal_CommissionAgentMonthly),
		Entry("Value is string, 12 - Works",
			&types.AttributeValueMemberN{Value: "12"}, Deal_Interest),
		Entry("Value is string, 13 - Works",
			&types.AttributeValueMemberN{Value: "13"}, Deal_BuyCancelled),
		Entry("Value is string, 14 - Works",
			&types.AttributeValueMemberN{Value: "14"}, Deal_SellCancelled),
		Entry("Value is string, 15 - Works",
			&types.AttributeValueMemberN{Value: "15"}, Deal_Dividend),
		Entry("Value is string, 16 - Works",
			&types.AttributeValueMemberN{Value: "16"}, Deal_Franked),
		Entry("Value is string, 17 - Works",
			&types.AttributeValueMemberN{Value: "17"}, Deal_Tax))

	// Test that attempting to deserialize a data.Deal.Type will fial and return an error if the value
	// cannot be converted to either the name value or integer value of the enum option
	It("Scan - Value is nil - Error", func() {

		// Attempt to convert a fake string value into a data.Deal.Type; this should return an error
		var enum *Deal_Type
		err := enum.Scan(nil)

		// Verify the error
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of %!q(<nil>) had an invalid type of <nil>"))
		Expect(enum).Should(BeNil())
	})

	// Test the conditions under which values should be convertible to a data.Deal.Type
	DescribeTable("Scan Tests",
		func(value interface{}, shouldBe Deal_Type) {

			// Attempt to convert the value into a data.Deal.Type; this should not fail
			var enum Deal_Type
			err := enum.Scan(value)

			// Verify that the deserialization was successful
			Expect(err).ShouldNot(HaveOccurred())
			Expect(enum).Should(Equal(shouldBe))
		},
		Entry("Buy - Works", "Buy", Deal_Buy),
		Entry("Sell - Works", "Sell", Deal_Sell),
		Entry("Balance - Works", "Balance", Deal_Balance),
		Entry("Credit - Works", "Credit", Deal_Credit),
		Entry("Charge - Works", "Charge", Deal_Charge),
		Entry("Correction - Works", "Correction", Deal_Correction),
		Entry("Bonus - Works", "Bonus", Deal_Bonus),
		Entry("Commission", "Commission", Deal_Commission),
		Entry("CommissionDaily", "CommissionDaily", Deal_CommissionDaily),
		Entry("CommissionMonthly", "CommissionMonthly", Deal_CommissionMonthly),
		Entry("CommissionAgentDaily", "CommissionAgentDaily", Deal_CommissionAgentDaily),
		Entry("CommissionAgentMonthly", "CommissionAgentMonthly", Deal_CommissionAgentMonthly),
		Entry("Interest", "Interest", Deal_Interest),
		Entry("BuyCancelled", "BuyCancelled", Deal_BuyCancelled),
		Entry("SellCancelled", "SellCancelled", Deal_SellCancelled),
		Entry("Dividend", "Dividend", Deal_Dividend),
		Entry("Franked", "Franked", Deal_Franked),
		Entry("Tax", "Tax", Deal_Tax),
		Entry("0 - Works", 0, Deal_Buy),
		Entry("1 - Works", 1, Deal_Sell),
		Entry("2 - Works", 2, Deal_Balance),
		Entry("3 - Works", 3, Deal_Credit),
		Entry("4 - Works", 4, Deal_Charge),
		Entry("5 - Works", 5, Deal_Correction),
		Entry("6 - Works", 6, Deal_Bonus),
		Entry("7 - Works", 7, Deal_Commission),
		Entry("8 - Works", 8, Deal_CommissionDaily),
		Entry("9 - Works", 9, Deal_CommissionMonthly),
		Entry("10 - Works", 10, Deal_CommissionAgentDaily),
		Entry("11 - Works", 11, Deal_CommissionAgentMonthly),
		Entry("12 - Works", 12, Deal_Interest),
		Entry("13 - Works", 13, Deal_BuyCancelled),
		Entry("14 - Works", 14, Deal_SellCancelled),
		Entry("15 - Works", 15, Deal_Dividend),
		Entry("16 - Works", 16, Deal_Franked),
		Entry("17 - Works", 17, Deal_Tax))
})

var _ = Describe("data.Deal.Entry Marshal/Unmarshal Tests", func() {

	// Test that converting the data.Deal.Entry enum to JSON works for all values
	DescribeTable("MarshalJSON Tests",
		func(enum Deal_Entry, value string) {
			data, err := json.Marshal(enum)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(string(data)).Should(Equal(value))
		},
		Entry("In - Works", Deal_In, "\"In\""),
		Entry("Out - Works", Deal_Out, "\"Out\""),
		Entry("Reverse - Works", Deal_Reverse, "\"Reverse\""),
		Entry("OutBy - Works", Deal_OutBy, "\"OutBy\""))

	// Test that converting the data.Deal.Entry enum to a CSV column works for all values
	DescribeTable("MarshalCSV Tests",
		func(enum Deal_Entry, value string) {
			data, err := enum.MarshalCSV()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(string(data)).Should(Equal(value))
		},
		Entry("In - Works", Deal_In, "In"),
		Entry("Out - Works", Deal_Out, "Out"),
		Entry("Reverse - Works", Deal_Reverse, "Reverse"),
		Entry("OutBy - Works", Deal_OutBy, "OutBy"))

	// Test that converting the data.Deal.Entry enum to a YAML node works for all values
	DescribeTable("MarshalYAML - Works",
		func(enum Deal_Entry, value string) {
			data, err := enum.MarshalYAML()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(data).Should(Equal(value))
		},
		Entry("In - Works", Deal_In, "In"),
		Entry("Out - Works", Deal_Out, "Out"),
		Entry("Reverse - Works", Deal_Reverse, "Reverse"),
		Entry("OutBy - Works", Deal_OutBy, "OutBy"))

	// Test that converting the data.Deal.Entry enum to a DynamoDB AttributeVAlue works for all values
	DescribeTable("MarshalDynamoDBAttributeValue - Works",
		func(enum Deal_Entry, value string) {
			data, err := attributevalue.Marshal(enum)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(data.(*types.AttributeValueMemberS).Value).Should(Equal(value))
		},
		Entry("In - Works", Deal_In, "In"),
		Entry("Out - Works", Deal_Out, "Out"),
		Entry("Reverse - Works", Deal_Reverse, "Reverse"),
		Entry("OutBy - Works", Deal_OutBy, "OutBy"))

	// Test that converting the data.Deal.Entry enum to an SQL value for all values
	DescribeTable("Value Tests",
		func(enum Deal_Entry, value string) {
			data, err := enum.Value()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(data).Should(Equal(value))
		},
		Entry("In - Works", Deal_In, "In"),
		Entry("Out - Works", Deal_Out, "Out"),
		Entry("Reverse - Works", Deal_Reverse, "Reverse"),
		Entry("OutBy - Works", Deal_OutBy, "OutBy"))

	// Test that attempting to deserialize a data.Deal.Entry will fail and return an error if the value
	// cannot be deserialized from a JSON value to a string
	It("UnmarshalJSON fails - Error", func() {

		// Attempt to convert a non-parseable string value into a data.Deal.Entry; this should return an error
		enum := new(Deal_Entry)
		err := enum.UnmarshalJSON([]byte("derp"))

		// Verify the error
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of \"derp\" cannot be mapped to a data.Deal_Entry"))
	})

	// Test that attempting to deserialize a data.Deal.Entry will fail and return an error if the value
	// cannot be converted to either the name value or integer value of the enum option
	It("UnmarshalJSON - Value is invalid - Error", func() {

		// Attempt to convert a fake string value into a data.Deal.Entry; this should return an error
		enum := new(Deal_Entry)
		err := enum.UnmarshalJSON([]byte("\"derp\""))

		// Verify the error
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of \"derp\" cannot be mapped to a data.Deal_Entry"))
	})

	// Test the conditions under which values should be convertible to a data.Deal.Entry
	DescribeTable("UnmarshalJSON Tests",
		func(value string, shouldBe Deal_Entry) {

			// Attempt to convert the string value into a data.Deal.Entry; this should not fail
			var enum Deal_Entry
			err := enum.UnmarshalJSON([]byte(value))

			// Verify that the deserialization was successful
			Expect(err).ShouldNot(HaveOccurred())
			Expect(enum).Should(Equal(shouldBe))
		},
		Entry("InOut - Works", "\"InOut\"", Deal_Reverse),
		Entry("In - Works", "\"In\"", Deal_In),
		Entry("Out - Works", "\"Out\"", Deal_Out),
		Entry("Reverse - Works", "\"Reverse\"", Deal_Reverse),
		Entry("OutBy - Works", "\"OutBy\"", Deal_OutBy),
		Entry("0 - Works", "\"0\"", Deal_In),
		Entry("1 - Works", "\"1\"", Deal_Out),
		Entry("2 - Works", "\"2\"", Deal_Reverse),
		Entry("3 - Works", "\"3\"", Deal_OutBy))

	// Test that attempting to deserialize a data.Deal.Entry will fail and return an error if the value
	// cannot be converted to either the name value or integer value of the enum option
	It("UnmarshalCSV - Value is empty - Error", func() {

		// Attempt to convert a fake string value into a data.Deal.Entry; this should return an error
		enum := new(Deal_Entry)
		err := enum.UnmarshalCSV("")

		// Verify the error
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of \"\" cannot be mapped to a data.Deal_Entry"))
	})

	// Test the conditions under which values should be convertible to a data.Deal.Entry
	DescribeTable("UnmarshalCSV Tests",
		func(value string, shouldBe Deal_Entry) {

			// Attempt to convert the value into a data.Deal.Entry; this should not fail
			var enum Deal_Entry
			err := enum.UnmarshalCSV(value)

			// Verify that the deserialization was successful
			Expect(err).ShouldNot(HaveOccurred())
			Expect(enum).Should(Equal(shouldBe))
		},
		Entry("InOut - Works", "InOut", Deal_Reverse),
		Entry("In - Works", "In", Deal_In),
		Entry("Out - Works", "Out", Deal_Out),
		Entry("Reverse - Works", "Reverse", Deal_Reverse),
		Entry("OutBy - Works", "OutBy", Deal_OutBy),
		Entry("0 - Works", "0", Deal_In),
		Entry("1 - Works", "1", Deal_Out),
		Entry("2 - Works", "2", Deal_Reverse),
		Entry("3 - Works", "3", Deal_OutBy))

	// Test that attempting to deserialize a data.Deal.Entry will fail and return an error if the YAML
	// node does not represent a scalar value
	It("UnmarshalYAML - Node type is not scalar - Error", func() {
		enum := new(Deal_Entry)
		err := enum.UnmarshalYAML(&yaml.Node{Kind: yaml.AliasNode})
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("YAML node had an invalid kind (expected scalar value)"))
	})

	// Test that attempting to deserialize a data.Deal.Entry will fail and return an error if the YAML
	// node value cannot be converted to either the name value or integer value of the enum option
	It("UnmarshalYAML - Parse fails - Error", func() {
		enum := new(Deal_Entry)
		err := enum.UnmarshalYAML(&yaml.Node{Kind: yaml.ScalarNode, Value: "derp"})
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of \"derp\" cannot be mapped to a data.Deal_Entry"))
	})

	// Test the conditions under which YAML node values should be convertible to a data.Deal.Entry
	DescribeTable("UnmarshalYAML Tests",
		func(value string, shouldBe Deal_Entry) {
			var enum Deal_Entry
			err := enum.UnmarshalYAML(&yaml.Node{Kind: yaml.ScalarNode, Value: value})
			Expect(err).ShouldNot(HaveOccurred())
			Expect(enum).Should(Equal(shouldBe))
		},
		Entry("InOut - Works", "InOut", Deal_Reverse),
		Entry("In - Works", "In", Deal_In),
		Entry("Out - Works", "Out", Deal_Out),
		Entry("Reverse - Works", "Reverse", Deal_Reverse),
		Entry("OutBy - Works", "OutBy", Deal_OutBy),
		Entry("0 - Works", "0", Deal_In),
		Entry("1 - Works", "1", Deal_Out),
		Entry("2 - Works", "2", Deal_Reverse),
		Entry("3 - Works", "3", Deal_OutBy))

	// Tests that, if the attribute type submitted to UnmarshalDynamoDBAttributeValue is not one we
	// recognize, then the function will return an error
	It("UnmarshalDynamoDBAttributeValue - AttributeValue type invalid - Error", func() {
		enum := new(Deal_Entry)
		err := attributevalue.Unmarshal(&types.AttributeValueMemberBOOL{Value: true}, &enum)
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("Attribute value of *types.AttributeValueMemberBOOL could not be converted to a data.Deal.Entry"))
	})

	// Tests that, if time parsing fails, then calling UnmarshalDynamoDBAttributeValue will return an error
	It("UnmarshalDynamoDBAttributeValue - Parse fails - Error", func() {
		enum := new(Deal_Entry)
		err := attributevalue.Unmarshal(&types.AttributeValueMemberS{Value: "derp"}, &enum)
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of \"derp\" cannot be mapped to a data.Deal_Entry"))
	})

	// Tests the conditions under which UnmarshalDynamoDBAttributeValue is called and no error is generated
	DescribeTable("UnmarshalDynamoDBAttributeValue - AttributeValue Conditions",
		func(value types.AttributeValue, expected Deal_Entry) {
			var enum Deal_Entry
			err := attributevalue.Unmarshal(value, &enum)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(enum).Should(Equal(expected))
		},
		Entry("Value is []bytes, InOut - Works",
			&types.AttributeValueMemberB{Value: []byte("InOut")}, Deal_Reverse),
		Entry("Value is []bytes, In - Works",
			&types.AttributeValueMemberB{Value: []byte("In")}, Deal_In),
		Entry("Value is []bytes, Out - Works",
			&types.AttributeValueMemberB{Value: []byte("Out")}, Deal_Out),
		Entry("Value is []bytes, Reverse - Works",
			&types.AttributeValueMemberB{Value: []byte("Reverse")}, Deal_Reverse),
		Entry("Value is []bytes, OutBy - Works",
			&types.AttributeValueMemberB{Value: []byte("OutBy")}, Deal_OutBy),
		Entry("Value is []bytes, 0 - Works",
			&types.AttributeValueMemberB{Value: []byte("0")}, Deal_In),
		Entry("Value is []bytes, 1 - Works",
			&types.AttributeValueMemberB{Value: []byte("1")}, Deal_Out),
		Entry("Value is []bytes, 2 - Works",
			&types.AttributeValueMemberB{Value: []byte("2")}, Deal_Reverse),
		Entry("Value is []bytes, 3 - Works",
			&types.AttributeValueMemberB{Value: []byte("3")}, Deal_OutBy),
		Entry("Value is int, 0 - Works",
			&types.AttributeValueMemberN{Value: "0"}, Deal_In),
		Entry("Value is int, 1 - Works",
			&types.AttributeValueMemberN{Value: "1"}, Deal_Out),
		Entry("Value is int, 2 - Works",
			&types.AttributeValueMemberN{Value: "2"}, Deal_Reverse),
		Entry("Value is int, 3 - Works",
			&types.AttributeValueMemberN{Value: "3"}, Deal_OutBy),
		Entry("Value is NULL - Works", new(types.AttributeValueMemberNULL), Deal_Entry(0)),
		Entry("Value is string, InOut - Works",
			&types.AttributeValueMemberS{Value: "InOut"}, Deal_Reverse),
		Entry("Value is string, In - Works",
			&types.AttributeValueMemberN{Value: "In"}, Deal_In),
		Entry("Value is string, Out - Works",
			&types.AttributeValueMemberN{Value: "Out"}, Deal_Out),
		Entry("Value is string, Reverse - Works",
			&types.AttributeValueMemberN{Value: "Reverse"}, Deal_Reverse),
		Entry("Value is string, OutBy - Works",
			&types.AttributeValueMemberN{Value: "OutBy"}, Deal_OutBy),
		Entry("Value is string, 0 - Works",
			&types.AttributeValueMemberN{Value: "0"}, Deal_In),
		Entry("Value is string, 1 - Works",
			&types.AttributeValueMemberN{Value: "1"}, Deal_Out),
		Entry("Value is string, 2 - Works",
			&types.AttributeValueMemberN{Value: "2"}, Deal_Reverse),
		Entry("Value is string, 3 - Works",
			&types.AttributeValueMemberN{Value: "3"}, Deal_OutBy))

	// Test that attempting to deserialize a data.Deal.Entry will fial and return an error if the value
	// cannot be converted to either the name value or integer value of the enum option
	It("Scan - Value is nil - Error", func() {

		// Attempt to convert a fake string value into a data.Deal.Entry; this should return an error
		var enum *Deal_Entry
		err := enum.Scan(nil)

		// Verify the error
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of %!q(<nil>) had an invalid type of <nil>"))
		Expect(enum).Should(BeNil())
	})

	// Test the conditions under which values should be convertible to a data.Deal.Entry
	DescribeTable("Scan Tests",
		func(value interface{}, shouldBe Deal_Entry) {

			// Attempt to convert the value into a data.Deal.Entry; this should not fail
			var enum Deal_Entry
			err := enum.Scan(value)

			// Verify that the deserialization was successful
			Expect(err).ShouldNot(HaveOccurred())
			Expect(enum).Should(Equal(shouldBe))
		},
		Entry("InOut - Works", "InOut", Deal_Reverse),
		Entry("In - Works", "In", Deal_In),
		Entry("Out - Works", "Out", Deal_Out),
		Entry("Reverse - Works", "Reverse", Deal_Reverse),
		Entry("OutBy - Works", "OutBy", Deal_OutBy),
		Entry("0 - Works", 0, Deal_In),
		Entry("1 - Works", 1, Deal_Out),
		Entry("2 - Works", 2, Deal_Reverse),
		Entry("3 - Works", 3, Deal_OutBy))
})

var _ = Describe("data.Deal.Reason Marshal/Unmarshal Tests", func() {

	// Test that converting the data.Deal.Reason enum to JSON works for all values
	DescribeTable("MarshalJSON Tests",
		func(enum Deal_Reason, value string) {
			data, err := json.Marshal(enum)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(string(data)).Should(Equal(value))
		},
		Entry("Client - Works", Deal_Client, "\"Client\""),
		Entry("Mobile - Works", Deal_Mobile, "\"Mobile\""),
		Entry("Web - Works", Deal_Web, "\"Web\""),
		Entry("Strategy - Works", Deal_Strategy, "\"Strategy\""),
		Entry("StopLoss - Works", Deal_StopLoss, "\"SL\""),
		Entry("TakeProfit - Works", Deal_TakeProfit, "\"TP\""),
		Entry("StopOut - Works", Deal_StopOut, "\"SO\""),
		Entry("Rollover - Works", Deal_Rollover, "\"Rollover\""),
		Entry("Margin - Works", Deal_Margin, "\"Margin\""),
		Entry("Split - Works", Deal_Split, "\"Split\""))

	// Test that converting the data.Deal.Reason enum to a CSV column works for all values
	DescribeTable("MarshalCSV Tests",
		func(enum Deal_Reason, value string) {
			data, err := enum.MarshalCSV()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(string(data)).Should(Equal(value))
		},
		Entry("Client - Works", Deal_Client, "Client"),
		Entry("Mobile - Works", Deal_Mobile, "Mobile"),
		Entry("Web - Works", Deal_Web, "Web"),
		Entry("Strategy - Works", Deal_Strategy, "Strategy"),
		Entry("StopLoss - Works", Deal_StopLoss, "SL"),
		Entry("TakeProfit - Works", Deal_TakeProfit, "TP"),
		Entry("StopOut - Works", Deal_StopOut, "SO"),
		Entry("Rollover - Works", Deal_Rollover, "Rollover"),
		Entry("Margin - Works", Deal_Margin, "Margin"),
		Entry("Split - Works", Deal_Split, "Split"))

	// Test that converting the data.Deal.Reason enum to a YAML node works for all values
	DescribeTable("MarshalYAML - Works",
		func(enum Deal_Reason, value string) {
			data, err := enum.MarshalYAML()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(data).Should(Equal(value))
		},
		Entry("Client - Works", Deal_Client, "Client"),
		Entry("Mobile - Works", Deal_Mobile, "Mobile"),
		Entry("Web - Works", Deal_Web, "Web"),
		Entry("Strategy - Works", Deal_Strategy, "Strategy"),
		Entry("StopLoss - Works", Deal_StopLoss, "SL"),
		Entry("TakeProfit - Works", Deal_TakeProfit, "TP"),
		Entry("StopOut - Works", Deal_StopOut, "SO"),
		Entry("Rollover - Works", Deal_Rollover, "Rollover"),
		Entry("Margin - Works", Deal_Margin, "Margin"),
		Entry("Split - Works", Deal_Split, "Split"))

	// Test that converting the data.Deal.Reason enum to a DynamoDB AttributeVAlue works for all values
	DescribeTable("MarshalDynamoDBAttributeValue - Works",
		func(enum Deal_Reason, value string) {
			data, err := attributevalue.Marshal(enum)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(data.(*types.AttributeValueMemberS).Value).Should(Equal(value))
		},
		Entry("Client - Works", Deal_Client, "Client"),
		Entry("Mobile - Works", Deal_Mobile, "Mobile"),
		Entry("Web - Works", Deal_Web, "Web"),
		Entry("Strategy - Works", Deal_Strategy, "Strategy"),
		Entry("StopLoss - Works", Deal_StopLoss, "SL"),
		Entry("TakeProfit - Works", Deal_TakeProfit, "TP"),
		Entry("StopOut - Works", Deal_StopOut, "SO"),
		Entry("Rollover - Works", Deal_Rollover, "Rollover"),
		Entry("Margin - Works", Deal_Margin, "Margin"),
		Entry("Split - Works", Deal_Split, "Split"))

	// Test that converting the data.Deal.Reason enum to an SQL value for all values
	DescribeTable("Value Tests",
		func(enum Deal_Reason, value string) {
			data, err := enum.Value()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(data).Should(Equal(value))
		},
		Entry("Client - Works", Deal_Client, "Client"),
		Entry("Mobile - Works", Deal_Mobile, "Mobile"),
		Entry("Web - Works", Deal_Web, "Web"),
		Entry("Strategy - Works", Deal_Strategy, "Strategy"),
		Entry("StopLoss - Works", Deal_StopLoss, "SL"),
		Entry("TakeProfit - Works", Deal_TakeProfit, "TP"),
		Entry("StopOut - Works", Deal_StopOut, "SO"),
		Entry("Rollover - Works", Deal_Rollover, "Rollover"),
		Entry("Margin - Works", Deal_Margin, "Margin"),
		Entry("Split - Works", Deal_Split, "Split"))

	// Test that attempting to deserialize a data.Deal.Reason will fail and return an error if the value
	// cannot be deserialized from a JSON value to a string
	It("UnmarshalJSON fails - Error", func() {

		// Attempt to convert a non-parseable string value into a data.Deal.Reason; this should return an error
		enum := new(Deal_Reason)
		err := enum.UnmarshalJSON([]byte("derp"))

		// Verify the error
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of \"derp\" cannot be mapped to a data.Deal_Reason"))
	})

	// Test that attempting to deserialize a data.Deal.Reason will fail and return an error if the value
	// cannot be converted to either the name value or integer value of the enum option
	It("UnmarshalJSON - Value is invalid - Error", func() {

		// Attempt to convert a fake string value into a data.Deal.Reason; this should return an error
		enum := new(Deal_Reason)
		err := enum.UnmarshalJSON([]byte("\"derp\""))

		// Verify the error
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of \"derp\" cannot be mapped to a data.Deal_Reason"))
	})

	// Test the conditions under which values should be convertible to a data.Deal.Reason
	DescribeTable("UnmarshalJSON Tests",
		func(value string, shouldBe Deal_Reason) {

			// Attempt to convert the string value into a data.Deal.Reason; this should not fail
			var enum Deal_Reason
			err := enum.UnmarshalJSON([]byte(value))

			// Verify that the deserialization was successful
			Expect(err).ShouldNot(HaveOccurred())
			Expect(enum).Should(Equal(shouldBe))
		},
		Entry("SL - Works", "\"SL\"", Deal_StopLoss),
		Entry("TP - Works", "\"TP\"", Deal_TakeProfit),
		Entry("SO - Works", "\"SO\"", Deal_StopOut),
		Entry("Client - Works", "\"Client\"", Deal_Client),
		Entry("Mobile - Works", "\"Mobile\"", Deal_Mobile),
		Entry("Web - Works", "\"Web\"", Deal_Web),
		Entry("Strategy - Works", "\"Strategy\"", Deal_Strategy),
		Entry("StopLoss - Works", "\"StopLoss\"", Deal_StopLoss),
		Entry("TakeProfit - Works", "\"TakeProfit\"", Deal_TakeProfit),
		Entry("StopOut - Works", "\"StopOut\"", Deal_StopOut),
		Entry("Rollover", "\"Rollover\"", Deal_Rollover),
		Entry("Margin", "\"Margin\"", Deal_Margin),
		Entry("Split", "\"Split\"", Deal_Split),
		Entry("0 - Works", "\"0\"", Deal_Client),
		Entry("1 - Works", "\"1\"", Deal_Mobile),
		Entry("2 - Works", "\"2\"", Deal_Web),
		Entry("3 - Works", "\"3\"", Deal_Strategy),
		Entry("4 - Works", "\"4\"", Deal_StopLoss),
		Entry("5 - Works", "\"5\"", Deal_TakeProfit),
		Entry("6 - Works", "\"6\"", Deal_StopOut),
		Entry("7 - Works", "\"7\"", Deal_Rollover),
		Entry("8 - Works", "\"8\"", Deal_Margin),
		Entry("9 - Works", "\"9\"", Deal_Split))

	// Test that attempting to deserialize a data.Deal.Reason will fail and return an error if the value
	// cannot be converted to either the name value or integer value of the enum option
	It("UnmarshalCSV - Value is empty - Error", func() {

		// Attempt to convert a fake string value into a data.Deal.Reason; this should return an error
		enum := new(Deal_Reason)
		err := enum.UnmarshalCSV("")

		// Verify the error
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of \"\" cannot be mapped to a data.Deal_Reason"))
	})

	// Test the conditions under which values should be convertible to a data.Deal.Reason
	DescribeTable("UnmarshalCSV Tests",
		func(value string, shouldBe Deal_Reason) {

			// Attempt to convert the value into a data.Deal.Reason; this should not fail
			var enum Deal_Reason
			err := enum.UnmarshalCSV(value)

			// Verify that the deserialization was successful
			Expect(err).ShouldNot(HaveOccurred())
			Expect(enum).Should(Equal(shouldBe))
		},
		Entry("SL - Works", "SL", Deal_StopLoss),
		Entry("TP - Works", "TP", Deal_TakeProfit),
		Entry("SO - Works", "SO", Deal_StopOut),
		Entry("Client - Works", "Client", Deal_Client),
		Entry("Mobile - Works", "Mobile", Deal_Mobile),
		Entry("Web - Works", "Web", Deal_Web),
		Entry("Strategy - Works", "Strategy", Deal_Strategy),
		Entry("StopLoss - Works", "StopLoss", Deal_StopLoss),
		Entry("TakeProfit - Works", "TakeProfit", Deal_TakeProfit),
		Entry("StopOut - Works", "StopOut", Deal_StopOut),
		Entry("Rollover", "Rollover", Deal_Rollover),
		Entry("Margin", "Margin", Deal_Margin),
		Entry("Split", "Split", Deal_Split),
		Entry("0 - Works", "0", Deal_Client),
		Entry("1 - Works", "1", Deal_Mobile),
		Entry("2 - Works", "2", Deal_Web),
		Entry("3 - Works", "3", Deal_Strategy),
		Entry("4 - Works", "4", Deal_StopLoss),
		Entry("5 - Works", "5", Deal_TakeProfit),
		Entry("6 - Works", "6", Deal_StopOut),
		Entry("7 - Works", "7", Deal_Rollover),
		Entry("8 - Works", "8", Deal_Margin),
		Entry("9 - Works", "9", Deal_Split))

	// Test that attempting to deserialize a data.Deal.Reason will fail and return an error if the YAML
	// node does not represent a scalar value
	It("UnmarshalYAML - Node type is not scalar - Error", func() {
		enum := new(Deal_Reason)
		err := enum.UnmarshalYAML(&yaml.Node{Kind: yaml.AliasNode})
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("YAML node had an invalid kind (expected scalar value)"))
	})

	// Test that attempting to deserialize a data.Deal.Reason will fail and return an error if the YAML
	// node value cannot be converted to either the name value or integer value of the enum option
	It("UnmarshalYAML - Parse fails - Error", func() {
		enum := new(Deal_Reason)
		err := enum.UnmarshalYAML(&yaml.Node{Kind: yaml.ScalarNode, Value: "derp"})
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of \"derp\" cannot be mapped to a data.Deal_Reason"))
	})

	// Test the conditions under which YAML node values should be convertible to a data.Deal.Reason
	DescribeTable("UnmarshalYAML Tests",
		func(value string, shouldBe Deal_Reason) {
			var enum Deal_Reason
			err := enum.UnmarshalYAML(&yaml.Node{Kind: yaml.ScalarNode, Value: value})
			Expect(err).ShouldNot(HaveOccurred())
			Expect(enum).Should(Equal(shouldBe))
		},
		Entry("SL - Works", "SL", Deal_StopLoss),
		Entry("TP - Works", "TP", Deal_TakeProfit),
		Entry("SO - Works", "SO", Deal_StopOut),
		Entry("Client - Works", "Client", Deal_Client),
		Entry("Mobile - Works", "Mobile", Deal_Mobile),
		Entry("Web - Works", "Web", Deal_Web),
		Entry("Strategy - Works", "Strategy", Deal_Strategy),
		Entry("StopLoss - Works", "StopLoss", Deal_StopLoss),
		Entry("TakeProfit - Works", "TakeProfit", Deal_TakeProfit),
		Entry("StopOut - Works", "StopOut", Deal_StopOut),
		Entry("Rollover", "Rollover", Deal_Rollover),
		Entry("Margin", "Margin", Deal_Margin),
		Entry("Split", "Split", Deal_Split),
		Entry("0 - Works", "0", Deal_Client),
		Entry("1 - Works", "1", Deal_Mobile),
		Entry("2 - Works", "2", Deal_Web),
		Entry("3 - Works", "3", Deal_Strategy),
		Entry("4 - Works", "4", Deal_StopLoss),
		Entry("5 - Works", "5", Deal_TakeProfit),
		Entry("6 - Works", "6", Deal_StopOut),
		Entry("7 - Works", "7", Deal_Rollover),
		Entry("8 - Works", "8", Deal_Margin),
		Entry("9 - Works", "9", Deal_Split))

	// Tests that, if the attribute type submitted to UnmarshalDynamoDBAttributeValue is not one we
	// recognize, then the function will return an error
	It("UnmarshalDynamoDBAttributeValue - AttributeValue type invalid - Error", func() {
		enum := new(Deal_Reason)
		err := attributevalue.Unmarshal(&types.AttributeValueMemberBOOL{Value: true}, &enum)
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("Attribute value of *types.AttributeValueMemberBOOL could not be converted to a data.Deal.Reason"))
	})

	// Tests that, if time parsing fails, then calling UnmarshalDynamoDBAttributeValue will return an error
	It("UnmarshalDynamoDBAttributeValue - Parse fails - Error", func() {
		enum := new(Deal_Reason)
		err := attributevalue.Unmarshal(&types.AttributeValueMemberS{Value: "derp"}, &enum)
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of \"derp\" cannot be mapped to a data.Deal_Reason"))
	})

	// Tests the conditions under which UnmarshalDynamoDBAttributeValue is called and no error is generated
	DescribeTable("UnmarshalDynamoDBAttributeValue - AttributeValue Conditions",
		func(value types.AttributeValue, expected Deal_Reason) {
			var enum Deal_Reason
			err := attributevalue.Unmarshal(value, &enum)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(enum).Should(Equal(expected))
		},
		Entry("Value is []bytes, SL - Works",
			&types.AttributeValueMemberB{Value: []byte("SL")}, Deal_StopLoss),
		Entry("Value is []bytes, TP - Works",
			&types.AttributeValueMemberB{Value: []byte("TP")}, Deal_TakeProfit),
		Entry("Value is []bytes, SO - Works",
			&types.AttributeValueMemberB{Value: []byte("SO")}, Deal_StopOut),
		Entry("Value is []bytes, Client - Works",
			&types.AttributeValueMemberB{Value: []byte("Client")}, Deal_Client),
		Entry("Value is []bytes, Mobile - Works",
			&types.AttributeValueMemberB{Value: []byte("Mobile")}, Deal_Mobile),
		Entry("Value is []bytes, Web - Works",
			&types.AttributeValueMemberB{Value: []byte("Web")}, Deal_Web),
		Entry("Value is []bytes, Strategy - Works",
			&types.AttributeValueMemberB{Value: []byte("Strategy")}, Deal_Strategy),
		Entry("Value is []bytes, StopLoss - Works",
			&types.AttributeValueMemberB{Value: []byte("StopLoss")}, Deal_StopLoss),
		Entry("Value is []bytes, TakeProfit - Works",
			&types.AttributeValueMemberB{Value: []byte("TakeProfit")}, Deal_TakeProfit),
		Entry("Value is []bytes, StopOut - Works",
			&types.AttributeValueMemberB{Value: []byte("StopOut")}, Deal_StopOut),
		Entry("Value is []bytes, Rollover - Works",
			&types.AttributeValueMemberB{Value: []byte("Rollover")}, Deal_Rollover),
		Entry("Value is []bytes, Margin - Works",
			&types.AttributeValueMemberB{Value: []byte("Margin")}, Deal_Margin),
		Entry("Value is []bytes, Split - Works",
			&types.AttributeValueMemberB{Value: []byte("Split")}, Deal_Split),
		Entry("Value is []bytes, 0 - Works",
			&types.AttributeValueMemberB{Value: []byte("0")}, Deal_Client),
		Entry("Value is []bytes, 1 - Works",
			&types.AttributeValueMemberB{Value: []byte("1")}, Deal_Mobile),
		Entry("Value is []bytes, 2 - Works",
			&types.AttributeValueMemberB{Value: []byte("2")}, Deal_Web),
		Entry("Value is []bytes, 3 - Works",
			&types.AttributeValueMemberB{Value: []byte("3")}, Deal_Strategy),
		Entry("Value is []bytes, 4 - Works",
			&types.AttributeValueMemberB{Value: []byte("4")}, Deal_StopLoss),
		Entry("Value is []bytes, 5 - Works",
			&types.AttributeValueMemberB{Value: []byte("5")}, Deal_TakeProfit),
		Entry("Value is []bytes, 6 - Works",
			&types.AttributeValueMemberB{Value: []byte("6")}, Deal_StopOut),
		Entry("Value is []bytes, 7 - Works",
			&types.AttributeValueMemberB{Value: []byte("7")}, Deal_Rollover),
		Entry("Value is []bytes, 8 - Works",
			&types.AttributeValueMemberB{Value: []byte("8")}, Deal_Margin),
		Entry("Value is []bytes, 9 - Works",
			&types.AttributeValueMemberB{Value: []byte("9")}, Deal_Split),
		Entry("Value is int, 0 - Works",
			&types.AttributeValueMemberN{Value: "0"}, Deal_Client),
		Entry("Value is int, 1 - Works",
			&types.AttributeValueMemberN{Value: "1"}, Deal_Mobile),
		Entry("Value is int, 2 - Works",
			&types.AttributeValueMemberN{Value: "2"}, Deal_Web),
		Entry("Value is int, 3 - Works",
			&types.AttributeValueMemberN{Value: "3"}, Deal_Strategy),
		Entry("Value is int, 4 - Works",
			&types.AttributeValueMemberN{Value: "4"}, Deal_StopLoss),
		Entry("Value is int, 5 - Works",
			&types.AttributeValueMemberN{Value: "5"}, Deal_TakeProfit),
		Entry("Value is int, 6 - Works",
			&types.AttributeValueMemberN{Value: "6"}, Deal_StopOut),
		Entry("Value is int, 7 - Works",
			&types.AttributeValueMemberN{Value: "7"}, Deal_Rollover),
		Entry("Value is int, 8 - Works",
			&types.AttributeValueMemberN{Value: "8"}, Deal_Margin),
		Entry("Value is int, 9 - Works",
			&types.AttributeValueMemberN{Value: "9"}, Deal_Split),
		Entry("Value is NULL - Works", new(types.AttributeValueMemberNULL), Deal_Reason(0)),
		Entry("Value is string, SL - Works",
			&types.AttributeValueMemberS{Value: "SL"}, Deal_StopLoss),
		Entry("Value is string, TP - Works",
			&types.AttributeValueMemberS{Value: "TP"}, Deal_TakeProfit),
		Entry("Value is string, SO - Works",
			&types.AttributeValueMemberS{Value: "SO"}, Deal_StopOut),
		Entry("Value is string, Client - Works",
			&types.AttributeValueMemberN{Value: "Client"}, Deal_Client),
		Entry("Value is string, Mobile - Works",
			&types.AttributeValueMemberN{Value: "Mobile"}, Deal_Mobile),
		Entry("Value is string, Web - Works",
			&types.AttributeValueMemberN{Value: "Web"}, Deal_Web),
		Entry("Value is string, Strategy - Works",
			&types.AttributeValueMemberN{Value: "Strategy"}, Deal_Strategy),
		Entry("Value is string, StopLoss - Works",
			&types.AttributeValueMemberN{Value: "StopLoss"}, Deal_StopLoss),
		Entry("Value is string, TakeProfit - Works",
			&types.AttributeValueMemberN{Value: "TakeProfit"}, Deal_TakeProfit),
		Entry("Value is string, StopOut - Works",
			&types.AttributeValueMemberN{Value: "StopOut"}, Deal_StopOut),
		Entry("Value is string, Rollover - Works",
			&types.AttributeValueMemberN{Value: "Rollover"}, Deal_Rollover),
		Entry("Value is string, Margin - Works",
			&types.AttributeValueMemberN{Value: "Margin"}, Deal_Margin),
		Entry("Value is string, Split - Works",
			&types.AttributeValueMemberN{Value: "Split"}, Deal_Split),
		Entry("Value is string, 0 - Works",
			&types.AttributeValueMemberN{Value: "0"}, Deal_Client),
		Entry("Value is string, 1 - Works",
			&types.AttributeValueMemberN{Value: "1"}, Deal_Mobile),
		Entry("Value is string, 2 - Works",
			&types.AttributeValueMemberN{Value: "2"}, Deal_Web),
		Entry("Value is string, 3 - Works",
			&types.AttributeValueMemberN{Value: "3"}, Deal_Strategy),
		Entry("Value is string, 4 - Works",
			&types.AttributeValueMemberN{Value: "4"}, Deal_StopLoss),
		Entry("Value is string, 5 - Works",
			&types.AttributeValueMemberN{Value: "5"}, Deal_TakeProfit),
		Entry("Value is string, 6 - Works",
			&types.AttributeValueMemberN{Value: "6"}, Deal_StopOut),
		Entry("Value is string, 7 - Works",
			&types.AttributeValueMemberN{Value: "7"}, Deal_Rollover),
		Entry("Value is string, 8 - Works",
			&types.AttributeValueMemberN{Value: "8"}, Deal_Margin),
		Entry("Value is string, 9 - Works",
			&types.AttributeValueMemberN{Value: "9"}, Deal_Split))

	// Test that attempting to deserialize a data.Deal.Reason will fial and return an error if the value
	// cannot be converted to either the name value or integer value of the enum option
	It("Scan - Value is nil - Error", func() {

		// Attempt to convert a fake string value into a data.Deal.Reason; this should return an error
		var enum *Deal_Reason
		err := enum.Scan(nil)

		// Verify the error
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of %!q(<nil>) had an invalid type of <nil>"))
		Expect(enum).Should(BeNil())
	})

	// Test the conditions under which values should be convertible to a data.Deal.Reason
	DescribeTable("Scan Tests",
		func(value interface{}, shouldBe Deal_Reason) {

			// Attempt to convert the value into a data.Deal.Reason; this should not fail
			var enum Deal_Reason
			err := enum.Scan(value)

			// Verify that the deserialization was successful
			Expect(err).ShouldNot(HaveOccurred())
			Expect(enum).Should(Equal(shouldBe))
		},
		Entry("SL - Works", "SL", Deal_StopLoss),
		Entry("TP - Works", "TP", Deal_TakeProfit),
		Entry("SO - Works", "SO", Deal_StopOut),
		Entry("Client - Works", "Client", Deal_Client),
		Entry("Mobile - Works", "Mobile", Deal_Mobile),
		Entry("Web - Works", "Web", Deal_Web),
		Entry("Strategy - Works", "Strategy", Deal_Strategy),
		Entry("StopLoss - Works", "StopLoss", Deal_StopLoss),
		Entry("TakeProfit - Works", "TakeProfit", Deal_TakeProfit),
		Entry("StopOut - Works", "StopOut", Deal_StopOut),
		Entry("Rollover", "Rollover", Deal_Rollover),
		Entry("Margin", "Margin", Deal_Margin),
		Entry("Split", "Split", Deal_Split),
		Entry("0 - Works", 0, Deal_Client),
		Entry("1 - Works", 1, Deal_Mobile),
		Entry("2 - Works", 2, Deal_Web),
		Entry("3 - Works", 3, Deal_Strategy),
		Entry("4 - Works", 4, Deal_StopLoss),
		Entry("5 - Works", 5, Deal_TakeProfit),
		Entry("6 - Works", 6, Deal_StopOut),
		Entry("7 - Works", 7, Deal_Rollover),
		Entry("8 - Works", 8, Deal_Margin),
		Entry("9 - Works", 9, Deal_Split))
})

var _ = Describe("data.Order.Type Marshal/Unmarshal Tests", func() {

	// Test that converting the data.Order.Type enum to JSON works for all values
	DescribeTable("MarshalJSON Tests",
		func(enum Order_Type, value string) {
			data, err := json.Marshal(enum)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(string(data)).Should(Equal(value))
		},
		Entry("Buy - Works", Order_Buy, "\"Buy\""),
		Entry("Sell - Works", Order_Sell, "\"Sell\""),
		Entry("BuyLimit - Works", Order_BuyLimit, "\"BuyLimit\""),
		Entry("SellLimit - Works", Order_SellLimit, "\"SellLimit\""),
		Entry("BuyStop - Works", Order_BuyStop, "\"BuyStop\""),
		Entry("SellStop - Works", Order_SellStop, "\"SellStop\""),
		Entry("BuyStopLimit - Works", Order_BuyStopLimit, "\"BuyStopLimit\""),
		Entry("SellStopLimit - Works", Order_SellStopLimit, "\"SellStopLimit\""),
		Entry("ClosedBy - Works", Order_ClosedBy, "\"ClosedBy\""))

	// Test that converting the data.Order.Type enum to a CSV column works for all values
	DescribeTable("MarshalCSV Tests",
		func(enum Order_Type, value string) {
			data, err := enum.MarshalCSV()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(string(data)).Should(Equal(value))
		},
		Entry("Buy - Works", Order_Buy, "Buy"),
		Entry("Sell - Works", Order_Sell, "Sell"),
		Entry("BuyLimit - Works", Order_BuyLimit, "BuyLimit"),
		Entry("SellLimit - Works", Order_SellLimit, "SellLimit"),
		Entry("BuyStop - Works", Order_BuyStop, "BuyStop"),
		Entry("SellStop - Works", Order_SellStop, "SellStop"),
		Entry("BuyStopLimit - Works", Order_BuyStopLimit, "BuyStopLimit"),
		Entry("SellStopLimit - Works", Order_SellStopLimit, "SellStopLimit"),
		Entry("ClosedBy - Works", Order_ClosedBy, "ClosedBy"))

	// Test that converting the data.Order.Type enum to a YAML node works for all values
	DescribeTable("MarshalYAML - Works",
		func(enum Order_Type, value string) {
			data, err := enum.MarshalYAML()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(data).Should(Equal(value))
		},
		Entry("Buy - Works", Order_Buy, "Buy"),
		Entry("Sell - Works", Order_Sell, "Sell"),
		Entry("BuyLimit - Works", Order_BuyLimit, "BuyLimit"),
		Entry("SellLimit - Works", Order_SellLimit, "SellLimit"),
		Entry("BuyStop - Works", Order_BuyStop, "BuyStop"),
		Entry("SellStop - Works", Order_SellStop, "SellStop"),
		Entry("BuyStopLimit - Works", Order_BuyStopLimit, "BuyStopLimit"),
		Entry("SellStopLimit - Works", Order_SellStopLimit, "SellStopLimit"),
		Entry("ClosedBy - Works", Order_ClosedBy, "ClosedBy"))

	// Test that converting the data.Order.Type enum to a DynamoDB AttributeVAlue works for all values
	DescribeTable("MarshalDynamoDBAttributeValue - Works",
		func(enum Order_Type, value string) {
			data, err := attributevalue.Marshal(enum)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(data.(*types.AttributeValueMemberS).Value).Should(Equal(value))
		},
		Entry("Buy - Works", Order_Buy, "Buy"),
		Entry("Sell - Works", Order_Sell, "Sell"),
		Entry("BuyLimit - Works", Order_BuyLimit, "BuyLimit"),
		Entry("SellLimit - Works", Order_SellLimit, "SellLimit"),
		Entry("BuyStop - Works", Order_BuyStop, "BuyStop"),
		Entry("SellStop - Works", Order_SellStop, "SellStop"),
		Entry("BuyStopLimit - Works", Order_BuyStopLimit, "BuyStopLimit"),
		Entry("SellStopLimit - Works", Order_SellStopLimit, "SellStopLimit"),
		Entry("ClosedBy - Works", Order_ClosedBy, "ClosedBy"))

	// Test that converting the data.Order.Type enum to an SQL value for all values
	DescribeTable("Value Tests",
		func(enum Order_Type, value string) {
			data, err := enum.Value()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(data).Should(Equal(value))
		},
		Entry("Buy - Works", Order_Buy, "Buy"),
		Entry("Sell - Works", Order_Sell, "Sell"),
		Entry("BuyLimit - Works", Order_BuyLimit, "BuyLimit"),
		Entry("SellLimit - Works", Order_SellLimit, "SellLimit"),
		Entry("BuyStop - Works", Order_BuyStop, "BuyStop"),
		Entry("SellStop - Works", Order_SellStop, "SellStop"),
		Entry("BuyStopLimit - Works", Order_BuyStopLimit, "BuyStopLimit"),
		Entry("SellStopLimit - Works", Order_SellStopLimit, "SellStopLimit"),
		Entry("ClosedBy - Works", Order_ClosedBy, "ClosedBy"))

	// Test that attempting to deserialize a data.Order.Type will fail and return an error if the value
	// cannot be deserialized from a JSON value to a string
	It("UnmarshalJSON fails - Error", func() {

		// Attempt to convert a non-parseable string value into a data.Order.Type; this should return an error
		enum := new(Order_Type)
		err := enum.UnmarshalJSON([]byte("derp"))

		// Verify the error
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of \"derp\" cannot be mapped to a data.Order_Type"))
	})

	// Test that attempting to deserialize a data.Order.Type will fail and return an error if the value
	// cannot be converted to either the name value or integer value of the enum option
	It("UnmarshalJSON - Value is invalid - Error", func() {

		// Attempt to convert a fake string value into a data.Order.Type; this should return an error
		enum := new(Order_Type)
		err := enum.UnmarshalJSON([]byte("\"derp\""))

		// Verify the error
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of \"derp\" cannot be mapped to a data.Order_Type"))
	})

	// Test the conditions under which values should be convertible to a data.Order.Type
	DescribeTable("UnmarshalJSON Tests",
		func(value string, shouldBe Order_Type) {

			// Attempt to convert the string value into a data.Order.Type; this should not fail
			var enum Order_Type
			err := enum.UnmarshalJSON([]byte(value))

			// Verify that the deserialization was successful
			Expect(err).ShouldNot(HaveOccurred())
			Expect(enum).Should(Equal(shouldBe))
		},
		Entry("Buy - Works", "\"Buy\"", Order_Buy),
		Entry("Sell - Works", "\"Sell\"", Order_Sell),
		Entry("BuyLimit - Works", "\"BuyLimit\"", Order_BuyLimit),
		Entry("SellLimit - Works", "\"SellLimit\"", Order_SellLimit),
		Entry("BuyStop - Works", "\"BuyStop\"", Order_BuyStop),
		Entry("SellStop - Works", "\"SellStop\"", Order_SellStop),
		Entry("BuyStopLimit - Works", "\"BuyStopLimit\"", Order_BuyStopLimit),
		Entry("SellStopLimit", "\"SellStopLimit\"", Order_SellStopLimit),
		Entry("ClosedBy", "\"ClosedBy\"", Order_ClosedBy),
		Entry("0 - Works", "\"0\"", Order_Buy),
		Entry("1 - Works", "\"1\"", Order_Sell),
		Entry("2 - Works", "\"2\"", Order_BuyLimit),
		Entry("3 - Works", "\"3\"", Order_SellLimit),
		Entry("4 - Works", "\"4\"", Order_BuyStop),
		Entry("5 - Works", "\"5\"", Order_SellStop),
		Entry("6 - Works", "\"6\"", Order_BuyStopLimit),
		Entry("7 - Works", "\"7\"", Order_SellStopLimit),
		Entry("8 - Works", "\"8\"", Order_ClosedBy))

	// Test that attempting to deserialize a data.Order.Type will fail and return an error if the value
	// cannot be converted to either the name value or integer value of the enum option
	It("UnmarshalCSV - Value is empty - Error", func() {

		// Attempt to convert a fake string value into a data.Order.Type; this should return an error
		enum := new(Order_Type)
		err := enum.UnmarshalCSV("")

		// Verify the error
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of \"\" cannot be mapped to a data.Order_Type"))
	})

	// Test the conditions under which values should be convertible to a data.Order.Type
	DescribeTable("UnmarshalCSV Tests",
		func(value string, shouldBe Order_Type) {

			// Attempt to convert the value into a data.Order.Type; this should not fail
			var enum Order_Type
			err := enum.UnmarshalCSV(value)

			// Verify that the deserialization was successful
			Expect(err).ShouldNot(HaveOccurred())
			Expect(enum).Should(Equal(shouldBe))
		},
		Entry("Buy - Works", "Buy", Order_Buy),
		Entry("Sell - Works", "Sell", Order_Sell),
		Entry("BuyLimit - Works", "BuyLimit", Order_BuyLimit),
		Entry("SellLimit - Works", "SellLimit", Order_SellLimit),
		Entry("BuyStop - Works", "BuyStop", Order_BuyStop),
		Entry("SellStop - Works", "SellStop", Order_SellStop),
		Entry("BuyStopLimit - Works", "BuyStopLimit", Order_BuyStopLimit),
		Entry("SellStopLimit", "SellStopLimit", Order_SellStopLimit),
		Entry("ClosedBy", "ClosedBy", Order_ClosedBy),
		Entry("0 - Works", "0", Order_Buy),
		Entry("1 - Works", "1", Order_Sell),
		Entry("2 - Works", "2", Order_BuyLimit),
		Entry("3 - Works", "3", Order_SellLimit),
		Entry("4 - Works", "4", Order_BuyStop),
		Entry("5 - Works", "5", Order_SellStop),
		Entry("6 - Works", "6", Order_BuyStopLimit),
		Entry("7 - Works", "7", Order_SellStopLimit),
		Entry("8 - Works", "8", Order_ClosedBy))

	// Test that attempting to deserialize a data.Order.Type will fail and return an error if the YAML
	// node does not represent a scalar value
	It("UnmarshalYAML - Node type is not scalar - Error", func() {
		enum := new(Order_Type)
		err := enum.UnmarshalYAML(&yaml.Node{Kind: yaml.AliasNode})
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("YAML node had an invalid kind (expected scalar value)"))
	})

	// Test that attempting to deserialize a data.Order.Type will fail and return an error if the YAML
	// node value cannot be converted to either the name value or integer value of the enum option
	It("UnmarshalYAML - Parse fails - Error", func() {
		enum := new(Order_Type)
		err := enum.UnmarshalYAML(&yaml.Node{Kind: yaml.ScalarNode, Value: "derp"})
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of \"derp\" cannot be mapped to a data.Order_Type"))
	})

	// Test the conditions under which YAML node values should be convertible to a data.Order.Type
	DescribeTable("UnmarshalYAML Tests",
		func(value string, shouldBe Order_Type) {
			var enum Order_Type
			err := enum.UnmarshalYAML(&yaml.Node{Kind: yaml.ScalarNode, Value: value})
			Expect(err).ShouldNot(HaveOccurred())
			Expect(enum).Should(Equal(shouldBe))
		},
		Entry("Buy - Works", "Buy", Order_Buy),
		Entry("Sell - Works", "Sell", Order_Sell),
		Entry("BuyLimit - Works", "BuyLimit", Order_BuyLimit),
		Entry("SellLimit - Works", "SellLimit", Order_SellLimit),
		Entry("BuyStop - Works", "BuyStop", Order_BuyStop),
		Entry("SellStop - Works", "SellStop", Order_SellStop),
		Entry("BuyStopLimit - Works", "BuyStopLimit", Order_BuyStopLimit),
		Entry("SellStopLimit", "SellStopLimit", Order_SellStopLimit),
		Entry("ClosedBy", "ClosedBy", Order_ClosedBy),
		Entry("0 - Works", "0", Order_Buy),
		Entry("1 - Works", "1", Order_Sell),
		Entry("2 - Works", "2", Order_BuyLimit),
		Entry("3 - Works", "3", Order_SellLimit),
		Entry("4 - Works", "4", Order_BuyStop),
		Entry("5 - Works", "5", Order_SellStop),
		Entry("6 - Works", "6", Order_BuyStopLimit),
		Entry("7 - Works", "7", Order_SellStopLimit),
		Entry("8 - Works", "8", Order_ClosedBy))

	// Tests that, if the attribute type submitted to UnmarshalDynamoDBAttributeValue is not one we
	// recognize, then the function will return an error
	It("UnmarshalDynamoDBAttributeValue - AttributeValue type invalid - Error", func() {
		enum := new(Order_Type)
		err := attributevalue.Unmarshal(&types.AttributeValueMemberBOOL{Value: true}, &enum)
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("Attribute value of *types.AttributeValueMemberBOOL could not be converted to a data.Order.Type"))
	})

	// Tests that, if time parsing fails, then calling UnmarshalDynamoDBAttributeValue will return an error
	It("UnmarshalDynamoDBAttributeValue - Parse fails - Error", func() {
		enum := new(Order_Type)
		err := attributevalue.Unmarshal(&types.AttributeValueMemberS{Value: "derp"}, &enum)
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of \"derp\" cannot be mapped to a data.Order_Type"))
	})

	// Tests the conditions under which UnmarshalDynamoDBAttributeValue is called and no error is generated
	DescribeTable("UnmarshalDynamoDBAttributeValue - AttributeValue Conditions",
		func(value types.AttributeValue, expected Order_Type) {
			var enum Order_Type
			err := attributevalue.Unmarshal(value, &enum)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(enum).Should(Equal(expected))
		},
		Entry("Value is []bytes, Buy - Works",
			&types.AttributeValueMemberB{Value: []byte("Buy")}, Order_Buy),
		Entry("Value is []bytes, Sell - Works",
			&types.AttributeValueMemberB{Value: []byte("Sell")}, Order_Sell),
		Entry("Value is []bytes, BuyLimit - Works",
			&types.AttributeValueMemberB{Value: []byte("BuyLimit")}, Order_BuyLimit),
		Entry("Value is []bytes, SellLimit - Works",
			&types.AttributeValueMemberB{Value: []byte("SellLimit")}, Order_SellLimit),
		Entry("Value is []bytes, BuyStop - Works",
			&types.AttributeValueMemberB{Value: []byte("BuyStop")}, Order_BuyStop),
		Entry("Value is []bytes, SellStop - Works",
			&types.AttributeValueMemberB{Value: []byte("SellStop")}, Order_SellStop),
		Entry("Value is []bytes, BuyStopLimit - Works",
			&types.AttributeValueMemberB{Value: []byte("BuyStopLimit")}, Order_BuyStopLimit),
		Entry("Value is []bytes, SellStopLimit - Works",
			&types.AttributeValueMemberB{Value: []byte("SellStopLimit")}, Order_SellStopLimit),
		Entry("Value is []bytes, ClosedBy - Works",
			&types.AttributeValueMemberB{Value: []byte("ClosedBy")}, Order_ClosedBy),
		Entry("Value is []bytes, 0 - Works",
			&types.AttributeValueMemberB{Value: []byte("0")}, Order_Buy),
		Entry("Value is []bytes, 1 - Works",
			&types.AttributeValueMemberB{Value: []byte("1")}, Order_Sell),
		Entry("Value is []bytes, 2 - Works",
			&types.AttributeValueMemberB{Value: []byte("2")}, Order_BuyLimit),
		Entry("Value is []bytes, 3 - Works",
			&types.AttributeValueMemberB{Value: []byte("3")}, Order_SellLimit),
		Entry("Value is []bytes, 4 - Works",
			&types.AttributeValueMemberB{Value: []byte("4")}, Order_BuyStop),
		Entry("Value is []bytes, 5 - Works",
			&types.AttributeValueMemberB{Value: []byte("5")}, Order_SellStop),
		Entry("Value is []bytes, 6 - Works",
			&types.AttributeValueMemberB{Value: []byte("6")}, Order_BuyStopLimit),
		Entry("Value is []bytes, 7 - Works",
			&types.AttributeValueMemberB{Value: []byte("7")}, Order_SellStopLimit),
		Entry("Value is []bytes, 8 - Works",
			&types.AttributeValueMemberB{Value: []byte("8")}, Order_ClosedBy),
		Entry("Value is int, 0 - Works",
			&types.AttributeValueMemberN{Value: "0"}, Order_Buy),
		Entry("Value is int, 1 - Works",
			&types.AttributeValueMemberN{Value: "1"}, Order_Sell),
		Entry("Value is int, 2 - Works",
			&types.AttributeValueMemberN{Value: "2"}, Order_BuyLimit),
		Entry("Value is int, 3 - Works",
			&types.AttributeValueMemberN{Value: "3"}, Order_SellLimit),
		Entry("Value is int, 4 - Works",
			&types.AttributeValueMemberN{Value: "4"}, Order_BuyStop),
		Entry("Value is int, 5 - Works",
			&types.AttributeValueMemberN{Value: "5"}, Order_SellStop),
		Entry("Value is int, 6 - Works",
			&types.AttributeValueMemberN{Value: "6"}, Order_BuyStopLimit),
		Entry("Value is int, 7 - Works",
			&types.AttributeValueMemberN{Value: "7"}, Order_SellStopLimit),
		Entry("Value is int, 8 - Works",
			&types.AttributeValueMemberN{Value: "8"}, Order_ClosedBy),
		Entry("Value is NULL - Works", new(types.AttributeValueMemberNULL), Order_Type(0)),
		Entry("Value is string, Buy - Works",
			&types.AttributeValueMemberN{Value: "Buy"}, Order_Buy),
		Entry("Value is string, Sell - Works",
			&types.AttributeValueMemberN{Value: "Sell"}, Order_Sell),
		Entry("Value is string, BuyLimit - Works",
			&types.AttributeValueMemberN{Value: "BuyLimit"}, Order_BuyLimit),
		Entry("Value is string, SellLimit - Works",
			&types.AttributeValueMemberN{Value: "SellLimit"}, Order_SellLimit),
		Entry("Value is string, BuyStop - Works",
			&types.AttributeValueMemberN{Value: "BuyStop"}, Order_BuyStop),
		Entry("Value is string, SellStop - Works",
			&types.AttributeValueMemberN{Value: "SellStop"}, Order_SellStop),
		Entry("Value is string, BuyStopLimit - Works",
			&types.AttributeValueMemberN{Value: "BuyStopLimit"}, Order_BuyStopLimit),
		Entry("Value is string, SellStopLimit - Works",
			&types.AttributeValueMemberN{Value: "SellStopLimit"}, Order_SellStopLimit),
		Entry("Value is string, ClosedBy - Works",
			&types.AttributeValueMemberN{Value: "ClosedBy"}, Order_ClosedBy),
		Entry("Value is string, 0 - Works",
			&types.AttributeValueMemberN{Value: "0"}, Order_Buy),
		Entry("Value is string, 1 - Works",
			&types.AttributeValueMemberN{Value: "1"}, Order_Sell),
		Entry("Value is string, 2 - Works",
			&types.AttributeValueMemberN{Value: "2"}, Order_BuyLimit),
		Entry("Value is string, 3 - Works",
			&types.AttributeValueMemberN{Value: "3"}, Order_SellLimit),
		Entry("Value is string, 4 - Works",
			&types.AttributeValueMemberN{Value: "4"}, Order_BuyStop),
		Entry("Value is string, 5 - Works",
			&types.AttributeValueMemberN{Value: "5"}, Order_SellStop),
		Entry("Value is string, 6 - Works",
			&types.AttributeValueMemberN{Value: "6"}, Order_BuyStopLimit),
		Entry("Value is string, 7 - Works",
			&types.AttributeValueMemberN{Value: "7"}, Order_SellStopLimit),
		Entry("Value is string, 8 - Works",
			&types.AttributeValueMemberN{Value: "8"}, Order_ClosedBy))

	// Test that attempting to deserialize a data.Order.Type will fial and return an error if the value
	// cannot be converted to either the name value or integer value of the enum option
	It("Scan - Value is nil - Error", func() {

		// Attempt to convert a fake string value into a data.Order.Type; this should return an error
		var enum *Order_Type
		err := enum.Scan(nil)

		// Verify the error
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of %!q(<nil>) had an invalid type of <nil>"))
		Expect(enum).Should(BeNil())
	})

	// Test the conditions under which values should be convertible to a data.Order.Type
	DescribeTable("Scan Tests",
		func(value interface{}, shouldBe Order_Type) {

			// Attempt to convert the value into a data.Order.Type; this should not fail
			var enum Order_Type
			err := enum.Scan(value)

			// Verify that the deserialization was successful
			Expect(err).ShouldNot(HaveOccurred())
			Expect(enum).Should(Equal(shouldBe))
		},
		Entry("Buy - Works", "Buy", Order_Buy),
		Entry("Sell - Works", "Sell", Order_Sell),
		Entry("BuyLimit - Works", "BuyLimit", Order_BuyLimit),
		Entry("SellLimit - Works", "SellLimit", Order_SellLimit),
		Entry("BuyStop - Works", "BuyStop", Order_BuyStop),
		Entry("SellStop - Works", "SellStop", Order_SellStop),
		Entry("BuyStopLimit - Works", "BuyStopLimit", Order_BuyStopLimit),
		Entry("SellStopLimit", "SellStopLimit", Order_SellStopLimit),
		Entry("ClosedBy", "ClosedBy", Order_ClosedBy),
		Entry("0 - Works", 0, Order_Buy),
		Entry("1 - Works", 1, Order_Sell),
		Entry("2 - Works", 2, Order_BuyLimit),
		Entry("3 - Works", 3, Order_SellLimit),
		Entry("4 - Works", 4, Order_BuyStop),
		Entry("5 - Works", 5, Order_SellStop),
		Entry("6 - Works", 6, Order_BuyStopLimit),
		Entry("7 - Works", 7, Order_SellStopLimit),
		Entry("8 - Works", 8, Order_ClosedBy))
})
