package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	Mongo *mongo.Client
}

var Databse = DB{}

func Init() DB {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017/restful"))

	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(ctx)

	Databse.Mongo = client

	return Databse

}
