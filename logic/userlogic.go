/**
 * @Author: Tobin
 * @Description:
 * @File:  UserService
 * @Version: 1.0.0
 * @Date: 2021/3/22 7:46 下午
 */

package logic

import (
	"errors"
	"fmt"
	srv "gin-boot/context"
	"gin-boot/databases/mysql"
	"gin-boot/models"
	"gin-boot/utils"
)

type UserLogic struct {
	srv *srv.ServiceContext
}

func NewUserLogic(srv *srv.ServiceContext) *UserLogic {
	return &UserLogic{
		srv: srv,
	}
}

// CreateUser 创建用户
func (u *UserLogic) CreateUser(userDao *models.User) (*models.User, error) {
	userDao.ID = utils.GenerateIntId()
	db := mysql.GetDB()

	if err := db.Create(&userDao).Error; err != nil {
		fmt.Println("插入失败", err)
		return nil, err
	}
	return userDao, nil
}

// 通过id查询用户

func (u *UserLogic) GetUser(uid int64) (*models.User, error) {

	db := mysql.GetDB()
	userDao := &models.User{}
	db.Where(`id=?`, uid).Find(userDao)

	if userDao.ID == 0 {
		return nil, errors.New(`data not found`)
	}
	return userDao, nil
}

// 修改记录
func (u *UserLogic) UpdateUser(userDao *models.User) error {
	db := mysql.GetDB()
	if err := db.Where(`id=?`, userDao.ID).Save(&userDao).Error; err != nil {
		return err
	}
	return nil

}

// 删除记录
func (u *UserLogic) DeleteUser(uid int64) error {
	db := mysql.GetDB()

	if err := db.Where("id = ?", uid).Delete(&models.User{}).Error; err != nil {
		return err
	}
	return nil
}

func (u *UserLogic) GetUserByUsername(username string) *models.User {
	userDao := &models.User{}
	mysql.GetDB().Where(`username=?`, username).Find(userDao)
	if userDao.ID == 0 {
		return nil
	}
	return userDao

}
