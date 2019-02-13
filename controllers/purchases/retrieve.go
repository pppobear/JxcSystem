package purchases

import (
	"github.com/gin-gonic/gin"

	"pppobear.cn/jxc-backend/models"
	"pppobear.cn/jxc-backend/module/errno"
	"pppobear.cn/jxc-backend/module/handler"
)

func Retrieve(c *gin.Context) {
	id := c.Param("id")
	pur, err := models.RetrievePurchase(id)
	if err != nil {
		handler.SendResponse(c, errno.ErrRecordNotFound, nil)
		return
	}
	handler.SendResponse(c, nil, pur)
}
