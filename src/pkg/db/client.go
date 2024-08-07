package db

import (
	"context"
	"storage/src/pkg/logger"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var mongoClient *mongo.Client

func Setup() {
	log := logger.Logger
	mongoUri := viper.GetString("mongo.uri")
	log.Debug("Connecting to MongoDB", "uri", mongoUri)

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
		log.Info("Connected to MongoDB")
	}

}
