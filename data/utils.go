package data

import (
	"database/sql/driver"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/xefino/protobuf-gen-go/utils"
	"gopkg.in/yaml.v3"
)

// FillPolicyAlternates contains alternate values for the data.FillPolicy enum
var FillPolicyAlternates = map[string]FillPolicy{
	"FOK": FillPolicy_FillOrKill,
	"IOC": FillPolicy_ImmediateOrCancel,
}

// FillPolicyMapping contains alternate names for the data.FillPolicy enum
var FillPolicyMapping = map[FillPolicy]string{
	FillPolicy_FillOrKill:        "FOK",
	FillPolicy_ImmediateOrCancel: "IOC",
}

// ExpirationPolicyAlternates contains alternate values for the data.ExpirationPolicy enum
var ExpirationPolicyAlternates = map[string]ExpirationPolicy{
	"GTC": ExpirationPolicy_UntilCancelled,
}

// ExpirationPolicyMapping contains alternate names for the data.ExpirationPolicy enum
var ExpirationPolicyMapping = map[ExpirationPolicy]string{
	ExpirationPolicy_UntilCancelled: "GTC",
}

// DealEntryAlternates contains alternate values for the data.Deal.Entry enum
var DealEntryAlternates = map[string]Deal_Entry{
	"InOut": Deal_Reverse,
}

// DealReasonAlternates contains alternate values for the data.Deal.Reason enum
var DealReasonAlternates = map[string]Deal_Reason{
	"SL": Deal_StopLoss,
	"TP": Deal_TakeProfit,
	"SO": Deal_StopOut,
}

// DealReasonMapping contains alternate names for the data.Deal.Reason enum
var DealReasonMapping = map[Deal_Reason]string{
	Deal_StopLoss:   "SL",
	Deal_TakeProfit: "TP",
	Deal_StopOut:    "SO",
}

// OrderReasonAlternates contains alternate values for the data.Order.Reason enum
var OrderReasonAlternates = map[string]Order_Reason{
	"SL": Order_StopLoss,
	"TP": Order_TakeProfit,
	"SO": Order_StopOut,
}

// OrderReasonMapping contains alternate names for the data.Order.Reason enum
var OrderReasonMapping = map[Order_Reason]string{
	Order_StopLoss:   "SL",
	Order_TakeProfit: "TP",
	Order_StopOut:    "SO",
}

// MarshalJSON converts a data.FillPolicy value to a JSON value
func (enum FillPolicy) MarshalJSON() ([]byte, error) {
	return []byte(utils.MarshalString(enum, FillPolicy_name, FillPolicyMapping, true)), nil
}

// MarshalCSV converts a data.FillPolicy value to CSV cell value
func (enum FillPolicy) MarshalCSV() (string, error) {
	return utils.MarshalString(enum, FillPolicy_name, FillPolicyMapping, false), nil
}

// MarshalYAML converts a data.FillPolicy value to a YAML node value
func (enum FillPolicy) MarshalYAML() (interface{}, error) {
	return utils.MarshalString(enum, FillPolicy_name, FillPolicyMapping, false), nil
}

// MarshalDynamoDBAttributeValue converts a data.FillPolicy value to a DynamoDB AttributeValue
func (enum FillPolicy) MarshalDynamoDBAttributeValue() (types.AttributeValue, error) {
	return &types.AttributeValueMemberS{Value: utils.MarshalString(enum, FillPolicy_name, FillPolicyMapping, false)}, nil
}

// Value converts a data.FillPolicy value to an SQL driver value
func (enum FillPolicy) Value() (driver.Value, error) {
	return utils.MarshalString(enum, FillPolicy_name, FillPolicyMapping, false), nil
}

// UnmarshalJSON attempts to convert a JSON value to a new data.FillPolicy value
func (enum *FillPolicy) UnmarshalJSON(raw []byte) error {
	return utils.UnmarshalValue(raw, FillPolicy_value, FillPolicyAlternates, enum)
}

// UnmarshalCSV attempts to convert a CSV cell value to a new data.FillPolicy value
func (enum *FillPolicy) UnmarshalCSV(raw string) error {
	return utils.UnmarshalString(raw, FillPolicy_value, FillPolicyAlternates, enum)
}

// UnmarshalYAML attempts to convert a YAML node to a new data.FillPolicy value
func (enum *FillPolicy) UnmarshalYAML(value *yaml.Node) error {
	if value.Kind != yaml.ScalarNode {
		return fmt.Errorf("YAML node had an invalid kind (expected scalar value)")
	} else {
		return utils.UnmarshalString(value.Value, FillPolicy_value, FillPolicyAlternates, enum)
	}
}

