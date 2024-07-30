package db

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var mongoClient *mongo.Client

func Setup() {
	mongoUri := os.Getenv("MONGO_URI")

	ctx := context.Background()
	mdb, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI(mongoUri),
	)

	mongoClient = mdb
	if err = mongoClient.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	} else {
		fmt.Println("Connected to MongoDB!")
	}

}
