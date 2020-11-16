package database

import (
	"graphdb/pkg/data"
)

type Service interface {
	Clear() error
	AddNode(node data.Node) error
	AddConnection(connection data.Connection) error
	GetNodeByName(name string) (data.Node, error)
	GetConnectedNodesWithType(fromID int, connectionType string) ([]data.Node, error)
}
