// Package main provides main  
package main

import (
	"fmt"
	"log"
	"strings"

	"satoru-itadori/go-graph-seeder/config"
	"satoru-itadori/go-graph-seeder/gremlin"
)

func main() {
	// Load the configuration
	template, err := config.LoadTemplate("template.json")
	if err != nil {
		log.Fatalf("Failed to load template: %v", err)
	}

	// Connect to Gremlin server
	g, connection, err := gremlin.NewConnection("ws://localhost:8182/gremlin")
	if err != nil {
		log.Fatalf("Failed to connect to Gremlin server: %v", err)
	}
	defer connection.Close()

	results := make(map[string]int)
	createdNodes := make(map[string][]interface{})

	// Create start nodes
	for _, startNode := range template.StartNodes {
		nodeIDs, err := gremlin.CreateNodesInBatches(g, startNode)
		if err != nil {
			log.Fatalf("Failed to create start nodes: %v", err)
		}
		label := startNode.NodeProperties["label"]
		results[label] += startNode.NodeCount
		createdNodes[label] = append(createdNodes[label], nodeIDs...)
	}

	// Create end nodes
	for _, endNode := range template.EndNodes {
		nodeIDs, err := gremlin.CreateNodesInBatches(g, endNode)
		if err != nil {
			log.Fatalf("Failed to create end nodes: %v", err)
		}
		label := endNode.NodeProperties["label"]
		results[label] += endNode.NodeCount
		createdNodes[label] = append(createdNodes[label], nodeIDs...)
	}

	// Create relationships
	for _, startNode := range template.StartNodes {
		startLabel := startNode.NodeProperties["label"]
		for _, relationship := range startNode.Relationships {
			endLabel := relationship.EndNode
			edgeLabel := relationship.Edge

			// Connect start nodes to end nodes
			for _, startID := range createdNodes[startLabel] {
				for _, endID := range createdNodes[endLabel] {
					err := gremlin.ConnectNodes(g, startID, endID, edgeLabel)
					if err != nil {
						log.Printf(
							"Failed to connect %s to %s with edge %s: %v",
							startLabel,
							endLabel,
							edgeLabel,
							err,
						)
					}
				}
			}
		}
	}

	// Print the results
	column1 := "Vertex"
	column2 := "# Created"
	col1Width := len(column1)
	col2Width := len(column2)

	for label, count := range results {
		if len(label) > col1Width {
			col1Width = len(label)
		}
		if len(fmt.Sprintf("%d", count)) > col2Width {
			col2Width = len(fmt.Sprintf("%d", count))
		}
	}

	// Draw the table
	fmt.Printf("\n%-*s %-*s\n", col1Width, column1, col2Width, column2)
	fmt.Println(strings.Repeat("-", col1Width+col2Width+1))
	for label, count := range results {
		fmt.Printf("%-*s %-*d\n", col1Width, label, col2Width, count)
	}
	fmt.Println()
	log.Printf("All nodes created successfully.")
}
