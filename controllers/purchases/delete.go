package purchases

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"pppobear.cn/jxc-backend/models"
	"pppobear.cn/jxc-backend/module/errno"
	"pppobear.cn/jxc-backend/module/handler"
)

func Delete(c *gin.Context) {
	idStr := c.Param("id")
	if err := models.DeletePurchases(idStr); err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}
	handler.SendResponse(c, nil, "Successfully deleted.")
}

func DeleteDetail(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	gid := c.Param("gid")
	if len(gid) == 0 {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	if err := models.DeletePurchasesDetail(id, gid); err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}
	handler.SendResponse(c, nil, "Successfully deleted.")
}
