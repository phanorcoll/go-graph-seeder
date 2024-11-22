// Package generator provides generator  
package generator

import (
	"fmt"
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

func init() {
	gofakeit.Seed(0)
}

// GenerateValue generates random data based on the property type.
func GenerateValue(propType string) interface{} {
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
	default:
		return gofakeit.Sentence(2)
	}
}

// GenerateNodes creates nodes with properties based on the template.
func GenerateNodes(properties map[string]string, start, end int) []map[string]interface{} {
	nodes := make([]map[string]interface{}, 0, end-start)
	for i := start; i < end; i++ {
		node := map[string]interface{}{}
		for key, propType := range properties {
			if key == "label" {
				node["label"] = propType
			} else if key == "account_type" {
				node[key] = propType
			} else {
				node[key] = GenerateValue(propType)
			}
		}
		nodes = append(nodes, node)
	}
	return nodes
}

// GeneratePersistentIDs generates a set of identifiers including email, name combination, and a UUID.
func GeneratePersistentIDs(email, firstName, lastName string) []interface{} {
	persistentIDs := []interface{}{
		email, // Add the email
		fmt.Sprintf("%s_%s", firstName, lastName), // Add first and last name combination
		gofakeit.UUID(), // Add a random UUID
	}
	return persistentIDs
}
