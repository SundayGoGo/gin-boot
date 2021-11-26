/**
 * @Author: Tobin
 * @Description:
 * @File:  UserService
 * @Version: 1.0.0
 * @Date: 2021/3/22 7:46 下午
 */

package services

import (
	"gin-web/context"
	"gin-web/models"
)

type UserService struct {
	mac *context.MapperContext
}

func NewUserService(mac *context.MapperContext) *UserService {
	return &UserService{mac: mac}
}

// CreateUser 创建用户
func (u *UserService) CreateUser(userDao models.User) (*models.User, error) {

	user, err := u.mac.UserMapper.CreateUser(&userDao)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// 通过id查询用户

func (u *UserService) GetUser(uid int64) (*models.User, error) {

	user, err := u.mac.UserMapper.GetUserById(uid)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// 修改记录
func (u *UserService) UpdateUser(userDao *models.User) error {
	err := u.mac.UserMapper.UpdateUser(userDao)
	if err != nil {
		return err
	}
	return nil

}

// 删除记录
func (u *UserService) DeleteUser(uid int64) error {
	err := u.mac.UserMapper.DeleteUser(uid)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserService) GetUserByUsername(username string) *models.User {
	user, err := u.mac.UserMapper.GetUserByUsername(username)
	if err != nil {
		return nil
	}

	return user

}
