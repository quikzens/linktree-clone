package db

import (
	"context"
	"log"

	"linktree-clone/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Mongo Collection
var (
	UserColl = DB.Collection("user")
)

// DB Instance
var (
	DB = connectDB()
)

// connectDB open mongodb connection, check the opened connection, and return the db client
func connectDB() *mongo.Database {
	uri := config.DBSource
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("successfully connecting to database")
	return client.Database(config.DBName)
}
