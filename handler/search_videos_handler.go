package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type SearchVideosHandlerInterface struct {}

func (s *SearchVideosHandlerInterface) SearchVideosHandler(e echo.Context)error{
	topic := e.QueryParam("topic")
	page := e.QueryParam("page")
	limit := e.QueryParam("limit")
	x:=fmt.Sprintf("Hello, searching videos from MongoDB with search: %s, %s, %s", topic, page, limit)
	return e.JSON(http.StatusOK, x)
}