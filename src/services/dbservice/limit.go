package dbservice

import (
	"context"
	"storage/src/constants"
	"storage/src/pkg/db"
	"storage/src/pkg/logger"
	"time"
)

type Limit struct {
	ContractAddress string    `json:"contractAddress,omitempty" bson:"contractAddress,omitempty"`
	ExtraStorage    int64     `json:"extraStorage,omitempty" bson:"extraStorage,omitempty"`
	StorageLimit    int64     `json:"storageLimit,omitempty" bson:"storageLimit,omitempty"`
	StorageUse      int64     `json:"storageUse,omitempty" bson:"storageUse,omitempty"`
	Unit            string    `json:"unit,omitempty" bson:"unit,omitempty"`
	Timestamp       time.Time `json:"timeStamp,omitempty" bson:"timeStamp,omitempty"`
}

func new(contract string) Limit {
	return Limit{
		ContractAddress: contract,
		StorageLimit:    200000000,
		Unit:            "bytes",
		Timestamp:       time.Now().UTC(),
	}
}

func UpdateLimit(ctx context.Context, contractAddress string, size int64) error {
	collection := db.GetCollection(constants.Limit)
	filter := map[string]interface{}{
		"contractAddress": contractAddress,
	}
	update := map[string]interface{}{
		"$set": map[string]interface{}{
			"timeStamp": time.Now().UTC(),
		},
		"$inc": map[string]interface{}{
			"storageUse": size,
		},
		"$setOnInsert": map[string]interface{}{
			"contractAddress": contractAddress,
			"storageLimit":    200000000,
			"unit":            "bytes",
		},
	}
	_, err := collection.FindOneAndUpdate(ctx, filter, update, true)
	return err
}

func GetLimit(ctx context.Context, contractAddress string) Limit {
	log := logger.GetContextLogger(ctx)
	collection := db.GetCollection(constants.Limit)
	filter := map[string]interface{}{
		"contractAddress": contractAddress,
	}
	var limit Limit
	err := collection.FindOne(ctx, filter, &limit)
	if err != nil {
		log.Info("limit not found")
		return new(contractAddress)
	}
	return limit
}
