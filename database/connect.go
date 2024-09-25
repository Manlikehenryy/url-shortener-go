package database

import (
	"context"
	"log"

	"github.com/manlikehenryy/url-shortener-go/configs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var DB *mongo.Database

// Initialize the MongoDB client
func Connect() {
    mongoURI := configs.Env.MONGO_DB_URI
    if mongoURI == "" {
        log.Fatal("MONGO_DB_URI environment variable is not set")
    }

    clientOptions := options.Client().ApplyURI(mongoURI)
    var err error
    Client, err = mongo.Connect(context.Background(), clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    // Verify the connection
    err = Client.Ping(context.Background(), nil)
    if err != nil {
        log.Fatal(err)
    }

    DB = Client.Database("go_gin_task_manager")

    log.Println("Connected to MongoDB")
}
