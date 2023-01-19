package v1

import (
	"database/sql/driver"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/xefino/protobuf-gen-go/utils"
	"gopkg.in/yaml.v3"
)

// AllowedTradesAlternates contains alternate values for the filtering.v1.AllowedTrades enum
var AllowedTradesAlternates = map[string]AllowedTrades{
	"all":  AllowedTrades_AllAllowed,
	"ALL":  AllowedTrades_AllAllowed,
	"buy":  AllowedTrades_BuysAllowed,
	"BUY":  AllowedTrades_BuysAllowed,
	"sell": AllowedTrades_SellsAllowed,
	"SELL": AllowedTrades_SellsAllowed,
}

// AllowedTradesMapping contains alternate names for the filtering.v1.AllowedTrades enum
var AllowedTradesMapping = map[AllowedTrades]string{
	AllowedTrades_AllAllowed:   "ALL",
	AllowedTrades_BuysAllowed:  "BUY",
	AllowedTrades_SellsAllowed: "SELL",
}

// MarshalJSON converts a filtering.v1.AllowedTrades value to a JSON value
func (enum AllowedTrades) MarshalJSON() ([]byte, error) {
	return []byte(utils.MarshalString(enum, AllowedTrades_name, AllowedTradesMapping, true)), nil
}

// MarshalCSV converts a filtering.v1.AllowedTrades value to CSV cell value
func (enum AllowedTrades) MarshalCSV() (string, error) {
	return utils.MarshalString(enum, AllowedTrades_name, AllowedTradesMapping, false), nil
}

// MarshalYAML converts a filtering.v1.AllowedTrades value to a YAML node value
func (enum AllowedTrades) MarshalYAML() (interface{}, error) {
	return utils.MarshalString(enum, AllowedTrades_name, AllowedTradesMapping, false), nil
}

// MarshalDynamoDBAttributeValue converts a filtering.v1.AllowedTrades value to a DynamoDB AttributeValue
func (enum AllowedTrades) MarshalDynamoDBAttributeValue() (types.AttributeValue, error) {
	return &types.AttributeValueMemberS{Value: utils.MarshalString(enum, AllowedTrades_name, AllowedTradesMapping, false)}, nil
}

// Value converts a filtering.v1.AllowedTrades value to an SQL driver value
func (enum AllowedTrades) Value() (driver.Value, error) {
	return utils.MarshalString(enum, AllowedTrades_name, AllowedTradesMapping, false), nil
}

// UnmarshalJSON attempts to convert a JSON value to a new filtering.v1.AllowedTrades value
func (enum *AllowedTrades) UnmarshalJSON(raw []byte) error {
	return utils.UnmarshalValue(raw, AllowedTrades_value, AllowedTradesAlternates, enum)
}

// UnmarshalCSV attempts to convert a CSV cell value to a new filtering.v1.AllowedTrades value
func (enum *AllowedTrades) UnmarshalCSV(raw string) error {
	return utils.UnmarshalString(raw, AllowedTrades_value, AllowedTradesAlternates, enum)
}

// UnmarshalYAML attempts to convert a YAML node to a new filtering.v1.AllowedTrades value
func (enum *AllowedTrades) UnmarshalYAML(value *yaml.Node) error {
	if value.Kind != yaml.ScalarNode {
		return fmt.Errorf("YAML node had an invalid kind (expected scalar value)")
	} else {
		return utils.UnmarshalString(value.Value, AllowedTrades_value, AllowedTradesAlternates, enum)
	}
}

// UnmarshalDynamoDBAttributeValue attempts to convert a DynamoDB AttributeVAlue to a filtering.v1.AllowedTrades
// value. This function can handle []bytes, numerics, or strings. If the AttributeValue is NULL then
// the AllowedTrades value will not be modified.
func (enum *AllowedTrades) UnmarshalDynamoDBAttributeValue(value types.AttributeValue) error {
	switch casted := value.(type) {
	case *types.AttributeValueMemberB:
		return utils.UnmarshalValue(casted.Value, AllowedTrades_value, AllowedTradesAlternates, enum)
	case *types.AttributeValueMemberN:
		return utils.UnmarshalString(casted.Value, AllowedTrades_value, AllowedTradesAlternates, enum)
	case *types.AttributeValueMemberNULL:
		return nil
	case *types.AttributeValueMemberS:
		return utils.UnmarshalString(casted.Value, AllowedTrades_value, AllowedTradesAlternates, enum)
	default:
		return fmt.Errorf("Attribute value of %T could not be converted to a filtering.v1.AllowedTrades", value)
	}
}

// Scan attempts to convert an SQL driver value to a new filtering.v1.AllowedTrades value
func (enum *AllowedTrades) Scan(value interface{}) error {
	return utils.ScanValue(value, AllowedTrades_value, AllowedTradesAlternates, enum)
}
