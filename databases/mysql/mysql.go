/**
 * @Author: Tobin
 * @Description:
 * @File:  mysql
 * @Version: 1.0.0
 * @Date: 2021/3/22 7:43 下午
 */

package mysql

import (
	"fmt"
	"gin-web/enums/mysqlEnum"
	"gin-web/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // 包装
)

var _db *gorm.DB // ‼️ *****全局变量禁止直接操作数据库，使用时需实例化临时变量操作数据库 eg: Mysql.getDB()获取db对象即可。**** ‼️

func init() {

	var err error
	connArgs := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local&timeout=10s", mysqlEnum.User, mysqlEnum.Password, mysqlEnum.Host, mysqlEnum.Port, mysqlEnum.DB)

	_db, err = gorm.Open("mysql", connArgs)
	if err != nil {

		fmt.Printf("mysql connect error %v", err)
	}
	if _db.Error != nil {
		fmt.Printf("database error %v", _db.Error)
	} else {
		//_db.LogMode(true) // 控制台打印sql
		_db.SingularTable(true) // 取消后缀s
		//设置数据库连接池参数
		_db.DB().SetMaxOpenConns(100)   //设置数据库连接池最大连接数
		_db.DB().SetMaxIdleConns(20)    //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭。
		_db.AutoMigrate(&models.User{}) // 自动创建表结构
	}

}

func GetDB() *gorm.DB {
	return _db
}
