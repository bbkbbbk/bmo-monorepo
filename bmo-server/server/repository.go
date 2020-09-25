package server

import (
	"context"
	"github.com/sirupsen/logrus"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	proto "github.com/bbkbbbk/bmo-monorepo/bmo-server/pkg/proto/v1"
)

const (
	collectionCardSet = "card-sets"
	collectionQuizzes = "quizzes"
)

type Repository interface {
	InsertCardSet(ctx context.Context, cs *proto.CardSet) (string, error)
	FetchCardSets(ctx context.Context) ([]*proto.CardSet, error)
	FetchCardSetByID(ctx context.Context, id string) (*proto.CardSet, error)
}

type repository struct {
	db *mongo.Database
}

type SearchCondition struct {
	Filter bson.D
	Sort bson.D
}

func NewRepository(db *mongo.Database) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) InsertCardSet(ctx context.Context, cs *proto.CardSet) (string, error) {
	doc, err := bson.Marshal(&cs)
	if err != nil {
		return "", errors.Wrap(err, "[r.InsertCardSet]: unable to marshal a card set")
	}

	result, err := r.db.Collection(collectionCardSet).InsertOne(ctx, doc)
	if err != nil {
		return "", errors.Wrap(err, "[r.InsertCardSet]: failed to insert a card set")
	}

	oid, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", errors.New("[r.InsertCardSet]: unable to parsed object id")
	}

	return oid.String(), nil
}

func (r *repository) FetchCardSetByID(ctx context.Context, id string) (*proto.CardSet, error) {
	cf := CardSetFilter{
		ID: id,
	}
	filter := cf.ToBson()

	var cardSet proto.CardSet
	err := r.db.Collection(collectionCardSet).FindOne(ctx, filter).Decode(&cardSet)
	if err != nil {
		return nil, errors.Wrapf(err, "[r.FetchCardSetByID]: unable to retrieve card set with filter %v", filter)
	}

	return &cardSet, nil
}

func (r *repository) FetchCardSets(ctx context.Context) ([]*proto.CardSet, error) {
	cur, err := r.db.Collection(collectionCardSet).Find(ctx, bson.M{})
	if err != nil {
		return nil, errors.Wrapf(err, "[FetchCardSets]: unable to find card sets")
	}
	defer func(ctx context.Context, cur *mongo.Cursor) {
		if err = cur.Close(ctx); err != nil {
			logrus.Warnf("[r.FetchCardSets]: unable to close cursor %v", err)
		}
	}(ctx, cur)

	cardSets := []*proto.CardSet{}
	for cur.Next(ctx) {
		var cs proto.CardSet
		if err := cur.Decode(&cs); err != nil {
			return nil, errors.Wrap(err, "[r.FetchCardSets]: unable to decode card set")
		}

		cardSets = append(cardSets, &cs)
	}

	if err := cur.Err(); err != nil {
		return nil, errors.Wrapf(err, "[r.FetchCardSets]: unable to get current cursor")
	}

	return cardSets, nil
}