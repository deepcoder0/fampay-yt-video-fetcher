package database

import (
	"context"
	"errors"
	"fampay-yt-video-fetcher/constants"
	"fampay-yt-video-fetcher/models"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetVideoDetailsByTitle(videoTitle string) (models.Video, error) {
	video := models.Video{}
	filter := bson.M{"videoTitle": videoTitle}
	err := DB.Database(constants.DATABASE_NAME).Collection(constants.COLLECTION_NAME).FindOne(context.TODO(), filter).Decode(&video)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			msg := fmt.Sprintf("Document not found in mongoDB! VideoTitle: %s", videoTitle)
			return models.Video{}, errors.New(msg)
		}
		return models.Video{}, err
	}
	return video, nil

}

func SaveVideoInDB(video models.Video) error {
	_, err := DB.Database(constants.DATABASE_NAME).Collection(constants.COLLECTION_NAME).InsertOne(context.TODO(), video)
	if err != nil {
		return err
	}
	return nil
}

