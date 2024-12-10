package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model

	Username string `gorm:"type:varchar(20);not null" json:"username"`
	Pwd      string `gorm:"type:varchar(255);not null" json:"pwd"`
	Role     int    `gorm:"type:int" json:"role"`
	Avatar   string `gorm:"type:varchar(255)" json:"avatar"`
}
