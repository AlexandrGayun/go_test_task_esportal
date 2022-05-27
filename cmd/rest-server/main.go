package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/AlexandrGayun/go_test_task_esportal/config"
	mariadbconfig "github.com/AlexandrGayun/go_test_task_esportal/config/db"
	"github.com/AlexandrGayun/go_test_task_esportal/internal/mariadb/tweet"
	"github.com/AlexandrGayun/go_test_task_esportal/internal/mariadb/tweet/db"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"time"
)

type serverConfig struct {
	db     *sql.DB
	router *gin.Engine
}

func newServerConfig(db *sql.DB, router *gin.Engine) *serverConfig {
	return &serverConfig{db: db, router: router}
}

// "manager" struct could include all things we need for working with Tweet instance - logger/memcache/query handler etc
// but in this simple solution we will use only db repo
type tweetManager struct {
	repo *tweet.Tweet
}

func newTweetManager(repo *tweet.Tweet) *tweetManager {
	return &tweetManager{repo: repo}
}

type getTweetsOfFollowedRequestUri struct {
	ID int64 `uri:"user_id" binding:"required,min=1"`
}

type getTweetsOfFollowedRequestQuery struct {
	PageId   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=1,max=10"`
}

func errorResponse(c *gin.Context, err error) {
	c.IndentedJSON(http.StatusBadRequest, gin.H{
		"error_message": fmt.Sprintf("%s", err),
	})
}

func (mngr tweetManager) getTweetsOfFollowed(c *gin.Context) {
	var reqUri getTweetsOfFollowedRequestUri
	var reqQuery getTweetsOfFollowedRequestQuery
	if err := c.ShouldBindUri(&reqUri); err != nil {
		errorResponse(c, err)
		return
	}
	if err := c.ShouldBindQuery(&reqQuery); err != nil {
		errorResponse(c, err)
		return
	}
	args := db.GetTweetsOfFollowedParams{
		FollowerID: reqUri.ID,
		Limit:      reqQuery.PageSize,
		Offset:     (reqQuery.PageId - 1) * reqQuery.PageSize,
	}

	result, err := mngr.repo.GetTweetsOfFollowed(context.Background(), args)
	if err != nil {
		errorResponse(c, err)
	} else {
		c.IndentedJSON(http.StatusOK, result)
	}
}

func main() {
	config.LoadEnv()
	dbSettings := mariadbconfig.GetDSN("")
	db, err := sql.Open("mysql", dbSettings)
	if err != nil {
		log.Fatalln("can not connect to db ", err)
	}
	router := gin.Default()
	srv := newServerConfig(db, router)
	srv.setupDB().handleRequests().start()

}

func (srv *serverConfig) setupDB() *serverConfig {
	srv.db.SetConnMaxLifetime(time.Minute * 3)
	srv.db.SetMaxOpenConns(10)
	srv.db.SetMaxIdleConns(10)
	return srv
}

func (srv *serverConfig) handleRequests() *serverConfig {
	//setup Tweet repo for this route
	repo := tweet.NewTweet(srv.db)
	mngr := newTweetManager(repo)
	srv.router.GET("api/tweets/followed_by/:user_id", mngr.getTweetsOfFollowed)
	return srv
}

func (srv *serverConfig) start() {
	//gin.SetMode(gin.ReleaseMode)
	err := srv.router.Run(":8080")
	if err != nil {
		log.Fatalln("cannot run the server", err)
	}
}
