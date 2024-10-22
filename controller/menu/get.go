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
	id := c.Param("id")
	var menu models.Menu
	if db.Mysql.Model(&models.Menu{}).Where("menu_id = ?", id).First(&menu).RowsAffected != 1 {
		response.Fail(c, errors.WithCode(code.ErrMenuNotFound, "菜单（id："+id+"）不存在"))
		return
	}
	response.Success(c, menu, "查询菜单成功")
}
