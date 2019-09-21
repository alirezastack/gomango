package main

import (
	"fmt"
	"strings"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	log "github.com/sirupsen/logrus"

	proto "tomato/proto"
	"tomato/store"

	"github.com/micro/go-micro/service/grpc"
	"go.mongodb.org/mongo-driver/mongo"
)

const serviceName = "TOMATO"

var MongoCollection *mongo.Collection

func Capitalize(s string) string {
	loweredVal := strings.ToLower(s)
	toTitle := strings.Title(loweredVal)
	return toTitle
}

// go run main.go mango.go OR -> go run .
func main() {
	log.SetLevel(log.DebugLevel)

	MongoCollection = store.MongoConnect(serviceName)

	// TODO add service field to all logs by default
	log.WithFields(log.Fields{
		"service": serviceName,
	}).Debug(fmt.Sprintf("%s service is starting up...", Capitalize(serviceName)))

	service := grpc.NewService(
		micro.Name(strings.ToLower(serviceName)),
		micro.Server(
			server.NewServer(
				server.Name(strings.ToLower(serviceName)),
				server.Address(":4500"),
			),
		),
	)

	service.Init()

	if err := proto.RegisterTomatoServiceHandler(service.Server(), new(TomatoService)); err != nil {
		panic(err)
	}

	if err := service.Run(); err != nil {
		panic(err)
	}
}
