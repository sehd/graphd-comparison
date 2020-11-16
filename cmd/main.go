package main

import (
	"fmt"
	"graphdb/internal/random"
	"graphdb/internal/timer"
	"graphdb/pkg/data"
	"graphdb/pkg/database"
	"math/rand"
	"time"
)

const (
	nodeSize       = 1000
	connectionSize = 10000
)

func main() {
	nodes := createNodes(nodeSize)
	connections := createConnections(nodeSize, connectionSize)
	var databaseService database.Service
	runTimedOperations(nodes, connections, databaseService)
}

func runTimedOperations(nodes []data.Node, connections []data.Connection, dbService database.Service) ([]time.Duration, error) {
	res := []time.Duration{}

	t, err := addNodesTimed(nodes, dbService)
	if err != nil {
		return nil, err
	}
	res = append(res, t)

	t, err = addConnectionsTimed(connections, dbService)
	if err != nil {
		return nil, err
	}
	res = append(res, t)

	t, err = getRandomNodesByNameTimed(nodes, dbService, 1000)
	if err != nil {
		return nil, err
	}
	res = append(res, t)

	return res, nil
}

func addNodesTimed(nodes []data.Node, dbService database.Service) (time.Duration, error) {
	return timer.Timed(func() error {
		for _, node := range nodes {
			if err := dbService.AddNode(node); err != nil {
				return err
			}
		}
		return nil
	})
}

func addConnectionsTimed(connections []data.Connection, dbService database.Service) (time.Duration, error) {
	return timer.Timed(func() error {
		for _, connection := range connections {
			if err := dbService.AddConnection(connection); err != nil {
				return err
			}
		}
		return nil
	})
}

func getRandomNodesByNameTimed(nodes []data.Node, dbService database.Service, queryCount int) (time.Duration, error) {
	randomNames := make([]string, queryCount)
	for i, nodeID := range random.GetRandomSampleWithReplacement(len(nodes), queryCount) {
		randomNames[i] = nodes[nodeID].Name
	}
	return timer.Timed(func() error {
		for _, randomName := range randomNames {
			node, err := dbService.GetNodeByName(randomName)
			if err != nil {
				return err
			}
			if node.Name != randomName {
				return fmt.Errorf("Query returned wrong result. Expected %s Got %s", randomName, node.Name)
			}
		}
		return nil
	})
}

func getConnectedNodesByTypeTimed() (time.Duration, error) {

}

func createNodes(size int) []data.Node {
	res := make([]data.Node, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		res[i] = data.Node{
			ID:    i,
			Name:  random.RandomWord(2, 10),
			Value: rand.Float64(),
		}
	}
	return res
}

func createConnections(nodeSize int, connectionSize int) []data.Connection {
	res := make([]data.Connection, connectionSize)
	types := [...]string{"Type 1", "Type 2", "Type 3"}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < connectionSize; i++ {
		res[i] = data.Connection{
			ID:   i,
			Type: types[rand.Intn(3)],
			From: rand.Intn(nodeSize),
			To:   rand.Intn(nodeSize),
		}
	}
	return res
}
