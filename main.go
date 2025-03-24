package main

import (
	"usermanager/dao"
	"usermanager/pkg/logger"
	"usermanager/router"
)

func main() {
	logger.Write(dao.Db, "mydebug")

	err := router.Router().Run(":8088")
	if err != nil {
		logger.Error(map[string]interface{}{"main error": err.Error()})
		return
	}
}
