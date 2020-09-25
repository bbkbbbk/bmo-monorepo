package main

import (
	"net"
	"os"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"

	pkgMongo "github.com/bbkbbbk/bmo-monorepo/bmo-server/pkg/mongo"
	proto "github.com/bbkbbbk/bmo-monorepo/bmo-server/pkg/proto/v1"
	"github.com/bbkbbbk/bmo-monorepo/bmo-server/server"
)

var (
	db *mongo.Database
)

func init() {
	db = pkgMongo.NewMongo(pkgMongo.Config{
		Database: os.Getenv("MONGO_DATABASE"),
		Host: os.Getenv("MONGO_HOST"),
		Username: os.Getenv("MONGO_USERNAME"),
		Password: os.Getenv("MONGO_PASSWORD"),
	})
}

func main() {
	logrus.Info("[main]: initializing server")

	repo := server.NewRepository(db)
	service := server.NewService(repo)
	bmoServiceServer := server.NewBMOServiceServer(service)

	lis, err := net.Listen("tcp", os.Getenv("APP_PORT"))
	if err != nil {
		logrus.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterBMOServiceServer(s, bmoServiceServer)
	if err := s.Serve(lis); err != nil {
		logrus.Fatalf("failed to serve: %v", err)
	}
}