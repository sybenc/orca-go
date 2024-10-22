package menu

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"orca/models"
	"orca/pkg/code"
	"orca/pkg/db"
	"orca/pkg/errors"
	"orca/pkg/response"
)

func (m *menuController) Delete(c *gin.Context) {
	var menu []models.Menu
	ids := c.QueryArray("ids")
	if len(ids) == 0 {
		response.Fail(c, errors.WithCode(code.ErrValidate, "无效的菜单ID"))
		return
	}
	if db.Mysql.Model(&models.Menu{}).Where("menu_id in ?", ids).Find(&menu).RowsAffected != int64(len(ids)) {
		response.Fail(c, errors.WithCode(code.ErrValidate, "存在无效的菜单ID"))
		return
	}
	err := db.Mysql.Transaction(func(tx *gorm.DB) error {
		if err := db.Mysql.Where("menu_id in ?", ids).Delete(&models.Menu{}).Error; err != nil {
			return errors.WithCode(code.ErrValidate, "删除菜单时，发生错误")
		}
		return nil
	})

	if err != nil {
		response.Fail(c, err)
	}

	response.Success(c, nil, "删除菜单成功")
}
