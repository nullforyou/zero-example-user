package user

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"go-common/tool"
	"go-common/utils/xerr"
	"gorm.io/gorm"
	"user/cmd/api/internal/svc"
	"user/cmd/api/internal/types"
	"user/cmd/dao/query"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	query.SetDefault(l.svcCtx.DbEngine)
	ctx := context.Background()
	memberDao := query.Member
	memberModel, err := memberDao.WithContext(ctx).Where(memberDao.Mobile.Eq(req.Mobile)).First()
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound){
		return nil, xerr.NewBusinessError(xerr.SetCode("MemberNotExistsError"))
	}

	if memberModel.Password != tool.CalcMD5(req.Password) {
		return nil, xerr.NewBusinessError(xerr.SetCode("LoginError"), xerr.SetMsg("账号或密码错误"))
	}

	jwtToken, err := tool.GetJwtToken(l.svcCtx.Config.Jwt.AccessSecret, l.svcCtx.Config.Jwt.AccessExpire, memberModel.ID)
	if err != nil {
		return nil, xerr.NewBusinessError(xerr.SetCode("GenerateJWTError"), xerr.SetMsg("生成JWT错误"))
	}
	return &types.LoginResp{
		Id: memberModel.ID,
		Mobile: memberModel.Mobile,
		Nickname: memberModel.Nickname,
		AccessToken: jwtToken,
		AccessExpire: l.svcCtx.Config.Jwt.AccessExpire,
	}, nil


}
