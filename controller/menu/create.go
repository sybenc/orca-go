package menu

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"orca/models"
	"orca/pkg/code"
	"orca/pkg/db"
	errors "orca/pkg/erorrs"
	"orca/pkg/response"
)

func (m *menuController) Create(c *gin.Context) {
	var menu models.Menu
	if err := c.ShouldBind(&menu); err != nil {
		response.Fail(c, errors.WithCode(code.ErrBind, "创建菜单时数据绑定错误"))
		return
	}

	if db.Mysql.Model(&models.Menu{}).
		Where("code=? or label=?", menu.Code, menu.Label).First(&menu).RowsAffected > 0 {
		response.Fail(c, errors.WithCode(code.ErrMenuAlreadyExist, "创建菜单时资源发生冲突"))
		return
	}

	if menu.Type != models.EnumMenuTypeMenu {
		menu.Component = nil
		menu.ParentID = nil
		menu.Route = nil
	}
	err := db.Mysql.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&models.Menu{}).Create(&menu).Error; err != nil {
			return errors.WithCode(code.ErrInternalServer, "将菜单数据插入到数据库时发生错误")
		}
		return nil
	})
	if err != nil {
		response.Fail(c, err)
		return
	}
	response.Success(c, nil, "创建菜单成功")
}
