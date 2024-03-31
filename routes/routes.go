package routes

import (
	"context"
	"net/http"

	"fampay-yt-video-fetcher/database"
	"fampay-yt-video-fetcher/handler"

	"github.com/labstack/echo/v4"
)

func RegisterCommonRoutes(commonRoutesGroup *echo.Group) {
	commonRoutesGroup.GET("/healthcheck", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome To Youtube Fetcher API")
	})

	commonRoutesGroup.GET("/dbReadiness", func(c echo.Context) error {
		err := database.DB.Ping(context.Background(), nil)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "Unable to connect to DB!")
		}
		return c.JSON(http.StatusOK, "Database Ready to Use")
	})

}

func RegisterFetchVideosRoutes(fetchVideosRoutesGroup *echo.Group){
	fetchVideosHandler := handler.FetchVideosHandlerInterface{}
	fetchVideosRoutesGroup.GET("/:topic", fetchVideosHandler.FetchVideosHandler)
}

func RegisterSearchVideosRoutes(searchVideosRoutesGroup *echo.Group){
	searchVideosHandler := handler.SearchVideosHandlerInterface{}
	searchVideosRoutesGroup.GET("", searchVideosHandler.GetVideosHandler)
	searchVideosRoutesGroup.GET("/:query", searchVideosHandler.SearchVideosQueryHandler)
}