// UnmarshalDynamoDBAttributeValue attempts to convert a DynamoDB AttributeVAlue to a data.FillPolicy
// value. This function can handle []bytes, numerics, or strings. If the AttributeValue is NULL then
// the FillPolicy value will not be modified.
func (enum *FillPolicy) UnmarshalDynamoDBAttributeValue(value types.AttributeValue) error {
	switch casted := value.(type) {
	case *types.AttributeValueMemberB:
		return utils.UnmarshalValue(casted.Value, FillPolicy_value, FillPolicyAlternates, enum)
	case *types.AttributeValueMemberN:
		return utils.UnmarshalString(casted.Value, FillPolicy_value, FillPolicyAlternates, enum)
	case *types.AttributeValueMemberNULL:
		return nil
	case *types.AttributeValueMemberS:
		return utils.UnmarshalString(casted.Value, FillPolicy_value, FillPolicyAlternates, enum)
	default:
		return fmt.Errorf("Attribute value of %T could not be converted to a data.FillPolicy", value)
	}
}

// Scan attempts to convert an SQL driver value to a new data.FillPolicy value
func (enum *FillPolicy) Scan(value interface{}) error {
	return utils.ScanValue(value, FillPolicy_value, FillPolicyAlternates, enum)
}

// MarshalJSON converts a data.ExpirationPolicy value to a JSON value
func (enum ExpirationPolicy) MarshalJSON() ([]byte, error) {
	return []byte(utils.MarshalString(enum, ExpirationPolicy_name, ExpirationPolicyMapping, true)), nil
}

// MarshalCSV converts a data.ExpirationPolicy value to CSV cell value
func (enum ExpirationPolicy) MarshalCSV() (string, error) {
	return utils.MarshalString(enum, ExpirationPolicy_name, ExpirationPolicyMapping, false), nil
}

// MarshalYAML converts a data.ExpirationPolicy value to a YAML node value
func (enum ExpirationPolicy) MarshalYAML() (interface{}, error) {
	return utils.MarshalString(enum, ExpirationPolicy_name, ExpirationPolicyMapping, false), nil
}

// MarshalDynamoDBAttributeValue converts a data.ExpirationPolicy value to a DynamoDB AttributeValue
func (enum ExpirationPolicy) MarshalDynamoDBAttributeValue() (types.AttributeValue, error) {
	return &types.AttributeValueMemberS{Value: utils.MarshalString(enum, ExpirationPolicy_name, ExpirationPolicyMapping, false)}, nil
}

// Value converts a data.ExpirationPolicy value to an SQL driver value
func (enum ExpirationPolicy) Value() (driver.Value, error) {
	return utils.MarshalString(enum, ExpirationPolicy_name, ExpirationPolicyMapping, false), nil
}

// UnmarshalJSON attempts to convert a JSON value to a new data.ExpirationPolicy value
func (enum *ExpirationPolicy) UnmarshalJSON(raw []byte) error {
	return utils.UnmarshalValue(raw, ExpirationPolicy_value, ExpirationPolicyAlternates, enum)
}

// UnmarshalCSV attempts to convert a CSV cell value to a new data.ExpirationPolicy value
func (enum *ExpirationPolicy) UnmarshalCSV(raw string) error {
	return utils.UnmarshalString(raw, ExpirationPolicy_value, ExpirationPolicyAlternates, enum)
}

// UnmarshalYAML attempts to convert a YAML node to a new data.ExpirationPolicy value
func (enum *ExpirationPolicy) UnmarshalYAML(value *yaml.Node) error {
	if value.Kind != yaml.ScalarNode {
		return fmt.Errorf("YAML node had an invalid kind (expected scalar value)")
	} else {
		return utils.UnmarshalString(value.Value, ExpirationPolicy_value, ExpirationPolicyAlternates, enum)
	}
}

// UnmarshalDynamoDBAttributeValue attempts to convert a DynamoDB AttributeVAlue to a data.ExpirationPolicy
// value. This function can handle []bytes, numerics, or strings. If the AttributeValue is NULL then
// the ExpirationPolicy value will not be modified.
func (enum *ExpirationPolicy) UnmarshalDynamoDBAttributeValue(value types.AttributeValue) error {
	switch casted := value.(type) {
	case *types.AttributeValueMemberB:
		return utils.UnmarshalValue(casted.Value, ExpirationPolicy_value, ExpirationPolicyAlternates, enum)
	case *types.AttributeValueMemberN:
		return utils.UnmarshalString(casted.Value, ExpirationPolicy_value, ExpirationPolicyAlternates, enum)
	case *types.AttributeValueMemberNULL:
		return nil
	case *types.AttributeValueMemberS:
		return utils.UnmarshalString(casted.Value, ExpirationPolicy_value, ExpirationPolicyAlternates, enum)
	default:
		return fmt.Errorf("Attribute value of %T could not be converted to a data.ExpirationPolicy", value)
	}
}

// Scan attempts to convert an SQL driver value to a new data.ExpirationPolicy value
func (enum *ExpirationPolicy) Scan(value interface{}) error {
	return utils.ScanValue(value, ExpirationPolicy_value, ExpirationPolicyAlternates, enum)
}

// MarshalJSON converts a data.Deal.Type value to a JSON value
func (enum Deal_Type) MarshalJSON() ([]byte, error) {
	return []byte(utils.MarshalString(enum, Deal_Type_name, utils.Ignore, true)), nil
}

// MarshalCSV converts a data.Deal.Type value to CSV cell value
func (enum Deal_Type) MarshalCSV() (string, error) {
	return utils.MarshalString(enum, Deal_Type_name, utils.Ignore, false), nil
}

