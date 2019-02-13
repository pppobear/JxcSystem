package staff

import (
	"github.com/gin-gonic/gin"

	"pppobear.cn/jxc-backend/models"
	"pppobear.cn/jxc-backend/module/errno"
	"pppobear.cn/jxc-backend/module/handler"
)

func DeleteStaff(c *gin.Context) {
	id := c.Param("id")
	if len(id) == 0 {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	if err := models.DeleteStaff(id); err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}
	handler.SendResponse(c, nil, "Successfully deleted.")
}

func DeleteSpe(c *gin.Context) {
	idStr := c.Param("id")
	if err := models.DeleteSpecialty(idStr); err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}
	handler.SendResponse(c, nil, "Successfully deleted.")
}

func DeleteDep(c *gin.Context) {
	idStr := c.Param("id")
	if err := models.DeleteDepartment(idStr); err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}
	handler.SendResponse(c, nil, "Successfully deleted.")
}
