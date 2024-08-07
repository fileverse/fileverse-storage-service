package db

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var mongoClient *mongo.Client

func Setup() {
	mongoUri := viper.GetString("mongo.uri")
	slog.Info("Connecting to MongoDB", "uri", mongoUri)

	ctx := context.Background()
	mdb, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI(mongoUri),
	)

	if err != nil {
		panic(err)
	}

	mongoClient = mdb
	if err = mongoClient.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	} else {
		fmt.Println("Connected to MongoDB!")
	}

}
