package graphql

import (
	"io/ioutil"
	"log"
	"reflect"

	"github.com/baba2k/graphql-crud-server/server"
	"github.com/baba2k/graphql-crud-server/storage"
)

func LoadSchemaFromFile(filename string) (error, string) {
	schemaBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	schema := string(schemaBytes)
	return err, schema
}

func StartServer(addr, schema string, db interface{}) {
	switch dbType := db.(type) {
	case storage.MongoDB:
		server.StartServer(addr, schema, dbType)
	default:
		log.Fatal("unknown database type: " + reflect.TypeOf(db).Name())
	}
}
