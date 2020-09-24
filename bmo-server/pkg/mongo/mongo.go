package dependency

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoConfig struct {
	Database string
	Host string
	Username string
	Password string
}

func NewMongo(c MongoConfig) *mongo.Database {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:27017", c.Username, c.Password, c.Host)

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		logrus.Errorf("Unable to connect database %v", err)
	}

	return client.Database(c.Database)
}