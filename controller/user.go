package controller

import (
	"strconv"
	"usermanager/module"

	"github.com/gin-gonic/gin"
)

type UserController struct {
}

type UserSearch struct {
	Id       int    `json:"id"`
	UserName string `json:"name"`
}

type UserPut struct {
	UserName string `json:"name"`
	Age      uint8  `json:"age"`
}

func (u UserController) GetUser(c *gin.Context) {
	defer ErrorRecover(c)
	id := c.Param("id")
	//name := c.PostForm("name")
	//name := c.PostFormMap("name")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}
	userInfo := module.GetUserById(idInt)
	Success(c, userInfo, 1)
}

func (u UserController) PutUser(c *gin.Context) {
	defer ErrorRecover(c)
	// param := make(map[string]interface{})
	param := &UserPut{}
	err := c.ShouldBindJSON(&param)
	if err != nil {
		Fail(c, err.Error())
		return
	}
	// fmt.Println(param)
	uid := module.AddUser(param.UserName, param.Age)
	Success(c, uid, 1)
}

func (u UserController) SearchUser(c *gin.Context) {
	defer ErrorRecover(c)
	param := &UserSearch{}
	err := c.ShouldBindJSON(&param)
	if err != nil {
		Fail(c, err.Error())
		return
	}
	Success(c, "search success", 1)
}

func (u UserController) GetUserFail(c *gin.Context) {
	Fail(c, "未找到用户")
}
