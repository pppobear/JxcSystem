package purchases

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"pppobear.cn/jxc-backend/models"
	"pppobear.cn/jxc-backend/module/errno"
	"pppobear.cn/jxc-backend/module/handler"
)

func Update(c *gin.Context) {
	var r handler.CreatePurSalRequest
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	// Validate the data.
	if err := r.Validate(); err != nil {
		handler.SendResponse(c, errno.ErrValidation, nil)
		return
	}

	pur := models.PurchasesModel{
		Id:         id,
		SupplierId: r.CustomerId,
		BuyerId:    r.StaffId,
		Datetime:   r.Datetime,
	}

	if err := pur.Update(); err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}
	handler.SendResponse(c, nil, pur)
}

func UpdateDetail(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	goodsId := c.Param("gid")
	var r handler.UpdateDetailRequest
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	purD := &models.PurchasesDetailModel{PurchaseID: id, GoodsId: goodsId, Number: r.Number}
	if err := purD.Update(); err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}
	handler.SendResponse(c, nil, purD)
}
