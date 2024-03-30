package main

import (
	"fampay-yt-video-fetcher/database"
	"fampay-yt-video-fetcher/routes"
	"fampay-yt-video-fetcher/util"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	config, err := util.LoadConfig(".")
    if err != nil {
        log.Fatal("cannot load config:", err)
    }

	database.ConnectDB(&config)
	// coll, err := client.Database("sample_mflix").Collection("movies").InsertOne(context.TODO(), book)
	// title := "Back to the Future"
	// fmt.Printf("Inserted document with _id: %v\n", coll.InsertedID)
	// if err != nil {
	// 	fmt.Print(err)
	// }

	commonRoutesGroup := app.Group("")
	fetchVideosRoutesGroup := app.Group("/fetch")
	searchVideosRoutesGroup := app.Group("/search")
	
	routes.RegisterCommonRoutes(commonRoutesGroup)
	routes.RegisterFetchVideosRoutes(fetchVideosRoutesGroup)
	routes.RegisterSearchVideosRoutes(searchVideosRoutesGroup)

	// app.Routes()
	
	// app.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello, World!")
	// })
	app.Logger.Fatal(app.Start(config.ServerAddress))
}