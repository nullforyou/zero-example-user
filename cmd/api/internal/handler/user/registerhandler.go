package user

import (
	"go-zero-base/utils/response"
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
			response.ValidateErrOrResponse(r, w, err, svcCtx.Validator.Trans)
			return
		}
		//验证参数
		if err := svcCtx.Validator.Validate.StructCtx(r.Context(), req); err != nil {
			response.ValidateErrOrResponse(r, w, err, svcCtx.Validator.Trans)
			return
		}

		l := user.NewRegisterLogic(r.Context(), svcCtx)
		resp, err := l.Register(&req)
		response.Response(r, w, resp, err)
		return
	}
}
