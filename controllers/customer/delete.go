package customer

import (
	"github.com/gin-gonic/gin"

	"pppobear.cn/jxc-backend/models"
	"pppobear.cn/jxc-backend/module/errno"
	"pppobear.cn/jxc-backend/module/handler"
)

func Delete(c *gin.Context) {
	id := c.Param("id")
	if len(id) == 0 {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	if err := models.DeleteCustomer(id); err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}
	handler.SendResponse(c, nil, "Successfully deleted.")
}
