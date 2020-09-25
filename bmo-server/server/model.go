package server

import (
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CardSetFilter struct {
	ID string `bson:"_id"`
}

func (csf *CardSetFilter) ToBson() bson.M {
	filter := bson.M{}

	if csf.ID != "" {
		oid, err := primitive.ObjectIDFromHex(csf.ID)
		if err != nil {
			logrus.Warnf("[CardSetFilter.ToBson]: unable to get oid from string %v", err)
		}
		filter["_id "] = oid
	}

	return filter
}
