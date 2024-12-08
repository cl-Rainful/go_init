package model

import (
	"TestDemo/utils"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err error

func InitDB() {

	db, err = gorm.Open(utils.Db, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		utils.DbUser, utils.DbPwd, utils.DbHost, utils.DbPort, utils.DbName))
	if err != nil {
		fmt.Println("连接数据库失败, 请检查参数:", err)
	}

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(10 * time.Second)

}
