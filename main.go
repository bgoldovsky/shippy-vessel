package main

import (
	"fmt"
	"log"
	"os"

	pb "github.com/bgoldovsky/shippy-vessel/proto/vessel"
	"github.com/micro/go-micro"
)

const (
	defaultHost = "localhost:27017"
)

func main() {

	host := os.Getenv("DB_HOST")
	if host == "" {
		host = defaultHost
	}

	session, err := CreateSession(host)
	if err != nil {
		log.Fatalf("Could not connect to datastore with host %s - %v", host, err)
	}
	defer session.Close()

	srv := micro.NewService(
		micro.Name("shippy-vessel"),
		micro.Version("latest"),
	)

	srv.Init()

	//TODO: Complete DB init
	vessels := []*pb.Vessel{
		{Id: "vessel1", Name: "Bob", MaxWeight: 2000000, Capacity: 1000},
	}
	handler := &service{session}
	handler.GetRepo().Create(vessels[0])

	pb.RegisterVesselServiceHandler(srv.Server(), &service{session})
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
