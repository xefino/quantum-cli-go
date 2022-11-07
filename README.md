[![Unit Tests Status](https://github.com/xefino/quantum-cli-go/actions/workflows/test.yml/badge.svg)](https://github.com/xefino/quantum-cli-go/actions)

# quantum-cli-go
Public-facing definitions of interfaces which may be implemented by Quantum CLI developers so that they can be integrated with the Quantum CLI

## Guidelines

When developing code for this repository, the first few steps should be done inside the protobuf repository. However, once the code has been deployed to this repository, additional work may be required to fully develop the code resources necessary. This section contains guidelines pertinent to that extensibility.

### Generated Files

The files in this repository are generated. Therefore, any file with a `.pb.go` extension will be overwritten by subsequent releases. Therefore, under no circumstances should these files be changed. If changes are necessary, they can be done via the protobuf repository. Otherwise, extensions or utility functions may be written to add functionality as normal `.go` files will not be deleted.

### Releases

As this code is nearly entirely code-generated, updates to this repository are automatic. However, releases still need to be performed manually. Therefore, after changes have been pushed, please ensure that either a pre-release or production release is drafted so that the changes can be consumed by downstream services.

### Utils vs. Extensions

Packages in this repository typically contain one or both of a set of files, called `utils.go` and `extensions.go` respectively. The first of these, `utils.go`, is intended to include utility functions that would be useful for things like marshalling and unmarshalling. The other file, `extensions.go`, is intended to provide extra functionalilty that the standard protobuf files do not allow for. This might include things such as special representation code, conversion to other types, comparison, etc. The goal is for this functionality to be packaged together with the generated Go files so that data received from an RPC endpoint contains this functionality "out of the box".

### Enums

Often, the type of field we want to work with internally will differ from the type of the data being returned. As an example, we may receive a particular field in an integer format that would serve us better as an enum. When this happens, please add a custom JSON marshaller and unmarshaller for the enum. You can do this easily by following this pattern:

```
// JSON Marshaller
func (enum My_Enum) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", enum.String())), nil
}

// SQL Valuer
func (enum My_Enum) Value() (driver.Value, error) {
	return driver.Value(int(enum)), nil
}

// DynamoDB Marshaller
func (enum My_Enum) MarshalDynamoDBAttributeValue() (types.AttributeValue, error) {
	return &types.AttributeValueMemberS{
		Value: enum.String(),
	}, nil
}

// JSON Unmarshaller
func (enum *My_Enum) UnmarshalJSON(data []byte) error {
    return utils.UnmarshalValue(data, My_Enum_value, MyEnumAlternates, enum) // Define this or use utils.None
}

// SQL Scanner
func (enum *My_Enum) Scan(value interface{}) error {
    return utils.ScanValue(value, My_Enum_value, MyEnumAlternates, enum) // Define this or use utils.None
}

// DynamoDB Unmarshaller
func (enum *My_Enum) UnmarshalDynamoDBAttributeValue(value types.AttributeValue) error {
	switch casted := value.(type) {
	case *types.AttributeValueMemberB:
		return utils.UnmarshalValue(casted.Value, My_Enum_value, MyEnumAlternates, enum)
	case *types.AttributeValueMemberN:
		return utils.UnmarshalString(casted.Value, My_Enum_value, MyEnumAlternates, enum)
	case *types.AttributeValueMemberNULL:
		return nil
	case *types.AttributeValueMemberS:
		return utils.UnmarshalString(casted.Value, My_Enum_value, MyEnumAlternates, enum)
	default:
		return fmt.Errorf("Attribute value of %T could not be converted to a my.enum", value)
	}
}
```

In the above code snippets, the MyEnumAlternates is a user-defined `map[string]My_Enum` that contains alternative string representations for enum values and the codes to which they map. This is useful for allowing human-readable enum names, or for fixing naming conflicts, without hurting deserialization. All the functions in utils are located [here](https://github.com/xefino/protobuf-gen-go/blob/main
