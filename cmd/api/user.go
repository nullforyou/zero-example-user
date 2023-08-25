package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"go-zero-base/utils/response"
	"go-zero-base/utils/xerr"
	"net/http"
	"user/cmd/api/internal/config"
	"user/cmd/api/internal/handler"
	"user/cmd/api/internal/svc"
)

var configFile = flag.String("f", "etc/user-api-dev.yaml", "the config file")

func main() {
	flag.Parse()

	logx.DisableStat()

	var c config.Config
	conf.MustLoad(*configFile, &c, conf.UseEnv())

	server := rest.MustNewServer(c.RestConf, rest.WithUnauthorizedCallback(func(w http.ResponseWriter, r *http.Request, err error) {
		//JWT验证失败自定义处理
		response.Response(r, w, nil, xerr.NewBusinessError(xerr.SetCode(xerr.ErrorTokenExpire)))
	}))
	defer server.Stop()

	serverCtx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, serverCtx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
