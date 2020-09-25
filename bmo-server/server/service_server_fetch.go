package server

import (
	"context"

	"github.com/pkg/errors"

	proto "github.com/bbkbbbk/bmo-monorepo/bmo-server/pkg/proto/v1"
)

func (h *newBMOServiceServer) FetchCardSetByID(ctx context.Context, req *proto.FetchCardSetByIDRequest) (*proto.FetchCardSetByIDResponse, error) {
	id := req.GetId()

	cs, err := h.service.FetchCardSetByID(ctx, id)
	if err != nil {
		return nil, errors.Wrapf(err, "[h.FetchCardSetByID]: unable to find a card set id %v", id)
	}

	res := proto.FetchCardSetByIDResponse{
		CardSet: cs,
	}
	return &res, nil
}
