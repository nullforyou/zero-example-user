// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package userclient

import (
	"context"
	"user/cmd/rpc/types/user"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	IdReq         = user.IdReq
	UserInfoReply = user.UserInfoReply

	User interface {
		GetUser(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*UserInfoReply, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

func (m *defaultUser) GetUser(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*UserInfoReply, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.GetUser(ctx, in, opts...)
}
