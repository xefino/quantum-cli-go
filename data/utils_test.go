package data

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"gopkg.in/yaml.v3"
)

var _ = Describe("data.AllowedTrades Marshal/Unmarshal Tests", func() {

	// Test that converting the data.AllowedTrades enum to JSON works for all values
	DescribeTable("MarshalJSON Tests",
		func(enum AllowedTrades, value string) {
			data, err := json.Marshal(enum)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(string(data)).Should(Equal(value))
		},
		Entry("AllAllowed - Works", AllowedTrades_AllAllowed, "\"ALL\""),
		Entry("BuysAllowed - Works", AllowedTrades_BuysAllowed, "\"BUY\""),
		Entry("SellsAllowed - Works", AllowedTrades_SellsAllowed, "\"SELL\""))

	// Test that converting the data.AllowedTrades enum to a CSV column works for all values
	DescribeTable("MarshalCSV Tests",
		func(enum AllowedTrades, value string) {
			data, err := enum.MarshalCSV()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(string(data)).Should(Equal(value))
		},
		Entry("AllAllowed - Works", AllowedTrades_AllAllowed, "ALL"),
		Entry("BuysAllowed - Works", AllowedTrades_BuysAllowed, "BUY"),
		Entry("SellsAllowed - Works", AllowedTrades_SellsAllowed, "SELL"))

	// Test that converting the data.AllowedTrades enum to a YAML node works for all values
	DescribeTable("MarshalYAML - Works",
		func(enum AllowedTrades, value string) {
			data, err := enum.MarshalYAML()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(data).Should(Equal(value))
		},
		Entry("AllAllowed - Works", AllowedTrades_AllAllowed, "ALL"),
		Entry("BuysAllowed - Works", AllowedTrades_BuysAllowed, "BUY"),
		Entry("SellsAllowed - Works", AllowedTrades_SellsAllowed, "SELL"))

	// Test that converting the data.AllowedTrades enum to a DynamoDB AttributeVAlue works for all values
	DescribeTable("MarshalDynamoDBAttributeValue - Works",
		func(enum AllowedTrades, value string) {
			data, err := attributevalue.Marshal(enum)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(data.(*types.AttributeValueMemberS).Value).Should(Equal(value))
		},
		Entry("AllAllowed - Works", AllowedTrades_AllAllowed, "ALL"),
		Entry("BuysAllowed - Works", AllowedTrades_BuysAllowed, "BUY"),
		Entry("SellsAllowed - Works", AllowedTrades_SellsAllowed, "SELL"))

	// Test that converting the data.AllowedTrades enum to an SQL value for all values
	DescribeTable("Value Tests",
		func(enum AllowedTrades, value string) {
			data, err := enum.Value()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(data).Should(Equal(value))
		},
		Entry("AllAllowed - Works", AllowedTrades_AllAllowed, "ALL"),
		Entry("BuysAllowed - Works", AllowedTrades_BuysAllowed, "BUY"),
		Entry("SellsAllowed - Works", AllowedTrades_SellsAllowed, "SELL"))

	// Test that attempting to deserialize a data.AllowedTrades will fail and return an error if the value
	// cannot be deserialized from a JSON value to a string
	It("UnmarshalJSON fails - Error", func() {

		// Attempt to convert a non-parseable string value into a data.AllowedTrades; this should return an error
		enum := new(AllowedTrades)
		err := enum.UnmarshalJSON([]byte("derp"))

		// Verify the error
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of \"derp\" cannot be mapped to a v1.AllowedTrades"))
	})

	// Test that attempting to deserialize a data.AllowedTrades will fail and return an error if the value
	// cannot be converted to either the name value or integer value of the enum option
	It("UnmarshalJSON - Value is invalid - Error", func() {

		// Attempt to convert a fake string value into a data.AllowedTrades; this should return an error
		enum := new(AllowedTrades)
		err := enum.UnmarshalJSON([]byte("\"derp\""))

		// Verify the error
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of \"derp\" cannot be mapped to a v1.AllowedTrades"))
	})

	// Test the conditions under which values should be convertible to a data.AllowedTrades
	DescribeTable("UnmarshalJSON Tests",
		func(value string, shouldBe AllowedTrades) {

			// Attempt to convert the string value into a data.AllowedTrades; this should not fail
			var enum AllowedTrades
			err := enum.UnmarshalJSON([]byte(value))

			// Verify that the deserialization was successful
			Expect(err).ShouldNot(HaveOccurred())
			Expect(enum).Should(Equal(shouldBe))
		},
		Entry("all - Works", "\"all\"", AllowedTrades_AllAllowed),
		Entry("ALL - Works", "\"ALL\"", AllowedTrades_AllAllowed),
		Entry("buy - Works", "\"buy\"", AllowedTrades_BuysAllowed),
		Entry("BUY - Works", "\"BUY\"", AllowedTrades_BuysAllowed),
		Entry("sell - Works", "\"sell\"", AllowedTrades_SellsAllowed),
		Entry("SELL - Works", "\"SELL\"", AllowedTrades_SellsAllowed),
		Entry("AllAllowed - Works", "\"AllAllowed\"", AllowedTrades_AllAllowed),
		Entry("BuysAllowed - Works", "\"BuysAllowed\"", AllowedTrades_BuysAllowed),
		Entry("SellsAllowed - Works", "\"SellsAllowed\"", AllowedTrades_SellsAllowed),
		Entry("0 - Works", "\"0\"", AllowedTrades_AllAllowed),
		Entry("1 - Works", "\"1\"", AllowedTrades_BuysAllowed),
		Entry("2 - Works", "\"2\"", AllowedTrades_SellsAllowed))

	// Test that attempting to deserialize a data.AllowedTrades will fail and return an error if the value
	// cannot be converted to either the name value or integer value of the enum option
	It("UnmarshalCSV - Value is empty - Error", func() {

		// Attempt to convert a fake string value into a data.AllowedTrades; this should return an error
		enum := new(AllowedTrades)
		err := enum.UnmarshalCSV("")

		// Verify the error
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of \"\" cannot be mapped to a v1.AllowedTrades"))
	})

	// Test the conditions under which values should be convertible to a data.AllowedTrades
	DescribeTable("UnmarshalCSV Tests",
		func(value string, shouldBe AllowedTrades) {

			// Attempt to convert the value into a data.AllowedTrades; this should not fail
			var enum AllowedTrades
			err := enum.UnmarshalCSV(value)

			// Verify that the deserialization was successful
			Expect(err).ShouldNot(HaveOccurred())
			Expect(enum).Should(Equal(shouldBe))
		},
		Entry("all - Works", "all", AllowedTrades_AllAllowed),
		Entry("ALL - Works", "ALL", AllowedTrades_AllAllowed),
		Entry("buy - Works", "buy", AllowedTrades_BuysAllowed),
		Entry("BUY - Works", "BUY", AllowedTrades_BuysAllowed),
		Entry("sell - Works", "sell", AllowedTrades_SellsAllowed),
		Entry("SELL - Works", "SELL", AllowedTrades_SellsAllowed),
		Entry("AllAllowed - Works", "AllAllowed", AllowedTrades_AllAllowed),
		Entry("BuysAllowed - Works", "BuysAllowed", AllowedTrades_BuysAllowed),
		Entry("SellsAllowed - Works", "SellsAllowed", AllowedTrades_SellsAllowed),
		Entry("0 - Works", "0", AllowedTrades_AllAllowed),
		Entry("1 - Works", "1", AllowedTrades_BuysAllowed),
		Entry("2 - Works", "2", AllowedTrades_SellsAllowed))

	// Test that attempting to deserialize a data.AllowedTrades will fail and return an error if the YAML
	// node does not represent a scalar value
	It("UnmarshalYAML - Node type is not scalar - Error", func() {
		enum := new(AllowedTrades)
		err := enum.UnmarshalYAML(&yaml.Node{Kind: yaml.AliasNode})
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("YAML node had an invalid kind (expected scalar value)"))
	})

	// Test that attempting to deserialize a data.AllowedTrades will fail and return an error if the YAML
	// node value cannot be converted to either the name value or integer value of the enum option
	It("UnmarshalYAML - Parse fails - Error", func() {
		enum := new(AllowedTrades)
		err := enum.UnmarshalYAML(&yaml.Node{Kind: yaml.ScalarNode, Value: "derp"})
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of \"derp\" cannot be mapped to a v1.AllowedTrades"))
	})

	// Test the conditions under which YAML node values should be convertible to a data.AllowedTrades
	DescribeTable("UnmarshalYAML Tests",
		func(value string, shouldBe AllowedTrades) {
			var enum AllowedTrades
			err := enum.UnmarshalYAML(&yaml.Node{Kind: yaml.ScalarNode, Value: value})
			Expect(err).ShouldNot(HaveOccurred())
			Expect(enum).Should(Equal(shouldBe))
		},
		Entry("all - Works", "all", AllowedTrades_AllAllowed),
		Entry("ALL - Works", "ALL", AllowedTrades_AllAllowed),
		Entry("buy - Works", "buy", AllowedTrades_BuysAllowed),
		Entry("BUY - Works", "BUY", AllowedTrades_BuysAllowed),
		Entry("sell - Works", "sell", AllowedTrades_SellsAllowed),
		Entry("SELL - Works", "SELL", AllowedTrades_SellsAllowed),
		Entry("AllAllowed - Works", "AllAllowed", AllowedTrades_AllAllowed),
		Entry("BuysAllowed - Works", "BuysAllowed", AllowedTrades_BuysAllowed),
		Entry("SellsAllowed - Works", "SellsAllowed", AllowedTrades_SellsAllowed),
		Entry("0 - Works", "0", AllowedTrades_AllAllowed),
		Entry("1 - Works", "1", AllowedTrades_BuysAllowed),
		Entry("2 - Works", "2", AllowedTrades_SellsAllowed))

	// Tests that, if the attribute type submitted to UnmarshalDynamoDBAttributeValue is not one we
	// recognize, then the function will return an error
	It("UnmarshalDynamoDBAttributeValue - AttributeValue type invalid - Error", func() {
		enum := new(AllowedTrades)
		err := attributevalue.Unmarshal(&types.AttributeValueMemberBOOL{Value: true}, &enum)
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("Attribute value of *types.AttributeValueMemberBOOL could not be converted to a data.AllowedTrades"))
	})

	// Tests that, if time parsing fails, then calling UnmarshalDynamoDBAttributeValue will return an error
	It("UnmarshalDynamoDBAttributeValue - Parse fails - Error", func() {
		enum := new(AllowedTrades)
		err := attributevalue.Unmarshal(&types.AttributeValueMemberS{Value: "derp"}, &enum)
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of \"derp\" cannot be mapped to a v1.AllowedTrades"))
	})

	// Tests the conditions under which UnmarshalDynamoDBAttributeValue is called and no error is generated
	DescribeTable("UnmarshalDynamoDBAttributeValue - AttributeValue Conditions",
		func(value types.AttributeValue, expected AllowedTrades) {
			var enum AllowedTrades
			err := attributevalue.Unmarshal(value, &enum)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(enum).Should(Equal(expected))
		},
		Entry("value is []bytes, all - Works",
			&types.AttributeValueMemberB{Value: []byte("all")}, AllowedTrades_AllAllowed),
		Entry("Value is []bytes, ALL - Works",
			&types.AttributeValueMemberB{Value: []byte("ALL")}, AllowedTrades_AllAllowed),
		Entry("Value is []bytes, buy - Works",
			&types.AttributeValueMemberB{Value: []byte("buy")}, AllowedTrades_BuysAllowed),
		Entry("Value is []bytes, BUY - Works",
			&types.AttributeValueMemberB{Value: []byte("BUY")}, AllowedTrades_BuysAllowed),
		Entry("Value is []bytes, sell - Works",
			&types.AttributeValueMemberB{Value: []byte("sell")}, AllowedTrades_SellsAllowed),
		Entry("Value is []bytes, SELL - Works",
			&types.AttributeValueMemberB{Value: []byte("SELL")}, AllowedTrades_SellsAllowed),
		Entry("Value is []bytes, AllAllowed - Works",
			&types.AttributeValueMemberB{Value: []byte("AllAllowed")}, AllowedTrades_AllAllowed),
		Entry("Value is []bytes, BuysAllowed - Works",
			&types.AttributeValueMemberB{Value: []byte("BuysAllowed")}, AllowedTrades_BuysAllowed),
		Entry("Value is []bytes, SellsAllowed - Works",
			&types.AttributeValueMemberB{Value: []byte("SellsAllowed")}, AllowedTrades_SellsAllowed),
		Entry("Value is []bytes, 0 - Works",
			&types.AttributeValueMemberB{Value: []byte("0")}, AllowedTrades_AllAllowed),
		Entry("Value is []bytes, 1 - Works",
			&types.AttributeValueMemberB{Value: []byte("1")}, AllowedTrades_BuysAllowed),
		Entry("Value is []bytes, 2 - Works",
			&types.AttributeValueMemberB{Value: []byte("2")}, AllowedTrades_SellsAllowed),
		Entry("Value is int, 0 - Works",
			&types.AttributeValueMemberN{Value: "0"}, AllowedTrades_AllAllowed),
		Entry("Value is int, 1 - Works",
			&types.AttributeValueMemberN{Value: "1"}, AllowedTrades_BuysAllowed),
		Entry("Value is int, 2 - Works",
			&types.AttributeValueMemberN{Value: "2"}, AllowedTrades_SellsAllowed),
		Entry("Value is NULL - Works", new(types.AttributeValueMemberNULL), AllowedTrades(0)),
		Entry("Value is string, all - Works",
			&types.AttributeValueMemberS{Value: "all"}, AllowedTrades_AllAllowed),
		Entry("Value is string, ALL - Works",
			&types.AttributeValueMemberS{Value: "ALL"}, AllowedTrades_AllAllowed),
		Entry("Value is string, buy - Works",
			&types.AttributeValueMemberS{Value: "buy"}, AllowedTrades_BuysAllowed),
		Entry("Value is string, BUY - Works",
			&types.AttributeValueMemberS{Value: "BUY"}, AllowedTrades_BuysAllowed),
		Entry("Value is string, sell - Works",
			&types.AttributeValueMemberS{Value: "sell"}, AllowedTrades_SellsAllowed),
		Entry("Value is string, SELL - Works",
			&types.AttributeValueMemberS{Value: "SELL"}, AllowedTrades_SellsAllowed),
		Entry("Value is string, AllAllowed - Works",
			&types.AttributeValueMemberS{Value: "AllAllowed"}, AllowedTrades_AllAllowed),
		Entry("Value is string, BuysAllowed - Works",
			&types.AttributeValueMemberS{Value: "BuysAllowed"}, AllowedTrades_BuysAllowed),
		Entry("Value is string, SellsAllowed - Works",
			&types.AttributeValueMemberS{Value: "SellsAllowed"}, AllowedTrades_SellsAllowed),
		Entry("Value is string, 0 - Works",
			&types.AttributeValueMemberS{Value: "0"}, AllowedTrades_AllAllowed),
		Entry("Value is string, 1 - Works",
			&types.AttributeValueMemberS{Value: "1"}, AllowedTrades_BuysAllowed),
		Entry("Value is string, 2 - Works",
			&types.AttributeValueMemberS{Value: "2"}, AllowedTrades_SellsAllowed))

	// Test that attempting to deserialize a data.AllowedTrades will fial and return an error if the value
	// cannot be converted to either the name value or integer value of the enum option
	It("Scan - Value is nil - Error", func() {

		// Attempt to convert a fake string value into a data.AllowedTrades; this should return an error
		var enum *AllowedTrades
		err := enum.Scan(nil)

		// Verify the error
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of %!q(<nil>) had an invalid type of <nil>"))
		Expect(enum).Should(BeNil())
	})

	// Test the conditions under which values should be convertible to a data.AllowedTrades
	DescribeTable("Scan Tests",
		func(value interface{}, shouldBe AllowedTrades) {

			// Attempt to convert the value into a data.AllowedTrades; this should not fail
			var enum AllowedTrades
			err := enum.Scan(value)

			// Verify that the deserialization was successful
			Expect(err).ShouldNot(HaveOccurred())
			Expect(enum).Should(Equal(shouldBe))
		},
		Entry("all - Works", "all", AllowedTrades_AllAllowed),
		Entry("ALL - Works", "ALL", AllowedTrades_AllAllowed),
		Entry("buy - Works", "buy", AllowedTrades_BuysAllowed),
		Entry("BUY - Works", "BUY", AllowedTrades_BuysAllowed),
		Entry("sell - Works", "sell", AllowedTrades_SellsAllowed),
		Entry("SELL - Works", "SELL", AllowedTrades_SellsAllowed),
		Entry("AllAllowed - Works", "AllAllowed", AllowedTrades_AllAllowed),
		Entry("BuysAllowed - Works", "BuysAllowed", AllowedTrades_BuysAllowed),
		Entry("SellsAllowed - Works", "SellsAllowed", AllowedTrades_SellsAllowed),
		Entry("0 - Works", 0, AllowedTrades_AllAllowed),
		Entry("1 - Works", 1, AllowedTrades_BuysAllowed),
		Entry("2 - Works", 2, AllowedTrades_SellsAllowed))
})

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
		Entry("Commission - Works", "\"Commission\"", Deal_Commission),
		Entry("CommissionDaily - Works", "\"CommissionDaily\"", Deal_CommissionDaily),
		Entry("CommissionMonthly - Works", "\"CommissionMonthly\"", Deal_CommissionMonthly),
		Entry("CommissionAgentDaily - Works", "\"CommissionAgentDaily\"", Deal_CommissionAgentDaily),
		Entry("CommissionAgentMonthly - Works", "\"CommissionAgentMonthly\"", Deal_CommissionAgentMonthly),
		Entry("Interest - Works", "\"Interest\"", Deal_Interest),
		Entry("BuyCancelled - Works", "\"BuyCancelled\"", Deal_BuyCancelled),
		Entry("SellCancelled - Works", "\"SellCancelled\"", Deal_SellCancelled),
		Entry("Dividend - Works", "\"Dividend\"", Deal_Dividend),
		Entry("Franked - Works", "\"Franked\"", Deal_Franked),
		Entry("Tax - Works", "\"Tax\"", Deal_Tax),
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
		Entry("Commission - Works", "Commission", Deal_Commission),
		Entry("CommissionDaily - Works", "CommissionDaily", Deal_CommissionDaily),
		Entry("CommissionMonthly - Works", "CommissionMonthly", Deal_CommissionMonthly),
		Entry("CommissionAgentDaily - Works", "CommissionAgentDaily", Deal_CommissionAgentDaily),
		Entry("CommissionAgentMonthly - Works", "CommissionAgentMonthly", Deal_CommissionAgentMonthly),
		Entry("Interest - Works", "Interest", Deal_Interest),
		Entry("BuyCancelled - Works", "BuyCancelled", Deal_BuyCancelled),
		Entry("SellCancelled - Works", "SellCancelled", Deal_SellCancelled),
		Entry("Dividend - Works", "Dividend", Deal_Dividend),
		Entry("Franked - Works", "Franked", Deal_Franked),
		Entry("Tax - Works", "Tax", Deal_Tax),
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
		Entry("Commission - Works", "Commission", Deal_Commission),
		Entry("CommissionDaily - Works", "CommissionDaily", Deal_CommissionDaily),
		Entry("CommissionMonthly - Works", "CommissionMonthly", Deal_CommissionMonthly),
		Entry("CommissionAgentDaily - Works", "CommissionAgentDaily", Deal_CommissionAgentDaily),
		Entry("CommissionAgentMonthly - Works", "CommissionAgentMonthly", Deal_CommissionAgentMonthly),
		Entry("Interest - Works", "Interest", Deal_Interest),
		Entry("BuyCancelled - Works", "BuyCancelled", Deal_BuyCancelled),
		Entry("SellCancelled - Works", "SellCancelled", Deal_SellCancelled),
		Entry("Dividend - Works", "Dividend", Deal_Dividend),
		Entry("Franked - Works", "Franked", Deal_Franked),
		Entry("Tax - Works", "Tax", Deal_Tax),
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
			&types.AttributeValueMemberS{Value: "Buy"}, Deal_Buy),
		Entry("Value is string, Sell - Works",
			&types.AttributeValueMemberS{Value: "Sell"}, Deal_Sell),
		Entry("Value is string, Balance - Works",
			&types.AttributeValueMemberS{Value: "Balance"}, Deal_Balance),
		Entry("Value is string, Credit - Works",
			&types.AttributeValueMemberS{Value: "Credit"}, Deal_Credit),
		Entry("Value is string, Charge - Works",
			&types.AttributeValueMemberS{Value: "Charge"}, Deal_Charge),
		Entry("Value is string, Correction - Works",
			&types.AttributeValueMemberS{Value: "Correction"}, Deal_Correction),
		Entry("Value is string, Bonus - Works",
			&types.AttributeValueMemberS{Value: "Bonus"}, Deal_Bonus),
		Entry("Value is string, Commission - Works",
			&types.AttributeValueMemberS{Value: "Commission"}, Deal_Commission),
		Entry("Value is string, CommissionDaily - Works",
			&types.AttributeValueMemberS{Value: "CommissionDaily"}, Deal_CommissionDaily),
		Entry("Value is string, CommissionMonthly - Works",
			&types.AttributeValueMemberS{Value: "CommissionMonthly"}, Deal_CommissionMonthly),
		Entry("Value is string, CommissionAgentDaily - Works",
			&types.AttributeValueMemberS{Value: "CommissionAgentDaily"}, Deal_CommissionAgentDaily),
		Entry("Value is string, CommissionAgentMonthly - Works",
			&types.AttributeValueMemberS{Value: "CommissionAgentMonthly"}, Deal_CommissionAgentMonthly),
		Entry("Value is string, Interest - Works",
			&types.AttributeValueMemberS{Value: "Interest"}, Deal_Interest),
		Entry("Value is string, BuyCancelled - Works",
			&types.AttributeValueMemberS{Value: "BuyCancelled"}, Deal_BuyCancelled),
		Entry("Value is string, SellCancelled - Works",
			&types.AttributeValueMemberS{Value: "SellCancelled"}, Deal_SellCancelled),
		Entry("Value is string, Dividend - Works",
			&types.AttributeValueMemberS{Value: "Dividend"}, Deal_Dividend),
		Entry("Value is string, Franked - Works",
			&types.AttributeValueMemberS{Value: "Franked"}, Deal_Franked),
		Entry("Value is string, Tax - Works",
			&types.AttributeValueMemberS{Value: "Tax"}, Deal_Tax),
		Entry("Value is string, 0 - Works",
			&types.AttributeValueMemberS{Value: "0"}, Deal_Buy),
		Entry("Value is string, 1 - Works",
			&types.AttributeValueMemberS{Value: "1"}, Deal_Sell),
		Entry("Value is string, 2 - Works",
			&types.AttributeValueMemberS{Value: "2"}, Deal_Balance),
		Entry("Value is string, 3 - Works",
			&types.AttributeValueMemberS{Value: "3"}, Deal_Credit),
		Entry("Value is string, 4 - Works",
			&types.AttributeValueMemberS{Value: "4"}, Deal_Charge),
		Entry("Value is string, 5 - Works",
			&types.AttributeValueMemberS{Value: "5"}, Deal_Correction),
		Entry("Value is string, 6 - Works",
			&types.AttributeValueMemberS{Value: "6"}, Deal_Bonus),
		Entry("Value is string, 7 - Works",
			&types.AttributeValueMemberS{Value: "7"}, Deal_Commission),
		Entry("Value is string, 8 - Works",
			&types.AttributeValueMemberS{Value: "8"}, Deal_CommissionDaily),
		Entry("Value is string, 9 - Works",
			&types.AttributeValueMemberS{Value: "9"}, Deal_CommissionMonthly),
		Entry("Value is string, 10 - Works",
			&types.AttributeValueMemberS{Value: "10"}, Deal_CommissionAgentDaily),
		Entry("Value is string, 11 - Works",
			&types.AttributeValueMemberS{Value: "11"}, Deal_CommissionAgentMonthly),
		Entry("Value is string, 12 - Works",
			&types.AttributeValueMemberS{Value: "12"}, Deal_Interest),
		Entry("Value is string, 13 - Works",
			&types.AttributeValueMemberS{Value: "13"}, Deal_BuyCancelled),
		Entry("Value is string, 14 - Works",
			&types.AttributeValueMemberS{Value: "14"}, Deal_SellCancelled),
		Entry("Value is string, 15 - Works",
			&types.AttributeValueMemberS{Value: "15"}, Deal_Dividend),
		Entry("Value is string, 16 - Works",
			&types.AttributeValueMemberS{Value: "16"}, Deal_Franked),
		Entry("Value is string, 17 - Works",
			&types.AttributeValueMemberS{Value: "17"}, Deal_Tax))

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
		Entry("Commission - Works", "Commission", Deal_Commission),
		Entry("CommissionDaily - Works", "CommissionDaily", Deal_CommissionDaily),
		Entry("CommissionMonthly - Works", "CommissionMonthly", Deal_CommissionMonthly),
		Entry("CommissionAgentDaily - Works", "CommissionAgentDaily", Deal_CommissionAgentDaily),
		Entry("CommissionAgentMonthly - Works", "CommissionAgentMonthly", Deal_CommissionAgentMonthly),
		Entry("Interest - Works", "Interest", Deal_Interest),
		Entry("BuyCancelled - Works", "BuyCancelled", Deal_BuyCancelled),
		Entry("SellCancelled - Works", "SellCancelled", Deal_SellCancelled),
		Entry("Dividend - Works", "Dividend", Deal_Dividend),
		Entry("Franked - Works", "Franked", Deal_Franked),
		Entry("Tax - Works", "Tax", Deal_Tax),
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
			&types.AttributeValueMemberS{Value: "In"}, Deal_In),
		Entry("Value is string, Out - Works",
			&types.AttributeValueMemberS{Value: "Out"}, Deal_Out),
		Entry("Value is string, Reverse - Works",
			&types.AttributeValueMemberS{Value: "Reverse"}, Deal_Reverse),
		Entry("Value is string, OutBy - Works",
			&types.AttributeValueMemberS{Value: "OutBy"}, Deal_OutBy),
		Entry("Value is string, 0 - Works",
			&types.AttributeValueMemberS{Value: "0"}, Deal_In),
		Entry("Value is string, 1 - Works",
			&types.AttributeValueMemberS{Value: "1"}, Deal_Out),
		Entry("Value is string, 2 - Works",
			&types.AttributeValueMemberS{Value: "2"}, Deal_Reverse),
		Entry("Value is string, 3 - Works",
			&types.AttributeValueMemberS{Value: "3"}, Deal_OutBy))

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
		Entry("Rollover - Works", "\"Rollover\"", Deal_Rollover),
		Entry("Margin - Works", "\"Margin\"", Deal_Margin),
		Entry("Split - Works", "\"Split\"", Deal_Split),
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
		Entry("Rollover - Works", "Rollover", Deal_Rollover),
		Entry("Margin - Works", "Margin", Deal_Margin),
		Entry("Split - Works", "Split", Deal_Split),
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
		Entry("Rollover - Works", "Rollover", Deal_Rollover),
		Entry("Margin - Works", "Margin", Deal_Margin),
		Entry("Split - Works", "Split", Deal_Split),
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
			&types.AttributeValueMemberS{Value: "Client"}, Deal_Client),
		Entry("Value is string, Mobile - Works",
			&types.AttributeValueMemberS{Value: "Mobile"}, Deal_Mobile),
		Entry("Value is string, Web - Works",
			&types.AttributeValueMemberS{Value: "Web"}, Deal_Web),
		Entry("Value is string, Strategy - Works",
			&types.AttributeValueMemberS{Value: "Strategy"}, Deal_Strategy),
		Entry("Value is string, StopLoss - Works",
			&types.AttributeValueMemberS{Value: "StopLoss"}, Deal_StopLoss),
		Entry("Value is string, TakeProfit - Works",
			&types.AttributeValueMemberS{Value: "TakeProfit"}, Deal_TakeProfit),
		Entry("Value is string, StopOut - Works",
			&types.AttributeValueMemberS{Value: "StopOut"}, Deal_StopOut),
		Entry("Value is string, Rollover - Works",
			&types.AttributeValueMemberS{Value: "Rollover"}, Deal_Rollover),
		Entry("Value is string, Margin - Works",
			&types.AttributeValueMemberS{Value: "Margin"}, Deal_Margin),
		Entry("Value is string, Split - Works",
			&types.AttributeValueMemberS{Value: "Split"}, Deal_Split),
		Entry("Value is string, 0 - Works",
			&types.AttributeValueMemberS{Value: "0"}, Deal_Client),
		Entry("Value is string, 1 - Works",
			&types.AttributeValueMemberS{Value: "1"}, Deal_Mobile),
		Entry("Value is string, 2 - Works",
			&types.AttributeValueMemberS{Value: "2"}, Deal_Web),
		Entry("Value is string, 3 - Works",
			&types.AttributeValueMemberS{Value: "3"}, Deal_Strategy),
		Entry("Value is string, 4 - Works",
			&types.AttributeValueMemberS{Value: "4"}, Deal_StopLoss),
		Entry("Value is string, 5 - Works",
			&types.AttributeValueMemberS{Value: "5"}, Deal_TakeProfit),
		Entry("Value is string, 6 - Works",
			&types.AttributeValueMemberS{Value: "6"}, Deal_StopOut),
		Entry("Value is string, 7 - Works",
			&types.AttributeValueMemberS{Value: "7"}, Deal_Rollover),
		Entry("Value is string, 8 - Works",
			&types.AttributeValueMemberS{Value: "8"}, Deal_Margin),
		Entry("Value is string, 9 - Works",
			&types.AttributeValueMemberS{Value: "9"}, Deal_Split))

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
		Entry("Rollover - Works", "Rollover", Deal_Rollover),
		Entry("Margin - Works", "Margin", Deal_Margin),
		Entry("Split - Works", "Split", Deal_Split),
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
		Entry("SellStopLimit - Works", "\"SellStopLimit\"", Order_SellStopLimit),
		Entry("ClosedBy - Works", "\"ClosedBy\"", Order_ClosedBy),
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
		Entry("SellStopLimit - Works", "SellStopLimit", Order_SellStopLimit),
		Entry("ClosedBy - Works", "ClosedBy", Order_ClosedBy),
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
		Entry("SellStopLimit - Works", "SellStopLimit", Order_SellStopLimit),
		Entry("ClosedBy - Works", "ClosedBy", Order_ClosedBy),
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
			&types.AttributeValueMemberS{Value: "Buy"}, Order_Buy),
		Entry("Value is string, Sell - Works",
			&types.AttributeValueMemberS{Value: "Sell"}, Order_Sell),
		Entry("Value is string, BuyLimit - Works",
			&types.AttributeValueMemberS{Value: "BuyLimit"}, Order_BuyLimit),
		Entry("Value is string, SellLimit - Works",
			&types.AttributeValueMemberS{Value: "SellLimit"}, Order_SellLimit),
		Entry("Value is string, BuyStop - Works",
			&types.AttributeValueMemberS{Value: "BuyStop"}, Order_BuyStop),
		Entry("Value is string, SellStop - Works",
			&types.AttributeValueMemberS{Value: "SellStop"}, Order_SellStop),
		Entry("Value is string, BuyStopLimit - Works",
			&types.AttributeValueMemberS{Value: "BuyStopLimit"}, Order_BuyStopLimit),
		Entry("Value is string, SellStopLimit - Works",
			&types.AttributeValueMemberS{Value: "SellStopLimit"}, Order_SellStopLimit),
		Entry("Value is string, ClosedBy - Works",
			&types.AttributeValueMemberS{Value: "ClosedBy"}, Order_ClosedBy),
		Entry("Value is string, 0 - Works",
			&types.AttributeValueMemberS{Value: "0"}, Order_Buy),
		Entry("Value is string, 1 - Works",
			&types.AttributeValueMemberS{Value: "1"}, Order_Sell),
		Entry("Value is string, 2 - Works",
			&types.AttributeValueMemberS{Value: "2"}, Order_BuyLimit),
		Entry("Value is string, 3 - Works",
			&types.AttributeValueMemberS{Value: "3"}, Order_SellLimit),
		Entry("Value is string, 4 - Works",
			&types.AttributeValueMemberS{Value: "4"}, Order_BuyStop),
		Entry("Value is string, 5 - Works",
			&types.AttributeValueMemberS{Value: "5"}, Order_SellStop),
		Entry("Value is string, 6 - Works",
			&types.AttributeValueMemberS{Value: "6"}, Order_BuyStopLimit),
		Entry("Value is string, 7 - Works",
			&types.AttributeValueMemberS{Value: "7"}, Order_SellStopLimit),
		Entry("Value is string, 8 - Works",
			&types.AttributeValueMemberS{Value: "8"}, Order_ClosedBy))

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
		Entry("SellStopLimit - Works", "SellStopLimit", Order_SellStopLimit),
		Entry("ClosedBy - Works", "ClosedBy", Order_ClosedBy),
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

