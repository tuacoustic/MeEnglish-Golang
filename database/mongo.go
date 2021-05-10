package database

import (
	"context"
	"me-english/utils/console"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Mongo_Connect() (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://admin:gkitchenDev2021@10.0.1.4"))
	if err != nil {
		console.Error("Connect Mongo Database err: ", err)
		return nil, err
	}
	return client, nil
}
