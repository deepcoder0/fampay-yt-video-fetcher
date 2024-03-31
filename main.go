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

	err = database.ConnectDB(&config)
	if err != nil {
		log.Fatal("OOPS! Unable to Connect with DB")
		return 
	}

	commonRoutesGroup := app.Group("")
	fetchVideosRoutesGroup := app.Group("/fetch")
	searchVideosRoutesGroup := app.Group("/search")
	
	routes.RegisterCommonRoutes(commonRoutesGroup)
	routes.RegisterFetchVideosRoutes(fetchVideosRoutesGroup)
	routes.RegisterSearchVideosRoutes(searchVideosRoutesGroup)

	app.Logger.Fatal(app.Start(config.ServerAddress))
}