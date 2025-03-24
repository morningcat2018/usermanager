package main

import (
	"fmt"
	"usermanager/router"
)

func main() {
	err := router.Router().Run(":8088")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
