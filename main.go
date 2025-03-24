package main

import (
	"usermanager/pkg/logger"
	"usermanager/router"
)

func main() {

	err := router.Router().Run(":8088")
	if err != nil {
		logger.Error(map[string]interface{}{"main error": err.Error()})
		return
	}
}
