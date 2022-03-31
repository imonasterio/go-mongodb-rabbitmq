package server

import (
	"github.com/imonasterio/go-mongodb-rabbitmq/cmd/api/handlers"
	"github.com/labstack/echo/v4"
)

func InitServer() {

	e := echo.New()

	e.GET("/tweets", handlers.GetTweetsEndpoint)
	e.GET("/search", handlers.SearchTweetsEndpoint)
	e.POST("/tweet", handlers.InsertTwitterEndopoint)

	e.Logger.Fatal(e.Start(":1323"))

}
