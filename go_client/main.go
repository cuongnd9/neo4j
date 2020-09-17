package main

import (
	"fmt"

	"github.com/103cuong/uid"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

func main() {
	configForNeo4j := func(conf *neo4j.Config) {
		conf.Encrypted = false
	}

	driver, err := neo4j.NewDriver("bolt://0.0.0.0:7687", neo4j.BasicAuth("neo4j", "123456", ""), configForNeo4j)
	if err != nil {
		panic("connect neo4j failed")
	}
	defer driver.Close()

	// For multidatabase support, set sessionConfig.DatabaseName to requested database
	sessionConfig := neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite}
	session, err := driver.NewSession(sessionConfig)
	if err != nil {
		panic("config session failed")
	}
	defer session.Close()

	result, err := session.Run("CREATE (c:Cat { id: $id, name: $name, color: $color }) RETURN c.id, c.name, c.color", map[string]interface{}{
		"id":    uid.Uid(16),
		"name":  "Kiki",
		"color": "white",
	})
	if err != nil {
		panic("create cat failed")
	}

	for result.Next() {
		fmt.Printf("😻: id=%s, name=%s, color=%s\n", result.Record().GetByIndex(0).(string), result.Record().GetByIndex(1).(string), result.Record().GetByIndex(2).(string))
	}

	fmt.Printf("result error: %v", result.Err())
}
