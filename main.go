/**
 * @Author: Tobin
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2021/3/22 7:33 下午
 */

package main

import (
	"gin-boot/config"
	"gin-boot/databases/mysql"
	"gin-boot/router"
	"github.com/jinzhu/gorm"

	"github.com/sirupsen/logrus"
)

func main() {

	defer func() {
		if err := recover(); err != nil {
			logrus.Info("Catch Error:", err)
		}
	}()

	// 1. 加载配置文件
	loadConfig := config.LoadConfig()

	// 2. 初始化数据库
	mysql.LoadMysql(loadConfig)

	defer func(db *gorm.DB) {
		var err = db.Close()
		if err != nil {

		}
	}(mysql.GetDB())

	router.InitRouter(loadConfig)
}
