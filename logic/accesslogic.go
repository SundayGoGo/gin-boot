package logic

import (
	"bytes"
	"fmt"
	srv "gin-boot/context"
	"gin-boot/databases/mysql"
	"gin-boot/models"
	"gin-boot/utils"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
)

type AccessLogic struct {
	srv *srv.ServiceContext
}

func NewAccessLogic(srv *srv.ServiceContext) *AccessLogic {
	return &AccessLogic{
		srv: srv,
	}
}

func (u *AccessLogic) CreateAccess(access *models.Access) (*models.Access, error) {
	access.ID = utils.GenerateIntId()
	db := mysql.GetDB()
	if err := db.Create(&access).Error; err != nil {
		fmt.Println("插入失败", err)
		return nil, err
	}
	return access, nil
}

func (u *AccessLogic) UpdateAccess(access *models.Access) error {
	db := mysql.GetDB()
	if err := db.Where(`id=?`, access.ID).Save(&access).Error; err != nil {
		return err
	}
	return nil
}

//

func (u *AccessLogic) GetList() *[]models.Access {
	accessList := make([]models.Access, 0)
	mysql.GetDB().Order(`created_at desc`).Find(&accessList)

	return &accessList
}

// 通过ip地址获取归属地 https://whois.pconline.com.cn/ipJson.jsp?ip=60.210.239.234&json=true

func (u *AccessLogic) GetRemoteAddress(ip string) []byte {
	res, err := utils.HTTPGet(fmt.Sprintf(`https://whois.pconline.com.cn/ipJson.jsp?ip=%s&json=true`, ip))

	if err != nil {
		return nil
	}
	// gbk to utf8
	reader := transform.NewReader(bytes.NewReader(res), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil
	}
	return d
}
