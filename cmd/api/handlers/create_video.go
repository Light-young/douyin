package handlers

import (
	"context"
	"douyin/v1/cmd/api/rpc"
	"douyin/v1/kitex_gen/video"
	"douyin/v1/pkg/constants"
	"douyin/v1/pkg/errno"
	"fmt"
	"path/filepath"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func PublishVideo(c *gin.Context) {
	// todo 根据token解析出userId
	claims := jwt.ExtractClaims(c)
	userID := int64(claims[constants.IdentityKey].(float64))

	// token := c.Query("token")
	titleStr := c.Query("title")
	if titleStr == "" {
		SendCreateVideoResponse(c, errno.ParamErr)
		return
	}
	data, err := c.FormFile("data")
	if err != nil {
		SendCreateVideoResponse(c, err)
		return
	}
	filename := filepath.Base(data.Filename)
	finalName := fmt.Sprintf(filename)
	// 存储视频文件
	saveFile := filepath.Join("./cmd/api/static/videos/", finalName)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		SendCreateVideoResponse(c, err)
		return
	}
	newVideo := video.Video{
		AuthorId:      userID,
		PlayUrl:       constants.PlayURL + filename,
		CoverUrl:      constants.CoverURL,
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
		Title:         titleStr,
	}
	if err := rpc.CreateVideo(context.Background(), &newVideo); err != nil {
		SendCreateVideoResponse(c, err)
	}
	SendCreateVideoResponse(c, errno.Success)
}