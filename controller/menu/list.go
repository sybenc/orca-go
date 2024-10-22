package menu

import (
	"github.com/gin-gonic/gin"
	"orca/models"
	"orca/pkg/code"
	"orca/pkg/db"
	"orca/pkg/errors"
	"orca/pkg/response"
	"strconv"
)

func (m *menuController) List(c *gin.Context) {
	var menuList models.MenuList
	code_ := c.Query("code")
	label := c.Query("label")
	type_ := c.Query("type")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	query := db.Mysql.Table("menu").Model(&models.Menu{}).
		Where("code like ?", "%"+code_+"%").
		Where("label like ?", "%"+label+"%").
		Where("type like ?", "%"+type_+"%")

	if err := query.Count(&menuList.Total).Error; err != nil {
		response.Fail(c, errors.WithCode(code.ErrInternalServer, "查询菜单总数时发生错误："))
	}

	if err := query.Offset(offset).Limit(limit).Find(&menuList.Items).Error; err != nil {
		response.Fail(c, errors.WithCode(code.ErrInternalServer, "查询菜单列表时发生错误："))
		return
	}

	response.Success(c, menuList, "查询菜单列表成功")
}
