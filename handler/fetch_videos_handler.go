package handler

import (
	"fampay-yt-video-fetcher/service"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type FetchVideosHandlerInterface struct {}

func (s *FetchVideosHandlerInterface) FetchVideosHandler(e echo.Context)error{
	topic := e.Param("topic")
	service.FetchVideosFromYoutubeAPI(topic)
	x:=fmt.Sprintf("Hello, fetching videos from YoutubeAPI with search: %s", topic)
	return e.JSON(http.StatusOK, x)
}