package purchases

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"pppobear.cn/jxc-backend/models"
	"pppobear.cn/jxc-backend/module/errno"
	"pppobear.cn/jxc-backend/module/handler"
)

func Create(c *gin.Context) {
	var r handler.CreatePurSalRequest
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	// Validate the data.
	if err := r.Validate(); err != nil {
		handler.SendResponse(c, errno.ErrValidation, nil)
		return
	}

	if len(r.Details) > 0 {
		for _, d := range r.Details {
			if err := d.Validate(); err != nil {
				handler.SendResponse(c, errno.ErrValidation, nil)
				return
			}
		}
	}

	purDetails := make([]models.PurchasesDetailModel, 0)

	for _, d := range r.Details {
		purDetails = append(purDetails, models.PurchasesDetailModel{
			UnitPrice: d.UnitPrice,
			GoodsId:   d.GoodsId,
			Number:    d.Number,
		})
	}

	pur := models.PurchasesModel{
		Datetime:   r.Datetime,
		BuyerId:    r.StaffId,
		SupplierId: r.CustomerId,
		Details:    purDetails,
	}

	// Insert the user to the database.
	if err := pur.Create(); err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	// Show the user information.
	handler.SendResponse(c, nil, pur)
}

func AddDetails(c *gin.Context) {
	var r handler.AddDetailRequest
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	r.Id = id
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	// Validate the data.
	if err := r.Validate(); err != nil {
		handler.SendResponse(c, errno.ErrValidation, nil)
		return
	}

	purDetail := models.PurchasesDetailModel{
		PurchaseID: r.Id,
		UnitPrice:  r.UnitPrice,
		GoodsId:    r.GoodsId,
		Number:     r.Number,
	}

	if err := purDetail.Create(); err != nil {
		if strings.Index(err.Error(), "PRIMARY KEY constraint") != -1 {
			handler.SendResponse(c, errno.ErrDataExist, nil)
		} else {
			handler.SendResponse(c, errno.ErrDatabase, nil)
		}
		return
	}

	handler.SendResponse(c, nil, purDetail)
}
