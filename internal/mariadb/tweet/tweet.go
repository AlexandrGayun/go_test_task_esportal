package tweet

import (
	"context"
	"github.com/AlexandrGayun/go_test_task_esportal/internal/mariadb/tweet/db"
)

//Tweet represents the repository used for interacting with Task records
type Tweet struct {
	q *db.Queries
}

//NewTweet instantiates new Tweet repository
func NewTweet(d db.DBTX) *Tweet {
	return &Tweet{
		q: db.New(d),
	}
}

func (t *Tweet) GetTweetsOfFollowed(ctx context.Context, params db.GetTweetsOfFollowedParams) ([]db.GetTweetsOfFollowedRow, error) {
	res, err := t.q.GetTweetsOfFollowed(ctx, params)
	if err != nil {
		return nil, err
	}
	return res, nil
	//return groupByUsers(res), nil
}

//Group tweets by users, excluding duplicate information about tweet author username/firstname/lastname etc

//type tweetsOfFollowedAPIResponse struct {
//	Username  string             `json:"username"`
//	FirstName string             `json:"first_name"`
//	LastName  string             `json:"last_name"`
//	Tweets    []tweetAPIResponse `json:"tweets"`
//}
//
//type tweetAPIResponse struct {
//	TweetText       string    `json:"tweet_text"`
//	PublicationDate time.Time `json:"publication_date"`
//}
//
//func groupByUsers(data []db.GetTweetsOfFollowedRow) []*tweetsOfFollowedAPIResponse {
//	var res []*tweetsOfFollowedAPIResponse
//	checkedIndexes := make(map[int]uint8)
//	for i := 0; i < len(data); i++ {
//		if _, ok := checkedIndexes[i]; ok {
//			continue
//		}
//		user := &tweetsOfFollowedAPIResponse{Username: data[i].Username,
//			FirstName: data[i].FirstName,
//			LastName:  data[i].LastName,
//			Tweets:    make([]tweetAPIResponse, 0),
//		}
//		res = append(res, user)
//		for j := 0; j < len(data); j++ {
//			_, checked := checkedIndexes[j]
//			if !checked && data[i].ID == data[j].ID {
//				tweet := tweetAPIResponse{TweetText: data[j].TweetText, PublicationDate: data[j].PublicationDate}
//				user.Tweets = append(user.Tweets, tweet)
//				checkedIndexes[j] = 1
//			}
//		}
//	}
//	return res
//}
