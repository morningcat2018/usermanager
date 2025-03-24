package module

import (
	"fmt"
	"usermanager/dao"
	"usermanager/pkg/logger"

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

func init() {
	dao.GetDB().AutoMigrate(&UserModule{})
}

func AddUser(name string, age uint8) uint {
	logger.Debug(map[string]interface{}{"AddUser name": name, "age": age})
	user := UserModule{Name: name, Age: age}

	result := dao.GetDB().Create(&user) // 通过数据的指针来创建
	if result.Error != nil {
		logger.Error(map[string]interface{}{"AddUser error": result.Error})
		panic("add user error")
	}
	return user.ID
}

func GetUserById(id int) (u UserModule) {
	err := dao.GetDB().Where("id=?", id).First(&u).Error
	if err != nil {
		panic(fmt.Sprintf("未找到id:%d", id))
	}
	return
}
