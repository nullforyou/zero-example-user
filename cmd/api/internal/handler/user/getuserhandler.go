package user

import (
	"go-common/utils/response"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"user/cmd/api/internal/logic/user"
	"user/cmd/api/internal/svc"
	"user/cmd/api/internal/types"
)

func GetUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			response.ParamErrorResponse(r, w, err)
			return
		}

		l := user.NewGetUserLogic(r.Context(), svcCtx)
		resp, err := l.GetUser(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			response.Response(r, w, resp)
		}
	}
}
