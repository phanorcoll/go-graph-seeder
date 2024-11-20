// Package config provides config  
package config

import (
	"encoding/json"
	"os"
)

// Relationship struct  
type Relationship struct {
	Edge    string `json:"edge"`
	EndNode string `json:"endNode"`
}

// NodeTemplate struct  
type NodeTemplate struct {
	NodeCount      int               `json:"nodeCount"`
	Relationships  []Relationship    `json:"relationships"`
	NodeProperties map[string]string `json:"nodeProperties"`
}

// Template struct  
type Template struct {
	StartNodes []NodeTemplate `json:"startNodes"`
	EndNodes   []NodeTemplate `json:"endNodes"`
}

// LoadTemplate loads and parses the JSON configuration file.
func LoadTemplate(filePath string) (Template, error) {
	var template Template
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		return template, err
	}

	if err := json.Unmarshal(fileData, &template); err != nil {
		return template, err
	}

	return template, nil
}
