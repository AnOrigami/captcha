package util

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Newdb(ctx context.Context) *mongo.Database {
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://192.168.17.131:27017"))
	db := client.Database("captcha")
	return db
}
