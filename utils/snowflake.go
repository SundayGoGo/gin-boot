/**
 * @Author: Tobin
 * @Description:
 * @File:  Snowflake
 * @Version: 1.0.0
 * @Date: 2021/3/19 10:30 上午
 */

package utils

import "github.com/GUAIK-ORG/go-snowflake/snowflake"

func GenerateUintId() uint {
	res, _ := snowflake.NewSnowflake(int64(0), int64(0))
	// ......
	// (s *Snowflake) NextVal() int64
	// 返回1 (int64): 唯一ID
	return uint(res.NextVal())
}
func GenerateIntId() int64 {
	res, _ := snowflake.NewSnowflake(int64(0), int64(0))
	// ......
	// (s *Snowflake) NextVal() int64
	// 返回1 (int64): 唯一ID
	return res.NextVal()
}
