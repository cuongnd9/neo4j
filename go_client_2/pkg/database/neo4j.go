package database

import (
	"github.com/103cuong/neo4j/go_client_crud/configs"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

var Session neo4j.Session

func ConnectToDB() (neo4j.Driver, error) {
	var (
		driver neo4j.Driver
		err error
	)
	neo4jConfig := configs.BuildNeo4jConfig()
	if driver, err = neo4j.NewDriver(
		neo4jConfig.Target,
		neo4j.BasicAuth(neo4jConfig.Username, neo4jConfig.Password, neo4jConfig.Realm),
		func(conf *neo4j.Config) {
			conf.Encrypted = false
		},
	); err != nil {
		return nil, err
	}
	if Session, err = driver.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeWrite,
	}); err != nil {
		return nil, err
	}
	return driver, nil
}
