package service

import (
	"context"
	"encoding/json"
	"fampay-yt-video-fetcher/constants"
	"fampay-yt-video-fetcher/database"
	"fampay-yt-video-fetcher/models"
	"fmt"

	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetVideosPaginated(page, pageSize int, topic string) ([]models.Video, error) {
	videos := []models.Video{}
	sort := bson.M{"publishingDate": -1}
	log.Info("HIiiiiiiiiiii2")
	res, err := database.DB.Database(constants.DATABASE_NAME).Collection(constants.COLLECTION_NAME).
		Find(context.TODO(), bson.D{}, options.Find().SetSort(sort).SetSkip(int64((page-1)*pageSize)).SetLimit(int64(pageSize)))
	if err != nil {
		log.Info(err.Error())
		return nil, err
	}
	log.Info(res)
	for res.Next(context.Background()) {
        var result models.Video
        if err := res.Decode(&result); err != nil {
            log.Fatal(err)
        }
        videos = append(videos, result)
    }

    // Check for cursor errors
    if err := res.Err(); err != nil {
        log.Fatal(err)
    }

    // Convert the results to JSON
    jsonResponse, err := json.Marshal(videos)
    if err != nil {
        log.Fatal(err)
    }

    // Print or return the JSON response
    fmt.Println(string(jsonResponse))

	return videos, nil
}

func SearchVideoQuery(searchQuery string, page, pageSize int64) ([]models.Video, error) {
	videos := []models.Video{}
	regex := bson.M{"$regex": searchQuery, "$options": "i"}
	// define the filter to search multiple fields using fuzzy logic
	filter := bson.M{
		"$or": []bson.M{
			{"title": regex},
			{"description": regex},
		},
	}

	sort := bson.M{"publishingDate": -1}

	res, err := database.DB.Database(constants.DATABASE_NAME).Collection(constants.COLLECTION_NAME).
		Find(context.TODO(), filter, options.Find().SetSort(sort).SetSkip((page-1)*pageSize).SetLimit(pageSize))
	if err != nil {
		return nil, err
	}

	for res.Next(context.Background()) {
        var result models.Video
        if err := res.Decode(&result); err != nil {
            log.Fatal(err)
        }
        videos = append(videos, result)
    }

    // Check for cursor errors
    if err := res.Err(); err != nil {
        log.Fatal(err)
    }

    // Convert the results to JSON
    jsonResponse, err := json.Marshal(videos)
    if err != nil {
        log.Fatal(err)
    }

    // Print or return the JSON response
    fmt.Println(string(jsonResponse))


	return videos, nil
}
