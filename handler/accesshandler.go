package handler

import (
	"encoding/json"
	"gin-boot/bean"
	srv "gin-boot/context"
	"gin-boot/logic"
	"gin-boot/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type AccessHandler struct {
	srv *srv.ServiceContext
}

func NewAccessHandler(srv *srv.ServiceContext) *AccessHandler {
	return &AccessHandler{
		srv: srv,
	}
}

//  获取客户端IP地址
func (u *AccessHandler) CreateAccess(ctx *gin.Context) {
	accessLogic := logic.NewAccessLogic(u.srv)
	ip := ctx.ClientIP()
	if ip == `::1` {
		ip = "127.0.0.1"
	}

	access := models.Access{
		Ip:      ip,
		Pro:     "",
		City:    "",
		RawJson: "",
		Addr:    "",
	}
	createAccess, err := accessLogic.CreateAccess(&access)
	if err != nil {
		logrus.Info(err.Error())
		ctx.JSON(400, ``)
		return
	}

	go func() {
		address := accessLogic.GetRemoteAddress(ip)
		m := make(map[string]string)
		err := json.Unmarshal(address, &m)
		if err != nil {
			logrus.Info("读取失败", err.Error())
			return
		}
		createAccess.Pro = m["pro"]
		createAccess.City = m["city"]
		createAccess.Addr = m["addr"]
		marshal, err := json.Marshal(&m) // 转为json字符串
		createAccess.RawJson = string(marshal)
		err = accessLogic.UpdateAccess(createAccess)
		if err != nil {
			logrus.Info("存储失败", err.Error())
			return
		}
		logrus.Info("异步获取客户端IP信息")
	}()
	ctx.JSON(200, "success")
}

func (u *AccessHandler) GetAccessList(ctx *gin.Context) {
	accessLogic := logic.NewAccessLogic(u.srv)
	list := accessLogic.GetList()
	response := bean.ResponseBean{
		Code:  0,
		Msg:   "success",
		Data:  list,
		Total: len(*list),
	} // 返回结构体
	ctx.JSON(200, response)
	return
}