var _ = Describe("data.Order.Status Marshal/Unmarshal Tests", func() {

	// Test that converting the data.Order.Status enum to JSON works for all values
	DescribeTable("MarshalJSON Tests",
		func(enum Order_Status, value string) {
			data, err := json.Marshal(enum)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(string(data)).Should(Equal(value))
		},
		Entry("Started - Works", Order_Started, "\"Started\""),
		Entry("Placed - Works", Order_Placed, "\"Placed\""),
		Entry("Cancelled - Works", Order_Cancelled, "\"Cancelled\""),
		Entry("Partial - Works", Order_Partial, "\"Partial\""),
		Entry("Filled - Works", Order_Filled, "\"Filled\""),
		Entry("Rejected - Works", Order_Rejected, "\"Rejected\""),
		Entry("Expired - Works", Order_Expired, "\"Expired\""),
		Entry("RequestAdd - Works", Order_RequestAdd, "\"RequestAdd\""),
		Entry("RequestModify - Works", Order_RequestModify, "\"RequestModify\""),
		Entry("RequestCancel - Works", Order_RequestCancel, "\"RequestCancel\""))

	// Test that converting the data.Order.Status enum to a CSV column works for all values
	DescribeTable("MarshalCSV Tests",
		func(enum Order_Status, value string) {
			data, err := enum.MarshalCSV()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(string(data)).Should(Equal(value))
		},
		Entry("Started - Works", Order_Started, "Started"),
		Entry("Placed - Works", Order_Placed, "Placed"),
		Entry("Cancelled - Works", Order_Cancelled, "Cancelled"),
		Entry("Partial - Works", Order_Partial, "Partial"),
		Entry("Filled - Works", Order_Filled, "Filled"),
		Entry("Rejected - Works", Order_Rejected, "Rejected"),
		Entry("Expired - Works", Order_Expired, "Expired"),
		Entry("RequestAdd - Works", Order_RequestAdd, "RequestAdd"),
		Entry("RequestModify - Works", Order_RequestModify, "RequestModify"),
		Entry("RequestCancel - Works", Order_RequestCancel, "RequestCancel"))

	// Test that converting the data.Order.Status enum to a YAML node works for all values
	DescribeTable("MarshalYAML - Works",
		func(enum Order_Status, value string) {
			data, err := enum.MarshalYAML()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(data).Should(Equal(value))
		},
		Entry("Started - Works", Order_Started, "Started"),
		Entry("Placed - Works", Order_Placed, "Placed"),
		Entry("Cancelled - Works", Order_Cancelled, "Cancelled"),
		Entry("Partial - Works", Order_Partial, "Partial"),
		Entry("Filled - Works", Order_Filled, "Filled"),
		Entry("Rejected - Works", Order_Rejected, "Rejected"),
		Entry("Expired - Works", Order_Expired, "Expired"),
		Entry("RequestAdd - Works", Order_RequestAdd, "RequestAdd"),
		Entry("RequestModify - Works", Order_RequestModify, "RequestModify"),
		Entry("RequestCancel - Works", Order_RequestCancel, "RequestCancel"))

	// Test that converting the data.Order.Status enum to a DynamoDB AttributeVAlue works for all values
	DescribeTable("MarshalDynamoDBAttributeValue - Works",
		func(enum Order_Status, value string) {
			data, err := attributevalue.Marshal(enum)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(data.(*types.AttributeValueMemberS).Value).Should(Equal(value))
		},
		Entry("Started - Works", Order_Started, "Started"),
		Entry("Placed - Works", Order_Placed, "Placed"),
		Entry("Cancelled - Works", Order_Cancelled, "Cancelled"),
		Entry("Partial - Works", Order_Partial, "Partial"),
		Entry("Filled - Works", Order_Filled, "Filled"),
		Entry("Rejected - Works", Order_Rejected, "Rejected"),
		Entry("Expired - Works", Order_Expired, "Expired"),
		Entry("RequestAdd - Works", Order_RequestAdd, "RequestAdd"),
		Entry("RequestModify - Works", Order_RequestModify, "RequestModify"),
		Entry("RequestCancel - Works", Order_RequestCancel, "RequestCancel"))

	// Test that converting the data.Order.Status enum to an SQL value for all values
	DescribeTable("Value Tests",
		func(enum Order_Status, value string) {
			data, err := enum.Value()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(data).Should(Equal(value))
		},
		Entry("Started - Works", Order_Started, "Started"),
		Entry("Placed - Works", Order_Placed, "Placed"),
		Entry("Cancelled - Works", Order_Cancelled, "Cancelled"),
		Entry("Partial - Works", Order_Partial, "Partial"),
		Entry("Filled - Works", Order_Filled, "Filled"),
		Entry("Rejected - Works", Order_Rejected, "Rejected"),
		Entry("Expired - Works", Order_Expired, "Expired"),
		Entry("RequestAdd - Works", Order_RequestAdd, "RequestAdd"),
		Entry("RequestModify - Works", Order_RequestModify, "RequestModify"),
		Entry("RequestCancel - Works", Order_RequestCancel, "RequestCancel"))

	// Test that attempting to deserialize a data.Order.Status will fail and return an error if the value
	// cannot be deserialized from a JSON value to a string
	It("UnmarshalJSON fails - Error", func() {

		// Attempt to convert a non-parseable string value into a data.Order.Status; this should return an error
		enum := new(Order_Status)
		err := enum.UnmarshalJSON([]byte("derp"))

		// Verify the error
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of \"derp\" cannot be mapped to a data.Order_Status"))
	})

	// Test that attempting to deserialize a data.Order.Status will fail and return an error if the value
	// cannot be converted to either the name value or integer value of the enum option
	It("UnmarshalJSON - Value is invalid - Error", func() {

		// Attempt to convert a fake string value into a data.Order.Status; this should return an error
		enum := new(Order_Status)
		err := enum.UnmarshalJSON([]byte("\"derp\""))

		// Verify the error
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of \"derp\" cannot be mapped to a data.Order_Status"))
	})

	// Test the conditions under which values should be convertible to a data.Order.Status
	DescribeTable("UnmarshalJSON Tests",
		func(value string, shouldBe Order_Status) {

			// Attempt to convert the string value into a data.Order.Status; this should not fail
			var enum Order_Status
			err := enum.UnmarshalJSON([]byte(value))

			// Verify that the deserialization was successful
			Expect(err).ShouldNot(HaveOccurred())
			Expect(enum).Should(Equal(shouldBe))
		},
		Entry("Started - Works", "\"Started\"", Order_Started),
		Entry("Placed - Works", "\"Placed\"", Order_Placed),
		Entry("Cancelled - Works", "\"Cancelled\"", Order_Cancelled),
		Entry("Partial - Works", "\"Partial\"", Order_Partial),
		Entry("Filled - Works", "\"Filled\"", Order_Filled),
		Entry("Rejected - Works", "\"Rejected\"", Order_Rejected),
		Entry("Expired - Works", "\"Expired\"", Order_Expired),
		Entry("RequestAdd - Works", "\"RequestAdd\"", Order_RequestAdd),
		Entry("RequestModify - Works", "\"RequestModify\"", Order_RequestModify),
		Entry("RequestCancel - Works", "\"RequestCancel\"", Order_RequestCancel),
		Entry("0 - Works", "\"0\"", Order_Started),
		Entry("1 - Works", "\"1\"", Order_Placed),
		Entry("2 - Works", "\"2\"", Order_Cancelled),
		Entry("3 - Works", "\"3\"", Order_Partial),
		Entry("4 - Works", "\"4\"", Order_Filled),
		Entry("5 - Works", "\"5\"", Order_Rejected),
		Entry("6 - Works", "\"6\"", Order_Expired),
		Entry("7 - Works", "\"7\"", Order_RequestAdd),
		Entry("8 - Works", "\"8\"", Order_RequestModify),
		Entry("9 - Works", "\"9\"", Order_RequestCancel))

	// Test that attempting to deserialize a data.Order.Status will fail and return an error if the value
	// cannot be converted to either the name value or integer value of the enum option
	It("UnmarshalCSV - Value is empty - Error", func() {

		// Attempt to convert a fake string value into a data.Order.Status; this should return an error
		enum := new(Order_Status)
		err := enum.UnmarshalCSV("")

		// Verify the error
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of \"\" cannot be mapped to a data.Order_Status"))
	})

	// Test the conditions under which values should be convertible to a data.Order.Status
	DescribeTable("UnmarshalCSV Tests",
		func(value string, shouldBe Order_Status) {

			// Attempt to convert the value into a data.Order.Status; this should not fail
			var enum Order_Status
			err := enum.UnmarshalCSV(value)

			// Verify that the deserialization was successful
			Expect(err).ShouldNot(HaveOccurred())
			Expect(enum).Should(Equal(shouldBe))
		},
		Entry("Started - Works", "Started", Order_Started),
		Entry("Placed - Works", "Placed", Order_Placed),
		Entry("Cancelled - Works", "Cancelled", Order_Cancelled),
		Entry("Partial - Works", "Partial", Order_Partial),
		Entry("Filled - Works", "Filled", Order_Filled),
		Entry("Rejected - Works", "Rejected", Order_Rejected),
		Entry("Expired - Works", "Expired", Order_Expired),
		Entry("RequestAdd - Works", "RequestAdd", Order_RequestAdd),
		Entry("RequestModify - Works", "RequestModify", Order_RequestModify),
		Entry("RequestCancel - Works", "RequestCancel", Order_RequestCancel),
		Entry("0 - Works", "0", Order_Started),
		Entry("1 - Works", "1", Order_Placed),
		Entry("2 - Works", "2", Order_Cancelled),
		Entry("3 - Works", "3", Order_Partial),
		Entry("4 - Works", "4", Order_Filled),
		Entry("5 - Works", "5", Order_Rejected),
		Entry("6 - Works", "6", Order_Expired),
		Entry("7 - Works", "7", Order_RequestAdd),
		Entry("8 - Works", "8", Order_RequestModify),
		Entry("9 - Works", "9", Order_RequestCancel))

	// Test that attempting to deserialize a data.Order.Status will fail and return an error if the YAML
	// node does not represent a scalar value
	It("UnmarshalYAML - Node type is not scalar - Error", func() {
		enum := new(Order_Status)
		err := enum.UnmarshalYAML(&yaml.Node{Kind: yaml.AliasNode})
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("YAML node had an invalid kind (expected scalar value)"))
	})

	// Test that attempting to deserialize a data.Order.Status will fail and return an error if the YAML
	// node value cannot be converted to either the name value or integer value of the enum option
	It("UnmarshalYAML - Parse fails - Error", func() {
		enum := new(Order_Status)
		err := enum.UnmarshalYAML(&yaml.Node{Kind: yaml.ScalarNode, Value: "derp"})
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of \"derp\" cannot be mapped to a data.Order_Status"))
	})

	// Test the conditions under which YAML node values should be convertible to a data.Order.Status
	DescribeTable("UnmarshalYAML Tests",
		func(value string, shouldBe Order_Status) {
			var enum Order_Status
			err := enum.UnmarshalYAML(&yaml.Node{Kind: yaml.ScalarNode, Value: value})
			Expect(err).ShouldNot(HaveOccurred())
			Expect(enum).Should(Equal(shouldBe))
		},
		Entry("Started - Works", "Started", Order_Started),
		Entry("Placed - Works", "Placed", Order_Placed),
		Entry("Cancelled - Works", "Cancelled", Order_Cancelled),
		Entry("Partial - Works", "Partial", Order_Partial),
		Entry("Filled - Works", "Filled", Order_Filled),
		Entry("Rejected - Works", "Rejected", Order_Rejected),
		Entry("Expired - Works", "Expired", Order_Expired),
		Entry("RequestAdd - Works", "RequestAdd", Order_RequestAdd),
		Entry("RequestModify - Works", "RequestModify", Order_RequestModify),
		Entry("RequestCancel - Works", "RequestCancel", Order_RequestCancel),
		Entry("0 - Works", "0", Order_Started),
		Entry("1 - Works", "1", Order_Placed),
		Entry("2 - Works", "2", Order_Cancelled),
		Entry("3 - Works", "3", Order_Partial),
		Entry("4 - Works", "4", Order_Filled),
		Entry("5 - Works", "5", Order_Rejected),
		Entry("6 - Works", "6", Order_Expired),
		Entry("7 - Works", "7", Order_RequestAdd),
		Entry("8 - Works", "8", Order_RequestModify),
		Entry("9 - Works", "9", Order_RequestCancel))

	// Tests that, if the attribute type submitted to UnmarshalDynamoDBAttributeValue is not one we
	// recognize, then the function will return an error
	It("UnmarshalDynamoDBAttributeValue - AttributeValue type invalid - Error", func() {
		enum := new(Order_Status)
		err := attributevalue.Unmarshal(&types.AttributeValueMemberBOOL{Value: true}, &enum)
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("Attribute value of *types.AttributeValueMemberBOOL could not be converted to a data.Order.Status"))
	})

	// Tests that, if time parsing fails, then calling UnmarshalDynamoDBAttributeValue will return an error
	It("UnmarshalDynamoDBAttributeValue - Parse fails - Error", func() {
		enum := new(Order_Status)
		err := attributevalue.Unmarshal(&types.AttributeValueMemberS{Value: "derp"}, &enum)
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of \"derp\" cannot be mapped to a data.Order_Status"))
	})

	// Tests the conditions under which UnmarshalDynamoDBAttributeValue is called and no error is generated
	DescribeTable("UnmarshalDynamoDBAttributeValue - AttributeValue Conditions",
		func(value types.AttributeValue, expected Order_Status) {
			var enum Order_Status
			err := attributevalue.Unmarshal(value, &enum)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(enum).Should(Equal(expected))
		},
		Entry("Value is []bytes, Started - Works",
			&types.AttributeValueMemberB{Value: []byte("Started")}, Order_Started),
		Entry("Value is []bytes, Placed - Works",
			&types.AttributeValueMemberB{Value: []byte("Placed")}, Order_Placed),
		Entry("Value is []bytes, Cancelled - Works",
			&types.AttributeValueMemberB{Value: []byte("Cancelled")}, Order_Cancelled),
		Entry("Value is []bytes, Partial - Works",
			&types.AttributeValueMemberB{Value: []byte("Partial")}, Order_Partial),
		Entry("Value is []bytes, Filled - Works",
			&types.AttributeValueMemberB{Value: []byte("Filled")}, Order_Filled),
		Entry("Value is []bytes, Rejected - Works",
			&types.AttributeValueMemberB{Value: []byte("Rejected")}, Order_Rejected),
		Entry("Value is []bytes, Expired - Works",
			&types.AttributeValueMemberB{Value: []byte("Expired")}, Order_Expired),
		Entry("Value is []bytes, RequestAdd - Works",
			&types.AttributeValueMemberB{Value: []byte("RequestAdd")}, Order_RequestAdd),
		Entry("Value is []bytes, RequestModify - Works",
			&types.AttributeValueMemberB{Value: []byte("RequestModify")}, Order_RequestModify),
		Entry("Value is []bytes, RequestCancel - Works",
			&types.AttributeValueMemberB{Value: []byte("RequestCancel")}, Order_RequestCancel),
		Entry("Value is []bytes, 0 - Works",
			&types.AttributeValueMemberB{Value: []byte("0")}, Order_Started),
		Entry("Value is []bytes, 1 - Works",
			&types.AttributeValueMemberB{Value: []byte("1")}, Order_Placed),
		Entry("Value is []bytes, 2 - Works",
			&types.AttributeValueMemberB{Value: []byte("2")}, Order_Cancelled),
		Entry("Value is []bytes, 3 - Works",
			&types.AttributeValueMemberB{Value: []byte("3")}, Order_Partial),
		Entry("Value is []bytes, 4 - Works",
			&types.AttributeValueMemberB{Value: []byte("4")}, Order_Filled),
		Entry("Value is []bytes, 5 - Works",
			&types.AttributeValueMemberB{Value: []byte("5")}, Order_Rejected),
		Entry("Value is []bytes, 6 - Works",
			&types.AttributeValueMemberB{Value: []byte("6")}, Order_Expired),
		Entry("Value is []bytes, 7 - Works",
			&types.AttributeValueMemberB{Value: []byte("7")}, Order_RequestAdd),
		Entry("Value is []bytes, 8 - Works",
			&types.AttributeValueMemberB{Value: []byte("8")}, Order_RequestModify),
		Entry("Value is []bytes, 9 - Works",
			&types.AttributeValueMemberB{Value: []byte("9")}, Order_RequestCancel),
		Entry("Value is int, 0 - Works",
			&types.AttributeValueMemberN{Value: "0"}, Order_Started),
		Entry("Value is int, 1 - Works",
			&types.AttributeValueMemberN{Value: "1"}, Order_Placed),
		Entry("Value is int, 2 - Works",
			&types.AttributeValueMemberN{Value: "2"}, Order_Cancelled),
		Entry("Value is int, 3 - Works",
			&types.AttributeValueMemberN{Value: "3"}, Order_Partial),
		Entry("Value is int, 4 - Works",
			&types.AttributeValueMemberN{Value: "4"}, Order_Filled),
		Entry("Value is int, 5 - Works",
			&types.AttributeValueMemberN{Value: "5"}, Order_Rejected),
		Entry("Value is int, 6 - Works",
			&types.AttributeValueMemberN{Value: "6"}, Order_Expired),
		Entry("Value is int, 7 - Works",
			&types.AttributeValueMemberN{Value: "7"}, Order_RequestAdd),
		Entry("Value is int, 8 - Works",
			&types.AttributeValueMemberN{Value: "8"}, Order_RequestModify),
		Entry("Value is int, 9 - Works",
			&types.AttributeValueMemberN{Value: "9"}, Order_RequestCancel),
		Entry("Value is NULL - Works", new(types.AttributeValueMemberNULL), Order_Status(0)),
		Entry("Value is string, Started - Works",
			&types.AttributeValueMemberS{Value: "Started"}, Order_Started),
		Entry("Value is string, Placed - Works",
			&types.AttributeValueMemberS{Value: "Placed"}, Order_Placed),
		Entry("Value is string, Cancelled - Works",
			&types.AttributeValueMemberS{Value: "Cancelled"}, Order_Cancelled),
		Entry("Value is string, Partial - Works",
			&types.AttributeValueMemberS{Value: "Partial"}, Order_Partial),
		Entry("Value is string, Filled - Works",
			&types.AttributeValueMemberS{Value: "Filled"}, Order_Filled),
		Entry("Value is string, Rejected - Works",
			&types.AttributeValueMemberS{Value: "Rejected"}, Order_Rejected),
		Entry("Value is string, Expired - Works",
			&types.AttributeValueMemberS{Value: "Expired"}, Order_Expired),
		Entry("Value is string, RequestAdd - Works",
			&types.AttributeValueMemberS{Value: "RequestAdd"}, Order_RequestAdd),
		Entry("Value is string, RequestModify - Works",
			&types.AttributeValueMemberS{Value: "RequestModify"}, Order_RequestModify),
		Entry("Value is string, RequestCancel - Works",
			&types.AttributeValueMemberS{Value: "RequestCancel"}, Order_RequestCancel),
		Entry("Value is string, 0 - Works",
			&types.AttributeValueMemberS{Value: "0"}, Order_Started),
		Entry("Value is string, 1 - Works",
			&types.AttributeValueMemberS{Value: "1"}, Order_Placed),
		Entry("Value is string, 2 - Works",
			&types.AttributeValueMemberS{Value: "2"}, Order_Cancelled),
		Entry("Value is string, 3 - Works",
			&types.AttributeValueMemberS{Value: "3"}, Order_Partial),
		Entry("Value is string, 4 - Works",
			&types.AttributeValueMemberS{Value: "4"}, Order_Filled),
		Entry("Value is string, 5 - Works",
			&types.AttributeValueMemberS{Value: "5"}, Order_Rejected),
		Entry("Value is string, 6 - Works",
			&types.AttributeValueMemberS{Value: "6"}, Order_Expired),
		Entry("Value is string, 7 - Works",
			&types.AttributeValueMemberS{Value: "7"}, Order_RequestAdd),
		Entry("Value is string, 8 - Works",
			&types.AttributeValueMemberS{Value: "8"}, Order_RequestModify),
		Entry("Value is string, 9 - Works",
			&types.AttributeValueMemberS{Value: "9"}, Order_RequestCancel))

	// Test that attempting to deserialize a data.Order.Status will fial and return an error if the value
	// cannot be converted to either the name value or integer value of the enum option
	It("Scan - Value is nil - Error", func() {

		// Attempt to convert a fake string value into a data.Order.Status; this should return an error
		var enum *Order_Status
		err := enum.Scan(nil)

		// Verify the error
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of %!q(<nil>) had an invalid type of <nil>"))
		Expect(enum).Should(BeNil())
	})

	// Test the conditions under which values should be convertible to a data.Order.Status
	DescribeTable("Scan Tests",
		func(value interface{}, shouldBe Order_Status) {

			// Attempt to convert the value into a data.Order.Status; this should not fail
			var enum Order_Status
			err := enum.Scan(value)

			// Verify that the deserialization was successful
			Expect(err).ShouldNot(HaveOccurred())
			Expect(enum).Should(Equal(shouldBe))
		},
		Entry("Started - Works", "Started", Order_Started),
		Entry("Placed - Works", "Placed", Order_Placed),
		Entry("Cancelled - Works", "Cancelled", Order_Cancelled),
		Entry("Partial - Works", "Partial", Order_Partial),
		Entry("Filled - Works", "Filled", Order_Filled),
		Entry("Rejected - Works", "Rejected", Order_Rejected),
		Entry("Expired - Works", "Expired", Order_Expired),
		Entry("RequestAdd - Works", "RequestAdd", Order_RequestAdd),
		Entry("RequestModify - Works", "RequestModify", Order_RequestModify),
		Entry("RequestCancel - Works", "RequestCancel", Order_RequestCancel),
		Entry("0 - Works", 0, Order_Started),
		Entry("1 - Works", 1, Order_Placed),
		Entry("2 - Works", 2, Order_Cancelled),
		Entry("3 - Works", 3, Order_Partial),
		Entry("4 - Works", 4, Order_Filled),
		Entry("5 - Works", 5, Order_Rejected),
		Entry("6 - Works", 6, Order_Expired),
		Entry("7 - Works", 7, Order_RequestAdd),
		Entry("8 - Works", 8, Order_RequestModify),
		Entry("9 - Works", 9, Order_RequestCancel))
})

