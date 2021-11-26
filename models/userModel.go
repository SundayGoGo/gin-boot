/**
 * @Author: Tobin
 * @Description:
 * @File:  UserModel
 * @Version: 1.0.0
 * @Date: 2021/3/22 7:45 下午
 */

package models

type User struct {
	BaseModel
	Username string `gorm:"unique;not null"`
	Password string `json:"-"`
	Salt     string
	Name     string
	Avatar   string
}

//gorm.Model                         //内嵌gorm.Model
//Name         string                  //名字
//Age          sql.NullInt64        //年龄 零值类型
//Birthday     *time.Time
//Email        string  `gorm:"type:varchar(100);unique_index"`  //结构体的tag
//Role         string  `gorm:"size:255"` // 设置字段大小为255
//MemberNumber *string `gorm:"unique;not null"` // 设置会员号（member number）唯一并且不为空
//Num          int     `gorm:"AUTO_INCREMENT"` // 设置 num 为自增类型
//Address      string  `gorm:"index:addr"` // 给address字段创建名为addr的索引
//IgnoreMe     int     `gorm:"-"` // 忽略本字段
