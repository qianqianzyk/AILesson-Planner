package main

import (
	"flag"
	"fmt"
	"github.com/qianqianzyk/AILesson-Planner/internal/config"
	"github.com/qianqianzyk/AILesson-Planner/internal/handler"
	"github.com/qianqianzyk/AILesson-Planner/internal/logs"
	"github.com/qianqianzyk/AILesson-Planner/internal/service"
	"github.com/qianqianzyk/AILesson-Planner/internal/svc"
	"github.com/qianqianzyk/AILesson-Planner/internal/utils"
	"github.com/unidoc/unioffice/v2/common/license"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/stat"
	"github.com/zeromicro/go-zero/rest"
	"log"
)

var configFile = flag.String("f", "etc/app.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	stat.DisableLog()

	if err := logs.ZapInit(c); err != nil {
		log.Fatal(err.Error())
	}

	server.Use(utils.CorsMiddleware)

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	service.ServiceInit(ctx.MySQLClient, ctx.RedisClient, ctx.ESClient, ctx.MinioClient, ctx.RestyClient, ctx.Neo4jDriver)
	service.ConfigInit(&c)

	err := license.SetMeteredKey(ctx.Config.Unidoc.APIKey)
	if err != nil {
		log.Fatalf("License load failed: %v", err)
	}

	service.StartScheduledCleanup()

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