// MarshalYAML converts a data.Deal.Type value to a YAML node value
func (enum Deal_Type) MarshalYAML() (interface{}, error) {
	return utils.MarshalString(enum, Deal_Type_name, utils.Ignore, false), nil
}

// MarshalDynamoDBAttributeValue converts a data.Deal.Type value to a DynamoDB AttributeValue
func (enum Deal_Type) MarshalDynamoDBAttributeValue() (types.AttributeValue, error) {
	return &types.AttributeValueMemberS{Value: utils.MarshalString(enum, Deal_Type_name, utils.Ignore, false)}, nil
}

// Value converts a data.Deal.Type value to an SQL driver value
func (enum Deal_Type) Value() (driver.Value, error) {
	return utils.MarshalString(enum, Deal_Type_name, utils.Ignore, false), nil
}

// UnmarshalJSON attempts to convert a JSON value to a new data.Deal.Type value
func (enum *Deal_Type) UnmarshalJSON(raw []byte) error {
	return utils.UnmarshalValue(raw, Deal_Type_value, utils.None, enum)
}

// UnmarshalCSV attempts to convert a CSV cell value to a new data.Deal.Type value
func (enum *Deal_Type) UnmarshalCSV(raw string) error {
	return utils.UnmarshalString(raw, Deal_Type_value, utils.None, enum)
}

// UnmarshalYAML attempts to convert a YAML node to a new data.Deal.Type value
func (enum *Deal_Type) UnmarshalYAML(value *yaml.Node) error {
	if value.Kind != yaml.ScalarNode {
		return fmt.Errorf("YAML node had an invalid kind (expected scalar value)")
	} else {
		return utils.UnmarshalString(value.Value, Deal_Type_value, utils.None, enum)
	}
}

// UnmarshalDynamoDBAttributeValue attempts to convert a DynamoDB AttributeVAlue to a data.Deal.Type
// value. This function can handle []bytes, numerics, or strings. If the AttributeValue is NULL then
// the Deal_Type value will not be modified.
func (enum *Deal_Type) UnmarshalDynamoDBAttributeValue(value types.AttributeValue) error {
	switch casted := value.(type) {
	case *types.AttributeValueMemberB:
		return utils.UnmarshalValue(casted.Value, Deal_Type_value, utils.None, enum)
	case *types.AttributeValueMemberN:
		return utils.UnmarshalString(casted.Value, Deal_Type_value, utils.None, enum)
	case *types.AttributeValueMemberNULL:
		return nil
	case *types.AttributeValueMemberS:
		return utils.UnmarshalString(casted.Value, Deal_Type_value, utils.None, enum)
	default:
		return fmt.Errorf("Attribute value of %T could not be converted to a data.Deal.Type", value)
	}
}

// Scan attempts to convert an SQL driver value to a new data.Deal.Type value
func (enum *Deal_Type) Scan(value interface{}) error {
	return utils.ScanValue(value, Deal_Type_value, utils.None, enum)
}

// MarshalJSON converts a data.Deal.Entry value to a JSON value
func (enum Deal_Entry) MarshalJSON() ([]byte, error) {
	return []byte(utils.MarshalString(enum, Deal_Entry_name, utils.Ignore, true)), nil
}

// MarshalCSV converts a data.Deal.Entry value to CSV cell value
func (enum Deal_Entry) MarshalCSV() (string, error) {
	return utils.MarshalString(enum, Deal_Entry_name, utils.Ignore, false), nil
}

// MarshalYAML converts a data.Deal.Entry value to a YAML node value
func (enum Deal_Entry) MarshalYAML() (interface{}, error) {
	return utils.MarshalString(enum, Deal_Entry_name, utils.Ignore, false), nil
}

// MarshalDynamoDBAttributeValue converts a data.Deal.Entry value to a DynamoDB AttributeValue
func (enum Deal_Entry) MarshalDynamoDBAttributeValue() (types.AttributeValue, error) {
	return &types.AttributeValueMemberS{Value: utils.MarshalString(enum, Deal_Entry_name, utils.Ignore, false)}, nil
}

// Value converts a data.Deal.Entry value to an SQL driver value
func (enum Deal_Entry) Value() (driver.Value, error) {
	return utils.MarshalString(enum, Deal_Entry_name, utils.Ignore, false), nil
}

// UnmarshalJSON attempts to convert a JSON value to a new data.Deal.Entry value
func (enum *Deal_Entry) UnmarshalJSON(raw []byte) error {
	return utils.UnmarshalValue(raw, Deal_Entry_value, DealEntryAlternates, enum)
}

// UnmarshalCSV attempts to convert a CSV cell value to a new data.Deal.Entry value
func (enum *Deal_Entry) UnmarshalCSV(raw string) error {
	return utils.UnmarshalString(raw, Deal_Entry_value, DealEntryAlternates, enum)
}

// UnmarshalYAML attempts to convert a YAML node to a new data.Deal.Entry value
func (enum *Deal_Entry) UnmarshalYAML(value *yaml.Node) error {
	if value.Kind != yaml.ScalarNode {
		return fmt.Errorf("YAML node had an invalid kind (expected scalar value)")
	} else {
		return utils.UnmarshalString(value.Value, Deal_Entry_value, DealEntryAlternates, enum)
	}
}

