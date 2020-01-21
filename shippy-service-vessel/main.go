package main

import (
	"context"
	"fmt"
	"log"
	"os"

	pb "github.com/girishg4t/shippy-service-vessel/proto/vessel"
	k8s "github.com/micro/examples/kubernetes/go/micro"
	"github.com/micro/go-micro"
)

const (
	defaultHost = "mongodb://localhost:27017"
)

func createDummyData(repo repository) {
	vessels := []*Vessel{
		{ID: "vessel001", Name: "Kane's Salty Secret", MaxWeight: 200000, Capacity: 500},
	}
	for _, v := range vessels {
		repo.Create(context.Background(), v)
	}
}

func main() {
	srv := k8s.NewService(
		micro.Name("shippy.service.vessel"),
	)

	srv.Init()

	uri := os.Getenv("DB_HOST")
	if uri == "" {
		uri = defaultHost
	}
	client, err := CreateClient(context.Background(), uri)
	if err != nil {
		log.Panic(err)
	}
	defer client.Disconnect(context.TODO())

	vesselCollection := client.Database("shippy").Collection("vessel")
	repository := &VesselRepository{
		vesselCollection,
	}

	createDummyData(repository)

	// Register our implementation with
	pb.RegisterVesselServiceHandler(srv.Server(), &handler{repository})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
