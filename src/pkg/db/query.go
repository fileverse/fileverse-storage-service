package db

import (
	"context"
	"storage/src/constants"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Collection struct {
	Db         string
	Collection string
	collection *mongo.Collection
}

func (c *Collection) init() {
	c.collection = mongoClient.Database(c.Db).Collection(c.Collection)
}

func (c *Collection) FindOne(ctx context.Context, query interface{}, resp interface{}) error {

	err := c.collection.FindOne(ctx, query).Decode(resp)
	return err
}

func (c *Collection) FindOneAndUpdate(ctx context.Context, filter interface{}, update interface{}, upsert bool) (interface{}, error) {
	options := &options.FindOneAndUpdateOptions{
		Upsert: &upsert,
	}

	var resp interface{}
	err := c.collection.FindOneAndUpdate(ctx, filter, update, options).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Collection) Find(ctx context.Context, query interface{}) (*mongo.Cursor, error) {
	resp, err := c.collection.Find(ctx, query)
	return resp, err
}

func (c *Collection) InsertOne(ctx context.Context, document interface{}) (*mongo.InsertOneResult, error) {
	resp, err := c.collection.InsertOne(ctx, document)
	return resp, err
}

func GetCollection(collection constants.CollectionName) *Collection {
	db := viper.GetString("mongo.db")
	c := &Collection{
		Db:         db,
		Collection: string(collection),
	}
	c.init()
	return c
}
