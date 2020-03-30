package server

import (
	"log"
	"net/http"

	s "github.com/baba2k/graphql-rungen/schema"
	"github.com/baba2k/graphql-rungen/storage"
	"github.com/graphql-go/handler"
)

func StartServer(addr, schema string, db storage.MongoDB) {
	parsedSchema, err := s.CreateSchema(schema, db)

	h := handler.New(&handler.Config{
		Schema:     &parsedSchema,
		Pretty:     true,
		GraphiQL:   false,
		Playground: true,
	})

	pattern := "/graphql"
	http.Handle(pattern, h)

	log.Println("start graphql server on " + addr + pattern)
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal("graphql server can not be started: " + err.Error())
	}
}
