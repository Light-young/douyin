package handlers

import (
	"context"
	"douyin/v1/cmd/api/vo"
	"fmt"
	"strconv"

	"douyin/v1/cmd/api/rpc"
	"douyin/v1/kitex_gen/user"
	"douyin/v1/pkg/constants"
	"douyin/v1/pkg/errno"
	"github.com/gin-gonic/gin"
)

// get follow list
func GetFollowList(c *gin.Context) {
	tokenId := vo.GetUserIdFromToken(c)
	fmt.Println("tokenId:", tokenId)
	toUserIDString := c.Query("user_id")
	toUserID, err := strconv.ParseInt(toUserIDString, 10, 64)
	req := &user.MGetUserRequest{UserId: tokenId, ToUserId: toUserID, ActionType: constants.QueryFollowList}
	user, err := rpc.MGetUser(context.Background(), req)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	// 这样的话返回的格式为:
	// "data": {
	//     "user_list": [
	//         {
	//             "note_id": 1,
	//             "user_id": 1,
	//             "user_name": "kinggo",
	//             "user_avatar": "test",
	//             "title": "test title",
	//             "content": "test content",
	//             "create_time": 1642525063
	//         }
	//     ],
	//     "total": 1
	// }
	// SendResponse(c, errno.Success, map[string]interface{}{"user_list": user})
	SendUserListResponse(c, errno.Success, user)
}

// get follower list
func GetFollowerList(c *gin.Context) {
	tokenId := vo.GetUserIdFromToken(c)
	fmt.Println("tokenId:", tokenId)
	toUserIDString := c.Query("user_id")
	toUserID, err := strconv.ParseInt(toUserIDString, 10, 64)

	req := &user.MGetUserRequest{UserId: tokenId, ToUserId: toUserID, ActionType: constants.QueryFollowerList}
	user, err := rpc.MGetUser(context.Background(), req)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendUserListResponse(c, errno.Success, user)
}
