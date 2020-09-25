package server

import (
	"context"

	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/timestamppb"

	proto "github.com/bbkbbbk/bmo-monorepo/bmo-server/pkg/proto/v1"
)

func (h *newBMOServiceServer) InsertCardSet(ctx context.Context, req *proto.InsertCardSetRequest) (*proto.InsertCardSetResponse, error) {
	for _, c := range req.GetCardSet().Cards {
		c.Weight = 0
	}

	req.CardSet.LastoOpened = timestamppb.Now()

	id, err := h.service.InsertCardSet(ctx, req.GetCardSet())
	if err != nil {
		return nil, errors.Wrap(err, "[h.InsertCardSet]: unable to insert a card set")
	}

	res := proto.InsertCardSetResponse{
		Id: id,
	}

	return &res, nil
}
