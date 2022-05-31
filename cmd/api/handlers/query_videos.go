package handlers

import (
	"context"
	"douyin/v1/cmd/api/rpc"
	"douyin/v1/pkg/errno"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func GetMyPublishVideoList(c *gin.Context) {
	token := c.Query("token")
	userIdStr := c.Query("user_id")
	if token == "" || userIdStr == "" {
		SendQueryByUserIdResponse(c, errno.ParamErr, nil)
		return
	}
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	videos, err := rpc.GetPublishVideoList(context.Background(), userId)
	if err != nil {
		SendQueryByUserIdResponse(c, errno.ConvertErr(err), nil)
		return
	}
	// todo fill videos with author
	SendQueryByUserIdResponse(c, errno.Success, videos)
}

func GetVideoFeed(c *gin.Context) {
	//token := c.Query("token")
	lastTimeStr := c.Query("latest_time")
	lastTime := time.Now().Unix()
	if lastTimeStr != "" {
		t, err := strconv.ParseInt(lastTimeStr, 10, 64)
		if err != nil {
			SendQueryByLastTimeResponse(c, errno.ConvertErr(err), nil, time.Now().Unix())
			return
		}
		if t != 0 {
			lastTime = t
		}
	}
	videos, nextTime, err := rpc.GetVideosFeed(context.Background(), lastTime)
	if err != nil {
		SendQueryByLastTimeResponse(c, err, nil, nextTime)
		return
	}
	// todo fill videos with author
	SendQueryByLastTimeResponse(c, errno.Success, videos, nextTime)
}
