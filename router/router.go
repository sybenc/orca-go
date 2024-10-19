package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"orca/middleware"
	"orca/pkg/code"
	"orca/pkg/erorrs"
)

func ABC() {
	err := errors.WithCode(code.Success, "Success")
	err = errors.WithMessage(err, "Success2")
	err = errors.Wrap(err, "mmmm")
	err = errors.WrapC(err, code.InternalServer, "Internal Server Error")
	fmt.Printf("%+v\n", err)

}

func Add(server *gin.Engine) {
	server.Use(middleware.Cors())
	server.Use(middleware.GinLogger(), middleware.GinRecovery(true))
}
