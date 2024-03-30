package main

import (
	"context"
	"fampay-yt-video-fetcher/util"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Book struct{
	ID int `json:"id"`
}

func main() {
	app := echo.New()

	config, err := util.LoadConfig(".")
    if err != nil {
        log.Fatal("cannot load config:", err)
    }

	mongoURI := "mongodb://mongodb:27017"
	clientOptions := options.Client().ApplyURI(mongoURI)

	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Println(err)
	}

	book := Book{
		ID: 1,
	}

	coll, err := client.Database("sample_mflix").Collection("movies").InsertOne(context.TODO(), book)
	// title := "Back to the Future"
	fmt.Printf("Inserted document with _id: %v\n", coll.InsertedID)
	if err != nil {
		fmt.Print(err)
	}
	
	
	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	app.Logger.Fatal(app.Start(config.ServerAddress))
}