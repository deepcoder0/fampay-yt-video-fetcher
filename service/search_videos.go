package service

import (
	"context"
	"fampay-yt-video-fetcher/constants"
	"fampay-yt-video-fetcher/database"
	"fampay-yt-video-fetcher/models"
	"fmt"
	"regexp"
	"strings"

	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetVideosPaginated(page, pageSize int, topic string) ([]models.Video, error) {
	videos := []models.Video{}
	sort := bson.M{"publishingDate": -1}
	filter := bson.M{}
	if topic != "" {
		filter = bson.M{"videoTopic": topic}
	}
	res, err := database.DB.Database(constants.DATABASE_NAME).Collection(constants.COLLECTION_NAME).
		Find(context.TODO(), filter, options.Find().SetSort(sort).SetSkip(int64((page-1)*pageSize)).SetLimit(int64(pageSize)))
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

	return videos, nil
}

func SearchVideoQuery(searchQuery string, page, pageSize int64) ([]models.Video, error) {
	videos := []models.Video{}

	// Split the string based on spaces
    keywords := strings.Split(searchQuery, " ")
    // Construct regex pattern to match all keywords simultaneously
    regexPattern := constructRegexPattern(keywords)
	log.Info(regexPattern)

	regexOperator := bson.M{"$regex": regexPattern, "$options": "i"}

	// the keywords can match in any of the title, description
	filter := bson.M{
		"$or": []bson.M{
			{"title": regexOperator},
			{"description": regexOperator},
		},
	}
	log.Info(filter)

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
	return videos, nil
}

func constructRegexPattern(keywords []string) string {
	// regexExpressionType = "(?=.*keyword1)(?=.*keyword2)(?=.*keyword3).+"
    escapedKeywords := make([]string, len(keywords))
    for i, keyword := range keywords {
        escapedKeywords[i] = fmt.Sprintf("(?=.*%s)", regexp.QuoteMeta(keyword))
    }
    regexPattern := strings.Join(escapedKeywords, "")
    regexPattern = regexPattern + ".+"
    return regexPattern
}