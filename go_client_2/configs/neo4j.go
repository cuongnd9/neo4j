package configs

import (
	"gopkg.in/mcuadros/go-defaults.v1"
	"os"
)

type Neo4jConfig struct {
	Target string `default:"bolt://0.0.0.0:7687"`
	Username string `default:"neo4j"`
	Password string `default:"123456"`
	Realm string `default:""`
}

func BuildNeo4jConfig() *Neo4jConfig {
	config := Neo4jConfig{
		Target:   os.Getenv("NEO4J_TARGET"),
		Username: os.Getenv("NEO4J_USERNAME"),
		Password: os.Getenv("NEO4J_PASSWORD"),
		Realm:    os.Getenv("NEO4J_REALM"),
	}
	defaults.SetDefaults(&config)
	return &config
}
