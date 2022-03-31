package handlers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/imonasterio/go-mongodb-rabbitmq/cmd/api/models"
	"github.com/imonasterio/go-mongodb-rabbitmq/cmd/api/mongodb"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetTweetsEndpoint(c echo.Context) error {
	var tweets []models.Tweet
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cursor, err := mongodb.Collection.Find(ctx, bson.M{})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Error: %s", err.Error()))
	}

	if err = cursor.All(ctx, &tweets); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error: %s", err.Error()))
	}

	return c.JSON(http.StatusOK, tweets)
}

func SearchTweetsEndpoint(c echo.Context) error {
	var tweets []models.Tweet
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	searchStage := bson.D{
		{"$search", bson.D{
			{"index", "synsearch"},
			{"text", bson.D{
				{"query", c.QueryParam("q")},
				{"path", "full_text"},
				{"synonyms", "slang"},
			}},
		}},
	}
	cursor, err := mongodb.Collection.Aggregate(ctx, mongo.Pipeline{searchStage})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error: %s", err.Error()))
	}
	if err = cursor.All(ctx, &tweets); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error: %s", err.Error()))

	}

	return c.JSON(http.StatusOK, tweets)
}

func InsertTwitterEndopoint(c echo.Context) error {
	var tweet models.Tweet
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	if err := c.Bind(&tweet); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Error: %s", err.Error()))
	}

	_, err := mongodb.Collection.InsertOne(ctx, tweet)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error: %s", err.Error()))

	}

	return c.JSON(http.StatusOK, "Succesfully")
}
