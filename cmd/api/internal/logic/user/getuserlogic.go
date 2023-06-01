package user

import (
	"context"
	"errors"
	"go-zero-base/utils/xerr"
	"gorm.io/gorm"
	"user/cmd/api/internal/svc"
	"user/cmd/api/internal/types"
	"user/cmd/dao/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLogic) GetUser(req *types.UserInfoReq) (resp *types.UserInfoResp, err error) {
	err = l.svcCtx.DbEngine.Model(model.Member{}).First(&resp, req.Id).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, xerr.NewBusinessError(xerr.SetCode(xerr.ErrorBusiness), xerr.SetMsg("用户不存在"))
	}
	return resp, nil
}
