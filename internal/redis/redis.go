package redis

import (
	"graphdb/pkg/database"

	"github.com/gomodule/redigo/redis"
	rg "github.com/redislabs/redisgraph-go"
)

type RedisDb struct {
}

func NewDatabaseService() database.Service {
	conn, _ := redis.Dial("tcp", "127.0.0.1:6379")
	defer conn.Close()

	graph := rg.GraphNew("social", conn)

	graph.Delete()

	john := rg.Node{
		Label: "person",
		Properties: map[string]interface{}{
			"name":   "John Doe",
			"age":    33,
			"gender": "male",
			"status": "single",
		},
	}
	graph.AddNode(&john)

	japan := rg.Node{
		Label: "country",
		Properties: map[string]interface{}{
			"name": "Japan",
		},
	}
	graph.AddNode(&japan)

	edge := rg.Edge{
		Source:      &john,
		Relation:    "visited",
		Destination: &japan,
	}
	graph.AddEdge(&edge)

	graph.Commit()
	return RedisDb{}
}
