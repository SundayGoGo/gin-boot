/**
 * @Author: Tobin
 * @Description:
 * @File:  UserController
 * @Version: 1.0.0
 * @Date: 2021/3/22 7:49 下午
 */

package controller

import (
	"fmt"
	"gin-web/bean"
	"gin-web/context"
	"gin-web/models"
	"gin-web/services"
	"gin-web/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	Mac *context.MapperContext
}

func NewUserController(mac *context.MapperContext) *UserController {
	return &UserController{Mac: mac}
}

// 创建用户

func (u *UserController) CreateUser(ctx *gin.Context) {
	userService := services.NewUserService(u.Mac)
	response := bean.ResponseBean{
		Code: 0,
	} // 返回结构体

	// 接收参数
	username := ctx.Query("username")
	password := ctx.Query("password")
	name := ctx.Query("name")
	fmt.Println(username)

	if username == "" || password == "" {
		response.Code = -1
		response.Msg = "The username and password cannot be empty"
		goto Error
	} else {
		// 首先查询username是否存在
		user := userService.GetUserByUsername(username)
		if user != nil {
			response.Code = -1
			response.Msg = "User already exists！"
			goto Error
		}

		userDao := models.User{}
		salt := utils.NewRandom4V()
		userDao.Username = username
		userDao.Salt = salt
		userDao.Password = utils.MakeHashCode(fmt.Sprintf(`%s%s`, password, salt))
		userDao.Name = name
		newUser, err := userService.CreateUser(userDao)
		if err != nil {
			return
		}
		response.Data = newUser
		ctx.JSON(http.StatusOK, response)
		return
	}

Error:
	ctx.JSON(http.StatusOK, response)

}
