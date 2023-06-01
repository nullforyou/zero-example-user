package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
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

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	serverCtx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, serverCtx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
