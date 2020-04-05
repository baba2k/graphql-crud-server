package graphql

import (
	"io/ioutil"

	"github.com/baba2k/graphql-crud-server/server"
	"github.com/baba2k/graphql-crud-server/storage"
)

func LoadSchemaFromFile(filename string) (string, error) {
	schemaBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	schema := string(schemaBytes)
	return schema, err
}

func StartServer(addr, schema string, db storage.Database) {
	server.StartServer(addr, schema, db)
}
