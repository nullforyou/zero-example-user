package user

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"go-common/custom_string"
	"go-common/utils/xerr"
	"user/cmd/api/internal/svc"
	"user/cmd/api/internal/types"
	"user/cmd/dao/model"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	var count int64
	l.svcCtx.DbEngine.Model(model.Member{}).Where("mobile = ?", req.Mobile).Count(&count)
	if count > 0 {
		return nil, xerr.NewBusinessError(xerr.SetCode("MobileExistsError"), xerr.SetMsg("手机号已存在"))
	}

	member := model.Member{}
	member.Mobile = req.Mobile
	member.Nickname = req.Mobile
	member.Password = custom_string.CalcMD5(req.Password)
	l.svcCtx.DbEngine.Create(&member)

	jwtToken, err := custom_string.GetJwtToken(l.svcCtx.Config.Jwt.AccessSecret, l.svcCtx.Config.Jwt.AccessExpire, member.ID)
	if err != nil {
		return nil, xerr.NewBusinessError(xerr.SetCode("GenerateJWTError"), xerr.SetMsg("生成JWT错误"))
	}
	return &types.RegisterResp{
		Id: member.ID,
		Mobile: member.Mobile,
		Nickname: member.Nickname,
		AccessToken: jwtToken,
		AccessExpire: l.svcCtx.Config.Jwt.AccessExpire,
	}, nil
}
