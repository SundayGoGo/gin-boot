/**
 * @Author: Tobin
 * @Description:
 * @File:  router
 * @Version: 1.0.0
 * @Date: 2021/3/22 7:42 下午
 */

package router

import (
	"gin-boot/context"
	"gin-boot/handler"
	"gin-boot/middlewares"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	router := gin.Default()
	// 要在路由组之前全局使用「跨域中间件」, 否则OPTIONS会返回404
	router.Use(middlewares.Cors())

	apiRouter := router.Group(``)
	//apiRouter.Use(middlewares.Jwt()) //加载jwt中间件
	// mapper 上下文
	//mapperContext := context.NewMapperContext()
	userRouter := apiRouter.Group("/user")
	serviceContext := srv.NewServiceContext()
	{
		userHandler := handler.NewUserHandler(serviceContext)
		userRouter.POST("", userHandler.CreateUser)
	}

	// 获取ip地址
	accessRouter := apiRouter.Group(`/access`)
	{
		accessHandler := handler.NewAccessHandler(serviceContext)
		accessRouter.GET(``, accessHandler.CreateAccess)
		accessRouter.GET(`/list`, accessHandler.GetAccessList)
	}
	_ = router.Run(":9300")
}
