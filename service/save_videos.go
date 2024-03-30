package service

import (
	"context"
	"fampay-yt-video-fetcher/database"
	"fampay-yt-video-fetcher/models"
	"fmt"
)

func SaveVideos(videos *[]models.Video){

	for _, video := range *videos{
		SaveSingleVideo(video)
	}

}

func SaveSingleVideo(video models.Video){
	_, err := database.DB.Database("fampay").Collection("videos").InsertOne(context.TODO(), video)
	if err != nil {
		fmt.Print(err)
	}
}