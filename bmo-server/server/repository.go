package server

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	collectionCardSet = "card-sets"
	collectionQuizzes = "quizzes"

	defaultTimeout = 30
)

type Repository interface {
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

func (r *repository) defaultContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), defaultTimeout*time.Second)
}

func (r *repository) InsertCardSet(cs CardSet) (string, error) {
	ctx, cancel := r.defaultContext()
	defer cancel()

	doc, err := bson.Marshal(cs)
	if err != nil {
		return "", errors.Wrap(err, "[InsertCardSet]: unable to marshal a card set")
	}

	result, err := r.db.Collection(collectionCardSet).InsertOne(ctx, doc)
	if err != nil {
		return "", errors.Wrap(err, "[InsertCardSet]: failed to insert a card set")
	}

	id := result.InsertedID.(string)

	return id, nil
}

func (r *repository) FetchCardSet(sc SearchCondition) (*CardSet, error) {
	ctx, cancel := r.defaultContext()
	defer cancel()

	var cardSet CardSet
	err := r.db.Collection(collectionCardSet).FindOne(ctx, sc.Filter).Decode(&cardSet)
	if err != nil {
		return nil, errors.Wrapf(err, "[FetchCardSet]: unable to retrieve card set with filter %v", sc.Filter)
	}

	return &cardSet, nil
}
