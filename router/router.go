package router

import (
	"github.com/gin-gonic/gin"
	"orca/controller/menu"
	"orca/middleware"
)

func Add(server *gin.Engine) {
	server.Use(middleware.Cors())
	server.Use(middleware.GinLogger(), middleware.GinRecovery(true))

	server.POST("/menu", menu.Controller.Create)
	server.GET("/menu/:code", menu.Controller.Get)
	server.GET("/menu", menu.Controller.List)
	server.DELETE("/menu", menu.Controller.Delete)
	server.PUT("/menu/:code", menu.Controller.Update)
}
