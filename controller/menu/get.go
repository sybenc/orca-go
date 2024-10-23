package menu

import (
	"github.com/gin-gonic/gin"
	"orca/models"
	"orca/pkg/code"
	"orca/pkg/db"
	"orca/pkg/errors"
	"orca/pkg/response"
)

func (m *menuController) Get(c *gin.Context) {
	code_ := c.Param("code")
	var menu models.Menu
	if db.Mysql.Model(&models.Menu{}).Where("code = ?", code_).First(&menu).RowsAffected != 1 {
		response.Fail(c, errors.WithCode(code.ErrMenuNotFound, "菜单（id："+code_+"）不存在"))
		return
	}
	response.Success(c, menu, "查询菜单成功")
}
