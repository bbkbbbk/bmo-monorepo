package server

import (
	"context"

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
