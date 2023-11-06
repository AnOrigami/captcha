package service

import "go.mongodb.org/mongo-driver/mongo"

type Service struct {
	DB             *mongo.Database
	CollectionUser *mongo.Collection
}

func Newservice(db *mongo.Database) *Service {
	collectionuser := db.Collection("user")
	return &Service{
		DB:             db,
		CollectionUser: collectionuser,
	}
}
