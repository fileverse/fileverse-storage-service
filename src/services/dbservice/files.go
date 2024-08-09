package dbservice

import (
	"context"
	"storage/src/constants"
	"storage/src/pkg/db"
	"time"
)

type Files struct {
	InvokerAddress  string    `json:"invokerAddress,omitempty" bson:"invokerAddress,omitempty"`
	ContractAddress string    `json:"contractAddress,omitempty" bson:"contractAddress,omitempty"`
	GatewayURL      string    `json:"gatewayUrl,omitempty" bson:"gatewayUrl,omitempty"`
	ChainID         string    `json:"chainId,omitempty" bson:"chainId,omitempty"`
	IpfsHash        string    `json:"ipfsHash,omitempty" bson:"ipfsHash,omitempty"`
	FileSize        int64     `json:"fileSize,omitempty" bson:"fileSize,omitempty"`
	Tags            []string  `json:"tags,omitempty" bson:"tags,omitempty"`
	Namespace       any       `json:"namespace,omitempty" bson:"namespace,omitempty"`
	Timestamp       time.Time `json:"timeStamp,omitempty" bson:"timeStamp,omitempty"`
}

func (f *Files) Create(ctx context.Context) error {
	collection := db.GetCollection(constants.Files)
	f.Timestamp = time.Now().UTC()
	_, err := collection.InsertOne(ctx, f)
	return err
}