// UnmarshalDynamoDBAttributeValue attempts to convert a DynamoDB AttributeVAlue to a data.Deal.Entry
// value. This function can handle []bytes, numerics, or strings. If the AttributeValue is NULL then
// the Deal_Entry value will not be modified.
func (enum *Deal_Entry) UnmarshalDynamoDBAttributeValue(value types.AttributeValue) error {
	switch casted := value.(type) {
	case *types.AttributeValueMemberB:
		return utils.UnmarshalValue(casted.Value, Deal_Entry_value, DealEntryAlternates, enum)
	case *types.AttributeValueMemberN:
		return utils.UnmarshalString(casted.Value, Deal_Entry_value, DealEntryAlternates, enum)
	case *types.AttributeValueMemberNULL:
		return nil
	case *types.AttributeValueMemberS:
		return utils.UnmarshalString(casted.Value, Deal_Entry_value, DealEntryAlternates, enum)
	default:
		return fmt.Errorf("Attribute value of %T could not be converted to a data.Deal.Entry", value)
	}
}

// Scan attempts to convert an SQL driver value to a new data.Deal.Entry value
func (enum *Deal_Entry) Scan(value interface{}) error {
	return utils.ScanValue(value, Deal_Entry_value, DealEntryAlternates, enum)
}

// MarshalJSON converts a data.Deal.Reason value to a JSON value
func (enum Deal_Reason) MarshalJSON() ([]byte, error) {
	return []byte(utils.MarshalString(enum, Deal_Reason_name, DealReasonMapping, true)), nil
}

// MarshalCSV converts a data.Deal.Reason value to CSV cell value
func (enum Deal_Reason) MarshalCSV() (string, error) {
	return utils.MarshalString(enum, Deal_Reason_name, DealReasonMapping, false), nil
}

// MarshalYAML converts a data.Deal.Reason value to a YAML node value
func (enum Deal_Reason) MarshalYAML() (interface{}, error) {
	return utils.MarshalString(enum, Deal_Reason_name, DealReasonMapping, false), nil
}

// MarshalDynamoDBAttributeValue converts a data.Deal.Reason value to a DynamoDB AttributeValue
func (enum Deal_Reason) MarshalDynamoDBAttributeValue() (types.AttributeValue, error) {
	return &types.AttributeValueMemberS{Value: utils.MarshalString(enum, Deal_Reason_name, DealReasonMapping, false)}, nil
}

// Value converts a data.Deal.Reason value to an SQL driver value
func (enum Deal_Reason) Value() (driver.Value, error) {
	return utils.MarshalString(enum, Deal_Reason_name, DealReasonMapping, false), nil
}

// UnmarshalJSON attempts to convert a JSON value to a new data.Deal.Reason value
func (enum *Deal_Reason) UnmarshalJSON(raw []byte) error {
	return utils.UnmarshalValue(raw, Deal_Reason_value, DealReasonAlternates, enum)
}

// UnmarshalCSV attempts to convert a CSV cell value to a new data.Deal.Reason value
func (enum *Deal_Reason) UnmarshalCSV(raw string) error {
	return utils.UnmarshalString(raw, Deal_Reason_value, DealReasonAlternates, enum)
}

// UnmarshalYAML attempts to convert a YAML node to a new data.Deal.Reason value
func (enum *Deal_Reason) UnmarshalYAML(value *yaml.Node) error {
	if value.Kind != yaml.ScalarNode {
		return fmt.Errorf("YAML node had an invalid kind (expected scalar value)")
	} else {
		return utils.UnmarshalString(value.Value, Deal_Reason_value, DealReasonAlternates, enum)
	}
}

// UnmarshalDynamoDBAttributeValue attempts to convert a DynamoDB AttributeVAlue to a data.Deal.Reason
// value. This function can handle []bytes, numerics, or strings. If the AttributeValue is NULL then
// the Deal_Reason value will not be modified.
func (enum *Deal_Reason) UnmarshalDynamoDBAttributeValue(value types.AttributeValue) error {
	switch casted := value.(type) {
	case *types.AttributeValueMemberB:
		return utils.UnmarshalValue(casted.Value, Deal_Reason_value, DealReasonAlternates, enum)
	case *types.AttributeValueMemberN:
		return utils.UnmarshalString(casted.Value, Deal_Reason_value, DealReasonAlternates, enum)
	case *types.AttributeValueMemberNULL:
		return nil
	case *types.AttributeValueMemberS:
		return utils.UnmarshalString(casted.Value, Deal_Reason_value, DealReasonAlternates, enum)
	default:
		return fmt.Errorf("Attribute value of %T could not be converted to a data.Deal.Reason", value)
	}
}

// Scan attempts to convert an SQL driver value to a new data.Deal.Reason value
func (enum *Deal_Reason) Scan(value interface{}) error {
	return utils.ScanValue(value, Deal_Reason_value, DealReasonAlternates, enum)
}

// MarshalJSON converts a data.Order.Type value to a JSON value
func (enum Order_Type) MarshalJSON() ([]byte, error) {
	return []byte(utils.MarshalString(enum, Order_Type_name, utils.Ignore, true)), nil
}

