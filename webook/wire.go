//go:build wireinject

package main

import (
	"github.com/ClearloveHn/webook/webook/internal/repository"
	"github.com/ClearloveHn/webook/webook/internal/repository/cache"
	"github.com/ClearloveHn/webook/webook/internal/repository/dao"
	"github.com/ClearloveHn/webook/webook/internal/service"
	"github.com/ClearloveHn/webook/webook/internal/web"
	"github.com/ClearloveHn/webook/webook/ioc"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func InitWebServer() *gin.Engine {
	wire.Build(
		// 第三方依赖
		ioc.InitRedis, ioc.InitDB,
		// DAO 部分
		dao.NewUserDAO,

		// cache 部分
		cache.NewCodeCache, cache.NewUserCache,

		// repository 部分
		repository.NewCachedUserRepository,
		repository.NewCodeRepository,

		// Service 部分
		ioc.InitSMSService,
		ioc.InitWechatService,
		service.NewUserService,
		service.NewCodeService,

		// handler 部分
		web.NewUserHandler,
		web.NewOAuth2WechatHandler,
		ioc.InitGinMiddlewares,
		ioc.InitWebServer,
	)
	return gin.Default()
}
