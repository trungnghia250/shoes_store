package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	Client   *mongo.Client
	Customer *mongo.Collection
	Order    *mongo.Collection
	Counter  *mongo.Collection
}

var DB Database

func connectMongo(url string) (*mongo.Client, error) {
	opt := options.Client()
	opt.ApplyURI(url)

	client, err := mongo.Connect(context.Background(), opt)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func ConnectDB() error {
	client, err := connectMongo("mongodb+srv://shoes_store:shoes_store@shoesstore.8gkpr.mongodb.net")
	if err != nil {
		return err
	}

	shoesStoreDB := client.Database("shoes_store")
	DB = Database{
		Client:   client,
		Customer: shoesStoreDB.Collection("customer"),
		Order:    shoesStoreDB.Collection("order"),
		Counter:  shoesStoreDB.Collection("counter"),
	}

	return nil
}