// MarshalCSV converts a data.Order.Type value to CSV cell value
func (enum Order_Type) MarshalCSV() (string, error) {
	return utils.MarshalString(enum, Order_Type_name, utils.Ignore, false), nil
}

// MarshalYAML converts a data.Order.Type value to a YAML node value
func (enum Order_Type) MarshalYAML() (interface{}, error) {
	return utils.MarshalString(enum, Order_Type_name, utils.Ignore, false), nil
}

// MarshalDynamoDBAttributeValue converts a data.Order.Type value to a DynamoDB AttributeValue
func (enum Order_Type) MarshalDynamoDBAttributeValue() (types.AttributeValue, error) {
	return &types.AttributeValueMemberS{Value: utils.MarshalString(enum, Order_Type_name, utils.Ignore, false)}, nil
}

// Value converts a data.Order.Type value to an SQL driver value
func (enum Order_Type) Value() (driver.Value, error) {
	return utils.MarshalString(enum, Order_Type_name, utils.Ignore, false), nil
}

// UnmarshalJSON attempts to convert a JSON value to a new data.Order.Type value
func (enum *Order_Type) UnmarshalJSON(raw []byte) error {
	return utils.UnmarshalValue(raw, Order_Type_value, utils.None, enum)
}

// UnmarshalCSV attempts to convert a CSV cell value to a new data.Order.Type value
func (enum *Order_Type) UnmarshalCSV(raw string) error {
	return utils.UnmarshalString(raw, Order_Type_value, utils.None, enum)
}

// UnmarshalYAML attempts to convert a YAML node to a new data.Order.Type value
func (enum *Order_Type) UnmarshalYAML(value *yaml.Node) error {
	if value.Kind != yaml.ScalarNode {
		return fmt.Errorf("YAML node had an invalid kind (expected scalar value)")
	} else {
		return utils.UnmarshalString(value.Value, Order_Type_value, utils.None, enum)
	}
}

// UnmarshalDynamoDBAttributeValue attempts to convert a DynamoDB AttributeVAlue to a data.Order.Type
// value. This function can handle []bytes, numerics, or strings. If the AttributeValue is NULL then
// the Order_Type value will not be modified.
func (enum *Order_Type) UnmarshalDynamoDBAttributeValue(value types.AttributeValue) error {
	switch casted := value.(type) {
	case *types.AttributeValueMemberB:
		return utils.UnmarshalValue(casted.Value, Order_Type_value, utils.None, enum)
	case *types.AttributeValueMemberN:
		return utils.UnmarshalString(casted.Value, Order_Type_value, utils.None, enum)
	case *types.AttributeValueMemberNULL:
		return nil
	case *types.AttributeValueMemberS:
		return utils.UnmarshalString(casted.Value, Order_Type_value, utils.None, enum)
	default:
		return fmt.Errorf("Attribute value of %T could not be converted to a data.Order.Type", value)
	}
}

// Scan attempts to convert an SQL driver value to a new data.Order.Type value
func (enum *Order_Type) Scan(value interface{}) error {
	return utils.ScanValue(value, Order_Type_value, utils.None, enum)
}

// MarshalJSON converts a data.Order.Status value to a JSON value
func (enum Order_Status) MarshalJSON() ([]byte, error) {
	return []byte(utils.MarshalString(enum, Order_Status_name, utils.Ignore, true)), nil
}

// MarshalCSV converts a data.Order.Status value to CSV cell value
func (enum Order_Status) MarshalCSV() (string, error) {
	return utils.MarshalString(enum, Order_Status_name, utils.Ignore, false), nil
}

// MarshalYAML converts a data.Order.Status value to a YAML node value
func (enum Order_Status) MarshalYAML() (interface{}, error) {
	return utils.MarshalString(enum, Order_Status_name, utils.Ignore, false), nil
}

// MarshalDynamoDBAttributeValue converts a data.Order.Status value to a DynamoDB AttributeValue
func (enum Order_Status) MarshalDynamoDBAttributeValue() (types.AttributeValue, error) {
	return &types.AttributeValueMemberS{Value: utils.MarshalString(enum, Order_Status_name, utils.Ignore, false)}, nil
}

// Value converts a data.Order.Status value to an SQL driver value
func (enum Order_Status) Value() (driver.Value, error) {
	return utils.MarshalString(enum, Order_Status_name, utils.Ignore, false), nil
}

// UnmarshalJSON attempts to convert a JSON value to a new data.Order.Status value
func (enum *Order_Status) UnmarshalJSON(raw []byte) error {
	return utils.UnmarshalValue(raw, Order_Status_value, utils.None, enum)
}

// UnmarshalCSV attempts to convert a CSV cell value to a new data.Order.Status value
func (enum *Order_Status) UnmarshalCSV(raw string) error {
	return utils.UnmarshalString(raw, Order_Status_value, utils.None, enum)
}