var _ = Describe("data.Order.Reason Marshal/Unmarshal Tests", func() {

	// Test that converting the data.Order.Reason enum to JSON works for all values
	DescribeTable("MarshalJSON Tests",
		func(enum Order_Reason, value string) {
			data, err := json.Marshal(enum)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(string(data)).Should(Equal(value))
		},
		Entry("Client - Works", Order_Client, "\"Client\""),
		Entry("Mobile - Works", Order_Mobile, "\"Mobile\""),
		Entry("Web - Works", Order_Web, "\"Web\""),
		Entry("Strategy - Works", Order_Strategy, "\"Strategy\""),
		Entry("StopLoss - Works", Order_StopLoss, "\"SL\""),
		Entry("TakeProfit - Works", Order_TakeProfit, "\"TP\""),
		Entry("StopOut - Works", Order_StopOut, "\"SO\""))

	// Test that converting the data.Order.Reason enum to a CSV column works for all values
	DescribeTable("MarshalCSV Tests",
		func(enum Order_Reason, value string) {
			data, err := enum.MarshalCSV()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(string(data)).Should(Equal(value))
		},
		Entry("Client - Works", Order_Client, "Client"),
		Entry("Mobile - Works", Order_Mobile, "Mobile"),
		Entry("Web - Works", Order_Web, "Web"),
		Entry("Strategy - Works", Order_Strategy, "Strategy"),
		Entry("StopLoss - Works", Order_StopLoss, "SL"),
		Entry("TakeProfit - Works", Order_TakeProfit, "TP"),
		Entry("StopOut - Works", Order_StopOut, "SO"))

	// Test that converting the data.Order.Reason enum to a YAML node works for all values
	DescribeTable("MarshalYAML - Works",
		func(enum Order_Reason, value string) {
			data, err := enum.MarshalYAML()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(data).Should(Equal(value))
		},
		Entry("Client - Works", Order_Client, "Client"),
		Entry("Mobile - Works", Order_Mobile, "Mobile"),
		Entry("Web - Works", Order_Web, "Web"),
		Entry("Strategy - Works", Order_Strategy, "Strategy"),
		Entry("StopLoss - Works", Order_StopLoss, "SL"),
		Entry("TakeProfit - Works", Order_TakeProfit, "TP"),
		Entry("StopOut - Works", Order_StopOut, "SO"))

	// Test that converting the data.Order.Reason enum to a DynamoDB AttributeVAlue works for all values
	DescribeTable("MarshalDynamoDBAttributeValue - Works",
		func(enum Order_Reason, value string) {
			data, err := attributevalue.Marshal(enum)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(data.(*types.AttributeValueMemberS).Value).Should(Equal(value))
		},
		Entry("Client - Works", Order_Client, "Client"),
		Entry("Mobile - Works", Order_Mobile, "Mobile"),
		Entry("Web - Works", Order_Web, "Web"),
		Entry("Strategy - Works", Order_Strategy, "Strategy"),
		Entry("StopLoss - Works", Order_StopLoss, "SL"),
		Entry("TakeProfit - Works", Order_TakeProfit, "TP"),
		Entry("StopOut - Works", Order_StopOut, "SO"))

	// Test that converting the data.Order.Reason enum to an SQL value for all values
	DescribeTable("Value Tests",
		func(enum Order_Reason, value string) {
			data, err := enum.Value()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(data).Should(Equal(value))
		},
		Entry("Client - Works", Order_Client, "Client"),
		Entry("Mobile - Works", Order_Mobile, "Mobile"),
		Entry("Web - Works", Order_Web, "Web"),
		Entry("Strategy - Works", Order_Strategy, "Strategy"),
		Entry("StopLoss - Works", Order_StopLoss, "SL"),
		Entry("TakeProfit - Works", Order_TakeProfit, "TP"),
		Entry("StopOut - Works", Order_StopOut, "SO"))

	// Test that attempting to deserialize a data.Order.Reason will fail and return an error if the value
	// cannot be deserialized from a JSON value to a string
	It("UnmarshalJSON fails - Error", func() {

		// Attempt to convert a non-parseable string value into a data.Order.Reason; this should return an error
		enum := new(Order_Reason)
		err := enum.UnmarshalJSON([]byte("derp"))

		// Verify the error
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of \"derp\" cannot be mapped to a data.Order_Reason"))
	})

	// Test that attempting to deserialize a data.Order.Reason will fail and return an error if the value
	// cannot be converted to either the name value or integer value of the enum option
	It("UnmarshalJSON - Value is invalid - Error", func() {

		// Attempt to convert a fake string value into a data.Order.Reason; this should return an error
		enum := new(Order_Reason)
		err := enum.UnmarshalJSON([]byte("\"derp\""))

		// Verify the error
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of \"derp\" cannot be mapped to a data.Order_Reason"))
	})

	// Test the conditions under which values should be convertible to a data.Order.Reason
	DescribeTable("UnmarshalJSON Tests",
		func(value string, shouldBe Order_Reason) {

			// Attempt to convert the string value into a data.Order.Reason; this should not fail
			var enum Order_Reason
			err := enum.UnmarshalJSON([]byte(value))

			// Verify that the deserialization was successful
			Expect(err).ShouldNot(HaveOccurred())
			Expect(enum).Should(Equal(shouldBe))
		},
		Entry("SL - Works", "\"SL\"", Order_StopLoss),
		Entry("TP - Works", "\"TP\"", Order_TakeProfit),
		Entry("SO - Works", "\"SO\"", Order_StopOut),
		Entry("Client - Works", "\"Client\"", Order_Client),
		Entry("Mobile - Works", "\"Mobile\"", Order_Mobile),
		Entry("Web - Works", "\"Web\"", Order_Web),
		Entry("Strategy - Works", "\"Strategy\"", Order_Strategy),
		Entry("StopLoss - Works", "\"StopLoss\"", Order_StopLoss),
		Entry("TakeProfit - Works", "\"TakeProfit\"", Order_TakeProfit),
		Entry("StopOut - Works", "\"StopOut\"", Order_StopOut),
		Entry("0 - Works", "\"0\"", Order_Client),
		Entry("1 - Works", "\"1\"", Order_Mobile),
		Entry("2 - Works", "\"2\"", Order_Web),
		Entry("3 - Works", "\"3\"", Order_Strategy),
		Entry("4 - Works", "\"4\"", Order_StopLoss),
		Entry("5 - Works", "\"5\"", Order_TakeProfit),
		Entry("6 - Works", "\"6\"", Order_StopOut))

	// Test that attempting to deserialize a data.Order.Reason will fail and return an error if the value
	// cannot be converted to either the name value or integer value of the enum option
	It("UnmarshalCSV - Value is empty - Error", func() {

		// Attempt to convert a fake string value into a data.Order.Reason; this should return an error
		enum := new(Order_Reason)
		err := enum.UnmarshalCSV("")

		// Verify the error
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of \"\" cannot be mapped to a data.Order_Reason"))
	})

	// Test the conditions under which values should be convertible to a data.Order.Reason
	DescribeTable("UnmarshalCSV Tests",
		func(value string, shouldBe Order_Reason) {

			// Attempt to convert the value into a data.Order.Reason; this should not fail
			var enum Order_Reason
			err := enum.UnmarshalCSV(value)

			// Verify that the deserialization was successful
			Expect(err).ShouldNot(HaveOccurred())
			Expect(enum).Should(Equal(shouldBe))
		},
		Entry("SL - Works", "SL", Order_StopLoss),
		Entry("TP - Works", "TP", Order_TakeProfit),
		Entry("SO - Works", "SO", Order_StopOut),
		Entry("Client - Works", "Client", Order_Client),
		Entry("Mobile - Works", "Mobile", Order_Mobile),
		Entry("Web - Works", "Web", Order_Web),
		Entry("Strategy - Works", "Strategy", Order_Strategy),
		Entry("StopLoss - Works", "StopLoss", Order_StopLoss),
		Entry("TakeProfit - Works", "TakeProfit", Order_TakeProfit),
		Entry("StopOut - Works", "StopOut", Order_StopOut),
		Entry("0 - Works", "0", Order_Client),
		Entry("1 - Works", "1", Order_Mobile),
		Entry("2 - Works", "2", Order_Web),
		Entry("3 - Works", "3", Order_Strategy),
		Entry("4 - Works", "4", Order_StopLoss),
		Entry("5 - Works", "5", Order_TakeProfit),
		Entry("6 - Works", "6", Order_StopOut))

	// Test that attempting to deserialize a data.Order.Reason will fail and return an error if the YAML
	// node does not represent a scalar value
	It("UnmarshalYAML - Node type is not scalar - Error", func() {
		enum := new(Order_Reason)
		err := enum.UnmarshalYAML(&yaml.Node{Kind: yaml.AliasNode})
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("YAML node had an invalid kind (expected scalar value)"))
	})

	// Test that attempting to deserialize a data.Order.Reason will fail and return an error if the YAML
	// node value cannot be converted to either the name value or integer value of the enum option
	It("UnmarshalYAML - Parse fails - Error", func() {
		enum := new(Order_Reason)
		err := enum.UnmarshalYAML(&yaml.Node{Kind: yaml.ScalarNode, Value: "derp"})
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of \"derp\" cannot be mapped to a data.Order_Reason"))
	})

	// Test the conditions under which YAML node values should be convertible to a data.Order.Reason
	DescribeTable("UnmarshalYAML Tests",
		func(value string, shouldBe Order_Reason) {
			var enum Order_Reason
			err := enum.UnmarshalYAML(&yaml.Node{Kind: yaml.ScalarNode, Value: value})
			Expect(err).ShouldNot(HaveOccurred())
			Expect(enum).Should(Equal(shouldBe))
		},
		Entry("SL - Works", "SL", Order_StopLoss),
		Entry("TP - Works", "TP", Order_TakeProfit),
		Entry("SO - Works", "SO", Order_StopOut),
		Entry("Client - Works", "Client", Order_Client),
		Entry("Mobile - Works", "Mobile", Order_Mobile),
		Entry("Web - Works", "Web", Order_Web),
		Entry("Strategy - Works", "Strategy", Order_Strategy),
		Entry("StopLoss - Works", "StopLoss", Order_StopLoss),
		Entry("TakeProfit - Works", "TakeProfit", Order_TakeProfit),
		Entry("StopOut - Works", "StopOut", Order_StopOut),
		Entry("0 - Works", "0", Order_Client),
		Entry("1 - Works", "1", Order_Mobile),
		Entry("2 - Works", "2", Order_Web),
		Entry("3 - Works", "3", Order_Strategy),
		Entry("4 - Works", "4", Order_StopLoss),
		Entry("5 - Works", "5", Order_TakeProfit),
		Entry("6 - Works", "6", Order_StopOut))

	// Tests that, if the attribute type submitted to UnmarshalDynamoDBAttributeValue is not one we
	// recognize, then the function will return an error
	It("UnmarshalDynamoDBAttributeValue - AttributeValue type invalid - Error", func() {
		enum := new(Order_Reason)
		err := attributevalue.Unmarshal(&types.AttributeValueMemberBOOL{Value: true}, &enum)
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("Attribute value of *types.AttributeValueMemberBOOL could not be converted to a data.Order.Reason"))
	})

	// Tests that, if time parsing fails, then calling UnmarshalDynamoDBAttributeValue will return an error
	It("UnmarshalDynamoDBAttributeValue - Parse fails - Error", func() {
		enum := new(Order_Reason)
		err := attributevalue.Unmarshal(&types.AttributeValueMemberS{Value: "derp"}, &enum)
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of \"derp\" cannot be mapped to a data.Order_Reason"))
	})

	// Tests the conditions under which UnmarshalDynamoDBAttributeValue is called and no error is generated
	DescribeTable("UnmarshalDynamoDBAttributeValue - AttributeValue Conditions",
		func(value types.AttributeValue, expected Order_Reason) {
			var enum Order_Reason
			err := attributevalue.Unmarshal(value, &enum)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(enum).Should(Equal(expected))
		},
		Entry("Value is []bytes, SL - Works",
			&types.AttributeValueMemberB{Value: []byte("SL")}, Order_StopLoss),
		Entry("Value is []bytes, TP - Works",
			&types.AttributeValueMemberB{Value: []byte("TP")}, Order_TakeProfit),
		Entry("Value is []bytes, SO - Works",
			&types.AttributeValueMemberB{Value: []byte("SO")}, Order_StopOut),
		Entry("Value is []bytes, Client - Works",
			&types.AttributeValueMemberB{Value: []byte("Client")}, Order_Client),
		Entry("Value is []bytes, Mobile - Works",
			&types.AttributeValueMemberB{Value: []byte("Mobile")}, Order_Mobile),
		Entry("Value is []bytes, Web - Works",
			&types.AttributeValueMemberB{Value: []byte("Web")}, Order_Web),
		Entry("Value is []bytes, Strategy - Works",
			&types.AttributeValueMemberB{Value: []byte("Strategy")}, Order_Strategy),
		Entry("Value is []bytes, StopLoss - Works",
			&types.AttributeValueMemberB{Value: []byte("StopLoss")}, Order_StopLoss),
		Entry("Value is []bytes, TakeProfit - Works",
			&types.AttributeValueMemberB{Value: []byte("TakeProfit")}, Order_TakeProfit),
		Entry("Value is []bytes, StopOut - Works",
			&types.AttributeValueMemberB{Value: []byte("StopOut")}, Order_StopOut),
		Entry("Value is []bytes, 0 - Works",
			&types.AttributeValueMemberB{Value: []byte("0")}, Order_Client),
		Entry("Value is []bytes, 1 - Works",
			&types.AttributeValueMemberB{Value: []byte("1")}, Order_Mobile),
		Entry("Value is []bytes, 2 - Works",
			&types.AttributeValueMemberB{Value: []byte("2")}, Order_Web),
		Entry("Value is []bytes, 3 - Works",
			&types.AttributeValueMemberB{Value: []byte("3")}, Order_Strategy),
		Entry("Value is []bytes, 4 - Works",
			&types.AttributeValueMemberB{Value: []byte("4")}, Order_StopLoss),
		Entry("Value is []bytes, 5 - Works",
			&types.AttributeValueMemberB{Value: []byte("5")}, Order_TakeProfit),
		Entry("Value is []bytes, 6 - Works",
			&types.AttributeValueMemberB{Value: []byte("6")}, Order_StopOut),
		Entry("Value is int, 0 - Works",
			&types.AttributeValueMemberN{Value: "0"}, Order_Client),
		Entry("Value is int, 1 - Works",
			&types.AttributeValueMemberN{Value: "1"}, Order_Mobile),
		Entry("Value is int, 2 - Works",
			&types.AttributeValueMemberN{Value: "2"}, Order_Web),
		Entry("Value is int, 3 - Works",
			&types.AttributeValueMemberN{Value: "3"}, Order_Strategy),
		Entry("Value is int, 4 - Works",
			&types.AttributeValueMemberN{Value: "4"}, Order_StopLoss),
		Entry("Value is int, 5 - Works",
			&types.AttributeValueMemberN{Value: "5"}, Order_TakeProfit),
		Entry("Value is int, 6 - Works",
			&types.AttributeValueMemberN{Value: "6"}, Order_StopOut),
		Entry("Value is NULL - Works", new(types.AttributeValueMemberNULL), Order_Reason(0)),
		Entry("Value is string, SL - Works",
			&types.AttributeValueMemberS{Value: "SL"}, Order_StopLoss),
		Entry("Value is string, TP - Works",
			&types.AttributeValueMemberS{Value: "TP"}, Order_TakeProfit),
		Entry("Value is string, SO - Works",
			&types.AttributeValueMemberS{Value: "SO"}, Order_StopOut),
		Entry("Value is string, Client - Works",
			&types.AttributeValueMemberS{Value: "Client"}, Order_Client),
		Entry("Value is string, Mobile - Works",
			&types.AttributeValueMemberS{Value: "Mobile"}, Order_Mobile),
		Entry("Value is string, Web - Works",
			&types.AttributeValueMemberS{Value: "Web"}, Order_Web),
		Entry("Value is string, Strategy - Works",
			&types.AttributeValueMemberS{Value: "Strategy"}, Order_Strategy),
		Entry("Value is string, StopLoss - Works",
			&types.AttributeValueMemberS{Value: "StopLoss"}, Order_StopLoss),
		Entry("Value is string, TakeProfit - Works",
			&types.AttributeValueMemberS{Value: "TakeProfit"}, Order_TakeProfit),
		Entry("Value is string, StopOut - Works",
			&types.AttributeValueMemberS{Value: "StopOut"}, Order_StopOut),
		Entry("Value is string, 0 - Works",
			&types.AttributeValueMemberS{Value: "0"}, Order_Client),
		Entry("Value is string, 1 - Works",
			&types.AttributeValueMemberS{Value: "1"}, Order_Mobile),
		Entry("Value is string, 2 - Works",
			&types.AttributeValueMemberS{Value: "2"}, Order_Web),
		Entry("Value is string, 3 - Works",
			&types.AttributeValueMemberS{Value: "3"}, Order_Strategy),
		Entry("Value is string, 4 - Works",
			&types.AttributeValueMemberS{Value: "4"}, Order_StopLoss),
		Entry("Value is string, 5 - Works",
			&types.AttributeValueMemberS{Value: "5"}, Order_TakeProfit),
		Entry("Value is string, 6 - Works",
			&types.AttributeValueMemberS{Value: "6"}, Order_StopOut))

	// Test that attempting to deserialize a data.Order.Reason will fial and return an error if the value
	// cannot be converted to either the name value or integer value of the enum option
	It("Scan - Value is nil - Error", func() {

		// Attempt to convert a fake string value into a data.Order.Reason; this should return an error
		var enum *Order_Reason
		err := enum.Scan(nil)

		// Verify the error
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of %!q(<nil>) had an invalid type of <nil>"))
		Expect(enum).Should(BeNil())
	})

	// Test the conditions under which values should be convertible to a data.Order.Reason
	DescribeTable("Scan Tests",
		func(value interface{}, shouldBe Order_Reason) {

			// Attempt to convert the value into a data.Order.Reason; this should not fail
			var enum Order_Reason
			err := enum.Scan(value)

			// Verify that the deserialization was successful
			Expect(err).ShouldNot(HaveOccurred())
			Expect(enum).Should(Equal(shouldBe))
		},
		Entry("SL - Works", "SL", Order_StopLoss),
		Entry("TP - Works", "TP", Order_TakeProfit),
		Entry("SO - Works", "SO", Order_StopOut),
		Entry("Client - Works", "Client", Order_Client),
		Entry("Mobile - Works", "Mobile", Order_Mobile),
		Entry("Web - Works", "Web", Order_Web),
		Entry("Strategy - Works", "Strategy", Order_Strategy),
		Entry("StopLoss - Works", "StopLoss", Order_StopLoss),
		Entry("TakeProfit - Works", "TakeProfit", Order_TakeProfit),
		Entry("StopOut - Works", "StopOut", Order_StopOut),
		Entry("0 - Works", 0, Order_Client),
		Entry("1 - Works", 1, Order_Mobile),
		Entry("2 - Works", 2, Order_Web),
		Entry("3 - Works", 3, Order_Strategy),
		Entry("4 - Works", 4, Order_StopLoss),
		Entry("5 - Works", 5, Order_TakeProfit),
		Entry("6 - Works", 6, Order_StopOut))
})

