package main

import (
	"github.com/103cuong/neo4j/go_client_crud/pkg/database"
	"github.com/103cuong/neo4j/go_client_crud/pkg/model"
	"github.com/103cuong/neo4j/go_client_crud/pkg/service"
	"github.com/103cuong/uid"
	"log"
	"sync"
	"syreclabs.com/go/faker"
)

func main()  {
	// connect to database.
	driver, err := database.ConnectToDB()
	if err != nil {
		log.Fatalln("connecting to Neo4j failed")
		panic("connecting to Neo4j failed")
	}
	// close driver & session after main function ends.
	defer driver.Close()
	defer database.Session.Close()

	c := make(chan error)
	mutex := &sync.Mutex{}

	// create new cat
	newCat := model.Cat{
		ID: uid.Uid(24),
		Name:   faker.Name().FirstName(),
		Color:  "white",
		Gender: "female",
	}
	go service.CreateCat(newCat, c, mutex)
	err = <- c

	if err != nil {
		log.Fatalf("cypher CRUD failed: %v", err)
		panic("cypher CRUD failed")
	}
}