// UnmarshalYAML attempts to convert a YAML node to a new data.Order.Status value
func (enum *Order_Status) UnmarshalYAML(value *yaml.Node) error {
	if value.Kind != yaml.ScalarNode {
		return fmt.Errorf("YAML node had an invalid kind (expected scalar value)")
	} else {
		return utils.UnmarshalString(value.Value, Order_Status_value, utils.None, enum)
	}
}

// UnmarshalDynamoDBAttributeValue attempts to convert a DynamoDB AttributeVAlue to a data.Order.Status
// value. This function can handle []bytes, numerics, or strings. If the AttributeValue is NULL then
// the Order_Status value will not be modified.
func (enum *Order_Status) UnmarshalDynamoDBAttributeValue(value types.AttributeValue) error {
	switch casted := value.(type) {
	case *types.AttributeValueMemberB:
		return utils.UnmarshalValue(casted.Value, Order_Status_value, utils.None, enum)
	case *types.AttributeValueMemberN:
		return utils.UnmarshalString(casted.Value, Order_Status_value, utils.None, enum)
	case *types.AttributeValueMemberNULL:
		return nil
	case *types.AttributeValueMemberS:
		return utils.UnmarshalString(casted.Value, Order_Status_value, utils.None, enum)
	default:
		return fmt.Errorf("Attribute value of %T could not be converted to a data.Order.Status", value)
	}
}

// Scan attempts to convert an SQL driver value to a new data.Order.Status value
func (enum *Order_Status) Scan(value interface{}) error {
	return utils.ScanValue(value, Order_Status_value, utils.None, enum)
}

// MarshalJSON converts a data.Order.Reason value to a JSON value
func (enum Order_Reason) MarshalJSON() ([]byte, error) {
	return []byte(utils.MarshalString(enum, Order_Reason_name, OrderReasonMapping, true)), nil
}

// MarshalCSV converts a data.Order.Reason value to CSV cell value
func (enum Order_Reason) MarshalCSV() (string, error) {
	return utils.MarshalString(enum, Order_Reason_name, OrderReasonMapping, false), nil
}

// MarshalYAML converts a data.Order.Reason value to a YAML node value
func (enum Order_Reason) MarshalYAML() (interface{}, error) {
	return utils.MarshalString(enum, Order_Reason_name, OrderReasonMapping, false), nil
}

// MarshalDynamoDBAttributeValue converts a data.Order.Reason value to a DynamoDB AttributeValue
func (enum Order_Reason) MarshalDynamoDBAttributeValue() (types.AttributeValue, error) {
	return &types.AttributeValueMemberS{Value: utils.MarshalString(enum, Order_Reason_name, OrderReasonMapping, false)}, nil
}

// Value converts a data.Order.Reason value to an SQL driver value
func (enum Order_Reason) Value() (driver.Value, error) {
	return utils.MarshalString(enum, Order_Reason_name, OrderReasonMapping, false), nil
}

// UnmarshalJSON attempts to convert a JSON value to a new data.Order.Reason value
func (enum *Order_Reason) UnmarshalJSON(raw []byte) error {
	return utils.UnmarshalValue(raw, Order_Reason_value, OrderReasonAlternates, enum)
}

// UnmarshalCSV attempts to convert a CSV cell value to a new data.Order.Reason value
func (enum *Order_Reason) UnmarshalCSV(raw string) error {
	return utils.UnmarshalString(raw, Order_Reason_value, OrderReasonAlternates, enum)
}

// UnmarshalYAML attempts to convert a YAML node to a new data.Order.Reason value
func (enum *Order_Reason) UnmarshalYAML(value *yaml.Node) error {
	if value.Kind != yaml.ScalarNode {
		return fmt.Errorf("YAML node had an invalid kind (expected scalar value)")
	} else {
		return utils.UnmarshalString(value.Value, Order_Reason_value, OrderReasonAlternates, enum)
	}
}

// UnmarshalDynamoDBAttributeValue attempts to convert a DynamoDB AttributeVAlue to a data.Order.Reason
// value. This function can handle []bytes, numerics, or strings. If the AttributeValue is NULL then
// the Order_Reason value will not be modified.
func (enum *Order_Reason) UnmarshalDynamoDBAttributeValue(value types.AttributeValue) error {
	switch casted := value.(type) {
	case *types.AttributeValueMemberB:
		return utils.UnmarshalValue(casted.Value, Order_Reason_value, OrderReasonAlternates, enum)
	case *types.AttributeValueMemberN:
		return utils.UnmarshalString(casted.Value, Order_Reason_value, OrderReasonAlternates, enum)
	case *types.AttributeValueMemberNULL:
		return nil
	case *types.AttributeValueMemberS:
		return utils.UnmarshalString(casted.Value, Order_Reason_value, OrderReasonAlternates, enum)
	default:
		return fmt.Errorf("Attribute value of %T could not be converted to a data.Order.Reason", value)
	}
}

// Scan attempts to convert an SQL driver value to a new data.Order.Reason value
func (enum *Order_Reason) Scan(value interface{}) error {
	return utils.ScanValue(value, Order_Reason_value, OrderReasonAlternates, enum)
}

// MarshalJSON converts a data.Position.Type value to a JSON value
func (enum Position_Type) MarshalJSON() ([]byte, error) {
	return []byte(utils.MarshalString(enum, Position_Type_name, utils.Ignore, true)), nil
}

