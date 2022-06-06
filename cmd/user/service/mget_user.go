package service

import (
	"context"
	"douyin/v1/cmd/user/dal/db"
	"douyin/v1/cmd/user/pack"
	"douyin/v1/kitex_gen/user"
)

type MGetUserService struct {
	ctx context.Context
}

// NewMGetUserService new MGetUserService
func NewMGetUserService(ctx context.Context) *MGetUserService {
	return &MGetUserService{ctx: ctx}
}

// MGetUser multiple get list of user info
func (s *MGetUserService) MGetUser(req *user.MGetUserRequest) ([]*user.User, error) {
	Users, err := db.MGetUsers(s.ctx, req)
	if err != nil {
		return nil, err
	}
	isFollowList, err1 := db.QueryFollowRelation(s.ctx, Users, req.UserId)
	if err1 != nil {
		return nil, err1
	}
	return pack.Users(Users, isFollowList), nil
	//if req.ActionType == constants.QueryFollowList {
	//	// 查询到的用户列表中的is_follow字段一定为true
	//	isFollowList, err1 := db.QueryFollowRelation(s.ctx, Users, req.UserId)
	//	if err1 != nil {
	//		return nil, err1
	//	}
	//	//for i := range isFollowList {
	//	//	fmt.Println("isFollowLIst[i]: ", isFollowList[i])
	//	//}
	//	return pack.Users(Users, isFollowList), nil
	//} else if req.ActionType == constants.QueryFollowerList {
	//	// 还得确认一下是否关注了该用户
	//	isFollowList, err1 := db.QueryFollowRelation(s.ctx, Users, req.UserId)
	//	if err1 != nil {
	//		return nil, err1
	//	}
	//	//for i := range isFollowList {
	//	//	fmt.Println("isFollowLIst[i]: ", isFollowList[i])
	//	//}
	//	return pack.Users(Users, isFollowList), nil
	//	// return pack.Users(Users, false), nil
	//} else {
	//	isFollowList := make([]bool, len(Users))
	//	for i := range Users {
	//		isFollowList[i] = false
	//	}
	//	return pack.Users(Users, isFollowList), nil // default
	//}
}
