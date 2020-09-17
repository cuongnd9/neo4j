# neo4j

Learning 📺 Neo4j

## Docker

```shell
docker run -p 7474:7474 -p 7687:7687 \
    --env NEO4J_AUTH=neo4j/123456 \
    --volume=$HOME/neo4j/data:/data \
    neo4j
```

**Connect to Neo4j**

![login](./.images/login.png)

**Result**

![login_result](./.images/login_result.png)

## Go client

```shell
go run go_client/main.go
```

## License

MIT © [Cuong Tran](https://github.com/103cuong)
