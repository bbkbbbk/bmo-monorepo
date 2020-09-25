package server

import (
	"context"
	proto "github.com/bbkbbbk/bmo-monorepo/bmo-server/pkg/proto/v1"
	"github.com/pkg/errors"
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

func (h *newBMOServiceServer) FetchCardSets(ctx context.Context, _ *proto.EmptyRequest) (*proto.FetchCardSetsResponse, error) {
	cs, err := h.service.FetchCardSets(ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "[h.FetchCardSets]: unable to find card sets")
	}

	res := proto.FetchCardSetsResponse{
		CardSets: cs,
	}

	return &res, nil
}
