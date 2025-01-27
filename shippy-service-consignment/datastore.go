// shippy-service-consignment/datastore.go
package main

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CreateClient -
func CreateClient(uri string) (*mongo.Client, error) {
	ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
	return mongo.Connect(ctx, options.Client().ApplyURI(uri))
}
