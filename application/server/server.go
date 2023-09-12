package main

import (
	"flag"
	"fmt"
	"gozero-login-pro/application/server/internal/config"
	"gozero-login-pro/application/server/internal/handler"
	"gozero-login-pro/application/server/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "application/server/etc/server-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	//logx.SetWriter(logx.NewWriter())

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
