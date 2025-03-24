package module

import (
	"fmt"
	"usermanager/dao"

	"gorm.io/gorm"
)

type UserModule struct {
	gorm.Model
	Name string `gorm:"type:varchar(255)"`
	Age  uint8  `gorm:"type:int"`
}

func (u UserModule) TableName() string {
	return "sys_user"
}

func AddUser(user *UserModule) uint {
	if user == nil {
		fmt.Println(user)
		panic("add user is null")
	}
	result := dao.Db.Create(user) // 通过数据的指针来创建
	if result.Error != nil {
		fmt.Println(result.Error)
		panic("add user error")
	}
	return user.ID
}

func GetUserById(id int) (u UserModule) {
	err := dao.Db.Where("id=?", id).First(&u).Error
	if err != nil {
		panic("未找到id")
	}
	return
}