var _ = Describe("data.Position.Type Marshal/Unmarshal Tests", func() {

	// Test that converting the data.Position.Type enum to JSON works for all values
	DescribeTable("MarshalJSON Tests",
		func(enum Position_Type, value string) {
			data, err := json.Marshal(enum)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(string(data)).Should(Equal(value))
		},
		Entry("Buy - Works", Position_Buy, "\"Buy\""),
		Entry("Sell - Works", Position_Sell, "\"Sell\""))

	// Test that converting the data.Position.Type enum to a CSV column works for all values
	DescribeTable("MarshalCSV Tests",
		func(enum Position_Type, value string) {
			data, err := enum.MarshalCSV()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(string(data)).Should(Equal(value))
		},
		Entry("Buy - Works", Position_Buy, "Buy"),
		Entry("Sell - Works", Position_Sell, "Sell"))

	// Test that converting the data.Position.Type enum to a YAML node works for all values
	DescribeTable("MarshalYAML - Works",
		func(enum Position_Type, value string) {
			data, err := enum.MarshalYAML()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(data).Should(Equal(value))
		},
		Entry("Buy - Works", Position_Buy, "Buy"),
		Entry("Sell - Works", Position_Sell, "Sell"))

	// Test that converting the data.Position.Type enum to a DynamoDB AttributeVAlue works for all values
	DescribeTable("MarshalDynamoDBAttributeValue - Works",
		func(enum Position_Type, value string) {
			data, err := attributevalue.Marshal(enum)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(data.(*types.AttributeValueMemberS).Value).Should(Equal(value))
		},
		Entry("Buy - Works", Position_Buy, "Buy"),
		Entry("Sell - Works", Position_Sell, "Sell"))

	// Test that converting the data.Position.Type enum to an SQL value for all values
	DescribeTable("Value Tests",
		func(enum Position_Type, value string) {
			data, err := enum.Value()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(data).Should(Equal(value))
		},
		Entry("Buy - Works", Position_Buy, "Buy"),
		Entry("Sell - Works", Position_Sell, "Sell"))

	// Test that attempting to deserialize a data.Position.Type will fail and return an error if the value
	// cannot be deserialized from a JSON value to a string
	It("UnmarshalJSON fails - Error", func() {

		// Attempt to convert a non-parseable string value into a data.Position.Type; this should return an error
		enum := new(Position_Type)
		err := enum.UnmarshalJSON([]byte("derp"))

		// Verify the error
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of \"derp\" cannot be mapped to a data.Position_Type"))
	})

	// Test that attempting to deserialize a data.Position.Type will fail and return an error if the value
	// cannot be converted to either the name value or integer value of the enum option
	It("UnmarshalJSON - Value is invalid - Error", func() {

		// Attempt to convert a fake string value into a data.Position.Type; this should return an error
		enum := new(Position_Type)
		err := enum.UnmarshalJSON([]byte("\"derp\""))

		// Verify the error
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of \"derp\" cannot be mapped to a data.Position_Type"))
	})

	// Test the conditions under which values should be convertible to a data.Position.Type
	DescribeTable("UnmarshalJSON Tests",
		func(value string, shouldBe Position_Type) {

			// Attempt to convert the string value into a data.Position.Type; this should not fail
			var enum Position_Type
			err := enum.UnmarshalJSON([]byte(value))

			// Verify that the deserialization was successful
			Expect(err).ShouldNot(HaveOccurred())
			Expect(enum).Should(Equal(shouldBe))
		},
		Entry("Buy - Works", "\"Buy\"", Position_Buy),
		Entry("Sell - Works", "\"Sell\"", Position_Sell),
		Entry("0 - Works", "\"0\"", Position_Buy),
		Entry("1 - Works", "\"1\"", Position_Sell))

	// Test that attempting to deserialize a data.Position.Type will fail and return an error if the value
	// cannot be converted to either the name value or integer value of the enum option
	It("UnmarshalCSV - Value is empty - Error", func() {

		// Attempt to convert a fake string value into a data.Position.Type; this should return an error
		enum := new(Position_Type)
		err := enum.UnmarshalCSV("")

		// Verify the error
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of \"\" cannot be mapped to a data.Position_Type"))
	})

	// Test the conditions under which values should be convertible to a data.Position.Type
	DescribeTable("UnmarshalCSV Tests",
		func(value string, shouldBe Position_Type) {

			// Attempt to convert the value into a data.Position.Type; this should not fail
			var enum Position_Type
			err := enum.UnmarshalCSV(value)

			// Verify that the deserialization was successful
			Expect(err).ShouldNot(HaveOccurred())
			Expect(enum).Should(Equal(shouldBe))
		},
		Entry("Buy - Works", "Buy", Position_Buy),
		Entry("Sell - Works", "Sell", Position_Sell),
		Entry("0 - Works", "0", Position_Buy),
		Entry("1 - Works", "1", Position_Sell))

	// Test that attempting to deserialize a data.Position.Type will fail and return an error if the YAML
	// node does not represent a scalar value
	It("UnmarshalYAML - Node type is not scalar - Error", func() {
		enum := new(Position_Type)
		err := enum.UnmarshalYAML(&yaml.Node{Kind: yaml.AliasNode})
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("YAML node had an invalid kind (expected scalar value)"))
	})

	// Test that attempting to deserialize a data.Position.Type will fail and return an error if the YAML
	// node value cannot be converted to either the name value or integer value of the enum option
	It("UnmarshalYAML - Parse fails - Error", func() {
		enum := new(Position_Type)
		err := enum.UnmarshalYAML(&yaml.Node{Kind: yaml.ScalarNode, Value: "derp"})
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of \"derp\" cannot be mapped to a data.Position_Type"))
	})

	// Test the conditions under which YAML node values should be convertible to a data.Position.Type
	DescribeTable("UnmarshalYAML Tests",
		func(value string, shouldBe Position_Type) {
			var enum Position_Type
			err := enum.UnmarshalYAML(&yaml.Node{Kind: yaml.ScalarNode, Value: value})
			Expect(err).ShouldNot(HaveOccurred())
			Expect(enum).Should(Equal(shouldBe))
		},
		Entry("Buy - Works", "Buy", Position_Buy),
		Entry("Sell - Works", "Sell", Position_Sell),
		Entry("0 - Works", "0", Position_Buy),
		Entry("1 - Works", "1", Position_Sell))

	// Tests that, if the attribute type submitted to UnmarshalDynamoDBAttributeValue is not one we
	// recognize, then the function will return an error
	It("UnmarshalDynamoDBAttributeValue - AttributeValue type invalid - Error", func() {
		enum := new(Position_Type)
		err := attributevalue.Unmarshal(&types.AttributeValueMemberBOOL{Value: true}, &enum)
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("Attribute value of *types.AttributeValueMemberBOOL could not be converted to a data.Position.Type"))
	})

	// Tests that, if time parsing fails, then calling UnmarshalDynamoDBAttributeValue will return an error
	It("UnmarshalDynamoDBAttributeValue - Parse fails - Error", func() {
		enum := new(Position_Type)
		err := attributevalue.Unmarshal(&types.AttributeValueMemberS{Value: "derp"}, &enum)
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of \"derp\" cannot be mapped to a data.Position_Type"))
	})

	// Tests the conditions under which UnmarshalDynamoDBAttributeValue is called and no error is generated
	DescribeTable("UnmarshalDynamoDBAttributeValue - AttributeValue Conditions",
		func(value types.AttributeValue, expected Position_Type) {
			var enum Position_Type
			err := attributevalue.Unmarshal(value, &enum)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(enum).Should(Equal(expected))
		},
		Entry("Value is []bytes, Buy - Works",
			&types.AttributeValueMemberB{Value: []byte("Buy")}, Position_Buy),
		Entry("Value is []bytes, Sell - Works",
			&types.AttributeValueMemberB{Value: []byte("Sell")}, Position_Sell),
		Entry("Value is []bytes, 0 - Works",
			&types.AttributeValueMemberB{Value: []byte("0")}, Position_Buy),
		Entry("Value is []bytes, 1 - Works",
			&types.AttributeValueMemberB{Value: []byte("1")}, Position_Sell),
		Entry("Value is int, 0 - Works",
			&types.AttributeValueMemberN{Value: "0"}, Position_Buy),
		Entry("Value is int, 1 - Works",
			&types.AttributeValueMemberN{Value: "1"}, Position_Sell),
		Entry("Value is NULL - Works", new(types.AttributeValueMemberNULL), Position_Type(0)),
		Entry("Value is string, Buy - Works",
			&types.AttributeValueMemberS{Value: "Buy"}, Position_Buy),
		Entry("Value is string, Sell - Works",
			&types.AttributeValueMemberS{Value: "Sell"}, Position_Sell),
		Entry("Value is string, 0 - Works",
			&types.AttributeValueMemberS{Value: "0"}, Position_Buy),
		Entry("Value is string, 1 - Works",
			&types.AttributeValueMemberS{Value: "1"}, Position_Sell))

	// Test that attempting to deserialize a data.Position.Type will fial and return an error if the value
	// cannot be converted to either the name value or integer value of the enum option
	It("Scan - Value is nil - Error", func() {

		// Attempt to convert a fake string value into a data.Position.Type; this should return an error
		var enum *Position_Type
		err := enum.Scan(nil)

		// Verify the error
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of %!q(<nil>) had an invalid type of <nil>"))
		Expect(enum).Should(BeNil())
	})

	// Test the conditions under which values should be convertible to a data.Position.Type
	DescribeTable("Scan Tests",
		func(value interface{}, shouldBe Position_Type) {

			// Attempt to convert the value into a data.Position.Type; this should not fail
			var enum Position_Type
			err := enum.Scan(value)

			// Verify that the deserialization was successful
			Expect(err).ShouldNot(HaveOccurred())
			Expect(enum).Should(Equal(shouldBe))
		},
		Entry("Buy - Works", "Buy", Position_Buy),
		Entry("Sell - Works", "Sell", Position_Sell),
		Entry("0 - Works", 0, Position_Buy),
		Entry("1 - Works", 1, Position_Sell))
})

