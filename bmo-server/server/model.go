package server

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CardSet struct {
	ID primitive.ObjectID
	UserID string
	Title string
	NumCards int32
	Cards []Card
	LastOpen *time.Time
}

type Card struct {
	Term       string
	Definition string
	Weight     int32
}
