// Code generated by Kitex v0.3.1. DO NOT EDIT.

package videoservice

import (
	"context"
	"douyin/v1/kitex_gen/video"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	GetPublishListByUser(ctx context.Context, userId int64, callOptions ...callopt.Option) (r *video.PublishListResponse, err error)
	GetVideosByLastTime(ctx context.Context, lastTime int64, callOptions ...callopt.Option) (r *video.VideoFeedResponse, err error)
	PublishVideo(ctx context.Context, publishedVideo *video.Video, callOptions ...callopt.Option) (r *video.BaseResp, err error)
	FavoriteByUser(ctx context.Context, request *video.FavoriteActionRequest, callOptions ...callopt.Option) (r *video.BaseResp, err error)
	GetFavoriteListBYUser(ctx context.Context, request *video.FavoriteListRequest, callOptions ...callopt.Option) (r *video.FavoriteListResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kVideoServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kVideoServiceClient struct {
	*kClient
}

func (p *kVideoServiceClient) GetPublishListByUser(ctx context.Context, userId int64, callOptions ...callopt.Option) (r *video.PublishListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetPublishListByUser(ctx, userId)
}

func (p *kVideoServiceClient) GetVideosByLastTime(ctx context.Context, lastTime int64, callOptions ...callopt.Option) (r *video.VideoFeedResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetVideosByLastTime(ctx, lastTime)
}

func (p *kVideoServiceClient) PublishVideo(ctx context.Context, publishedVideo *video.Video, callOptions ...callopt.Option) (r *video.BaseResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.PublishVideo(ctx, publishedVideo)
}

func (p *kVideoServiceClient) FavoriteByUser(ctx context.Context, request *video.FavoriteActionRequest, callOptions ...callopt.Option) (r *video.BaseResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FavoriteByUser(ctx, request)
}

func (p *kVideoServiceClient) GetFavoriteListBYUser(ctx context.Context, request *video.FavoriteListRequest, callOptions ...callopt.Option) (r *video.FavoriteListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetFavoriteListBYUser(ctx, request)
}
