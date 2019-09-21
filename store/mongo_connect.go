package store

import (
	"context"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strings"
)

func MongoConnect(serviceName string) *mongo.Collection {
	// create a new context
	mctx := context.Background()

	// create a mongo client
	client, err := mongo.NewClient(
		options.Client().ApplyURI("mongodb://admin:admin@localhost:27017/"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// connect to mongo
	if err := client.Connect(mctx); err != nil {
		log.Fatal(err)
	}

	// disconnects from mongo
	//defer client.Disconnect(mctx)

	db := client.Database(strings.ToLower(serviceName))
	col := db.Collection("question")
	return col
}
