package service

import (
	"encoding/json"
	"fampay-yt-video-fetcher/constants"
	"fampay-yt-video-fetcher/models"
	"fampay-yt-video-fetcher/util"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func FetchVideosFromYoutubeAPI(topic string) {
	config, err := util.LoadConfig(".")
    if err != nil {
        log.Fatal("cannot load config:", err)
    }
	apiKeys := config.APIKeysYT
	apiKeyIndex := 0

	go func() {
		for {
			apiKey := apiKeys[apiKeyIndex]
			currentTime := time.Now()
			publishedBefore := currentTime.Format(time.RFC3339)

			url := fmt.Sprintf("%s?key=%s&q=%s&type=video&order=date&part=snippet&publishedBefore=%s",constants.SEARCH, apiKey, topic, publishedBefore)

			videos, err := CallYouTubeAPI(url, topic)
			if err != nil {
				log.Printf("Error fetching videos: %v", err)
				// Switch to the next available API key if quota is exhausted
				apiKeyIndex = (apiKeyIndex + 1) % len(apiKeys)
				continue
			}

			SaveVideos(&videos)

			time.Sleep(10 * time.Second)
		}
	}()

}

func CallYouTubeAPI(url, topic string) ([]models.Video, error) {
	fmt.Println("The Video search endpoint : ", url)

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("YouTube API request failed with status code: %d", response.StatusCode)
	}
	// fmt.Print(response)
	var apiResponse models.YouTubeVideoResponse
	err = json.NewDecoder(response.Body).Decode(&apiResponse)
	if err != nil {
		return nil, err
	}

	var videos []models.Video
	for _, item := range apiResponse.Items {
		fmt.Println(item.Snippet)
		publishTime, err := time.Parse(time.RFC3339, item.Snippet.PublishTime)
		if err != nil {
			return nil, err
		}
 
		video := models.Video{
			VideoId: uuid.NewString(),
			VideoTopic: topic,
			VideoTitle: item.Snippet.Title,
			Description: item.Snippet.Description,
			PublishingDate:  publishTime,
			ThumbnailsUrl: item.Snippet.Thumbnails.Default.URL,
		}
		videos = append(videos, video)
	}

	return videos, nil
}