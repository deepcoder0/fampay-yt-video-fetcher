package database

import (
	"context"
	"fampay-yt-video-fetcher/util"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client

func ConnectDB(config *util.Config) {
	var err error
	mongoURI := config.MongoURI
	clientOptions := options.Client().ApplyURI(mongoURI)

	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	DB, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Println(err)
	}
	log.Println("🚀 Connected Successfully to the Database")
}

	
