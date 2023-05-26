package logic

import (
	"context"
	"user/cmd/dao/query"

	"user/cmd/rpc/internal/svc"
	"greet-pb/user/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *user.IdReq) (*user.UserInfoReply, error) {

	query.SetDefault(l.svcCtx.DbEngine)
	ctx := context.Background()
	memberDao := query.Member
	memberModel, err := memberDao.WithContext(ctx).Where(memberDao.ID.Eq(in.Id)).First()

	if err != nil {
		return nil, err
	}

	return &user.UserInfoReply{
		Id:     memberModel.ID,
		Mobile:   memberModel.Mobile,
		Nikename: memberModel.Nickname,
	}, nil


}