var _ = Describe("data.Position.Reason Marshal/Unmarshal Tests", func() {

	// Test that converting the data.Position.Reason enum to JSON works for all values
	DescribeTable("MarshalJSON Tests",
		func(enum Position_Reason, value string) {
			data, err := json.Marshal(enum)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(string(data)).Should(Equal(value))
		},
		Entry("Client - Works", Position_Client, "\"Client\""),
		Entry("Mobile - Works", Position_Mobile, "\"Mobile\""),
		Entry("Web - Works", Position_Web, "\"Web\""),
		Entry("Strategy - Works", Position_Strategy, "\"Strategy\""))

	// Test that converting the data.Position.Reason enum to a CSV column works for all values
	DescribeTable("MarshalCSV Tests",
		func(enum Position_Reason, value string) {
			data, err := enum.MarshalCSV()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(string(data)).Should(Equal(value))
		},
		Entry("Client - Works", Position_Client, "Client"),
		Entry("Mobile - Works", Position_Mobile, "Mobile"),
		Entry("Web - Works", Position_Web, "Web"),
		Entry("Strategy - Works", Position_Strategy, "Strategy"))

	// Test that converting the data.Position.Reason enum to a YAML node works for all values
	DescribeTable("MarshalYAML - Works",
		func(enum Position_Reason, value string) {
			data, err := enum.MarshalYAML()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(data).Should(Equal(value))
		},
		Entry("Client - Works", Position_Client, "Client"),
		Entry("Mobile - Works", Position_Mobile, "Mobile"),
		Entry("Web - Works", Position_Web, "Web"),
		Entry("Strategy - Works", Position_Strategy, "Strategy"))

	// Test that converting the data.Position.Reason enum to a DynamoDB AttributeVAlue works for all values
	DescribeTable("MarshalDynamoDBAttributeValue - Works",
		func(enum Position_Reason, value string) {
			data, err := attributevalue.Marshal(enum)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(data.(*types.AttributeValueMemberS).Value).Should(Equal(value))
		},
		Entry("Client - Works", Position_Client, "Client"),
		Entry("Mobile - Works", Position_Mobile, "Mobile"),
		Entry("Web - Works", Position_Web, "Web"),
		Entry("Strategy - Works", Position_Strategy, "Strategy"))

	// Test that converting the data.Position.Reason enum to an SQL value for all values
	DescribeTable("Value Tests",
		func(enum Position_Reason, value string) {
			data, err := enum.Value()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(data).Should(Equal(value))
		},
		Entry("Client - Works", Position_Client, "Client"),
		Entry("Mobile - Works", Position_Mobile, "Mobile"),
		Entry("Web - Works", Position_Web, "Web"),
		Entry("Strategy - Works", Position_Strategy, "Strategy"))

	// Test that attempting to deserialize a data.Position.Reason will fail and return an error if the value
	// cannot be deserialized from a JSON value to a string
	It("UnmarshalJSON fails - Error", func() {

		// Attempt to convert a non-parseable string value into a data.Position.Reason; this should return an error
		enum := new(Position_Reason)
		err := enum.UnmarshalJSON([]byte("derp"))

		// Verify the error
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of \"derp\" cannot be mapped to a data.Position_Reason"))
	})

	// Test that attempting to deserialize a data.Position.Reason will fail and return an error if the value
	// cannot be converted to either the name value or integer value of the enum option
	It("UnmarshalJSON - Value is invalid - Error", func() {

		// Attempt to convert a fake string value into a data.Position.Reason; this should return an error
		enum := new(Position_Reason)
		err := enum.UnmarshalJSON([]byte("\"derp\""))

		// Verify the error
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of \"derp\" cannot be mapped to a data.Position_Reason"))
	})

	// Test the conditions under which values should be convertible to a data.Position.Reason
	DescribeTable("UnmarshalJSON Tests",
		func(value string, shouldBe Position_Reason) {

			// Attempt to convert the string value into a data.Position.Reason; this should not fail
			var enum Position_Reason
			err := enum.UnmarshalJSON([]byte(value))

			// Verify that the deserialization was successful
			Expect(err).ShouldNot(HaveOccurred())
			Expect(enum).Should(Equal(shouldBe))
		},
		Entry("Client - Works", "\"Client\"", Position_Client),
		Entry("Mobile - Works", "\"Mobile\"", Position_Mobile),
		Entry("Web - Works", "\"Web\"", Position_Web),
		Entry("Strategy - Works", "\"Strategy\"", Position_Strategy),
		Entry("0 - Works", "\"0\"", Position_Client),
		Entry("1 - Works", "\"1\"", Position_Mobile),
		Entry("2 - Works", "\"2\"", Position_Web),
		Entry("3 - Works", "\"3\"", Position_Strategy))

	// Test that attempting to deserialize a data.Position.Reason will fail and return an error if the value
	// cannot be converted to either the name value or integer value of the enum option
	It("UnmarshalCSV - Value is empty - Error", func() {

		// Attempt to convert a fake string value into a data.Position.Reason; this should return an error
		enum := new(Position_Reason)
		err := enum.UnmarshalCSV("")

		// Verify the error
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of \"\" cannot be mapped to a data.Position_Reason"))
	})

	// Test the conditions under which values should be convertible to a data.Position.Reason
	DescribeTable("UnmarshalCSV Tests",
		func(value string, shouldBe Position_Reason) {

			// Attempt to convert the value into a data.Position.Reason; this should not fail
			var enum Position_Reason
			err := enum.UnmarshalCSV(value)

			// Verify that the deserialization was successful
			Expect(err).ShouldNot(HaveOccurred())
			Expect(enum).Should(Equal(shouldBe))
		},
		Entry("Client - Works", "Client", Position_Client),
		Entry("Mobile - Works", "Mobile", Position_Mobile),
		Entry("Web - Works", "Web", Position_Web),
		Entry("Strategy - Works", "Strategy", Position_Strategy),
		Entry("0 - Works", "0", Position_Client),
		Entry("1 - Works", "1", Position_Mobile),
		Entry("2 - Works", "2", Position_Web),
		Entry("3 - Works", "3", Position_Strategy))

	// Test that attempting to deserialize a data.Position.Reason will fail and return an error if the YAML
	// node does not represent a scalar value
	It("UnmarshalYAML - Node type is not scalar - Error", func() {
		enum := new(Position_Reason)
		err := enum.UnmarshalYAML(&yaml.Node{Kind: yaml.AliasNode})
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("YAML node had an invalid kind (expected scalar value)"))
	})

	// Test that attempting to deserialize a data.Position.Reason will fail and return an error if the YAML
	// node value cannot be converted to either the name value or integer value of the enum option
	It("UnmarshalYAML - Parse fails - Error", func() {
		enum := new(Position_Reason)
		err := enum.UnmarshalYAML(&yaml.Node{Kind: yaml.ScalarNode, Value: "derp"})
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of \"derp\" cannot be mapped to a data.Position_Reason"))
	})

	// Test the conditions under which YAML node values should be convertible to a data.Position.Reason
	DescribeTable("UnmarshalYAML Tests",
		func(value string, shouldBe Position_Reason) {
			var enum Position_Reason
			err := enum.UnmarshalYAML(&yaml.Node{Kind: yaml.ScalarNode, Value: value})
			Expect(err).ShouldNot(HaveOccurred())
			Expect(enum).Should(Equal(shouldBe))
		},
		Entry("Client - Works", "Client", Position_Client),
		Entry("Mobile - Works", "Mobile", Position_Mobile),
		Entry("Web - Works", "Web", Position_Web),
		Entry("Strategy - Works", "Strategy", Position_Strategy),
		Entry("0 - Works", "0", Position_Client),
		Entry("1 - Works", "1", Position_Mobile),
		Entry("2 - Works", "2", Position_Web),
		Entry("3 - Works", "3", Position_Strategy))

	// Tests that, if the attribute type submitted to UnmarshalDynamoDBAttributeValue is not one we
	// recognize, then the function will return an error
	It("UnmarshalDynamoDBAttributeValue - AttributeValue type invalid - Error", func() {
		enum := new(Position_Reason)
		err := attributevalue.Unmarshal(&types.AttributeValueMemberBOOL{Value: true}, &enum)
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("Attribute value of *types.AttributeValueMemberBOOL could not be converted to a data.Position.Reason"))
	})

	// Tests that, if time parsing fails, then calling UnmarshalDynamoDBAttributeValue will return an error
	It("UnmarshalDynamoDBAttributeValue - Parse fails - Error", func() {
		enum := new(Position_Reason)
		err := attributevalue.Unmarshal(&types.AttributeValueMemberS{Value: "derp"}, &enum)
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of \"derp\" cannot be mapped to a data.Position_Reason"))
	})

	// Tests the conditions under which UnmarshalDynamoDBAttributeValue is called and no error is generated
	DescribeTable("UnmarshalDynamoDBAttributeValue - AttributeValue Conditions",
		func(value types.AttributeValue, expected Position_Reason) {
			var enum Position_Reason
			err := attributevalue.Unmarshal(value, &enum)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(enum).Should(Equal(expected))
		},
		Entry("Value is []bytes, Client - Works",
			&types.AttributeValueMemberB{Value: []byte("Client")}, Position_Client),
		Entry("Value is []bytes, Mobile - Works",
			&types.AttributeValueMemberB{Value: []byte("Mobile")}, Position_Mobile),
		Entry("Value is []bytes, Web - Works",
			&types.AttributeValueMemberB{Value: []byte("Web")}, Position_Web),
		Entry("Value is []bytes, Strategy - Works",
			&types.AttributeValueMemberB{Value: []byte("Strategy")}, Position_Strategy),
		Entry("Value is []bytes, 0 - Works",
			&types.AttributeValueMemberB{Value: []byte("0")}, Position_Client),
		Entry("Value is []bytes, 1 - Works",
			&types.AttributeValueMemberB{Value: []byte("1")}, Position_Mobile),
		Entry("Value is []bytes, 2 - Works",
			&types.AttributeValueMemberB{Value: []byte("2")}, Position_Web),
		Entry("Value is []bytes, 3 - Works",
			&types.AttributeValueMemberB{Value: []byte("3")}, Position_Strategy),
		Entry("Value is int, 0 - Works",
			&types.AttributeValueMemberN{Value: "0"}, Position_Client),
		Entry("Value is int, 1 - Works",
			&types.AttributeValueMemberN{Value: "1"}, Position_Mobile),
		Entry("Value is int, 2 - Works",
			&types.AttributeValueMemberN{Value: "2"}, Position_Web),
		Entry("Value is int, 3 - Works",
			&types.AttributeValueMemberN{Value: "3"}, Position_Strategy),
		Entry("Value is NULL - Works", new(types.AttributeValueMemberNULL), Position_Reason(0)),
		Entry("Value is string, Client - Works",
			&types.AttributeValueMemberS{Value: "Client"}, Position_Client),
		Entry("Value is string, Mobile - Works",
			&types.AttributeValueMemberS{Value: "Mobile"}, Position_Mobile),
		Entry("Value is string, Web - Works",
			&types.AttributeValueMemberS{Value: "Web"}, Position_Web),
		Entry("Value is string, Strategy - Works",
			&types.AttributeValueMemberS{Value: "Strategy"}, Position_Strategy),
		Entry("Value is string, 0 - Works",
			&types.AttributeValueMemberS{Value: "0"}, Position_Client),
		Entry("Value is string, 1 - Works",
			&types.AttributeValueMemberS{Value: "1"}, Position_Mobile),
		Entry("Value is string, 2 - Works",
			&types.AttributeValueMemberS{Value: "2"}, Position_Web),
		Entry("Value is string, 3 - Works",
			&types.AttributeValueMemberS{Value: "3"}, Position_Strategy))

	// Test that attempting to deserialize a data.Position.Reason will fial and return an error if the value
	// cannot be converted to either the name value or integer value of the enum option
	It("Scan - Value is nil - Error", func() {

		// Attempt to convert a fake string value into a data.Position.Reason; this should return an error
		var enum *Position_Reason
		err := enum.Scan(nil)

		// Verify the error
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of %!q(<nil>) had an invalid type of <nil>"))
		Expect(enum).Should(BeNil())
	})

	// Test the conditions under which values should be convertible to a data.Position.Reason
	DescribeTable("Scan Tests",
		func(value interface{}, shouldBe Position_Reason) {

			// Attempt to convert the value into a data.Position.Reason; this should not fail
			var enum Position_Reason
			err := enum.Scan(value)

			// Verify that the deserialization was successful
			Expect(err).ShouldNot(HaveOccurred())
			Expect(enum).Should(Equal(shouldBe))
		},
		Entry("Client - Works", "Client", Position_Client),
		Entry("Mobile - Works", "Mobile", Position_Mobile),
		Entry("Web - Works", "Web", Position_Web),
		Entry("Strategy - Works", "Strategy", Position_Strategy),
		Entry("0 - Works", 0, Position_Client),
		Entry("1 - Works", 1, Position_Mobile),
		Entry("2 - Works", 2, Position_Web),
		Entry("3 - Works", 3, Position_Strategy))
})

