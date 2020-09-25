package mongo

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	Database string
	Host string
	Username string
	Password string
}

func NewMongo(c Config) *mongo.Database {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:27017/?authSource=bmo", c.Username, c.Password, c.Host)

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		logrus.Errorf("Unable to connect database %v", err)
	}
	db := client.Database(c.Database)

	return db
}