// MarshalCSV converts a data.Position.Type value to CSV cell value
func (enum Position_Type) MarshalCSV() (string, error) {
	return utils.MarshalString(enum, Position_Type_name, utils.Ignore, false), nil
}

// MarshalYAML converts a data.Position.Type value to a YAML node value
func (enum Position_Type) MarshalYAML() (interface{}, error) {
	return utils.MarshalString(enum, Position_Type_name, utils.Ignore, false), nil
}

// MarshalDynamoDBAttributeValue converts a data.Position.Type value to a DynamoDB AttributeValue
func (enum Position_Type) MarshalDynamoDBAttributeValue() (types.AttributeValue, error) {
	return &types.AttributeValueMemberS{Value: utils.MarshalString(enum, Position_Type_name, utils.Ignore, false)}, nil
}

// Value converts a data.Position.Type value to an SQL driver value
func (enum Position_Type) Value() (driver.Value, error) {
	return utils.MarshalString(enum, Position_Type_name, utils.Ignore, false), nil
}

// UnmarshalJSON attempts to convert a JSON value to a new data.Position.Type value
func (enum *Position_Type) UnmarshalJSON(raw []byte) error {
	return utils.UnmarshalValue(raw, Position_Type_value, utils.None, enum)
}

// UnmarshalCSV attempts to convert a CSV cell value to a new data.Position.Type value
func (enum *Position_Type) UnmarshalCSV(raw string) error {
	return utils.UnmarshalString(raw, Position_Type_value, utils.None, enum)
}

// UnmarshalYAML attempts to convert a YAML node to a new data.Position.Type value
func (enum *Position_Type) UnmarshalYAML(value *yaml.Node) error {
	if value.Kind != yaml.ScalarNode {
		return fmt.Errorf("YAML node had an invalid kind (expected scalar value)")
	} else {
		return utils.UnmarshalString(value.Value, Position_Type_value, utils.None, enum)
	}
}

// UnmarshalDynamoDBAttributeValue attempts to convert a DynamoDB AttributeVAlue to a data.Position.Type
// value. This function can handle []bytes, numerics, or strings. If the AttributeValue is NULL then
// the Position_Type value will not be modified.
func (enum *Position_Type) UnmarshalDynamoDBAttributeValue(value types.AttributeValue) error {
	switch casted := value.(type) {
	case *types.AttributeValueMemberB:
		return utils.UnmarshalValue(casted.Value, Position_Type_value, utils.None, enum)
	case *types.AttributeValueMemberN:
		return utils.UnmarshalString(casted.Value, Position_Type_value, utils.None, enum)
	case *types.AttributeValueMemberNULL:
		return nil
	case *types.AttributeValueMemberS:
		return utils.UnmarshalString(casted.Value, Position_Type_value, utils.None, enum)
	default:
		return fmt.Errorf("Attribute value of %T could not be converted to a data.Position.Type", value)
	}
}

// Scan attempts to convert an SQL driver value to a new data.Position.Type value
func (enum *Position_Type) Scan(value interface{}) error {
	return utils.ScanValue(value, Position_Type_value, utils.None, enum)
}

// MarshalJSON converts a data.Position.Reason value to a JSON value
func (enum Position_Reason) MarshalJSON() ([]byte, error) {
	return []byte(utils.MarshalString(enum, Position_Reason_name, utils.Ignore, true)), nil
}

// MarshalCSV converts a data.Position.Reason value to CSV cell value
func (enum Position_Reason) MarshalCSV() (string, error) {
	return utils.MarshalString(enum, Position_Reason_name, utils.Ignore, false), nil
}

// MarshalYAML converts a data.Position.Reason value to a YAML node value
func (enum Position_Reason) MarshalYAML() (interface{}, error) {
	return utils.MarshalString(enum, Position_Reason_name, utils.Ignore, false), nil
}

// MarshalDynamoDBAttributeValue converts a data.Position.Reason value to a DynamoDB AttributeValue
func (enum Position_Reason) MarshalDynamoDBAttributeValue() (types.AttributeValue, error) {
	return &types.AttributeValueMemberS{Value: utils.MarshalString(enum, Position_Reason_name, utils.Ignore, false)}, nil
}

// Value converts a data.Position.Reason value to an SQL driver value
func (enum Position_Reason) Value() (driver.Value, error) {
	return utils.MarshalString(enum, Position_Reason_name, utils.Ignore, false), nil
}

// UnmarshalJSON attempts to convert a JSON value to a new data.Position.Reason value
func (enum *Position_Reason) UnmarshalJSON(raw []byte) error {
	return utils.UnmarshalValue(raw, Position_Reason_value, utils.None, enum)
}

// UnmarshalCSV attempts to convert a CSV cell value to a new data.Position.Reason value
func (enum *Position_Reason) UnmarshalCSV(raw string) error {
	return utils.UnmarshalString(raw, Position_Reason_value, utils.None, enum)
}