var _ = Describe("data.TradeRequest.Action Marshal/Unmarshal Tests", func() {

	// Test that converting the data.TradeRequest.Action enum to JSON works for all values
	DescribeTable("MarshalJSON Tests",
		func(enum TradeRequest_Action, value string) {
			data, err := json.Marshal(enum)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(string(data)).Should(Equal(value))
		},
		Entry("Deal - Works", TradeRequest_Deal, "\"Deal\""),
		Entry("Pending - Works", TradeRequest_Pending, "\"Pending\""),
		Entry("SLTP - Works", TradeRequest_SLTP, "\"SLTP\""),
		Entry("Modify - Works", TradeRequest_Modify, "\"Modify\""),
		Entry("Remove - Works", TradeRequest_Remove, "\"Remove\""),
		Entry("CloseBy - Works", TradeRequest_CloseBy, "\"CloseBy\""))

	// Test that converting the data.TradeRequest.Action enum to a CSV column works for all values
	DescribeTable("MarshalCSV Tests",
		func(enum TradeRequest_Action, value string) {
			data, err := enum.MarshalCSV()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(string(data)).Should(Equal(value))
		},
		Entry("Deal - Works", TradeRequest_Deal, "Deal"),
		Entry("Pending - Works", TradeRequest_Pending, "Pending"),
		Entry("SLTP - Works", TradeRequest_SLTP, "SLTP"),
		Entry("Modify - Works", TradeRequest_Modify, "Modify"),
		Entry("Remove - Works", TradeRequest_Remove, "Remove"),
		Entry("CloseBy - Works", TradeRequest_CloseBy, "CloseBy"))

	// Test that converting the data.TradeRequest.Action enum to a YAML node works for all values
	DescribeTable("MarshalYAML - Works",
		func(enum TradeRequest_Action, value string) {
			data, err := enum.MarshalYAML()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(data).Should(Equal(value))
		},
		Entry("Deal - Works", TradeRequest_Deal, "Deal"),
		Entry("Pending - Works", TradeRequest_Pending, "Pending"),
		Entry("SLTP - Works", TradeRequest_SLTP, "SLTP"),
		Entry("Modify - Works", TradeRequest_Modify, "Modify"),
		Entry("Remove - Works", TradeRequest_Remove, "Remove"),
		Entry("CloseBy - Works", TradeRequest_CloseBy, "CloseBy"))

	// Test that converting the data.TradeRequest.Action enum to a DynamoDB AttributeVAlue works for all values
	DescribeTable("MarshalDynamoDBAttributeValue - Works",
		func(enum TradeRequest_Action, value string) {
			data, err := attributevalue.Marshal(enum)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(data.(*types.AttributeValueMemberS).Value).Should(Equal(value))
		},
		Entry("Deal - Works", TradeRequest_Deal, "Deal"),
		Entry("Pending - Works", TradeRequest_Pending, "Pending"),
		Entry("SLTP - Works", TradeRequest_SLTP, "SLTP"),
		Entry("Modify - Works", TradeRequest_Modify, "Modify"),
		Entry("Remove - Works", TradeRequest_Remove, "Remove"),
		Entry("CloseBy - Works", TradeRequest_CloseBy, "CloseBy"))

	// Test that converting the data.TradeRequest.Action enum to an SQL value for all values
	DescribeTable("Value Tests",
		func(enum TradeRequest_Action, value string) {
			data, err := enum.Value()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(data).Should(Equal(value))
		},
		Entry("Deal - Works", TradeRequest_Deal, "Deal"),
		Entry("Pending - Works", TradeRequest_Pending, "Pending"),
		Entry("SLTP - Works", TradeRequest_SLTP, "SLTP"),
		Entry("Modify - Works", TradeRequest_Modify, "Modify"),
		Entry("Remove - Works", TradeRequest_Remove, "Remove"),
		Entry("CloseBy - Works", TradeRequest_CloseBy, "CloseBy"))

	// Test that attempting to deserialize a data.TradeRequest.Action will fail and return an error
	// if the value cannot be deserialized from a JSON value to a string
	It("UnmarshalJSON fails - Error", func() {

		// Attempt to convert a non-parseable string value into a data.TradeRequest.Action; this should return an error
		enum := new(TradeRequest_Action)
		err := enum.UnmarshalJSON([]byte("derp"))

		// Verify the error
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of \"derp\" cannot be mapped to a data.TradeRequest_Action"))
	})

	// Test that attempting to deserialize a data.TradeRequest.Action will fail and return an error
	// if the value cannot be converted to either the name value or integer value of the enum option
	It("UnmarshalJSON - Value is invalid - Error", func() {

		// Attempt to convert a fake string value into a data.TradeRequest.Action; this should return an error
		enum := new(TradeRequest_Action)
		err := enum.UnmarshalJSON([]byte("\"derp\""))

		// Verify the error
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of \"derp\" cannot be mapped to a data.TradeRequest_Action"))
	})

	// Test the conditions under which values should be convertible to a data.TradeRequest.Action
	DescribeTable("UnmarshalJSON Tests",
		func(value string, shouldBe TradeRequest_Action) {

			// Attempt to convert the string value into a data.TradeRequest.Action; this should not fail
			var enum TradeRequest_Action
			err := enum.UnmarshalJSON([]byte(value))

			// Verify that the deserialization was successful
			Expect(err).ShouldNot(HaveOccurred())
			Expect(enum).Should(Equal(shouldBe))
		},
		Entry("Deal - Works", "\"Deal\"", TradeRequest_Deal),
		Entry("Pending - Works", "\"Pending\"", TradeRequest_Pending),
		Entry("SLTP - Works", "\"SLTP\"", TradeRequest_SLTP),
		Entry("Modify - Works", "\"Modify\"", TradeRequest_Modify),
		Entry("Remove - Works", "\"Remove\"", TradeRequest_Remove),
		Entry("CloseBy - Works", "\"CloseBy\"", TradeRequest_CloseBy),
		Entry("0 - Works", "\"0\"", TradeRequest_Deal),
		Entry("1 - Works", "\"1\"", TradeRequest_Pending),
		Entry("2 - Works", "\"2\"", TradeRequest_SLTP),
		Entry("3 - Works", "\"3\"", TradeRequest_Modify),
		Entry("4 - Works", "\"4\"", TradeRequest_Remove),
		Entry("5 - Works", "\"5\"", TradeRequest_CloseBy))

	// Test that attempting to deserialize a data.TradeRequest.Action will fail and return an error
	// if the value cannot be converted to either the name value or integer value of the enum option
	It("UnmarshalCSV - Value is empty - Error", func() {

		// Attempt to convert a fake string value into a data.TradeRequest.Action; this should return an error
		enum := new(TradeRequest_Action)
		err := enum.UnmarshalCSV("")

		// Verify the error
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of \"\" cannot be mapped to a data.TradeRequest_Action"))
	})

	// Test the conditions under which values should be convertible to a data.TradeRequest.Action
	DescribeTable("UnmarshalCSV Tests",
		func(value string, shouldBe TradeRequest_Action) {

			// Attempt to convert the value into a data.TradeRequest.Action; this should not fail
			var enum TradeRequest_Action
			err := enum.UnmarshalCSV(value)

			// Verify that the deserialization was successful
			Expect(err).ShouldNot(HaveOccurred())
			Expect(enum).Should(Equal(shouldBe))
		},
		Entry("Deal - Works", "Deal", TradeRequest_Deal),
		Entry("Pending - Works", "Pending", TradeRequest_Pending),
		Entry("SLTP - Works", "SLTP", TradeRequest_SLTP),
		Entry("Modify - Works", "Modify", TradeRequest_Modify),
		Entry("Remove - Works", "Remove", TradeRequest_Remove),
		Entry("CloseBy - Works", "CloseBy", TradeRequest_CloseBy),
		Entry("0 - Works", "0", TradeRequest_Deal),
		Entry("1 - Works", "1", TradeRequest_Pending),
		Entry("2 - Works", "2", TradeRequest_SLTP),
		Entry("3 - Works", "3", TradeRequest_Modify),
		Entry("4 - Works", "4", TradeRequest_Remove),
		Entry("5 - Works", "5", TradeRequest_CloseBy))

	// Test that attempting to deserialize a data.TradeRequest.Action will fail and return an error
	// if the YAML node does not represent a scalar value
	It("UnmarshalYAML - Node type is not scalar - Error", func() {
		enum := new(TradeRequest_Action)
		err := enum.UnmarshalYAML(&yaml.Node{Kind: yaml.AliasNode})
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("YAML node had an invalid kind (expected scalar value)"))
	})

	// Test that attempting to deserialize a data.TradeRequest.Action will fail and return an error if
	// the YAML node value cannot be converted to either the name value or integer value of the enum option
	It("UnmarshalYAML - Parse fails - Error", func() {
		enum := new(TradeRequest_Action)
		err := enum.UnmarshalYAML(&yaml.Node{Kind: yaml.ScalarNode, Value: "derp"})
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of \"derp\" cannot be mapped to a data.TradeRequest_Action"))
	})

	// Test the conditions under which YAML node values should be convertible to a data.TradeRequest.Action
	DescribeTable("UnmarshalYAML Tests",
		func(value string, shouldBe TradeRequest_Action) {
			var enum TradeRequest_Action
			err := enum.UnmarshalYAML(&yaml.Node{Kind: yaml.ScalarNode, Value: value})
			Expect(err).ShouldNot(HaveOccurred())
			Expect(enum).Should(Equal(shouldBe))
		},
		Entry("Deal - Works", "Deal", TradeRequest_Deal),
		Entry("Pending - Works", "Pending", TradeRequest_Pending),
		Entry("SLTP - Works", "SLTP", TradeRequest_SLTP),
		Entry("Modify - Works", "Modify", TradeRequest_Modify),
		Entry("Remove - Works", "Remove", TradeRequest_Remove),
		Entry("CloseBy - Works", "CloseBy", TradeRequest_CloseBy),
		Entry("0 - Works", "0", TradeRequest_Deal),
		Entry("1 - Works", "1", TradeRequest_Pending),
		Entry("2 - Works", "2", TradeRequest_SLTP),
		Entry("3 - Works", "3", TradeRequest_Modify),
		Entry("4 - Works", "4", TradeRequest_Remove),
		Entry("5 - Works", "5", TradeRequest_CloseBy))

	// Tests that, if the attribute type submitted to UnmarshalDynamoDBAttributeValue is not one we
	// recognize, then the function will return an error
	It("UnmarshalDynamoDBAttributeValue - AttributeValue type invalid - Error", func() {
		enum := new(TradeRequest_Action)
		err := attributevalue.Unmarshal(&types.AttributeValueMemberBOOL{Value: true}, &enum)
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("Attribute value of *types.AttributeValueMemberBOOL could not be converted to a data.TradeRequest.Action"))
	})

	// Tests that, if time parsing fails, then calling UnmarshalDynamoDBAttributeValue will return an error
	It("UnmarshalDynamoDBAttributeValue - Parse fails - Error", func() {
		enum := new(TradeRequest_Action)
		err := attributevalue.Unmarshal(&types.AttributeValueMemberS{Value: "derp"}, &enum)
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of \"derp\" cannot be mapped to a data.TradeRequest_Action"))
	})

	// Tests the conditions under which UnmarshalDynamoDBAttributeValue is called and no error is generated
	DescribeTable("UnmarshalDynamoDBAttributeValue - AttributeValue Conditions",
		func(value types.AttributeValue, expected TradeRequest_Action) {
			var enum TradeRequest_Action
			err := attributevalue.Unmarshal(value, &enum)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(enum).Should(Equal(expected))
		},
		Entry("Value is []bytes, Deal - Works",
			&types.AttributeValueMemberB{Value: []byte("Deal")}, TradeRequest_Deal),
		Entry("Value is []bytes, Pending - Works",
			&types.AttributeValueMemberB{Value: []byte("Pending")}, TradeRequest_Pending),
		Entry("Value is []bytes, SLTP - Works",
			&types.AttributeValueMemberB{Value: []byte("SLTP")}, TradeRequest_SLTP),
		Entry("Value is []bytes, Modify - Works",
			&types.AttributeValueMemberB{Value: []byte("Modify")}, TradeRequest_Modify),
		Entry("Value is []bytes, Remove - Works",
			&types.AttributeValueMemberB{Value: []byte("Remove")}, TradeRequest_Remove),
		Entry("Value is []bytes, CloseBy - Works",
			&types.AttributeValueMemberB{Value: []byte("CloseBy")}, TradeRequest_CloseBy),
		Entry("Value is []bytes, 0 - Works",
			&types.AttributeValueMemberB{Value: []byte("0")}, TradeRequest_Deal),
		Entry("Value is []bytes, 1 - Works",
			&types.AttributeValueMemberB{Value: []byte("1")}, TradeRequest_Pending),
		Entry("Value is []bytes, 2 - Works",
			&types.AttributeValueMemberB{Value: []byte("2")}, TradeRequest_SLTP),
		Entry("Value is []bytes, 3 - Works",
			&types.AttributeValueMemberB{Value: []byte("3")}, TradeRequest_Modify),
		Entry("Value is []bytes, 4 - Works",
			&types.AttributeValueMemberB{Value: []byte("4")}, TradeRequest_Remove),
		Entry("Value is []bytes, 5 - Works",
			&types.AttributeValueMemberB{Value: []byte("5")}, TradeRequest_CloseBy),
		Entry("Value is int, 0 - Works",
			&types.AttributeValueMemberN{Value: "0"}, TradeRequest_Deal),
		Entry("Value is int, 1 - Works",
			&types.AttributeValueMemberN{Value: "1"}, TradeRequest_Pending),
		Entry("Value is int, 2 - Works",
			&types.AttributeValueMemberN{Value: "2"}, TradeRequest_SLTP),
		Entry("Value is int, 3 - Works",
			&types.AttributeValueMemberN{Value: "3"}, TradeRequest_Modify),
		Entry("Value is int, 4 - Works",
			&types.AttributeValueMemberN{Value: "4"}, TradeRequest_Remove),
		Entry("Value is int, 5 - Works",
			&types.AttributeValueMemberN{Value: "5"}, TradeRequest_CloseBy),
		Entry("Value is NULL - Works", new(types.AttributeValueMemberNULL), TradeRequest_Action(0)),
		Entry("Value is string, Deal - Works",
			&types.AttributeValueMemberS{Value: "Deal"}, TradeRequest_Deal),
		Entry("Value is string, Pending - Works",
			&types.AttributeValueMemberS{Value: "Pending"}, TradeRequest_Pending),
		Entry("Value is string, SLTP - Works",
			&types.AttributeValueMemberS{Value: "SLTP"}, TradeRequest_SLTP),
		Entry("Value is string, Modify - Works",
			&types.AttributeValueMemberS{Value: "Modify"}, TradeRequest_Modify),
		Entry("Value is string, Remove - Works",
			&types.AttributeValueMemberS{Value: "Remove"}, TradeRequest_Remove),
		Entry("Value is string, CloseBy - Works",
			&types.AttributeValueMemberS{Value: "CloseBy"}, TradeRequest_CloseBy),
		Entry("Value is string, 0 - Works",
			&types.AttributeValueMemberS{Value: "0"}, TradeRequest_Deal),
		Entry("Value is string, 1 - Works",
			&types.AttributeValueMemberS{Value: "1"}, TradeRequest_Pending),
		Entry("Value is string, 2 - Works",
			&types.AttributeValueMemberS{Value: "2"}, TradeRequest_SLTP),
		Entry("Value is string, 3 - Works",
			&types.AttributeValueMemberS{Value: "3"}, TradeRequest_Modify),
		Entry("Value is string, 4 - Works",
			&types.AttributeValueMemberS{Value: "4"}, TradeRequest_Remove),
		Entry("Value is string, 5 - Works",
			&types.AttributeValueMemberS{Value: "5"}, TradeRequest_CloseBy))

	// Test that attempting to deserialize a data.TradeRequest.Action will fial and return an error
	// if the value cannot be converted to either the name value or integer value of the enum option
	It("Scan - Value is nil - Error", func() {

		// Attempt to convert a fake string value into a data.TradeRequest.Action; this should return an error
		var enum *TradeRequest_Action
		err := enum.Scan(nil)

		// Verify the error
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of %!q(<nil>) had an invalid type of <nil>"))
		Expect(enum).Should(BeNil())
	})

	// Test the conditions under which values should be convertible to a data.TradeRequest.Action
	DescribeTable("Scan Tests",
		func(value interface{}, shouldBe TradeRequest_Action) {

			// Attempt to convert the value into a data.TradeRequest.Action; this should not fail
			var enum TradeRequest_Action
			err := enum.Scan(value)

			// Verify that the deserialization was successful
			Expect(err).ShouldNot(HaveOccurred())
			Expect(enum).Should(Equal(shouldBe))
		},
		Entry("Deal - Works", "Deal", TradeRequest_Deal),
		Entry("Pending - Works", "Pending", TradeRequest_Pending),
		Entry("SLTP - Works", "SLTP", TradeRequest_SLTP),
		Entry("Modify - Works", "Modify", TradeRequest_Modify),
		Entry("Remove - Works", "Remove", TradeRequest_Remove),
		Entry("CloseBy - Works", "CloseBy", TradeRequest_CloseBy),
		Entry("0 - Works", 0, TradeRequest_Deal),
		Entry("1 - Works", 1, TradeRequest_Pending),
		Entry("2 - Works", 2, TradeRequest_SLTP),
		Entry("3 - Works", 3, TradeRequest_Modify),
		Entry("4 - Works", 4, TradeRequest_Remove),
		Entry("5 - Works", 5, TradeRequest_CloseBy))
})
