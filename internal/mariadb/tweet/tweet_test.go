package tweet

import (
	"context"
	"database/sql"
	"github.com/AlexandrGayun/go_test_task_esportal/config"
	mariadbconfig "github.com/AlexandrGayun/go_test_task_esportal/config/db"
	"github.com/AlexandrGayun/go_test_task_esportal/internal/mariadb/tweet/db"
	"log"
	"testing"
)

var TweetRepo *Tweet

func TestMain(m *testing.M) {
	config.LoadEnv("./../../../.env")
	dbSettings := mariadbconfig.GetDSN("test_")
	db, err := sql.Open("mysql", dbSettings)
	if err != nil {
		log.Println("can not connect to db ", err)
	}
	TweetRepo = NewTweet(db)
	m.Run()
}

func tweetChecker(t *testing.T, actualResult []db.GetTweetsOfFollowedRow, expectedRowCount int) {
	t.Helper()
	actualRowCount := len(actualResult)
	if actualRowCount != expectedRowCount {
		t.Errorf("Get wrong result, expected %v, got %v", expectedRowCount, actualRowCount)
	}
}

func TestTweet_GetTweetsOfFollowed(t *testing.T) {
	testCases := []struct {
		params                  db.GetTweetsOfFollowedParams
		shouldRetrieveRowsCount int
	}{
		{db.GetTweetsOfFollowedParams{FollowerID: 1, Limit: 10, Offset: 0}, 6},
		{db.GetTweetsOfFollowedParams{FollowerID: 1, Limit: 5, Offset: 2}, 4},
		{db.GetTweetsOfFollowedParams{FollowerID: 1, Limit: 5, Offset: 2}, 4},
		{db.GetTweetsOfFollowedParams{FollowerID: 2, Limit: 5, Offset: 0}, 2},
		{db.GetTweetsOfFollowedParams{FollowerID: 6, Limit: 5, Offset: 0}, 0},
	}
	for _, tc := range testCases {
		t.Run("Retrieving tweets", func(t *testing.T) {
			res, err := TweetRepo.GetTweetsOfFollowed(context.Background(), tc.params)
			if err != nil {
				t.Errorf("Unexpected error %q", err)
			}
			tweetChecker(t, res, tc.shouldRetrieveRowsCount)
		})
	}
}
