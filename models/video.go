package models

import "time"

type Thumbnail struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type Thumbnails struct {
	Default  Thumbnail `json:"default"`
	Medium   Thumbnail `json:"medium"`
	High     Thumbnail `json:"high"`
	Standard Thumbnail `json:"standard"`
	Maxres   Thumbnail `json:"maxres"`
}

type Video struct {
	VideoId        string     `bson:"_id,omitempty" json:"id"`
	VideoTopic     string     `bson:"videoTopic" json:"videoTopic"`
	VideoTitle     string     `bson:"videoTitle" json:"videoTitle"`
	Description    string     `bson:"description" json:"description"`
	PublishingDate time.Time  `bson:"publishingDate" json:"publishingDate"`
	ThumbnailsUrl  string `bson:"thumbnails" json:"thumbnails"`
}