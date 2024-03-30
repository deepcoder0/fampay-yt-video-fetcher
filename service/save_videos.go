package service

import (
	"fampay-yt-video-fetcher/database"
	"fampay-yt-video-fetcher/models"
	"fmt"
)

func SaveVideos(videos *[]models.Video){

	for _, video := range *videos{
		videoTitle := video.VideoTitle
		_, err := database.GetVideoDetailsByTitle(videoTitle)
		if err == nil {
			fmt.Println("Video already Present in DB, VideoTitle := ", videoTitle)
			continue
		}
		err = database.SaveVideoInDB(video)
		if err != nil {
			fmt.Println("Unable to save the video in DB, VideoTitle := ", videoTitle, err)
		}
		fmt.Println("Video inserted into mongoDB where _id := ", video.VideoId)
	}
}