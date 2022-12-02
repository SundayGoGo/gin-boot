/**
 * @Author: Tobin
 * @Description:
 * @File:  Model
 * @Version: 1.0.0
 * @Date: 2021/3/22 10:29 下午
 */

package models

import (
	"time"
)

/**
@see https://colobu.com/2017/06/21/json-tricks-in-Go/#%E7%94%A8%E5%AD%97%E7%AC%A6%E4%B8%B2%E4%BC%A0%E9%80%92%E6%95%B0%E5%AD%97
*/
type BaseModel struct {
	ID        int64      `gorm:"primary_key" json:"id,string"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt" sql:"index"`
}
