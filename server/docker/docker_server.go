package main

import (
	"context"
	"log"
	"os"
	"strings"
	"time"

	"github.com/baba2k/graphql-crud-server"
	"github.com/baba2k/graphql-crud-server/storage"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// load schema from file
	schema, err := graphql.LoadSchemaFromFile(os.Getenv("SCHEMA_FILE"))
	if err != nil {
		log.Fatal("can load schema: " + err.Error())
	}

	var db storage.Database

	switch strings.ToLower(os.Getenv("DATABASE_TYPE")) {
	case "mapdb":
		// create mapdb database interface
		db, err = storage.NewMapDB()
		if err != nil {
			log.Fatal("can not connect to database: " + err.Error())
		}
	case "mongodb":
		// create mongodb database interface
		opt := options.Client()
		mongodbHosts := os.Getenv("MONGODB_HOSTS")
		if len(mongodbHosts) > 0 {
			opt = opt.SetHosts(strings.Split(mongodbHosts, ","))
		} else {
			opt = opt.SetHosts([]string{"localhost"})
		}
		mongodbUsername := os.Getenv("MONGODB_USERNAME")
		if len(mongodbUsername) > 0 {
			opt = opt.SetAuth(options.Credential{Username: mongodbUsername})
		}
		mongodbPassword := os.Getenv("MONGODB_PASSWORD")
		if len(mongodbPassword) > 0 {
			if opt.Auth != nil {
				opt.Auth.Password = mongodbPassword
			} else {
				opt = opt.SetAuth(options.Credential{Password: mongodbPassword})
			}
		}
		mongodbDatabase := os.Getenv("MONGODB_DATABASE")
		if len(mongodbDatabase) == 0 {
			mongodbDatabase = "graphql"
		}
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		db, err = storage.NewMongoDB(ctx, opt, mongodbDatabase)
		if err != nil {
			log.Fatal("can not connect to database: " + err.Error())
		}
	}

	// start server
	graphql.StartServer(":8080", schema, db)
}
