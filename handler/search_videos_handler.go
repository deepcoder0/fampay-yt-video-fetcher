package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"fampay-yt-video-fetcher/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type SearchVideosHandlerInterface struct {}

func (s *SearchVideosHandlerInterface) GetVideosHandler(e echo.Context)error{
	topic := e.QueryParam("topic")
	page, _ := strconv.Atoi(e.QueryParam("page"))
    if page == 0 {
        page = 1 // Default page number
    }
    pageSize, _ := strconv.Atoi(e.QueryParam("pageSize"))
    if pageSize == 0 {
        pageSize = 10 // Default page size
    }
	msg := fmt.Sprintf("Hello, searching videos from MongoDB with queryParams: Topic: %s, Page: %d, PageSize: %d", topic, page, pageSize)
	log.Info(msg)
	videos, err := service.GetVideosPaginated(page, pageSize, topic)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err)
	}
	return e.JSON(http.StatusOK, videos)
}

func (s *SearchVideosHandlerInterface) SearchVideosQueryHandler(e echo.Context)error{
	query := e.Param("query")
	pageSize := 5 // keeping default values for this
	page := 1 // keeping default values for this
	videos, err :=service.SearchVideoQuery(query)
	x:=fmt.Sprintf("Hello, fetching videos from YoutubeAPI with search, Eheheheh: %s", query)
	// log.Info(x)
	if err != nil {
		log.Info(err)
		return e.JSON(http.StatusBadRequest, x) // c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}
	return e.JSON(http.StatusOK, videos)
}