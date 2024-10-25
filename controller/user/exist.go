package user

import (
	"github.com/gin-gonic/gin"
	"orca/models"
	"orca/pkg/code"
	"orca/pkg/db"
	"orca/pkg/errors"
	"orca/pkg/response"
)

func (uc *userCtrl) Exist(c *gin.Context) {
	username := c.Query("username")
	phone := c.Query("phone")
	email := c.Query("email")

	if username == "" && phone == "" && email == "" {
		response.Fail(c, errors.WithCode(code.ErrBadRequest, "请求不合法"))
		return
	}

	var user models.User
	if err := db.Mysql.Table("users").
		Select("username").
		Joins("join user_profile on user_profile.user_id = users.user_id").
		Where("username = ? or phone = ? or email = ?", username, phone, email).
		First(&user).Error; err != nil {
		response.Fail(c, errors.WithCode(code.ErrUserNotFound, "用户登陆时发生错误，用户名/手机/邮箱记录不存在"))
		return
	}

	response.Success(c, true, "用户存在")
}
