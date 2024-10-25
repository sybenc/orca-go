package router

import (
	"github.com/gin-gonic/gin"
	"orca/controller/menu"
	"orca/controller/user"
	"orca/middleware"
)

func Add(server *gin.Engine) {
	server.Use(middleware.Cors())
	server.Use(middleware.GinLogger(), middleware.GinRecovery(true))

	server.POST("/menu", menu.Ctrl.Create)
	server.GET("/menu/:code", menu.Ctrl.Get)
	server.GET("/menu", menu.Ctrl.List)
	server.DELETE("/menu", menu.Ctrl.Delete)
	server.PUT("/menu/:code", menu.Ctrl.Update)

	server.GET("/user", user.Ctrl.Exist)
}
