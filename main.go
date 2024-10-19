package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"orca/conf"
	"orca/router"
)

func init() {
	conf.InitConfig("dev")
}

func main() {
	gin.SetMode(conf.GetString("server.mode"))
	server := gin.Default()
	router.Add(server)
	router.ABC()
	err := server.Run(fmt.Sprintf(":%s", conf.GetString("server.port")))
	if err != nil {
		panic(fmt.Sprintf("Failed to start the service: %s", err.Error()))
	}
}
