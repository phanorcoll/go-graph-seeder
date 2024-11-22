// Package generator provides generator  
package generator

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

// EnumValue represents a single value in an enum with a weight.
type EnumValue struct {
	Value  string  `json:"value"`
	Weight float64 `json:"weight"`
}

// random is a local random generator instance.
var random = rand.New(rand.NewSource(time.Now().UnixNano()))

func init() {
	gofakeit.Seed(0)
}

// GenerateValue generates random data based on the property type.
func GenerateValue(propType string, enumValues []EnumValue) interface{} {
	switch propType {
	case "email":
		return gofakeit.Email()
	case "fname":
		return gofakeit.FirstName()
	case "lname":
		return gofakeit.LastName()
	case "address":
		return gofakeit.Address().Address
	case "city":
		return gofakeit.City()
	case "state":
		return gofakeit.State()
	case "country":
		return gofakeit.Country()
	case "postal":
		return gofakeit.Zip()
	case "phone":
		return gofakeit.Phone()
	case "job":
		return gofakeit.JobTitle()
	case "company":
		return gofakeit.Company()
	case "trade":
		return gofakeit.Word()
	case "trade sector":
		return gofakeit.Sentence(2)
	case "secondary_trade_sector":
		return gofakeit.Sentence(2)
	case "language":
		return gofakeit.Language()
	case "date":
		return gofakeit.Date().Format("2006-01-02T15:04:05Z")
	case "boolean":
		return gofakeit.Bool()
	case "name":
		return gofakeit.MovieName()
	case "size":
		return fmt.Sprintf("%d", gofakeit.Number(1, 1000))
	case "id":
		return gofakeit.UUID()
	case "timestamp":
		return time.Now().Unix()
	case "enum":
		return generateEnumValue(enumValues)
	default:
		return gofakeit.Sentence(2)
	}
}

// GenerateNodes creates nodes with properties based on the template.
func GenerateNodes(properties map[string]interface{}, start, end int) []map[string]interface{} {
	nodes := make([]map[string]interface{}, 0, end-start)

	for i := start; i < end; i++ {
		node := map[string]interface{}{}
		for key, prop := range properties {
			propMap, isMap := prop.(map[string]interface{})
			if isMap && propMap["type"] == "enum" {
				enumValues := parseEnumValues(propMap["values"])
				node[key] = GenerateValue("enum", enumValues)
			} else if key == "label" {
				node["label"] = prop
			} else if key == "account_type" {
				node[key] = prop
			} else {
				node[key] = GenerateValue(prop.(string), nil)
			}
		}
		nodes = append(nodes, node)
	}
	return nodes
}

// func GenerateNodes(properties map[string]interface{}, start, end int) []map[string]interface{} {
// 	nodes := make([]map[string]interface{}, 0, end-start)
// 	for i := start; i < end; i++ {
// 		node := map[string]interface{}{}
// 		for key, propType := range properties {
// 			propMap, isMap := propType.(map[string]interface{})
// 			if isMap && propMap["type"] == "enum" {
// 				// Handle enum type
// 				enumValues := parseEnumValues(propMap["values"])
// 				node[key] = GenerateValue("enum", enumValues)
// 			} else if key == "label" {
// 				node["label"] = propType
// 			} else if key == "account_type" {
// 				node[key] = propType
// 			} else {
// 				node[key] = GenerateValue(propType.(string), nil)
// 			}
// 		}
// 		nodes = append(nodes, node)
// 	}
// 	return nodes
// }

// GeneratePersistentIDs generates a set of identifiers including email, name combination, and a UUID.
func GeneratePersistentIDs(email, firstName, lastName string) []interface{} {
	persistentIDs := []interface{}{
		email, // Add the email
		fmt.Sprintf("%s_%s", firstName, lastName), // Add first and last name combination
		gofakeit.UUID(), // Add a random UUID
	}
	return persistentIDs
}

// generateEnumValue selects a random value from the enum based on its weight.
func generateEnumValue(enumValues []EnumValue) string {
	// Generate a weighted list where values appear proportional to their weights
	weightedValues := []string{}
	for _, enum := range enumValues {
		count := int(enum.Weight * 10) // Scale weight to create proportional entries
		for i := 0; i < count; i++ {
			weightedValues = append(weightedValues, enum.Value)
		}
	}

	// Pick a random value from the weighted list
	return weightedValues[random.Intn(len(weightedValues))]
}

// parseEnumValues parses the enum values from the configuration.
func parseEnumValues(values interface{}) []EnumValue {
	rawValues, ok := values.([]interface{})
	if !ok {
		return nil
	}

	enumValues := make([]EnumValue, 0, len(rawValues))
	for _, rawValue := range rawValues {
		valueMap, ok := rawValue.(map[string]interface{})
		if !ok {
			continue
		}
		enumValues = append(enumValues, EnumValue{
			Value:  valueMap["value"].(string),
			Weight: valueMap["weight"].(float64),
		})
	}
	return enumValues
}
