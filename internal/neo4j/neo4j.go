package neo4j

import (
	"graphdb/pkg/database"
)

type Neo4jDb struct {
}

func NewDatabaseService() database.Service {
	return Neo4jDb{}
}
