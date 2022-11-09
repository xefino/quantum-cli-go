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
