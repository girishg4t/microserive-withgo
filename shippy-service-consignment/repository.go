// shippy-service-consignment/repository.go
package main

import (
	"context"

	pb "github.com/girishg4t/shippy-service-consignment/proto/consignment"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type repository interface {
	Create(ctx context.Context, consignment *pb.Consignment) error
	GetAll(ctx context.Context) ([]*pb.Consignment, error)
}

// MongoRepository implementation
type MongoRepository struct {
	collection *mongo.Collection
}

// Create -
func (repository *MongoRepository) Create(
	ctx context.Context,
	consignment *pb.Consignment,
) error {
	res, err := repository.collection.InsertOne(ctx, consignment)
	if res != nil {
		return nil
	}
	return err
}

// GetAll -
func (repository *MongoRepository) GetAll(ctx context.Context) ([]*pb.Consignment, error) {

	cur, err := repository.collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())
	var consignments []*pb.Consignment
	for cur.Next(context.Background()) {
		// To decode into a struct, use cursor.Decode()
		var result *pb.Consignment
		err := cur.Decode(&result)
		if err != nil {
			return nil, err
		}
		// do something with result...
		consignments = append(consignments, result)
		// To get the raw bson bytes use cursor.Current
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}

	return consignments, err
}
