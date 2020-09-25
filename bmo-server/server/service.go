package server

import (
	"context"

	"github.com/pkg/errors"

	proto "github.com/bbkbbbk/bmo-monorepo/bmo-server/pkg/proto/v1"
)

type Service interface {
	InsertCardSet(ctx context.Context, cs *proto.CardSet) (string, error)
	FetchCardSetByID(ctx context.Context, id string) (*proto.CardSet, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{
		repo: r,
	}
}

func (s *service) InsertCardSet(ctx context.Context, cs *proto.CardSet) (string, error) {
	id, err := s.repo.InsertCardSet(ctx, cs)
	if err != nil {
		return "", errors.Wrap(err, "[s.InsertCardSet]: unable to insert a card set")
	}

	return id, nil
}

func (s *service) FetchCardSetByID(ctx context.Context, id string) (*proto.CardSet, error) {
	cs, err := s.repo.FetchCardSetByID(ctx, id)
	if err != nil {
		return nil, errors.Wrapf(err, "[s.FetchCardSetByID]: unable to find a card set id %v", id)
	}

	return cs, nil
}




