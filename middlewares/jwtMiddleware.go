/**
 * @Author: Tobin
 * @Description:
 * @File:  CorsMiddleware
 * @Version: 1.0.0
 * @Date: 2021/3/22 7:47 下午
 */

package middlewares

import (
	"gin-boot/enums/jwtEnum"
	utils "gin-boot/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.Request.Header.Get("Authorization")
		if authorization == "" || !strings.HasPrefix(authorization, jwtEnum.JwtBearer) {
			c.Abort()
			c.JSON(http.StatusUnauthorized, nil)
			return
		}

		// 进入过滤器
		_, err := utils.LoginUserGetJWTToken(strings.Replace(authorization, jwtEnum.JwtBearer, "", 1))
		if err != nil {
			c.Abort()
			c.JSON(http.StatusUnauthorized, nil)
			return
		}

		c.Next()

	}
}
