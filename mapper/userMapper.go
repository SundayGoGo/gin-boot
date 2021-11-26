/**
 * @Author: Tobin
 * @Description:
 * @File:  UserModel
 * @Version: 1.0.0
 * @Date: 2021/3/22 7:45 下午
 */

package mapper

import (
	"errors"
	"fmt"
	"gin-web/databases/mysql"
	"gin-web/models"
	"gin-web/utils"
)

type UserMapper interface {
	CreateUser(*models.User) (*models.User, error)
	GetUserById(uid int64) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	DeleteUser(uid int64) error
	UpdateUser(userDao *models.User) error
	GetUserQueryPage(username string, name string) *[]models.User
}

type imlUserMapper struct {
}

func NewUserMapper() UserMapper {
	return &imlUserMapper{}
}

func (u *imlUserMapper) GetUserByUsername(username string) (*models.User, error) {
	userDao := &models.User{}
	mysql.GetDB().Where(`username=?`, username).Find(userDao)
	if userDao.ID == 0 {
		return nil, errors.New(`data not found`)
	}
	return userDao, nil
}
func (u *imlUserMapper) CreateUser(userDao *models.User) (*models.User, error) {
	userDao.ID = utils.GenerateIntId()
	db := mysql.GetDB()

	if err := db.Create(userDao).Error; err != nil {
		fmt.Println("插入失败", err)
		return nil, err
	}
	return userDao, nil

}

// GetUserById 通过id 获取一条记录
func (u *imlUserMapper) GetUserById(uid int64) (*models.User, error) {
	db := mysql.GetDB()
	userDao := &models.User{}
	db.Where(`id=?`, uid).Find(userDao)

	if userDao.ID == 0 {
		return nil, errors.New(`data not found`)
	}
	return userDao, nil
}

// DeleteUser 删除用户
func (u *imlUserMapper) DeleteUser(uid int64) error {
	db := mysql.GetDB()

	if err := db.Where("id = ?", uid).Delete(&models.User{}).Error; err != nil {
		return err
	}
	return nil

}

// UpdateUser 修改记录
func (u *imlUserMapper) UpdateUser(userDao *models.User) error {
	db := mysql.GetDB()
	if err := db.Where(`id=?`, userDao.ID).Save(&userDao).Error; err != nil {
		return err
	}
	return nil
}

// GetUserQueryPage 获取user
func (u *imlUserMapper) GetUserQueryPage(username string, name string) *[]models.User {
	db := mysql.GetDB()
	if username != "" {
		db = db.Where(`username=?`, username)
	}
	if name != "" {
		db = db.Where(`name=?`, name)
	}
	userList := make([]models.User, 0)
	db.Order(`created_at desc`).Find(&userList)

	return &userList

}
