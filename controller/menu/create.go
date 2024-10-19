package menu

import (
	"github.com/gin-gonic/gin"
	"orca/models"
)

func (m *menuController) Create(c *gin.Context) {
	var menu models.Menu
	if err := c.ShouldBind(&menu); err != nil {
	}
}
