package sales

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

	sales := models.SalesModel{
		Id:           id,
		CustomerId:   r.CustomerId,
		SalesStaffId: r.StaffId,
		Datetime:     r.Datetime,
	}

	if err := sales.Update(); err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}
	handler.SendResponse(c, nil, sales)
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

	salesD := &models.SalesDetailModel{SalesId: id, GoodsId: goodsId, Number: r.Number}
	if err := salesD.Update(); err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}
	handler.SendResponse(c, nil, salesD)
}
