package model

import (
	"TestDemo/utils"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
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

	db.SingularTable(false)                          // 自动表名复数形式
	db.AutoMigrate(&User{}, &Article{}, &Category{}) // 自动创建表 --失败 还是自己建表把~ --改大写成功了 然后新版说是需要注意foreignkey之类的设置 在字段定义后面的~
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(10 * time.Second)
}
