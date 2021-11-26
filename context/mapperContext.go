/**
 * @Author: Tobin
 * @Description:
 * @File:  serviceContent
 * @Version: 1.0.0
 * @Date: 2021/11/26 10:39 下午
 */

package context

import (
	"gin-web/mapper"
)
// 注入mapper接口 供service使用

type MapperContext struct {
	UserMapper mapper.UserMapper
}

func NewMapperContext() *MapperContext {
	return &MapperContext{UserMapper: mapper.NewUserMapper()}
}


