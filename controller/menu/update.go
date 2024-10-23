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

func (m *menuController) Update(c *gin.Context) {
	var menu models.Menu
	code_ := c.Param("code")

	if db.Mysql.Model(&models.Menu{}).Where("code = ?", code_).First(&menu).RowsAffected == 0 {
		response.Fail(c, errors.WithCode(code.ErrMenuNotFound, "菜单未找到D"))
		return
	}

	if err := c.ShouldBind(&menu); err != nil {
		response.Fail(c, errors.WithCode(code.ErrBind, "更新菜单时，数据绑定错误"))
		return
	}

	if err := menu.Validate(); err != nil {
		response.Fail(c, errors.WithCode(code.ErrValidate, "更新菜单时，字段验证错误"))
		return
	}

	err := db.Mysql.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&models.Menu{}).Where("code = ?", code_).Save(&menu).Error; err != nil {
			return errors.WithCode(code.ErrInternalServer, "更新菜单失败")
		}
		return nil
	})

	if err != nil {
		response.Fail(c, err)
	}

	response.Success(c, nil, "更新菜单成功")
}
