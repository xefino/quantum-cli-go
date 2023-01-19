package v1

import (
	"encoding/json"
	"testing"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"gopkg.in/yaml.v3"
)

// Create a new test runner we'll use to test all the
// modules in the filtering.v1 package
func TestFilteringV1(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Filtering V1 Suite")
}

var _ = Describe("filtering.v1.AllowedTrades Marshal/Unmarshal Tests", func() {

	// Test that converting the filtering.v1.AllowedTrades enum to JSON works for all values
	DescribeTable("MarshalJSON Tests",
		func(enum AllowedTrades, value string) {
			data, err := json.Marshal(enum)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(string(data)).Should(Equal(value))
		},
		Entry("AllAllowed - Works", AllowedTrades_AllAllowed, "\"ALL\""),
		Entry("BuysAllowed - Works", AllowedTrades_BuysAllowed, "\"BUY\""),
		Entry("SellsAllowed - Works", AllowedTrades_SellsAllowed, "\"SELL\""))

	// Test that converting the filtering.v1.AllowedTrades enum to a CSV column works for all values
	DescribeTable("MarshalCSV Tests",
		func(enum AllowedTrades, value string) {
			data, err := enum.MarshalCSV()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(string(data)).Should(Equal(value))
		},
		Entry("AllAllowed - Works", AllowedTrades_AllAllowed, "ALL"),
		Entry("BuysAllowed - Works", AllowedTrades_BuysAllowed, "BUY"),
		Entry("SellsAllowed - Works", AllowedTrades_SellsAllowed, "SELL"))

	// Test that converting the filtering.v1.AllowedTrades enum to a YAML node works for all values
	DescribeTable("MarshalYAML - Works",
		func(enum AllowedTrades, value string) {
			data, err := enum.MarshalYAML()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(data).Should(Equal(value))
		},
		Entry("AllAllowed - Works", AllowedTrades_AllAllowed, "ALL"),
		Entry("BuysAllowed - Works", AllowedTrades_BuysAllowed, "BUY"),
		Entry("SellsAllowed - Works", AllowedTrades_SellsAllowed, "SELL"))

	// Test that converting the filtering.v1.AllowedTrades enum to a DynamoDB AttributeVAlue works for all values
	DescribeTable("MarshalDynamoDBAttributeValue - Works",
		func(enum AllowedTrades, value string) {
			data, err := attributevalue.Marshal(enum)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(data.(*types.AttributeValueMemberS).Value).Should(Equal(value))
		},
		Entry("AllAllowed - Works", AllowedTrades_AllAllowed, "ALL"),
		Entry("BuysAllowed - Works", AllowedTrades_BuysAllowed, "BUY"),
		Entry("SellsAllowed - Works", AllowedTrades_SellsAllowed, "SELL"))

	// Test that converting the filtering.v1.AllowedTrades enum to an SQL value for all values
	DescribeTable("Value Tests",
		func(enum AllowedTrades, value string) {
			data, err := enum.Value()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(data).Should(Equal(value))
		},
		Entry("AllAllowed - Works", AllowedTrades_AllAllowed, "ALL"),
		Entry("BuysAllowed - Works", AllowedTrades_BuysAllowed, "BUY"),
		Entry("SellsAllowed - Works", AllowedTrades_SellsAllowed, "SELL"))

	// Test that attempting to deserialize a filtering.v1.AllowedTrades will fail and return an error if the value
	// cannot be deserialized from a JSON value to a string
	It("UnmarshalJSON fails - Error", func() {

		// Attempt to convert a non-parseable string value into a filtering.v1.AllowedTrades; this should return an error
		enum := new(AllowedTrades)
		err := enum.UnmarshalJSON([]byte("derp"))

		// Verify the error
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of \"derp\" cannot be mapped to a v1.AllowedTrades"))
	})

	// Test that attempting to deserialize a filtering.v1.AllowedTrades will fail and return an error if the value
	// cannot be converted to either the name value or integer value of the enum option
	It("UnmarshalJSON - Value is invalid - Error", func() {

		// Attempt to convert a fake string value into a filtering.v1.AllowedTrades; this should return an error
		enum := new(AllowedTrades)
		err := enum.UnmarshalJSON([]byte("\"derp\""))

		// Verify the error
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of \"derp\" cannot be mapped to a v1.AllowedTrades"))
	})

	// Test the conditions under which values should be convertible to a filtering.v1.AllowedTrades
	DescribeTable("UnmarshalJSON Tests",
		func(value string, shouldBe AllowedTrades) {

			// Attempt to convert the string value into a filtering.v1.AllowedTrades; this should not fail
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

	// Test that attempting to deserialize a filtering.v1.AllowedTrades will fail and return an error if the value
	// cannot be converted to either the name value or integer value of the enum option
	It("UnmarshalCSV - Value is empty - Error", func() {

		// Attempt to convert a fake string value into a filtering.v1.AllowedTrades; this should return an error
		enum := new(AllowedTrades)
		err := enum.UnmarshalCSV("")

		// Verify the error
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of \"\" cannot be mapped to a v1.AllowedTrades"))
	})

	// Test the conditions under which values should be convertible to a filtering.v1.AllowedTrades
	DescribeTable("UnmarshalCSV Tests",
		func(value string, shouldBe AllowedTrades) {

			// Attempt to convert the value into a filtering.v1.AllowedTrades; this should not fail
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

	// Test that attempting to deserialize a filtering.v1.AllowedTrades will fail and return an error if the YAML
	// node does not represent a scalar value
	It("UnmarshalYAML - Node type is not scalar - Error", func() {
		enum := new(AllowedTrades)
		err := enum.UnmarshalYAML(&yaml.Node{Kind: yaml.AliasNode})
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("YAML node had an invalid kind (expected scalar value)"))
	})

	// Test that attempting to deserialize a filtering.v1.AllowedTrades will fail and return an error if the YAML
	// node value cannot be converted to either the name value or integer value of the enum option
	It("UnmarshalYAML - Parse fails - Error", func() {
		enum := new(AllowedTrades)
		err := enum.UnmarshalYAML(&yaml.Node{Kind: yaml.ScalarNode, Value: "derp"})
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of \"derp\" cannot be mapped to a v1.AllowedTrades"))
	})

	// Test the conditions under which YAML node values should be convertible to a filtering.v1.AllowedTrades
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
		Expect(err.Error()).Should(Equal("Attribute value of *types.AttributeValueMemberBOOL could not be converted to a filtering.v1.AllowedTrades"))
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

	// Test that attempting to deserialize a filtering.v1.AllowedTrades will fial and return an error if the value
	// cannot be converted to either the name value or integer value of the enum option
	It("Scan - Value is nil - Error", func() {

		// Attempt to convert a fake string value into a filtering.v1.AllowedTrades; this should return an error
		var enum *AllowedTrades
		err := enum.Scan(nil)

		// Verify the error
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("value of %!q(<nil>) had an invalid type of <nil>"))
		Expect(enum).Should(BeNil())
	})

	// Test the conditions under which values should be convertible to a filtering.v1.AllowedTrades
	DescribeTable("Scan Tests",
		func(value interface{}, shouldBe AllowedTrades) {

			// Attempt to convert the value into a filtering.v1.AllowedTrades; this should not fail
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
