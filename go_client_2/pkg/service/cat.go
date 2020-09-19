package service

import (
	"github.com/103cuong/neo4j/go_client_crud/pkg/database"
	"github.com/103cuong/neo4j/go_client_crud/pkg/model"
	"sync"
)

func CreateCat(cat model.Cat, c chan error, mutex *sync.Mutex)  {
	mutex.Lock()

	cypher := `
		CREATE (c: Cat{ id: $id, name: $name, color: $color, gender: $gender })
		RETURN c
	`
	result, err := database.Session.Run(cypher, map[string]interface{}{
		"id": cat.ID,
		"name": cat.Name,
		"color": cat.Color,
		"gender": cat.Gender,
	})
	if err != nil {
		c <- err
		return
	}
	mutex.Unlock()
	if err = result.Err(); err != nil {
		c <- err
		return
	}

	c <- nil
}
