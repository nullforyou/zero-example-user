package user

import (
	"go-common/utils/response"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"user/cmd/api/internal/logic/user"
	"user/cmd/api/internal/svc"
	"user/cmd/api/internal/types"
)

func RegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterReq
		if err := httpx.Parse(r, &req); err != nil {
			response.ParamErrorResponse(r, w, err)
			return
		}

		//验证参数
		if err := svcCtx.Validator.Validate.StructCtx(r.Context(), req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewRegisterLogic(r.Context(), svcCtx)
		resp, err := l.Register(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		} else {
			response.Response(r, w, resp)
			return
		}
	}
}
