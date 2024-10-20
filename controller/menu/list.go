package menu

import (
	"github.com/gin-gonic/gin"
	"orca/models"
	"orca/pkg/code"
	"orca/pkg/db"
	errors "orca/pkg/erorrs"
	"orca/pkg/response"
)

func (m *menuController) List(c *gin.Context) {
	var menuList models.MenuList

	if err := db.Mysql.Table("menu").Model(&models.Menu{}).Find(&menuList.Items).Count(&menuList.Total).Error; err != nil {
		response.Fail(c, errors.WithCode(code.ErrInternalServer, "查询菜单列表时发生错误："))
		return
	}
	response.Success(c, menuList, "查询菜单列表成功")
}
