package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
	"go-common/utils/response"
	"go-common/utils/xerr"
	"net/http"
	"reflect"
	"user/cmd/api/internal/config"
	"user/cmd/api/internal/handler"
	"user/cmd/api/internal/svc"
)

var configFile = flag.String("f", "etc/user-api.yaml", "the config file")

func main() {
	flag.Parse()

	logx.DisableStat()

	var c config.Config
	conf.MustLoad(*configFile, &c, conf.UseEnv())

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	serverCtx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, serverCtx)

	httpx.SetErrorHandlerCtx(func(ctx context.Context, err error) (int, interface{}) {
		logx.Debugf("进入错误处理程序",  reflect.TypeOf(err))
		switch e := err.(type) {
		case *xerr.BusinessError:
			return http.StatusBadRequest, &response.Body{Code: e.GetErrCode(), Message: e.GetErrMsg()}
		case validator.ValidationErrors:
			errs := err.(validator.ValidationErrors)
			var msg string
			for _, e := range errs {
				msg = e.Translate(serverCtx.Validator.Trans)
				break
			}
			return http.StatusBadRequest, &response.Body{Code: "ValidationError", Message: msg, Data: errs.Translate(serverCtx.Validator.Trans)}
		default:
			fmt.Println("type:unkown,value:", err)
			return http.StatusInternalServerError, nil
		}
	})

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