// UnmarshalYAML attempts to convert a YAML node to a new data.Position.Reason value
func (enum *Position_Reason) UnmarshalYAML(value *yaml.Node) error {
	if value.Kind != yaml.ScalarNode {
		return fmt.Errorf("YAML node had an invalid kind (expected scalar value)")
	} else {
		return utils.UnmarshalString(value.Value, Position_Reason_value, utils.None, enum)
	}
}

// UnmarshalDynamoDBAttributeValue attempts to convert a DynamoDB AttributeVAlue to a data.Position.Reason
// value. This function can handle []bytes, numerics, or strings. If the AttributeValue is NULL then
// the Position_Reason value will not be modified.
func (enum *Position_Reason) UnmarshalDynamoDBAttributeValue(value types.AttributeValue) error {
	switch casted := value.(type) {
	case *types.AttributeValueMemberB:
		return utils.UnmarshalValue(casted.Value, Position_Reason_value, utils.None, enum)
	case *types.AttributeValueMemberN:
		return utils.UnmarshalString(casted.Value, Position_Reason_value, utils.None, enum)
	case *types.AttributeValueMemberNULL:
		return nil
	case *types.AttributeValueMemberS:
		return utils.UnmarshalString(casted.Value, Position_Reason_value, utils.None, enum)
	default:
		return fmt.Errorf("Attribute value of %T could not be converted to a data.Position.Reason", value)
	}
}

// Scan attempts to convert an SQL driver value to a new data.Position.Reason value
func (enum *Position_Reason) Scan(value interface{}) error {
	return utils.ScanValue(value, Position_Reason_value, utils.None, enum)
}

// MarshalJSON converts a data.TradeRequest.Action value to a JSON value
func (enum TradeRequest_Action) MarshalJSON() ([]byte, error) {
	return []byte(utils.MarshalString(enum, TradeRequest_Action_name, utils.Ignore, true)), nil
}

// MarshalCSV converts a data.TradeRequest.Action value to CSV cell value
func (enum TradeRequest_Action) MarshalCSV() (string, error) {
	return utils.MarshalString(enum, TradeRequest_Action_name, utils.Ignore, false), nil
}

// MarshalYAML converts a data.TradeRequest.Action value to a YAML node value
func (enum TradeRequest_Action) MarshalYAML() (interface{}, error) {
	return utils.MarshalString(enum, TradeRequest_Action_name, utils.Ignore, false), nil
}

// MarshalDynamoDBAttributeValue converts a data.TradeRequest.Action value to a DynamoDB AttributeValue
func (enum TradeRequest_Action) MarshalDynamoDBAttributeValue() (types.AttributeValue, error) {
	return &types.AttributeValueMemberS{Value: utils.MarshalString(enum, TradeRequest_Action_name, utils.Ignore, false)}, nil
}

// Value converts a data.TradeRequest.Action value to an SQL driver value
func (enum TradeRequest_Action) Value() (driver.Value, error) {
	return utils.MarshalString(enum, TradeRequest_Action_name, utils.Ignore, false), nil
}

// UnmarshalJSON attempts to convert a JSON value to a new data.TradeRequest.Action value
func (enum *TradeRequest_Action) UnmarshalJSON(raw []byte) error {
	return utils.UnmarshalValue(raw, TradeRequest_Action_value, utils.None, enum)
}

// UnmarshalCSV attempts to convert a CSV cell value to a new data.TradeRequest.Action value
func (enum *TradeRequest_Action) UnmarshalCSV(raw string) error {
	return utils.UnmarshalString(raw, TradeRequest_Action_value, utils.None, enum)
}

// UnmarshalYAML attempts to convert a YAML node to a new data.TradeRequest.Action value
func (enum *TradeRequest_Action) UnmarshalYAML(value *yaml.Node) error {
	if value.Kind != yaml.ScalarNode {
		return fmt.Errorf("YAML node had an invalid kind (expected scalar value)")
	} else {
		return utils.UnmarshalString(value.Value, TradeRequest_Action_value, utils.None, enum)
	}
}

// UnmarshalDynamoDBAttributeValue attempts to convert a DynamoDB AttributeVAlue to a data.TradeRequest.Action
// value. This function can handle []bytes, numerics, or strings. If the AttributeValue is NULL then
// the TradeRequest_Action value will not be modified.
func (enum *TradeRequest_Action) UnmarshalDynamoDBAttributeValue(value types.AttributeValue) error {
	switch casted := value.(type) {
	case *types.AttributeValueMemberB:
		return utils.UnmarshalValue(casted.Value, TradeRequest_Action_value, utils.None, enum)
	case *types.AttributeValueMemberN:
		return utils.UnmarshalString(casted.Value, TradeRequest_Action_value, utils.None, enum)
	case *types.AttributeValueMemberNULL:
		return nil
	case *types.AttributeValueMemberS:
		return utils.UnmarshalString(casted.Value, TradeRequest_Action_value, utils.None, enum)
	default:
		return fmt.Errorf("Attribute value of %T could not be converted to a data.TradeRequest.Action", value)
	}
}

// Scan attempts to convert an SQL driver value to a new data.TradeRequest.Action value
func (enum *TradeRequest_Action) Scan(value interface{}) error {
	return utils.ScanValue(value, TradeRequest_Action_value, utils.None, enum)
}
