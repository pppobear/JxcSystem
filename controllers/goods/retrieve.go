package goods

import (
	"github.com/gin-gonic/gin"

	"pppobear.cn/jxc-backend/models"
	"pppobear.cn/jxc-backend/module/errno"
	"pppobear.cn/jxc-backend/module/handler"
)

func RetrieveGoods(c *gin.Context) {
	id := c.Param("id")
	goods, err := models.RetrieveGoods(id)
	if err != nil {
		handler.SendResponse(c, errno.ErrRecordNotFound, nil)
		return
	}
	handler.SendResponse(c, nil, goods)
}

func RetrieveInventory(c *gin.Context) {
	id := c.Param("id")
	goods, err := models.RetrieveInventory(id)
	if err != nil {
		handler.SendResponse(c, errno.ErrRecordNotFound, nil)
		return
	}
	handler.SendResponse(c, nil, goods)
}
