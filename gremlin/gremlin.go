// Package gremlin provides gremlin  
package gremlin

import (
	"fmt"
	"log"
	"phanorcoll/go-graph-seeder/generator"
	"sync"

	"phanorcoll/go-graph-seeder/config"

	gremlingo "github.com/apache/tinkerpop/gremlin-go/v3/driver"
)

// NewConnection creates a new connection to the Gremlin server.
func NewConnection(
	url string,
) (*gremlingo.GraphTraversalSource, *gremlingo.DriverRemoteConnection, error) {
	// Establish a remote connection
	driverRemoteConnection, err := gremlingo.NewDriverRemoteConnection(url)
	if err != nil {
		return nil, nil, err
	}

	// Create a traversal source with the remote connection
	g := gremlingo.Traversal_().WithRemote(driverRemoteConnection)
	return g, driverRemoteConnection, nil
}

// CreateNodesInBatches creates nodes in batches to optimize performance.
func CreateNodesInBatches(
	g *gremlingo.GraphTraversalSource,
	nodeTemplate config.NodeTemplate,
) ([]interface{}, error) {
	batchSize := 1000
	numBatches := nodeTemplate.NodeCount / batchSize
	if nodeTemplate.NodeCount%batchSize != 0 {
		numBatches++
	}

	var allNodeIDs []interface{}
	var batchWg sync.WaitGroup
	batchNodeIDsChan := make(chan []interface{}, numBatches)

	for batch := 0; batch < numBatches; batch++ {
		batchWg.Add(1)
		go func(batch int) {
			defer batchWg.Done()
			start := batch * batchSize
			end := start + batchSize
			if end > nodeTemplate.NodeCount {
				end = nodeTemplate.NodeCount
			}

			nodes := generator.GenerateNodes(nodeTemplate.NodeProperties, start, end)
			nodeIDs, err := AddNodesToGraph(g, nodes)
			if err != nil {
				log.Printf("Failed to add nodes in batch %d: %v", batch, err)
				return
			}
			batchNodeIDsChan <- nodeIDs
		}(batch)
	}

	// Wait for all batches to complete
	batchWg.Wait()
	close(batchNodeIDsChan)

	// Collect node IDs from all batches
	for batchNodeIDs := range batchNodeIDsChan {
		allNodeIDs = append(allNodeIDs, batchNodeIDs...)
	}

	return allNodeIDs, nil
}

// AddNodesToGraph sends a batch of node creation queries to the Gremlin server.
func AddNodesToGraph(
	g *gremlingo.GraphTraversalSource,
	nodes []map[string]interface{},
) ([]interface{}, error) {
	var nodeIDs []interface{}
	for _, node := range nodes {
		label := node["label"].(string)
		traversal := g.AddV(label)

		for key, value := range node {
			if key != "label" {
				traversal = traversal.Property(key, value)
			}
		}

		if label == "identity" {
			email, _ := node["email_address"].(string)
			firstName, _ := node["first_name"].(string)
			lastName, _ := node["last_name"].(string)

			persistentIDs := generator.GeneratePersistentIDs(email, firstName, lastName)
			for _, id := range persistentIDs {
				traversal = traversal.Property("persistent_ids", id)
			}
		}

		result, err := traversal.Next()
		if err != nil {
			return nil, fmt.Errorf("failed to add node with label '%s': %w", label, err)
		}

		vertex, ok := result.GetInterface().(*gremlingo.Vertex)
		if !ok {
			return nil, fmt.Errorf("unexpected result type: %T", result.GetInterface())
		}

		nodeIDs = append(nodeIDs, vertex.Id)
	}
	return nodeIDs, nil
}

// ConnectNodes creates an edge between two nodes with a specified label.
func ConnectNodes(
	g *gremlingo.GraphTraversalSource,
	startNodeID, endNodeID interface{},
	edgeLabel string,
) error {
	traversal := g.V(startNodeID).AddE(edgeLabel).To(g.V(endNodeID)).Iterate()
	err := <-traversal
	if err != nil {
		return fmt.Errorf("failed to connect nodes: %w", err)
	}
	return nil
}
