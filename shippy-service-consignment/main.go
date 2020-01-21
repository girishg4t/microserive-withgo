// shippy-service-consignment/main.go
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	pb "github.com/girishg4t/shippy-service-consignment/proto/consignment"
	vesselProto "github.com/girishg4t/shippy-service-vessel/proto/vessel"
	k8s "github.com/micro/examples/kubernetes/go/micro"
	"github.com/micro/go-micro"
)

const (
	port        = ":50051"
	defaultHost = "mongodb://localhost:27017"
)

func main() {
	// Set-up micro instance
	srv := k8s.NewService(
		micro.Name("shippy.service.consignment"),
	)

	srv.Init()

	uri := os.Getenv("DB_HOST")
	if uri == "" {
		uri = defaultHost
	}
	client, err := CreateClient(uri)
	if err != nil {
		log.Panic(err)
	}
	defer client.Disconnect(context.TODO())

	consignmentCollection := client.Database("shippy").Collection("consignments")

	repository := &MongoRepository{consignmentCollection}
	vesselClient := vesselProto.NewVesselServiceClient("shippy.service.vessel", srv.Client())
	h := &handler{repository, vesselClient}

	// Register handlers
	pb.RegisterShippingServiceHandler(srv.Server(), h)

	// Run the server
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
