/**
 * @Author: Tobin
 * @Description:
 * @File:  router
 * @Version: 1.0.0
 * @Date: 2021/3/22 7:42 下午
 */

package router

import (
	"gin-web/context"
	"gin-web/controller"
	"gin-web/middlewares"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	router := gin.Default()
	// 要在路由组之前全局使用「跨域中间件」, 否则OPTIONS会返回404
	router.Use(middlewares.Cors())

	apiRouter := router.Group(``)
	apiRouter.Use(middlewares.Jwt()) //加载jwt中间件
	// mapper 上下文
	mapperContext := context.NewMapperContext()
	userRouter := apiRouter.Group("/user")
	{
		userController := controller.NewUserController(mapperContext)
		userRouter.POST("", userController.CreateUser)
	}

	_ = router.Run(":9300")
}
