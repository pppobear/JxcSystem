package goods

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"pppobear.cn/jxc-backend/models"
	"pppobear.cn/jxc-backend/module/errno"
	"pppobear.cn/jxc-backend/module/handler"
	"pppobear.cn/jxc-backend/module/pagination"
)

type LGRequest struct {
	handler.ListRequest
	models.GoodsFilterRequest
}

type LIRequest struct {
	handler.ListRequest
	models.InventoryFilterRequest
}

func ListGoods(c *gin.Context) {
	var (
		lgr LGRequest
		req handler.ListRequest
		ft  models.GoodsFilterRequest
	)
	if err := c.MustBindWith(&lgr, binding.Form); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	req = lgr.ListRequest
	ft = lgr.GoodsFilterRequest

	resp, err := models.ListGoods(&req, &ft)
	if err == nil {
		pagination.SetPageNav(&resp, c)
	}

	handler.SendResponse(c, err, resp)
}

func ListInventories(c *gin.Context) {
	var (
		lir LIRequest
		req handler.ListRequest
		ft  models.InventoryFilterRequest
	)
	if err := c.MustBindWith(&lir, binding.Form); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	req = lir.ListRequest
	ft = lir.InventoryFilterRequest

	resp, err := models.ListInventory(&req, &ft)
	if err == nil {
		pagination.SetPageNav(&resp, c)
	}

	handler.SendResponse(c, err, resp)
}
