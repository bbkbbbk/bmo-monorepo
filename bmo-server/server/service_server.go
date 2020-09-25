package server

import (
	proto "github.com/bbkbbbk/bmo-monorepo/bmo-server/pkg/proto/v1"
)

type newBMOServiceServer struct {
	service Service
}

func NewBMOServiceServer(s Service) proto.BMOServiceServer {
	return &newBMOServiceServer{
		service: s,
	}
}